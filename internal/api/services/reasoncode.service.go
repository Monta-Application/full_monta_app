// COPYRIGHT(c) 2024 Trenova
//
// This file is part of Trenova.
//
// The Trenova software is licensed under the Business Source License 1.1. You are granted the right
// to copy, modify, and redistribute the software, but only for non-production use or with a total
// of less than three server instances. Starting from the Change Date (November 16, 2026), the
// software will be made available under version 2 or later of the GNU General Public License.
// If you use the software in violation of this license, your rights under the license will be
// terminated automatically. The software is provided "as is," and the Licensor disclaims all
// warranties and conditions. If you use this license's text or the "Business Source License" name
// and trademark, you must comply with the Licensor's covenants, which include specifying the
// Change License as the GPL Version 2.0 or a compatible license, specifying an Additional Use
// Grant, and not modifying the license in any other way.

package services

import (
	"context"
	"fmt"
	"strings"

	"github.com/emoss08/trenova/config"
	"github.com/emoss08/trenova/internal/api/common"
	"github.com/emoss08/trenova/internal/server"
	"github.com/emoss08/trenova/pkg/models"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

// ReasonCodeService handles business logic for ReasonCode
type ReasonCodeService struct {
	common.AuditableService
	logger *config.ServerLogger
}

// NewReasonCodeService creates a new instance of ReasonCodeService
func NewReasonCodeService(s *server.Server) *ReasonCodeService {
	return &ReasonCodeService{
		AuditableService: common.AuditableService{
			DB:           s.DB,
			AuditService: s.AuditService,
		},
		logger: s.Logger,
	}
}

// ReasonCodeQueryFilter defines the filter parameters for querying ReasonCode.
type ReasonCodeQueryFilter struct {
	Query          string
	OrganizationID uuid.UUID
	BusinessUnitID uuid.UUID
	Limit          int
	Offset         int
}

// filterQuery applies filters to the query
func (s ReasonCodeService) filterQuery(q *bun.SelectQuery, f *ReasonCodeQueryFilter) *bun.SelectQuery {
	q = q.Where("rc.organization_id = ?", f.OrganizationID).
		Where("rc.business_unit_id = ?", f.BusinessUnitID)

	if f.Query != "" {
		q = q.Where("rc.code = ? OR rc.description ILIKE ?", f.Query, "%"+strings.ToLower(f.Query)+"%")
	}

	q = q.OrderExpr("CASE WHEN rc.code = ? THEN 0 ELSE 1 END", f.Query).
		Order("rc.created_at DESC")

	return q.Limit(f.Limit).Offset(f.Offset)
}

// GetAll retrieves all ReasonCode based on the provided filter
func (s ReasonCodeService) GetAll(ctx context.Context, filter *ReasonCodeQueryFilter) ([]*models.ReasonCode, int, error) {
	var entities []*models.ReasonCode

	q := s.DB.NewSelect().Model(&entities)
	q = s.filterQuery(q, filter)

	count, err := q.ScanAndCount(ctx)
	if err != nil {
		s.logger.Error().Err(err).Msg("Failed to fetch ReasonCode")
		return nil, 0, fmt.Errorf("failed to fetch ReasonCode: %w", err)
	}

	return entities, count, nil
}

// Get retrieves a single ReasonCode by ID
func (s ReasonCodeService) Get(ctx context.Context, id, orgID, buID uuid.UUID) (*models.ReasonCode, error) {
	entity := new(models.ReasonCode)
	err := s.GetByID(ctx, id, orgID, buID, entity)
	if err != nil {
		s.logger.Error().Err(err).Msg("Failed to fetch ReasonCode")
		return nil, fmt.Errorf("failed to fetch ReasonCode: %w", err)
	}

	return entity, nil
}

// Create creates a new ReasonCode
func (s ReasonCodeService) Create(ctx context.Context, entity *models.ReasonCode, userID uuid.UUID) (*models.ReasonCode, error) {
	_, err := s.CreateWithAudit(ctx, entity, userID)
	if err != nil {
		s.logger.Error().Err(err).Msg("Failed to create ReasonCode")
		return nil, fmt.Errorf("failed to create ReasonCode: %w", err)
	}

	return entity, nil
}

// UpdateOne updates an existing ReasonCode
func (s ReasonCodeService) UpdateOne(ctx context.Context, entity *models.ReasonCode, userID uuid.UUID) (*models.ReasonCode, error) {
	err := s.UpdateWithAudit(ctx, entity, userID)
	if err != nil {
		s.logger.Error().Err(err).Msg("Failed to update ReasonCode")
		return nil, fmt.Errorf("failed to update ReasonCode: %w", err)
	}

	return entity, nil
}
