package services

import (
	"context"
	"strings"

	"github.com/emoss08/trenova/internal/server"
	"github.com/emoss08/trenova/pkg/gen"
	"github.com/emoss08/trenova/pkg/models"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/uptrace/bun"
)

type WorkerService struct {
	db      *bun.DB
	logger  *zerolog.Logger
	codeGen *gen.CodeGenerator
}

func NewWorkerService(s *server.Server) *WorkerService {
	return &WorkerService{
		db:      s.DB,
		logger:  s.Logger,
		codeGen: s.CodeGenerator,
	}
}

// QueryFilter defines the filter parameters for querying Worker
type WorkerQueryFilter struct {
	Query          string
	OrganizationID uuid.UUID
	BusinessUnitID uuid.UUID
	Limit          int
	Offset         int
}

func (s WorkerService) filterQuery(q *bun.SelectQuery, f *WorkerQueryFilter) *bun.SelectQuery {
	q = q.Where("wk.organization_id = ?", f.OrganizationID).
		Where("wk.business_unit_id = ?", f.BusinessUnitID)

	if f.Query != "" {
		q = q.Where("wk.code = ? OR wk.code ILIKE ?", f.Query, "%"+strings.ToLower(f.Query)+"%")
	}

	q = q.OrderExpr("CASE WHEN wk.code = ? THEN 0 ELSE 1 END", f.Query).
		Order("wk.created_at DESC")

	return q.Limit(f.Limit).Offset(f.Offset)
}

func (s WorkerService) GetAll(ctx context.Context, filter *WorkerQueryFilter) ([]*models.Worker, int, error) {
	var entities []*models.Worker

	q := s.db.NewSelect().
		Model(&entities).
		Relation("WorkerProfile")

	q = s.filterQuery(q, filter)

	count, err := q.ScanAndCount(ctx)
	if err != nil {
		s.logger.Error().Err(err).Msg("failed to get workers")
		return nil, 0, err
	}

	return entities, count, nil
}

func (s WorkerService) Get(ctx context.Context, id uuid.UUID, orgID, buID uuid.UUID) (*models.Worker, error) {
	entity := new(models.Worker)
	err := s.db.NewSelect().
		Model(entity).
		Where("wk.organization_id = ?", orgID).
		Where("wk.business_unit_id = ?", buID).
		Where("wk.id = ?", id).
		Scan(ctx)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (s WorkerService) Create(ctx context.Context, entity *models.Worker) (*models.Worker, error) {
	err := s.db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		// Query the master key generation entity.
		mkg, mErr := models.QueryWorkerMasterKeyGenerationByOrgID(ctx, s.db, entity.OrganizationID)
		if mErr != nil {
			return mErr
		}

		return entity.InsertWorker(ctx, tx, s.codeGen, mkg.Pattern)
	})
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (s WorkerService) UpdateOne(ctx context.Context, entity *models.Worker) (*models.Worker, error) {
	err := s.db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		return entity.UpdateWorker(ctx, tx)
	})
	if err != nil {
		return nil, err
	}

	return entity, nil
}
