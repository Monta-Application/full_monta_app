// Code generated by entc, DO NOT EDIT.

package shipmentcharges

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/emoss08/trenova/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldLTE(FieldID, id))
}

// BusinessUnitID applies equality check predicate on the "business_unit_id" field. It's identical to BusinessUnitIDEQ.
func BusinessUnitID(v uuid.UUID) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldEQ(FieldBusinessUnitID, v))
}

// OrganizationID applies equality check predicate on the "organization_id" field. It's identical to OrganizationIDEQ.
func OrganizationID(v uuid.UUID) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldEQ(FieldOrganizationID, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldEQ(FieldUpdatedAt, v))
}

// Version applies equality check predicate on the "version" field. It's identical to VersionEQ.
func Version(v int) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldEQ(FieldVersion, v))
}

// ShipmentID applies equality check predicate on the "shipment_id" field. It's identical to ShipmentIDEQ.
func ShipmentID(v uuid.UUID) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldEQ(FieldShipmentID, v))
}

// AccessorialChargeID applies equality check predicate on the "accessorial_charge_id" field. It's identical to AccessorialChargeIDEQ.
func AccessorialChargeID(v uuid.UUID) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldEQ(FieldAccessorialChargeID, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldEQ(FieldDescription, v))
}

// ChargeAmount applies equality check predicate on the "charge_amount" field. It's identical to ChargeAmountEQ.
func ChargeAmount(v float64) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldEQ(FieldChargeAmount, v))
}

// Units applies equality check predicate on the "units" field. It's identical to UnitsEQ.
func Units(v int) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldEQ(FieldUnits, v))
}

// SubTotal applies equality check predicate on the "sub_total" field. It's identical to SubTotalEQ.
func SubTotal(v float64) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldEQ(FieldSubTotal, v))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v uuid.UUID) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldEQ(FieldCreatedBy, v))
}

// BusinessUnitIDEQ applies the EQ predicate on the "business_unit_id" field.
func BusinessUnitIDEQ(v uuid.UUID) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldEQ(FieldBusinessUnitID, v))
}

// BusinessUnitIDNEQ applies the NEQ predicate on the "business_unit_id" field.
func BusinessUnitIDNEQ(v uuid.UUID) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldNEQ(FieldBusinessUnitID, v))
}

// BusinessUnitIDIn applies the In predicate on the "business_unit_id" field.
func BusinessUnitIDIn(vs ...uuid.UUID) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldIn(FieldBusinessUnitID, vs...))
}

// BusinessUnitIDNotIn applies the NotIn predicate on the "business_unit_id" field.
func BusinessUnitIDNotIn(vs ...uuid.UUID) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldNotIn(FieldBusinessUnitID, vs...))
}

// OrganizationIDEQ applies the EQ predicate on the "organization_id" field.
func OrganizationIDEQ(v uuid.UUID) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldEQ(FieldOrganizationID, v))
}

// OrganizationIDNEQ applies the NEQ predicate on the "organization_id" field.
func OrganizationIDNEQ(v uuid.UUID) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldNEQ(FieldOrganizationID, v))
}

// OrganizationIDIn applies the In predicate on the "organization_id" field.
func OrganizationIDIn(vs ...uuid.UUID) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldIn(FieldOrganizationID, vs...))
}

// OrganizationIDNotIn applies the NotIn predicate on the "organization_id" field.
func OrganizationIDNotIn(vs ...uuid.UUID) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldNotIn(FieldOrganizationID, vs...))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldLTE(FieldUpdatedAt, v))
}

// VersionEQ applies the EQ predicate on the "version" field.
func VersionEQ(v int) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldEQ(FieldVersion, v))
}

// VersionNEQ applies the NEQ predicate on the "version" field.
func VersionNEQ(v int) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldNEQ(FieldVersion, v))
}

// VersionIn applies the In predicate on the "version" field.
func VersionIn(vs ...int) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldIn(FieldVersion, vs...))
}

// VersionNotIn applies the NotIn predicate on the "version" field.
func VersionNotIn(vs ...int) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldNotIn(FieldVersion, vs...))
}

