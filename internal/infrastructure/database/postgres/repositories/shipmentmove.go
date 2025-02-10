package repositories

import (
	"context"
	"fmt"

	"github.com/emoss08/trenova/internal/core/domain/shipment"
	"github.com/emoss08/trenova/internal/core/ports/db"
	"github.com/emoss08/trenova/internal/core/ports/repositories"
	"github.com/emoss08/trenova/internal/pkg/errors"
	"github.com/emoss08/trenova/internal/pkg/logger"
	"github.com/rotisserie/eris"
	"github.com/rs/zerolog"
	"github.com/uptrace/bun"
	"go.uber.org/fx"
)

type ShipmentMoveRepositoryParams struct {
	fx.In

	DB     db.Connection
	Logger *logger.Logger
}

type shipmentMoveRepository struct {
	db db.Connection
	l  *zerolog.Logger
}

func NewShipmentMoveRepository(p ShipmentMoveRepositoryParams) repositories.ShipmentMoveRepository {
	log := p.Logger.With().
		Str("repository", "shipmentmove").
		Logger()

	return &shipmentMoveRepository{
		db: p.DB,
		l:  &log,
	}
}

func (sr *shipmentMoveRepository) GetByID(ctx context.Context, opts repositories.GetMoveByIDOptions) (*shipment.ShipmentMove, error) {
	dba, err := sr.db.DB(ctx)
	if err != nil {
		return nil, eris.Wrap(err, "get database connection")
	}

	log := sr.l.With().
		Str("operation", "GetByID").
		Str("moveID", opts.MoveID.String()).
		Str("orgID", opts.OrgID.String()).
		Str("buID", opts.BuID.String()).
		Logger()

	move := new(shipment.ShipmentMove)

	q := dba.NewSelect().Model(move).
		Where("sm.id = ?", opts.MoveID).
		Where("sm.organization_id = ?", opts.OrgID).
		Where("sm.business_unit_id = ?", opts.BuID)

	if err = q.Scan(ctx); err != nil {
		log.Error().Err(err).
			Interface("move", move).
			Msg("failed to get move by id")
		return nil, err
	}

	return move, nil
}

func (sr *shipmentMoveRepository) BulkUpdateStatus(ctx context.Context, req repositories.BulkUpdateMoveStatusRequest) ([]*shipment.ShipmentMove, error) {
	dba, err := sr.db.DB(ctx)
	if err != nil {
		return nil, eris.Wrap(err, "get database connection")
	}

	log := sr.l.With().
		Str("operation", "BulkUpdateStatus").
		Interface("moveIDs", req.MoveIDs).
		Str("status", string(req.Status)).
		Logger()

	moves := make([]*shipment.ShipmentMove, len(req.MoveIDs))
	results, err := dba.NewUpdate().
		Model(moves).
		Set("status = ?", req.Status).
		Where("sm.id IN (?)", bun.In(req.MoveIDs)).
		Returning("*").
		Exec(ctx)
	if err != nil {
		log.Error().Err(err).Msg("failed to bulk update move status")
		return nil, err
	}

	rows, err := results.RowsAffected()
	if err != nil {
		log.Error().Err(err).Msg("failed to get rows affected")
		return nil, err
	}

	if rows != int64(len(req.MoveIDs)) {
		return nil, errors.NewValidationError(
			"move.status",
			errors.ErrVersionMismatch,
			fmt.Sprintf("Version mismatch. The move (%s) has been updated since your last request.", moves[0].ID),
		)
	}

	return moves, nil
}

func (sr *shipmentMoveRepository) UpdateStatus(ctx context.Context, opts *repositories.UpdateMoveStatusRequest) (*shipment.ShipmentMove, error) {
	dba, err := sr.db.DB(ctx)
	if err != nil {
		return nil, eris.Wrap(err, "get database connection")
	}

	log := sr.l.With().
		Str("operation", "UpdateStatus").
		Str("moveID", opts.GetMoveOpts.MoveID.String()).
		Str("status", string(opts.Status)).
		Logger()

	// Get the move
	move, err := sr.GetByID(ctx, opts.GetMoveOpts)
	if err != nil {
		return nil, err
	}

	err = dba.RunInTx(ctx, nil, func(c context.Context, tx bun.Tx) error {
		// Update the move version
		ov := move.Version
		move.Version++

		results, rErr := tx.NewUpdate().Model(move).
			WherePK().
			Where("sm.version = ?", ov).
			Set("status = ?", opts.Status).
			Returning("*").
			Exec(c)
		if rErr != nil {
			log.Error().Err(rErr).
				Interface("move", move).
				Msg("failed to update move version")
			return rErr
		}

		rows, roErr := results.RowsAffected()
		if roErr != nil {
			log.Error().Err(roErr).
				Interface("move", move).
				Msg("failed to get rows affected")
			return roErr
		}

		if rows == 0 {
			return errors.NewValidationError(
				"move.version",
				errors.ErrVersionMismatch,
				fmt.Sprintf("Version mismatch. The move (%s) has been updated since your last request.", move.ID),
			)
		}

		return nil
	})
	if err != nil {
		log.Error().Err(err).
			Interface("move", move).
			Msg("failed to update move status")
		return nil, err
	}

	return move, nil
}

func (sr *shipmentMoveRepository) GetMovesByShipmentID(ctx context.Context, opts repositories.GetMovesByShipmentIDOptions) ([]*shipment.ShipmentMove, error) {
	dba, err := sr.db.DB(ctx)
	if err != nil {
		return nil, eris.Wrap(err, "get database connection")
	}

	log := sr.l.With().
		Str("operation", "GetMovesByShipmentID").
		Str("shipmentID", opts.ShipmentID.String()).
		Logger()

	moves := make([]*shipment.ShipmentMove, 0)

	q := dba.NewSelect().Model(&moves).
		Where("sm.shipment_id = ?", opts.ShipmentID).
		Where("sm.organization_id = ?", opts.OrgID).
		Where("sm.business_unit_id = ?", opts.BuID)

	if err = q.Scan(ctx); err != nil {
		log.Error().Err(err).Msg("failed to get moves by shipment id")
		return nil, err
	}

	return moves, nil
}

func (sr *shipmentMoveRepository) BulkInsert(ctx context.Context, moves []*shipment.ShipmentMove) ([]*shipment.ShipmentMove, error) {
	dba, err := sr.db.DB(ctx)
	if err != nil {
		return nil, eris.Wrap(err, "get database connection")
	}

	log := sr.l.With().
		Str("operation", "BulkInsert").
		Interface("moves", moves).
		Logger()

	err = dba.RunInTx(ctx, nil, func(c context.Context, tx bun.Tx) error {
		if _, err = tx.NewInsert().Model(&moves).Exec(c); err != nil {
			log.Error().Err(err).Msg("failed to bulk insert moves")
			return err
		}

		return nil
	})
	if err != nil {
		log.Error().Err(err).Msg("failed to bulk insert moves and stops")
		return nil, err
	}

	return moves, nil
}
