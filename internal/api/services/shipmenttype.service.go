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

// ShipmentTypeService handles business logic for ShipmentType
type ShipmentTypeService struct {
	db     *bun.DB
	logger *zerolog.Logger
}

// NewShipmentTypeService creates a new instance of ShipmentTypeService
func NewShipmentTypeService(s *server.Server) *ShipmentTypeService {
	return &ShipmentTypeService{
		db:     s.DB,
		logger: s.Logger,
	}
}

// QueryFilter defines the filter parameters for querying ShipmentType
type ShipmentTypeQueryFilter struct {
	Query          string
	OrganizationID uuid.UUID
	BusinessUnitID uuid.UUID
	Limit          int
	Offset         int
}

// filterQuery applies filters to the query
func (s ShipmentTypeService) filterQuery(q *bun.SelectQuery, f *ShipmentTypeQueryFilter) *bun.SelectQuery {
	q = q.Where("st.organization_id = ?", f.OrganizationID).
		Where("st.business_unit_id = ?", f.BusinessUnitID)

	if f.Query != "" {
		q = q.Where("st.code = ? OR st.description ILIKE ?", f.Query, "%"+strings.ToLower(f.Query)+"%")
	}

	q = q.OrderExpr("CASE WHEN st.code = ? THEN 0 ELSE 1 END", f.Query).
		Order("st.created_at DESC")

	return q.Limit(f.Limit).Offset(f.Offset)
}

// GetAll retrieves all ShipmentType based on the provided filter
func (s ShipmentTypeService) GetAll(ctx context.Context, filter *ShipmentTypeQueryFilter) ([]*models.ShipmentType, int, error) {
	var entities []*models.ShipmentType

	q := s.db.NewSelect().
		Model(&entities)

	q = s.filterQuery(q, filter)

	count, err := q.ScanAndCount(ctx)
	if err != nil {
		s.logger.Error().Err(err).Msg("Failed to fetch ShipmentType")
		return nil, 0, fmt.Errorf("failed to fetch ShipmentType: %w", err)
	}

	return entities, count, nil
}

// Get retrieves a single ShipmentType by ID
func (s ShipmentTypeService) Get(ctx context.Context, id, orgID, buID uuid.UUID) (*models.ShipmentType, error) {
	entity := new(models.ShipmentType)
	err := s.db.NewSelect().
		Model(entity).
		Where("st.organization_id = ?", orgID).
		Where("st.business_unit_id = ?", buID).
		Where("st.id = ?", id).
		Scan(ctx)
	if err != nil {
		s.logger.Error().Err(err).Msg("Failed to fetch ShipmentType")
		return nil, fmt.Errorf("failed to fetch ShipmentType: %w", err)
	}

	return entity, nil
}

// Create creates a new ShipmentType
func (s ShipmentTypeService) Create(ctx context.Context, entity *models.ShipmentType) (*models.ShipmentType, error) {
	err := s.db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		_, err := tx.NewInsert().
			Model(entity).
			Returning("*").
			Exec(ctx)
		return err
	})
	if err != nil {
		s.logger.Error().Err(err).Msg("Failed to create ShipmentType")
		return nil, fmt.Errorf("failed to create ShipmentType: %w", err)
	}

	return entity, nil
}

// UpdateOne updates an existing ShipmentType
func (s ShipmentTypeService) UpdateOne(ctx context.Context, entity *models.ShipmentType) (*models.ShipmentType, error) {
	err := s.db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		if err := entity.OptimisticUpdate(ctx, tx); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		s.logger.Error().Err(err).Msg("Failed to update ShipmentType")
		return nil, fmt.Errorf("failed to update ShipmentType: %w", err)
	}

	return entity, nil
}