// VersionGT applies the GT predicate on the "version" field.
func VersionGT(v int) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldGT(FieldVersion, v))
}

// VersionGTE applies the GTE predicate on the "version" field.
func VersionGTE(v int) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldGTE(FieldVersion, v))
}

// VersionLT applies the LT predicate on the "version" field.
func VersionLT(v int) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldLT(FieldVersion, v))
}

// VersionLTE applies the LTE predicate on the "version" field.
func VersionLTE(v int) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldLTE(FieldVersion, v))
}

// ShipmentIDEQ applies the EQ predicate on the "shipment_id" field.
func ShipmentIDEQ(v uuid.UUID) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldEQ(FieldShipmentID, v))
}

// ShipmentIDNEQ applies the NEQ predicate on the "shipment_id" field.
func ShipmentIDNEQ(v uuid.UUID) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldNEQ(FieldShipmentID, v))
}

// ShipmentIDIn applies the In predicate on the "shipment_id" field.
func ShipmentIDIn(vs ...uuid.UUID) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldIn(FieldShipmentID, vs...))
}

// ShipmentIDNotIn applies the NotIn predicate on the "shipment_id" field.
func ShipmentIDNotIn(vs ...uuid.UUID) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldNotIn(FieldShipmentID, vs...))
}

// AccessorialChargeIDEQ applies the EQ predicate on the "accessorial_charge_id" field.
func AccessorialChargeIDEQ(v uuid.UUID) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldEQ(FieldAccessorialChargeID, v))
}

// AccessorialChargeIDNEQ applies the NEQ predicate on the "accessorial_charge_id" field.
func AccessorialChargeIDNEQ(v uuid.UUID) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldNEQ(FieldAccessorialChargeID, v))
}

// AccessorialChargeIDIn applies the In predicate on the "accessorial_charge_id" field.
func AccessorialChargeIDIn(vs ...uuid.UUID) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldIn(FieldAccessorialChargeID, vs...))
}

// AccessorialChargeIDNotIn applies the NotIn predicate on the "accessorial_charge_id" field.
func AccessorialChargeIDNotIn(vs ...uuid.UUID) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldNotIn(FieldAccessorialChargeID, vs...))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldContainsFold(FieldDescription, v))
}

// ChargeAmountEQ applies the EQ predicate on the "charge_amount" field.
func ChargeAmountEQ(v float64) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldEQ(FieldChargeAmount, v))
}

// ChargeAmountNEQ applies the NEQ predicate on the "charge_amount" field.
func ChargeAmountNEQ(v float64) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldNEQ(FieldChargeAmount, v))
}

// ChargeAmountIn applies the In predicate on the "charge_amount" field.
func ChargeAmountIn(vs ...float64) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldIn(FieldChargeAmount, vs...))
}

// ChargeAmountNotIn applies the NotIn predicate on the "charge_amount" field.
func ChargeAmountNotIn(vs ...float64) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldNotIn(FieldChargeAmount, vs...))
}

// ChargeAmountGT applies the GT predicate on the "charge_amount" field.
func ChargeAmountGT(v float64) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldGT(FieldChargeAmount, v))
}

// ChargeAmountGTE applies the GTE predicate on the "charge_amount" field.
func ChargeAmountGTE(v float64) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldGTE(FieldChargeAmount, v))
}

// ChargeAmountLT applies the LT predicate on the "charge_amount" field.
func ChargeAmountLT(v float64) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldLT(FieldChargeAmount, v))
}

// ChargeAmountLTE applies the LTE predicate on the "charge_amount" field.
func ChargeAmountLTE(v float64) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldLTE(FieldChargeAmount, v))
}

// UnitsEQ applies the EQ predicate on the "units" field.
func UnitsEQ(v int) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldEQ(FieldUnits, v))
}

// UnitsNEQ applies the NEQ predicate on the "units" field.
func UnitsNEQ(v int) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldNEQ(FieldUnits, v))
}

// UnitsIn applies the In predicate on the "units" field.
func UnitsIn(vs ...int) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldIn(FieldUnits, vs...))
}

