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

package models

import (
	"context"
	"fmt"
	"time"

	"github.com/emoss08/trenova/pkg/validator"
	validation "github.com/go-ozzo/ozzo-validation/v4"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Commodity struct {
	bun.BaseModel `bun:"table:commodities,alias:com" json:"-"`

	ID            uuid.UUID `bun:",pk,type:uuid,default:uuid_generate_v4()" json:"id"`
	Name          string    `bun:"type:VARCHAR(100),notnull" json:"name" queryField:"true"`
	Status        string    `bun:"type:status_enum,notnull,default:'Active'" json:"status"`
	IsHazmat      bool      `bun:"type:boolean,notnull,default:false" json:"isHazmat"`
	UnitOfMeasure string    `bun:"type:VARCHAR(50),notnull" json:"unitOfMeasure"`
	MinTemp       *int16    `bun:"type:integer,nullzero" json:"minTemp"`
	MaxTemp       *int16    `bun:"type:integer,nullzero" json:"maxTemp"`
	Version       int64     `bun:"type:BIGINT" json:"version"`
	CreatedAt     time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"createdAt"`
	UpdatedAt     time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"updatedAt"`

	BusinessUnitID      uuid.UUID `bun:"type:uuid,notnull" json:"businessUnitId"`
	OrganizationID      uuid.UUID `bun:"type:uuid,notnull" json:"organizationId"`
	HazardousMaterialID uuid.UUID `bun:"type:uuid" json:"hazardousMaterialId"`

	BusinessUnit      *BusinessUnit      `bun:"rel:belongs-to,join:business_unit_id=id" json:"-"`
	Organization      *Organization      `bun:"rel:belongs-to,join:organization_id=id" json:"-"`
	HazardousMaterial *HazardousMaterial `bun:"rel:belongs-to,join:hazardous_material_id=id" json:"-"`
}

func (c Commodity) Validate() error {
	return validation.ValidateStruct(
		&c,
		validation.Field(&c.Name, validation.Required, validation.Length(1, 50)),
		validation.Field(&c.BusinessUnitID, validation.Required),
		validation.Field(&c.OrganizationID, validation.Required),
		validation.Field(&c.IsHazmat,
			validation.Required.Error("IsHazmat is required. Please Try again."),
			validation.When(c.HazardousMaterialID != uuid.Nil, validation.Required.Error("Hazardous Material is required when IsHazmat is true. Please try again."))),
	)
}

func (c *Commodity) BeforeUpdate(_ context.Context) error {
	c.Version++

	return nil
}

func (c *Commodity) OptimisticUpdate(ctx context.Context, tx bun.IDB) error {
	ov := c.Version

	if err := c.BeforeUpdate(ctx); err != nil {
		return err
	}

	result, err := tx.NewUpdate().
		Model(c).
		WherePK().
		Where("version = ?", ov).
		Returning("*").
		Exec(ctx)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return &validator.BusinessLogicError{
			Message: fmt.Sprintf("Version mismatch. The Commodity (ID: %s) has been updated by another user. Please refresh and try again.", c.ID),
		}
	}

	return nil
}

var _ bun.BeforeAppendModelHook = (*Commodity)(nil)

func (c *Commodity) BeforeAppendModel(_ context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		c.CreatedAt = time.Now()
	case *bun.UpdateQuery:
		c.UpdatedAt = time.Now()
	}
	return nil
}
