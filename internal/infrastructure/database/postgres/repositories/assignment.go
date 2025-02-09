package repositories

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/emoss08/trenova/internal/core/domain/shipment"
	"github.com/emoss08/trenova/internal/core/ports/db"
	"github.com/emoss08/trenova/internal/core/ports/repositories"
	"github.com/emoss08/trenova/internal/pkg/errors"
	"github.com/emoss08/trenova/internal/pkg/logger"
	"github.com/emoss08/trenova/pkg/types/pulid"
	"github.com/rotisserie/eris"
	"github.com/rs/zerolog"
	"github.com/uptrace/bun"
	"go.uber.org/fx"
)

type AssignmentRepositoryParams struct {
	fx.In

	DB           db.Connection
	MoveRepo     repositories.ShipmentMoveRepository
	ShipmentRepo repositories.ShipmentRepository
	Logger       *logger.Logger
}

type assignmentRepository struct {
	db           db.Connection
	moveRepo     repositories.ShipmentMoveRepository
	shipmentRepo repositories.ShipmentRepository
	l            *zerolog.Logger
}

func NewAssignmentRepository(p AssignmentRepositoryParams) repositories.AssignmentRepository {
	log := p.Logger.With().
		Str("repository", "assignment").
		Logger()

	return &assignmentRepository{
		db:           p.DB,
		moveRepo:     p.MoveRepo,
		shipmentRepo: p.ShipmentRepo,
		l:            &log,
	}
}

func (ar *assignmentRepository) GetByID(ctx context.Context, opts repositories.GetAssignmentByIDOptions) (*shipment.Assignment, error) {
	dba, err := ar.db.DB(ctx)
	if err != nil {
		return nil, eris.Wrap(err, "get database connection")
	}

	log := ar.l.With().
		Str("operation", "GetByID").
		Str("assignmentID", opts.ID.String()).
		Logger()

	entity := new(shipment.Assignment)

	err = dba.NewSelect().Model(entity).
		Where("a.id = ? AND a.organization_id = ? AND a.business_unit_id = ?",
			opts.ID, opts.OrganizationID, opts.BusinessUnitID).
		Scan(ctx)
	if err != nil {
		if eris.Is(err, sql.ErrNoRows) {
			return nil, errors.NewNotFoundError("Assignment not found within your organization")
		}
		log.Error().Err(err).Msg("failed to get assignment")
		return nil, err
	}

	return entity, nil
}

func (ar *assignmentRepository) BulkAssign(ctx context.Context, req *repositories.AssignmentRequest) ([]*shipment.Assignment, error) {
	dba, err := ar.db.DB(ctx)
	if err != nil {
		return nil, eris.Wrap(err, "get database connection")
	}

	log := ar.l.With().
		Str("operation", "BulkAssign").
		Str("shipmentID", req.ShipmentID.String()).
		Logger()

	shipmentMoves, err := ar.moveRepo.GetMovesByShipmentID(ctx, repositories.GetMovesByShipmentIDOptions{
		ShipmentID: req.ShipmentID,
		OrgID:      req.OrgID,
		BuID:       req.BuID,
	})
	if err != nil {
		log.Error().Err(err).Msg("failed to get shipment moves")
		return nil, err
	}

	assignments := ar.createAssignments(shipmentMoves, req)
	moveIDs := ar.extractMoveIDs(shipmentMoves)

	err = dba.RunInTx(ctx, nil, func(c context.Context, tx bun.Tx) error {
		return ar.processBulkAssignment(c, tx, assignments, moveIDs, req)
	})
	if err != nil {
		log.Error().Err(err).Msg("failed to process bulk assignment")
		return nil, err
	}

	return assignments, nil
}

