// Copyright (c) 2024 Trenova Technologies, LLC
//
// Licensed under the Business Source License 1.1 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://trenova.app/pricing/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// Key Terms:
// - Non-production use only
// - Change Date: 2026-11-16
// - Change License: GNU General Public License v2 or later
//
// For full license text, see the LICENSE file in the root directory.

package testutils

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/emoss08/trenova/migrate/migrations"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
	"github.com/uptrace/bun/migrate"
)

type TestDBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func getDBConfig() TestDBConfig {
	return TestDBConfig{
		Host:     getEnv("TEST_DB_HOST", "localhost"),
		Port:     getEnv("TEST_DB_PORT", "5432"),
		User:     getEnv("TEST_DB_USER", "postgres"),
		Password: getEnv("TEST_DB_PASSWORD", "postgres"),
		DBName:   getEnv("TEST_DB_NAME", "test_db"),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

// TestDB encapsulates the database connection and provides utility methods
type TestDB struct {
	DB   *bun.DB
	once sync.Once
}

// NewTestDB creates a new test database connection
func NewTestDB() (*TestDB, error) {
	config := getDBConfig()
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.User, config.Password, config.Host, config.Port, config.DBName)

	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	// Set connection pool parameters
	sqldb.SetMaxOpenConns(4)
	sqldb.SetMaxIdleConns(4)
	sqldb.SetConnMaxLifetime(time.Hour)

	db := bun.NewDB(sqldb, pgdialect.New())

	// Add query hook for debugging
	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.FromEnv("BUNDEBUG"),
	))

	return &TestDB{DB: db}, nil
}

// GetDB returns the bun.DB instance
func (tdb *TestDB) GetDB() *bun.DB {
	return tdb.DB
}

// InitSchema initializes the database schema
func (tdb *TestDB) InitSchema() error {
	var err error
	tdb.once.Do(func() {
		ctx := context.Background()
		migrator := migrate.NewMigrator(tdb.DB, migrations.Migrations)

		// Initialize the migration tables
		if err = migrator.Init(ctx); err != nil {
			err = fmt.Errorf("failed to initialize migrator: %v", err)
			return
		}

		if err = migrator.Lock(ctx); err != nil {
			err = fmt.Errorf("failed to acquire migration lock: %v", err)
			return
		}
		defer migrator.Unlock(ctx)

		group, migrateErr := migrator.Migrate(ctx)
		if migrateErr != nil {
			err = fmt.Errorf("failed to run migrations: %v", migrateErr)
			return
		}

		if group.IsZero() {
			fmt.Println("There are no new migrations to run (database is up to date)")
		} else {
			fmt.Printf("Applied %d migrations\n", len(group.Migrations))
		}
	})
	return err
}

func (tdb *TestDB) ResetDatabase() error {
	ctx := context.Background()

	// Drop all tables, including migration tables
	if err := tdb.DropAllTablesAndTypes(ctx); err != nil {
		return fmt.Errorf("failed to drop tables and types: %w", err)
	}

	// Recreate migration tables
	migrator := migrate.NewMigrator(tdb.DB, migrations.Migrations)
	if err := migrator.Init(ctx); err != nil {
		return fmt.Errorf("failed to initialize migrator: %w", err)
	}

	// Run migrations to recreate schema
	group, err := migrator.Migrate(ctx)
	if err != nil {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	if group.IsZero() {
		fmt.Println("No migrations were applied")
	} else {
		fmt.Printf("Applied %d migrations\n", len(group.Migrations))
	}

	return nil
}

func (tdb *TestDB) DropAllTablesAndTypes(ctx context.Context) error {
	query := `
	DO $$ 
	DECLARE 
		r RECORD;
	BEGIN
		-- Disable triggers
		EXECUTE 'SET session_replication_role = replica';

		-- Drop all tables
		FOR r IN (SELECT tablename FROM pg_tables WHERE schemaname = current_schema()) LOOP
			EXECUTE 'DROP TABLE IF EXISTS ' || quote_ident(r.tablename) || ' CASCADE';
		END LOOP;

		-- Drop all custom types
		FOR r IN (SELECT typname FROM pg_type WHERE typtype = 'e' AND typnamespace = (SELECT oid FROM pg_namespace WHERE nspname = current_schema())) LOOP
			EXECUTE 'DROP TYPE IF EXISTS ' || quote_ident(r.typname) || ' CASCADE';
		END LOOP;

		-- Re-enable triggers
		EXECUTE 'SET session_replication_role = DEFAULT';
	END $$;
	`

	_, err := tdb.DB.ExecContext(ctx, query)
	return err
}

func (tdb *TestDB) Close() error {
	return tdb.DB.Close()
}

// WithTransaction runs a function within a transaction
func (tdb *TestDB) WithTransaction(fn func(*bun.Tx) error) error {
	return tdb.DB.RunInTx(context.Background(), nil, func(ctx context.Context, tx bun.Tx) error {
		return fn(&tx)
	})
}

// SetupTestCase prepares the database for a test case
// SetupTestCase prepares the database for a test case
func SetupTestCase(t *testing.T) (*TestDB, func()) {
	t.Helper()

	testDB, err := NewTestDB()
	if err != nil {
		t.Fatalf("Failed to create test database: %v", err)
	}

	// Reset the database to a clean state
	if err := testDB.ResetDatabase(); err != nil {
		t.Fatalf("Failed to reset database: %v", err)
	}

	return testDB, func() {
		if err := testDB.Close(); err != nil {
			t.Errorf("Failed to close database connection: %v", err)
		}
	}
}
