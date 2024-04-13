// Code generated by entc, DO NOT EDIT.

package documentclassification

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/emoss08/trenova/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldLTE(FieldID, id))
}

// BusinessUnitID applies equality check predicate on the "business_unit_id" field. It's identical to BusinessUnitIDEQ.
func BusinessUnitID(v uuid.UUID) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldEQ(FieldBusinessUnitID, v))
}

// OrganizationID applies equality check predicate on the "organization_id" field. It's identical to OrganizationIDEQ.
func OrganizationID(v uuid.UUID) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldEQ(FieldOrganizationID, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldEQ(FieldUpdatedAt, v))
}

// Version applies equality check predicate on the "version" field. It's identical to VersionEQ.
func Version(v int) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldEQ(FieldVersion, v))
}

// Code applies equality check predicate on the "code" field. It's identical to CodeEQ.
func Code(v string) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldEQ(FieldCode, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldEQ(FieldDescription, v))
}

// Color applies equality check predicate on the "color" field. It's identical to ColorEQ.
func Color(v string) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldEQ(FieldColor, v))
}

// BusinessUnitIDEQ applies the EQ predicate on the "business_unit_id" field.
func BusinessUnitIDEQ(v uuid.UUID) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldEQ(FieldBusinessUnitID, v))
}

// BusinessUnitIDNEQ applies the NEQ predicate on the "business_unit_id" field.
func BusinessUnitIDNEQ(v uuid.UUID) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldNEQ(FieldBusinessUnitID, v))
}

// BusinessUnitIDIn applies the In predicate on the "business_unit_id" field.
func BusinessUnitIDIn(vs ...uuid.UUID) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldIn(FieldBusinessUnitID, vs...))
}

// BusinessUnitIDNotIn applies the NotIn predicate on the "business_unit_id" field.
func BusinessUnitIDNotIn(vs ...uuid.UUID) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldNotIn(FieldBusinessUnitID, vs...))
}

// OrganizationIDEQ applies the EQ predicate on the "organization_id" field.
func OrganizationIDEQ(v uuid.UUID) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldEQ(FieldOrganizationID, v))
}

// OrganizationIDNEQ applies the NEQ predicate on the "organization_id" field.
func OrganizationIDNEQ(v uuid.UUID) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldNEQ(FieldOrganizationID, v))
}

// OrganizationIDIn applies the In predicate on the "organization_id" field.
func OrganizationIDIn(vs ...uuid.UUID) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldIn(FieldOrganizationID, vs...))
}

// OrganizationIDNotIn applies the NotIn predicate on the "organization_id" field.
func OrganizationIDNotIn(vs ...uuid.UUID) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldNotIn(FieldOrganizationID, vs...))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldLTE(FieldUpdatedAt, v))
}

// VersionEQ applies the EQ predicate on the "version" field.
func VersionEQ(v int) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldEQ(FieldVersion, v))
}

// VersionNEQ applies the NEQ predicate on the "version" field.
func VersionNEQ(v int) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldNEQ(FieldVersion, v))
}

// VersionIn applies the In predicate on the "version" field.
func VersionIn(vs ...int) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldIn(FieldVersion, vs...))
}

// VersionNotIn applies the NotIn predicate on the "version" field.
func VersionNotIn(vs ...int) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldNotIn(FieldVersion, vs...))
}

// VersionGT applies the GT predicate on the "version" field.
func VersionGT(v int) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldGT(FieldVersion, v))
}

// VersionGTE applies the GTE predicate on the "version" field.
func VersionGTE(v int) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldGTE(FieldVersion, v))
}

// VersionLT applies the LT predicate on the "version" field.
func VersionLT(v int) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldLT(FieldVersion, v))
}

// VersionLTE applies the LTE predicate on the "version" field.
func VersionLTE(v int) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldLTE(FieldVersion, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldNotIn(FieldStatus, vs...))
}

// CodeEQ applies the EQ predicate on the "code" field.
func CodeEQ(v string) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldEQ(FieldCode, v))
}

// CodeNEQ applies the NEQ predicate on the "code" field.
func CodeNEQ(v string) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldNEQ(FieldCode, v))
}

// CodeIn applies the In predicate on the "code" field.
func CodeIn(vs ...string) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldIn(FieldCode, vs...))
}

// CodeNotIn applies the NotIn predicate on the "code" field.
func CodeNotIn(vs ...string) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldNotIn(FieldCode, vs...))
}

// CodeGT applies the GT predicate on the "code" field.
func CodeGT(v string) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldGT(FieldCode, v))
}

// CodeGTE applies the GTE predicate on the "code" field.
func CodeGTE(v string) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldGTE(FieldCode, v))
}

// CodeLT applies the LT predicate on the "code" field.
func CodeLT(v string) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldLT(FieldCode, v))
}