// SingleAssign implementation remains mostly the same but with improved error wrapping
func (ar *assignmentRepository) SingleAssign(ctx context.Context, a *shipment.Assignment) (*shipment.Assignment, error) {
	dba, err := ar.db.DB(ctx)
	if err != nil {
		return nil, eris.Wrap(err, "get database connection")
	}

	log := ar.l.With().
		Str("operation", "SingleAssign").
		Str("orgID", a.OrganizationID.String()).
		Logger()

	err = dba.RunInTx(ctx, nil, func(c context.Context, tx bun.Tx) error {
		if err = ar.processAssignment(c, tx, a); err != nil {
			return err
		}
		return ar.updateAssignmentStatuses(c, a)
	})
	if err != nil {
		log.Error().Err(err).Msg("failed to process single assignment")
		return nil, err
	}

	return a, nil
}

// Reassign updates an existing assignment for a shipment move.
func (ar *assignmentRepository) Reassign(ctx context.Context, a *shipment.Assignment) (*shipment.Assignment, error) {
	dba, err := ar.db.DB(ctx)
	if err != nil {
		return nil, eris.Wrap(err, "get database connection")
	}

	log := ar.l.With().
		Str("operation", "Reassign").
		Str("assignmentID", a.ID.String()).
		Logger()

	err = dba.RunInTx(ctx, nil, func(c context.Context, tx bun.Tx) error {
		return ar.processReassignment(c, tx, a)
	})
	if err != nil {
		log.Error().Err(err).Msg("failed to process reassignment")
		return nil, eris.Wrap(err, "process reassignment")
	}

	return a, nil
}

func (ar *assignmentRepository) createAssignments(moves []*shipment.ShipmentMove, req *repositories.AssignmentRequest) []*shipment.Assignment {
	assignments := make([]*shipment.Assignment, len(moves))
	for i, move := range moves {
		assignments[i] = &shipment.Assignment{
			ShipmentMoveID:    move.ID,
			OrganizationID:    req.OrgID,
			BusinessUnitID:    req.BuID,
			PrimaryWorkerID:   req.PrimaryWorkerID,
			TractorID:         req.TractorID,
			TrailerID:         req.TrailerID,
			SecondaryWorkerID: req.SecondaryWorkerID,
		}
	}
	return assignments
}

func (ar *assignmentRepository) extractMoveIDs(moves []*shipment.ShipmentMove) []pulid.ID {
	moveIDs := make([]pulid.ID, len(moves))
	for i, move := range moves {
		moveIDs[i] = move.ID
	}
	return moveIDs
}

func (ar *assignmentRepository) processBulkAssignment(ctx context.Context, tx bun.Tx, assignments []*shipment.Assignment,
	moveIDs []pulid.ID, req *repositories.AssignmentRequest,
) error {
	if err := tx.NewInsert().Model(assignments).Scan(ctx); err != nil {
		return err
	}

	if err := ar.updateMovesStatus(ctx, tx, moveIDs); err != nil {
		return err
	}

	if err := ar.updateShipmentStatus(ctx, req); err != nil {
		return err
	}

	return nil
}

func (ar *assignmentRepository) updateMovesStatus(ctx context.Context, tx bun.Tx, moveIDs []pulid.ID) error {
	_, err := tx.NewUpdate().
		Model((*shipment.ShipmentMove)(nil)).
		Set("status = ?", shipment.MoveStatusAssigned).
		Where("sm.id IN (?)", bun.In(moveIDs)).
		Exec(ctx)
	if err != nil {
		return eris.Wrap(err, "update move status")
	}
	return nil
}

func (ar *assignmentRepository) updateShipmentStatus(ctx context.Context, req *repositories.AssignmentRequest) error {
	_, err := ar.shipmentRepo.UpdateStatus(ctx, &repositories.UpdateShipmentStatusRequest{
		GetOpts: repositories.GetShipmentByIDOptions{
			ID:    req.ShipmentID,
			OrgID: req.OrgID,
			BuID:  req.BuID,
		},
		Status: shipment.StatusAssigned,
	})
	return err
}

