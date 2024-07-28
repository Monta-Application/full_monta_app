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

package fixtures

import (
	"context"

	"github.com/emoss08/trenova/pkg/audit"
	"github.com/emoss08/trenova/pkg/gen"
	"github.com/emoss08/trenova/pkg/models"
	"github.com/emoss08/trenova/pkg/models/property"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

func loadCustomers(ctx context.Context, db *bun.DB, gen *gen.CodeGenerator, auditService *audit.Service, user *models.User, orgID, buID uuid.UUID) error {
	count, err := db.NewSelect().Model((*models.Customer)(nil)).Count(ctx)
	if err != nil {
		return err
	}

	state := new(models.UsState)
	err = db.NewSelect().Model(state).Where("abbreviation = ?", "AL").Scan(ctx)
	if err != nil {
		return err
	}

	if count < 20 {
		for i := 0; i < 20; i++ {
			customer := models.Customer{
				BusinessUnitID:      buID,
				OrganizationID:      orgID,
				Status:              property.StatusActive,
				Name:                "TEST",
				AddressLine1:        "123 Main St",
				City:                "Minneapolis",
				StateID:             state.ID,
				PostalCode:          "55401",
				AutoMarkReadyToBill: true,
			}

			err = db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
				mkg, mErr := models.QueryCustomerMasterKeyGenerationByOrgID(ctx, db, orgID)
				if mErr != nil {
					return mErr
				}

				auditUser := audit.AuditUser{
					ID:       user.ID,
					Username: user.Username,
				}

				return customer.InsertWithCodeGen(ctx, tx, gen, mkg.Pattern, auditService, auditUser)
			})
			if err != nil {
				return err
			}
		}

		return nil
	}

	return nil
}
