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

package config

import (
	"crypto/rsa"
	"fmt"

	"github.com/fatih/color"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

var configEnv string // This will store the value of the environment flag

type FiberServer struct {
	// ListenAddress is the address that the server will listen on.
	ListenAddress string

	// Enable prefork on the server instance.
	//
	// When prefork is enabled, the server will fork itself into multiple processes to handle incoming requests.
	// This can be useful to take advantage of multiple CPU cores.
	//
	// When prefork is disabled, the server will run in a single process and will only be able to take advantage of a single CPU core.
	// By default, prefork is disabled.
	//
	// Read More: https://github.com/gofiber/fiber/issues/180
	EnablePrefork bool

	// Print out all the routes to the console.
	EnablePrintRoutes bool

	// Enable logging middleware on the server instance.
	//
	// Read More: https://docs.gofiber.io/contrib/fiberzap/
	EnableLoggingMiddleware bool

	// Enable helmet middleware on the server instance.
	//
	// Read More: https://docs.gofiber.io/api/middleware/helmet
	EnableHelmetMiddleware bool

	// Enable request ID middleware on the server instance.
	//
	// Read More: https://docs.gofiber.io/api/middleware/requestid
	EnableRequestIDMiddleware bool

	// Enable recover middleware on the server instance.
	//
	// Read More: https://docs.gofiber.io/api/middleware/recover
	EnableRecoverMiddleware bool

	// Enable CORS middleware on the server instance.
	//
	// Read More: https://docs.gofiber.io/api/middleware/cors
	EnableCORSMiddleware bool

	// Enable Idempotency middleware on the server instance.
	//
	// Read More: https://docs.gofiber.io/api/middleware/idempotency
	EnableIdempotencyMiddleware bool

	// Enable Prometheus middleware on the server instance.
	EnablePrometheusMiddleware bool
}

type Integration struct {
	// GenerateReportEndpoint is the URL of the endpoint that will generate a report.
	GenerateReportEndpoint string
}

type Meilisearch struct {
	// Host is the URL of the MeiliSearch server.
	Host string

	// Token is the API key used to authenticate with the MeiliSearch server.
	Token string
}

type Auth struct {
	// PrivateKey is the RSA private key used to sign JWT tokens.
	PrivateKey *rsa.PrivateKey

	// PublicKey is the RSA public key used to verify JWT tokens.
	PublicKey *rsa.PublicKey
}

type Cors struct {
	// Allowed Origins for Cors Middleware.
	// Example: "https://localhost:5173, https://localhost:4173"
	AllowedOrigins string

	// Allowed Headers for Cors Middleware.
	// Example: "Authorization, Origin, Content-Type, Accept, X-CSRF-Token, X-Idempotency-Key"
	AllowedHeaders string

	// Allowed Methods for Cors Middleware.
	// Example: "GET, POST, PUT, DELETE, OPTIONS"
	AllowedMethods string

	// AllowCredentials for Cors Middleware.
	AllowCredentials bool

	// MaxAge for Cors Middleware.
	MaxAge int
}

type Logger struct {
	// Level is the log level for the logger.
	Level zerolog.Level

	// PrettyPrintConsole will print the logs in a human-readable format.
	PrettyPrintConsole bool
}

type Minio struct {
	// Endpoint is the URL of the Minio server.
	Endpoint string `json:"-"`

	// AccessKey is the access key used to authenticate with the Minio server.
	AccessKey string `json:"-"`

	// SecretKey is the secret key used to authenticate with the Minio server.
	SecretKey string `json:"-"`

	// UseSSL is a flag to determine if the Minio client should use SSL.
	UseSSL bool
}

type KafkaServer struct {
	// Brokers is the list of Kafka brokers.
	// Example: "localhost:9092, localhost:9093"
	Broker string
}

type CasbinConfig struct {
	// ModelPath is the path to the Casbin model file.
	ModelPath string
}

type Server struct {
	// FiberServer contains configuration options for the Fiber server.
	Fiber FiberServer

	// Database contains configuration options for the database.
	DB Database

	// Mellisearch contains configuration options for the MeiliSearch server.
	// Meilisearch Meilisearch

	// Auth contains configuration options for the JWT authentication.
	Auth Auth

	// Cors contains configuration options for the CORS middleware.
	Cors Cors

	// Logger contains configuration options for the logger.
	Logger Logger

	// Minio contains configuration options for the Minio server.
	Minio Minio

	// Kafka contains configuration options for the Kafka server.
	Kafka KafkaServer

	// Integration contains configuration options for the integration services.
	Integration Integration

	// Casbin contains configuration options for the Casbin authorization library.
	Casbin CasbinConfig
}

func DefaultServiceConfigFromEnv() (Server, error) {
	v := viper.New()

	configName := "config"
	if configEnv != "" {
		configName = fmt.Sprintf("config.%s", configEnv)
	}
	v.SetConfigName(configName)
	v.SetConfigType("yaml")

	v.AddConfigPath(".")
	v.AddConfigPath("./config")
	v.AddConfigPath("$HOME/.trenova")
	v.AddConfigPath("/etc/trenova")

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return Server{}, fmt.Errorf("config file not found: %w", err)
		}
		return Server{}, fmt.Errorf("failed to read config file: %w", err)
	}

	configFile := v.ConfigFileUsed()

	c := color.New(color.FgCyan, color.Underline, color.Bold)
	c.Printf("Using configuration file: %s\n", configFile)

	var config Server
	if err := v.Unmarshal(&config); err != nil {
		return Server{}, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	config.Logger.Level = zerolog.Level(v.GetInt("logger.level"))
	config.DB.ConnMaxLifetime = v.GetDuration("db.connMaxLifetime")

	return config, nil
}

// SetConfigEnv sets the configuration environment
func SetConfigEnv(env string) {
	configEnv = env
}