func (ar *assignmentRepository) processAssignment(ctx context.Context, tx bun.Tx, a *shipment.Assignment) error {
	_, err := tx.NewInsert().Model(a).Exec(ctx)
	if err != nil {
		return eris.Wrap(err, "insert assignment")
	}
	return nil
}

func (ar *assignmentRepository) updateAssignmentStatuses(ctx context.Context, a *shipment.Assignment) error {
	move, err := ar.moveRepo.GetByID(ctx, repositories.GetMoveByIDOptions{
		MoveID: a.ShipmentMoveID,
		OrgID:  a.OrganizationID,
		BuID:   a.BusinessUnitID,
	})
	if err != nil {
		return err
	}

	// Update move status
	if err = ar.updateMoveStatus(ctx, a); err != nil {
		return err
	}

	// Update shipment status
	return ar.updateLinkedShipmentStatus(ctx, move.ShipmentID, a)
}

func (ar *assignmentRepository) updateMoveStatus(ctx context.Context, a *shipment.Assignment) error {
	_, err := ar.moveRepo.UpdateStatus(ctx, &repositories.UpdateMoveStatusRequest{
		GetMoveOpts: repositories.GetMoveByIDOptions{
			MoveID: a.ShipmentMoveID,
			OrgID:  a.OrganizationID,
			BuID:   a.BusinessUnitID,
		},
		Status: shipment.MoveStatusAssigned,
	})
	return err
}

func (ar *assignmentRepository) updateLinkedShipmentStatus(ctx context.Context, shipmentID pulid.ID, a *shipment.Assignment) error {
	// We need to check if the shipment has any other moves that are not assigned
	moves, err := ar.moveRepo.GetMovesByShipmentID(ctx, repositories.GetMovesByShipmentIDOptions{
		ShipmentID: shipmentID,
		OrgID:      a.OrganizationID,
		BuID:       a.BusinessUnitID,
	})
	if err != nil {
		return err
	}

	// If all moves are assigned, we can update the shipment status to assigned
	// Otherwise, we update the shipment status to partially assigned
	allAssigned := true
	for _, move := range moves {
		if move.Status != shipment.MoveStatusAssigned {
			allAssigned = false
			break
		}
	}

	var status shipment.Status
	if allAssigned {
		status = shipment.StatusAssigned
	} else {
		status = shipment.StatusPartiallyAssigned
	}

	_, err = ar.shipmentRepo.UpdateStatus(ctx, &repositories.UpdateShipmentStatusRequest{
		GetOpts: repositories.GetShipmentByIDOptions{
			ID:    shipmentID,
			OrgID: a.OrganizationID,
			BuID:  a.BusinessUnitID,
		},
		Status: status,
	})
	return err
}

func (ar *assignmentRepository) processReassignment(ctx context.Context, tx bun.Tx, a *shipment.Assignment) error {
	// Increment the version of the assignment
	originalVersion := a.Version
	a.Version++

	res, err := tx.NewUpdate().
		Model(a).
		Set("tractor_id = ?", a.TractorID).
		Set("trailer_id = ?", a.TrailerID).
		Set("primary_worker_id = ?", a.PrimaryWorkerID).
		Set("secondary_worker_id = ?", a.SecondaryWorkerID).
		Set("version = ?", a.Version).
		WhereGroup(" AND ", func(q *bun.UpdateQuery) *bun.UpdateQuery {
			return q.Where("a.id = ?", a.ID).
				Where("a.organization_id = ?", a.OrganizationID).
				Where("a.business_unit_id = ?", a.BusinessUnitID).
				Where("a.version = ?", originalVersion)
		}).
		Exec(ctx)
	if err != nil {
		return err
	}

	// Check if the update affected any rows
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.NewValidationError(
			"version",
			errors.ErrVersionMismatch,
			fmt.Sprintf("Version mismatch. The Assignment (%s) has either been updated or deleted since the last request.", a.ID),
		)
	}

	return nil
}
