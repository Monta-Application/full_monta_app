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

package services

import (
	"context"
	"fmt"
	"strings"

	"github.com/emoss08/trenova/internal/server"
	"github.com/emoss08/trenova/pkg/models"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/uptrace/bun"
)

// LocationCategoryService handles business logic for LocationCategory
type LocationCategoryService struct {
	db     *bun.DB
	logger *zerolog.Logger
}

// NewLocationCategoryService creates a new instance of LocationCategoryService
func NewLocationCategoryService(s *server.Server) *LocationCategoryService {
	return &LocationCategoryService{
		db:     s.DB,
		logger: s.Logger,
	}
}

// QueryFilter defines the filter parameters for querying LocationCategory
type LocationCategoryQueryFilter struct {
	Query          string
	OrganizationID uuid.UUID
	BusinessUnitID uuid.UUID
	Limit          int
	Offset         int
}

// filterQuery applies filters to the query
func (s LocationCategoryService) filterQuery(q *bun.SelectQuery, f *LocationCategoryQueryFilter) *bun.SelectQuery {
	q = q.Where("lc.organization_id = ?", f.OrganizationID).
		Where("lc.business_unit_id = ?", f.BusinessUnitID)

	if f.Query != "" {
		q = q.Where("lc.name = ? OR lc.description ILIKE ?", f.Query, "%"+strings.ToLower(f.Query)+"%")
	}

	q = q.OrderExpr("CASE WHEN lc.name = ? THEN 0 ELSE 1 END", f.Query).
		Order("lc.created_at DESC")

	return q.Limit(f.Limit).Offset(f.Offset)
}

// GetAll retrieves all LocationCategory based on the provided filter
func (s LocationCategoryService) GetAll(ctx context.Context, filter *LocationCategoryQueryFilter) ([]*models.LocationCategory, int, error) {
	var entities []*models.LocationCategory

	q := s.db.NewSelect().
		Model(&entities)

	q = s.filterQuery(q, filter)

	count, err := q.ScanAndCount(ctx)
	if err != nil {
		s.logger.Error().Err(err).Msg("Failed to fetch LocationCategory")
		return nil, 0, fmt.Errorf("failed to fetch LocationCategory: %w", err)
	}

	return entities, count, nil
}

// Get retrieves a single LocationCategory by ID
func (s LocationCategoryService) Get(ctx context.Context, id, orgID, buID uuid.UUID) (*models.LocationCategory, error) {
	entity := new(models.LocationCategory)
	err := s.db.NewSelect().
		Model(entity).
		Where("lc.organization_id = ?", orgID).
		Where("lc.business_unit_id = ?", buID).
		Where("lc.id = ?", id).
		Scan(ctx)
	if err != nil {
		s.logger.Error().Err(err).Msg("Failed to fetch LocationCategory")
		return nil, fmt.Errorf("failed to fetch LocationCategory: %w", err)
	}

	return entity, nil
}

// Create creates a new LocationCategory
func (s LocationCategoryService) Create(ctx context.Context, entity *models.LocationCategory) (*models.LocationCategory, error) {
	err := s.db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		_, err := tx.NewInsert().
			Model(entity).
			Returning("*").
			Exec(ctx)
		return err
	})
	if err != nil {
		s.logger.Error().Err(err).Msg("Failed to create LocationCategory")
		return nil, fmt.Errorf("failed to create LocationCategory: %w", err)
	}

	return entity, nil
}

// UpdateOne updates an existing LocationCategory
func (s LocationCategoryService) UpdateOne(ctx context.Context, entity *models.LocationCategory) (*models.LocationCategory, error) {
	err := s.db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		if err := entity.OptimisticUpdate(ctx, tx); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		s.logger.Error().Err(err).Msg("Failed to update LocationCategory")
		return nil, fmt.Errorf("failed to update LocationCategory: %w", err)
	}

	return entity, nil
}