// UnitsNotIn applies the NotIn predicate on the "units" field.
func UnitsNotIn(vs ...int) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldNotIn(FieldUnits, vs...))
}

// UnitsGT applies the GT predicate on the "units" field.
func UnitsGT(v int) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldGT(FieldUnits, v))
}

// UnitsGTE applies the GTE predicate on the "units" field.
func UnitsGTE(v int) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldGTE(FieldUnits, v))
}

// UnitsLT applies the LT predicate on the "units" field.
func UnitsLT(v int) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldLT(FieldUnits, v))
}

// UnitsLTE applies the LTE predicate on the "units" field.
func UnitsLTE(v int) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldLTE(FieldUnits, v))
}

// SubTotalEQ applies the EQ predicate on the "sub_total" field.
func SubTotalEQ(v float64) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldEQ(FieldSubTotal, v))
}

// SubTotalNEQ applies the NEQ predicate on the "sub_total" field.
func SubTotalNEQ(v float64) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldNEQ(FieldSubTotal, v))
}

// SubTotalIn applies the In predicate on the "sub_total" field.
func SubTotalIn(vs ...float64) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldIn(FieldSubTotal, vs...))
}

// SubTotalNotIn applies the NotIn predicate on the "sub_total" field.
func SubTotalNotIn(vs ...float64) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldNotIn(FieldSubTotal, vs...))
}

// SubTotalGT applies the GT predicate on the "sub_total" field.
func SubTotalGT(v float64) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldGT(FieldSubTotal, v))
}

// SubTotalGTE applies the GTE predicate on the "sub_total" field.
func SubTotalGTE(v float64) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldGTE(FieldSubTotal, v))
}

// SubTotalLT applies the LT predicate on the "sub_total" field.
func SubTotalLT(v float64) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldLT(FieldSubTotal, v))
}

// SubTotalLTE applies the LTE predicate on the "sub_total" field.
func SubTotalLTE(v float64) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldLTE(FieldSubTotal, v))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v uuid.UUID) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v uuid.UUID) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...uuid.UUID) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...uuid.UUID) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// HasBusinessUnit applies the HasEdge predicate on the "business_unit" edge.
func HasBusinessUnit() predicate.ShipmentCharges {
	return predicate.ShipmentCharges(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, BusinessUnitTable, BusinessUnitColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBusinessUnitWith applies the HasEdge predicate on the "business_unit" edge with a given conditions (other predicates).
func HasBusinessUnitWith(preds ...predicate.BusinessUnit) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(func(s *sql.Selector) {
		step := newBusinessUnitStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOrganization applies the HasEdge predicate on the "organization" edge.
func HasOrganization() predicate.ShipmentCharges {
	return predicate.ShipmentCharges(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, OrganizationTable, OrganizationColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOrganizationWith applies the HasEdge predicate on the "organization" edge with a given conditions (other predicates).
func HasOrganizationWith(preds ...predicate.Organization) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(func(s *sql.Selector) {
		step := newOrganizationStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasShipment applies the HasEdge predicate on the "shipment" edge.
func HasShipment() predicate.ShipmentCharges {
	return predicate.ShipmentCharges(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ShipmentTable, ShipmentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasShipmentWith applies the HasEdge predicate on the "shipment" edge with a given conditions (other predicates).
func HasShipmentWith(preds ...predicate.Shipment) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(func(s *sql.Selector) {
		step := newShipmentStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasAccessorialCharge applies the HasEdge predicate on the "accessorial_charge" edge.
func HasAccessorialCharge() predicate.ShipmentCharges {
	return predicate.ShipmentCharges(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AccessorialChargeTable, AccessorialChargeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAccessorialChargeWith applies the HasEdge predicate on the "accessorial_charge" edge with a given conditions (other predicates).
func HasAccessorialChargeWith(preds ...predicate.AccessorialCharge) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(func(s *sql.Selector) {
		step := newAccessorialChargeStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.ShipmentCharges {
	return predicate.ShipmentCharges(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ShipmentCharges) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ShipmentCharges) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.ShipmentCharges) predicate.ShipmentCharges {
	return predicate.ShipmentCharges(sql.NotPredicates(p))
}
