// Code generated by entc, DO NOT EDIT.

package locationcomment

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/emoss08/trenova/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.LocationComment {
	return predicate.LocationComment(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.LocationComment {
	return predicate.LocationComment(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.LocationComment {
	return predicate.LocationComment(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.LocationComment {
	return predicate.LocationComment(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.LocationComment {
	return predicate.LocationComment(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.LocationComment {
	return predicate.LocationComment(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.LocationComment {
	return predicate.LocationComment(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.LocationComment {
	return predicate.LocationComment(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.LocationComment {
	return predicate.LocationComment(sql.FieldLTE(FieldID, id))
}

// BusinessUnitID applies equality check predicate on the "business_unit_id" field. It's identical to BusinessUnitIDEQ.
func BusinessUnitID(v uuid.UUID) predicate.LocationComment {
	return predicate.LocationComment(sql.FieldEQ(FieldBusinessUnitID, v))
}

// OrganizationID applies equality check predicate on the "organization_id" field. It's identical to OrganizationIDEQ.
func OrganizationID(v uuid.UUID) predicate.LocationComment {
	return predicate.LocationComment(sql.FieldEQ(FieldOrganizationID, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.LocationComment {
	return predicate.LocationComment(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.LocationComment {
	return predicate.LocationComment(sql.FieldEQ(FieldUpdatedAt, v))
}

// Version applies equality check predicate on the "version" field. It's identical to VersionEQ.
func Version(v int) predicate.LocationComment {
	return predicate.LocationComment(sql.FieldEQ(FieldVersion, v))
}

// LocationID applies equality check predicate on the "location_id" field. It's identical to LocationIDEQ.
func LocationID(v uuid.UUID) predicate.LocationComment {
	return predicate.LocationComment(sql.FieldEQ(FieldLocationID, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v uuid.UUID) predicate.LocationComment {
	return predicate.LocationComment(sql.FieldEQ(FieldUserID, v))
}

// CommentTypeID applies equality check predicate on the "comment_type_id" field. It's identical to CommentTypeIDEQ.
func CommentTypeID(v uuid.UUID) predicate.LocationComment {
	return predicate.LocationComment(sql.FieldEQ(FieldCommentTypeID, v))
}

// Comment applies equality check predicate on the "comment" field. It's identical to CommentEQ.
func Comment(v string) predicate.LocationComment {
	return predicate.LocationComment(sql.FieldEQ(FieldComment, v))
}

// BusinessUnitIDEQ applies the EQ predicate on the "business_unit_id" field.
func BusinessUnitIDEQ(v uuid.UUID) predicate.LocationComment {
	return predicate.LocationComment(sql.FieldEQ(FieldBusinessUnitID, v))
}

// BusinessUnitIDNEQ applies the NEQ predicate on the "business_unit_id" field.
func BusinessUnitIDNEQ(v uuid.UUID) predicate.LocationComment {
	return predicate.LocationComment(sql.FieldNEQ(FieldBusinessUnitID, v))
}

// BusinessUnitIDIn applies the In predicate on the "business_unit_id" field.
func BusinessUnitIDIn(vs ...uuid.UUID) predicate.LocationComment {
	return predicate.LocationComment(sql.FieldIn(FieldBusinessUnitID, vs...))
}

// BusinessUnitIDNotIn applies the NotIn predicate on the "business_unit_id" field.
func BusinessUnitIDNotIn(vs ...uuid.UUID) predicate.LocationComment {
	return predicate.LocationComment(sql.FieldNotIn(FieldBusinessUnitID, vs...))
}

// OrganizationIDEQ applies the EQ predicate on the "organization_id" field.
func OrganizationIDEQ(v uuid.UUID) predicate.LocationComment {
	return predicate.LocationComment(sql.FieldEQ(FieldOrganizationID, v))
}

// OrganizationIDNEQ applies the NEQ predicate on the "organization_id" field.
func OrganizationIDNEQ(v uuid.UUID) predicate.LocationComment {
	return predicate.LocationComment(sql.FieldNEQ(FieldOrganizationID, v))
}

// OrganizationIDIn applies the In predicate on the "organization_id" field.
func OrganizationIDIn(vs ...uuid.UUID) predicate.LocationComment {
	return predicate.LocationComment(sql.FieldIn(FieldOrganizationID, vs...))
}

// OrganizationIDNotIn applies the NotIn predicate on the "organization_id" field.
func OrganizationIDNotIn(vs ...uuid.UUID) predicate.LocationComment {
	return predicate.LocationComment(sql.FieldNotIn(FieldOrganizationID, vs...))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.LocationComment {
	return predicate.LocationComment(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.LocationComment {
	return predicate.LocationComment(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.LocationComment {
	return predicate.LocationComment(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.LocationComment {
	return predicate.LocationComment(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.LocationComment {
	return predicate.LocationComment(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.LocationComment {
	return predicate.LocationComment(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.LocationComment {
	return predicate.LocationComment(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.LocationComment {
	return predicate.LocationComment(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.LocationComment {
	return predicate.LocationComment(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.LocationComment {
	return predicate.LocationComment(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.LocationComment {
	return predicate.LocationComment(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.LocationComment {
	return predicate.LocationComment(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.LocationComment {
	return predicate.LocationComment(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.LocationComment {
	return predicate.LocationComment(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.LocationComment {
	return predicate.LocationComment(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.LocationComment {
	return predicate.LocationComment(sql.FieldLTE(FieldUpdatedAt, v))
}

// VersionEQ applies the EQ predicate on the "version" field.
func VersionEQ(v int) predicate.LocationComment {
	return predicate.LocationComment(sql.FieldEQ(FieldVersion, v))
}

// VersionNEQ applies the NEQ predicate on the "version" field.
func VersionNEQ(v int) predicate.LocationComment {
	return predicate.LocationComment(sql.FieldNEQ(FieldVersion, v))
}

// VersionIn applies the In predicate on the "version" field.
func VersionIn(vs ...int) predicate.LocationComment {
	return predicate.LocationComment(sql.FieldIn(FieldVersion, vs...))
}

// VersionNotIn applies the NotIn predicate on the "version" field.
func VersionNotIn(vs ...int) predicate.LocationComment {
	return predicate.LocationComment(sql.FieldNotIn(FieldVersion, vs...))
}

// VersionGT applies the GT predicate on the "version" field.
func VersionGT(v int) predicate.LocationComment {
	return predicate.LocationComment(sql.FieldGT(FieldVersion, v))
}

// VersionGTE applies the GTE predicate on the "version" field.
func VersionGTE(v int) predicate.LocationComment {
	return predicate.LocationComment(sql.FieldGTE(FieldVersion, v))
}

// VersionLT applies the LT predicate on the "version" field.
func VersionLT(v int) predicate.LocationComment {
	return predicate.LocationComment(sql.FieldLT(FieldVersion, v))
}

// VersionLTE applies the LTE predicate on the "version" field.
func VersionLTE(v int) predicate.LocationComment {
	return predicate.LocationComment(sql.FieldLTE(FieldVersion, v))
}

// LocationIDEQ applies the EQ predicate on the "location_id" field.
func LocationIDEQ(v uuid.UUID) predicate.LocationComment {
	return predicate.LocationComment(sql.FieldEQ(FieldLocationID, v))
}

// LocationIDNEQ applies the NEQ predicate on the "location_id" field.
func LocationIDNEQ(v uuid.UUID) predicate.LocationComment {
	return predicate.LocationComment(sql.FieldNEQ(FieldLocationID, v))
}

// LocationIDIn applies the In predicate on the "location_id" field.
func LocationIDIn(vs ...uuid.UUID) predicate.LocationComment {
	return predicate.LocationComment(sql.FieldIn(FieldLocationID, vs...))
}

// LocationIDNotIn applies the NotIn predicate on the "location_id" field.
func LocationIDNotIn(vs ...uuid.UUID) predicate.LocationComment {
	return predicate.LocationComment(sql.FieldNotIn(FieldLocationID, vs...))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v uuid.UUID) predicate.LocationComment {
	return predicate.LocationComment(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v uuid.UUID) predicate.LocationComment {
	return predicate.LocationComment(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...uuid.UUID) predicate.LocationComment {
	return predicate.LocationComment(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...uuid.UUID) predicate.LocationComment {
	return predicate.LocationComment(sql.FieldNotIn(FieldUserID, vs...))
}

// CommentTypeIDEQ applies the EQ predicate on the "comment_type_id" field.
func CommentTypeIDEQ(v uuid.UUID) predicate.LocationComment {
	return predicate.LocationComment(sql.FieldEQ(FieldCommentTypeID, v))
}

// CommentTypeIDNEQ applies the NEQ predicate on the "comment_type_id" field.
func CommentTypeIDNEQ(v uuid.UUID) predicate.LocationComment {
	return predicate.LocationComment(sql.FieldNEQ(FieldCommentTypeID, v))
}

// CommentTypeIDIn applies the In predicate on the "comment_type_id" field.
func CommentTypeIDIn(vs ...uuid.UUID) predicate.LocationComment {
	return predicate.LocationComment(sql.FieldIn(FieldCommentTypeID, vs...))
}

// CommentTypeIDNotIn applies the NotIn predicate on the "comment_type_id" field.
func CommentTypeIDNotIn(vs ...uuid.UUID) predicate.LocationComment {
	return predicate.LocationComment(sql.FieldNotIn(FieldCommentTypeID, vs...))
}

// CommentEQ applies the EQ predicate on the "comment" field.
func CommentEQ(v string) predicate.LocationComment {
	return predicate.LocationComment(sql.FieldEQ(FieldComment, v))
}

// CommentNEQ applies the NEQ predicate on the "comment" field.
func CommentNEQ(v string) predicate.LocationComment {
	return predicate.LocationComment(sql.FieldNEQ(FieldComment, v))
}

// CommentIn applies the In predicate on the "comment" field.
func CommentIn(vs ...string) predicate.LocationComment {
	return predicate.LocationComment(sql.FieldIn(FieldComment, vs...))
}

// CommentNotIn applies the NotIn predicate on the "comment" field.
func CommentNotIn(vs ...string) predicate.LocationComment {
	return predicate.LocationComment(sql.FieldNotIn(FieldComment, vs...))
}

// CommentGT applies the GT predicate on the "comment" field.
func CommentGT(v string) predicate.LocationComment {
	return predicate.LocationComment(sql.FieldGT(FieldComment, v))
}

// CommentGTE applies the GTE predicate on the "comment" field.
func CommentGTE(v string) predicate.LocationComment {
	return predicate.LocationComment(sql.FieldGTE(FieldComment, v))
}

// CommentLT applies the LT predicate on the "comment" field.
func CommentLT(v string) predicate.LocationComment {
	return predicate.LocationComment(sql.FieldLT(FieldComment, v))
}

// CommentLTE applies the LTE predicate on the "comment" field.
func CommentLTE(v string) predicate.LocationComment {
	return predicate.LocationComment(sql.FieldLTE(FieldComment, v))
}

// CommentContains applies the Contains predicate on the "comment" field.
func CommentContains(v string) predicate.LocationComment {
	return predicate.LocationComment(sql.FieldContains(FieldComment, v))
}

// CommentHasPrefix applies the HasPrefix predicate on the "comment" field.
func CommentHasPrefix(v string) predicate.LocationComment {
	return predicate.LocationComment(sql.FieldHasPrefix(FieldComment, v))
}

// CommentHasSuffix applies the HasSuffix predicate on the "comment" field.
func CommentHasSuffix(v string) predicate.LocationComment {
	return predicate.LocationComment(sql.FieldHasSuffix(FieldComment, v))
}

// CommentEqualFold applies the EqualFold predicate on the "comment" field.
func CommentEqualFold(v string) predicate.LocationComment {
	return predicate.LocationComment(sql.FieldEqualFold(FieldComment, v))
}

// CommentContainsFold applies the ContainsFold predicate on the "comment" field.
func CommentContainsFold(v string) predicate.LocationComment {
	return predicate.LocationComment(sql.FieldContainsFold(FieldComment, v))
}

// HasBusinessUnit applies the HasEdge predicate on the "business_unit" edge.
func HasBusinessUnit() predicate.LocationComment {
	return predicate.LocationComment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, BusinessUnitTable, BusinessUnitColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBusinessUnitWith applies the HasEdge predicate on the "business_unit" edge with a given conditions (other predicates).
func HasBusinessUnitWith(preds ...predicate.BusinessUnit) predicate.LocationComment {
	return predicate.LocationComment(func(s *sql.Selector) {
		step := newBusinessUnitStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOrganization applies the HasEdge predicate on the "organization" edge.
func HasOrganization() predicate.LocationComment {
	return predicate.LocationComment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, OrganizationTable, OrganizationColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOrganizationWith applies the HasEdge predicate on the "organization" edge with a given conditions (other predicates).
func HasOrganizationWith(preds ...predicate.Organization) predicate.LocationComment {
	return predicate.LocationComment(func(s *sql.Selector) {
		step := newOrganizationStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasLocation applies the HasEdge predicate on the "location" edge.
func HasLocation() predicate.LocationComment {
	return predicate.LocationComment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, LocationTable, LocationColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasLocationWith applies the HasEdge predicate on the "location" edge with a given conditions (other predicates).
func HasLocationWith(preds ...predicate.Location) predicate.LocationComment {
	return predicate.LocationComment(func(s *sql.Selector) {
		step := newLocationStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.LocationComment {
	return predicate.LocationComment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.LocationComment {
	return predicate.LocationComment(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCommentType applies the HasEdge predicate on the "comment_type" edge.
func HasCommentType() predicate.LocationComment {
	return predicate.LocationComment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, CommentTypeTable, CommentTypeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCommentTypeWith applies the HasEdge predicate on the "comment_type" edge with a given conditions (other predicates).
func HasCommentTypeWith(preds ...predicate.CommentType) predicate.LocationComment {
	return predicate.LocationComment(func(s *sql.Selector) {
		step := newCommentTypeStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.LocationComment) predicate.LocationComment {
	return predicate.LocationComment(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.LocationComment) predicate.LocationComment {
	return predicate.LocationComment(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.LocationComment) predicate.LocationComment {
	return predicate.LocationComment(sql.NotPredicates(p))
}