// CodeLTE applies the LTE predicate on the "code" field.
func CodeLTE(v string) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldLTE(FieldCode, v))
}

// CodeContains applies the Contains predicate on the "code" field.
func CodeContains(v string) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldContains(FieldCode, v))
}

// CodeHasPrefix applies the HasPrefix predicate on the "code" field.
func CodeHasPrefix(v string) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldHasPrefix(FieldCode, v))
}

// CodeHasSuffix applies the HasSuffix predicate on the "code" field.
func CodeHasSuffix(v string) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldHasSuffix(FieldCode, v))
}

// CodeEqualFold applies the EqualFold predicate on the "code" field.
func CodeEqualFold(v string) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldEqualFold(FieldCode, v))
}

// CodeContainsFold applies the ContainsFold predicate on the "code" field.
func CodeContainsFold(v string) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldContainsFold(FieldCode, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldIsNull(FieldDescription))
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldNotNull(FieldDescription))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldContainsFold(FieldDescription, v))
}

// ColorEQ applies the EQ predicate on the "color" field.
func ColorEQ(v string) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldEQ(FieldColor, v))
}

// ColorNEQ applies the NEQ predicate on the "color" field.
func ColorNEQ(v string) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldNEQ(FieldColor, v))
}

// ColorIn applies the In predicate on the "color" field.
func ColorIn(vs ...string) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldIn(FieldColor, vs...))
}

// ColorNotIn applies the NotIn predicate on the "color" field.
func ColorNotIn(vs ...string) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldNotIn(FieldColor, vs...))
}

// ColorGT applies the GT predicate on the "color" field.
func ColorGT(v string) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldGT(FieldColor, v))
}

// ColorGTE applies the GTE predicate on the "color" field.
func ColorGTE(v string) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldGTE(FieldColor, v))
}

// ColorLT applies the LT predicate on the "color" field.
func ColorLT(v string) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldLT(FieldColor, v))
}

// ColorLTE applies the LTE predicate on the "color" field.
func ColorLTE(v string) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldLTE(FieldColor, v))
}

// ColorContains applies the Contains predicate on the "color" field.
func ColorContains(v string) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldContains(FieldColor, v))
}

// ColorHasPrefix applies the HasPrefix predicate on the "color" field.
func ColorHasPrefix(v string) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldHasPrefix(FieldColor, v))
}

// ColorHasSuffix applies the HasSuffix predicate on the "color" field.
func ColorHasSuffix(v string) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldHasSuffix(FieldColor, v))
}

// ColorIsNil applies the IsNil predicate on the "color" field.
func ColorIsNil() predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldIsNull(FieldColor))
}

// ColorNotNil applies the NotNil predicate on the "color" field.
func ColorNotNil() predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldNotNull(FieldColor))
}

// ColorEqualFold applies the EqualFold predicate on the "color" field.
func ColorEqualFold(v string) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldEqualFold(FieldColor, v))
}

// ColorContainsFold applies the ContainsFold predicate on the "color" field.
func ColorContainsFold(v string) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.FieldContainsFold(FieldColor, v))
}

// HasBusinessUnit applies the HasEdge predicate on the "business_unit" edge.
func HasBusinessUnit() predicate.DocumentClassification {
	return predicate.DocumentClassification(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, BusinessUnitTable, BusinessUnitColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBusinessUnitWith applies the HasEdge predicate on the "business_unit" edge with a given conditions (other predicates).
func HasBusinessUnitWith(preds ...predicate.BusinessUnit) predicate.DocumentClassification {
	return predicate.DocumentClassification(func(s *sql.Selector) {
		step := newBusinessUnitStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOrganization applies the HasEdge predicate on the "organization" edge.
func HasOrganization() predicate.DocumentClassification {
	return predicate.DocumentClassification(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, OrganizationTable, OrganizationColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOrganizationWith applies the HasEdge predicate on the "organization" edge with a given conditions (other predicates).
func HasOrganizationWith(preds ...predicate.Organization) predicate.DocumentClassification {
	return predicate.DocumentClassification(func(s *sql.Selector) {
		step := newOrganizationStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasShipmentDocumentation applies the HasEdge predicate on the "shipment_documentation" edge.
func HasShipmentDocumentation() predicate.DocumentClassification {
	return predicate.DocumentClassification(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ShipmentDocumentationTable, ShipmentDocumentationColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasShipmentDocumentationWith applies the HasEdge predicate on the "shipment_documentation" edge with a given conditions (other predicates).
func HasShipmentDocumentationWith(preds ...predicate.ShipmentDocumentation) predicate.DocumentClassification {
	return predicate.DocumentClassification(func(s *sql.Selector) {
		step := newShipmentDocumentationStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.DocumentClassification) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.DocumentClassification) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.DocumentClassification) predicate.DocumentClassification {
	return predicate.DocumentClassification(sql.NotPredicates(p))
}
