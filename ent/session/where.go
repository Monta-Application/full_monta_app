// Code generated by ent, DO NOT EDIT.

package session

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/emoss08/trenova/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Session {
	return predicate.Session(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Session {
	return predicate.Session(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Session {
	return predicate.Session(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Session {
	return predicate.Session(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Session {
	return predicate.Session(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Session {
	return predicate.Session(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Session {
	return predicate.Session(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Session {
	return predicate.Session(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Session {
	return predicate.Session(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.Session {
	return predicate.Session(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.Session {
	return predicate.Session(sql.FieldContainsFold(FieldID, id))
}

// Data applies equality check predicate on the "data" field. It's identical to DataEQ.
func Data(v string) predicate.Session {
	return predicate.Session(sql.FieldEQ(FieldData, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Session {
	return predicate.Session(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Session {
	return predicate.Session(sql.FieldEQ(FieldUpdatedAt, v))
}

// ExpiresAt applies equality check predicate on the "expires_at" field. It's identical to ExpiresAtEQ.
func ExpiresAt(v time.Time) predicate.Session {
	return predicate.Session(sql.FieldEQ(FieldExpiresAt, v))
}

// DataEQ applies the EQ predicate on the "data" field.
func DataEQ(v string) predicate.Session {
	return predicate.Session(sql.FieldEQ(FieldData, v))
}

// DataNEQ applies the NEQ predicate on the "data" field.
func DataNEQ(v string) predicate.Session {
	return predicate.Session(sql.FieldNEQ(FieldData, v))
}

// DataIn applies the In predicate on the "data" field.
func DataIn(vs ...string) predicate.Session {
	return predicate.Session(sql.FieldIn(FieldData, vs...))
}

// DataNotIn applies the NotIn predicate on the "data" field.
func DataNotIn(vs ...string) predicate.Session {
	return predicate.Session(sql.FieldNotIn(FieldData, vs...))
}

// DataGT applies the GT predicate on the "data" field.
func DataGT(v string) predicate.Session {
	return predicate.Session(sql.FieldGT(FieldData, v))
}

// DataGTE applies the GTE predicate on the "data" field.
func DataGTE(v string) predicate.Session {
	return predicate.Session(sql.FieldGTE(FieldData, v))
}

// DataLT applies the LT predicate on the "data" field.
func DataLT(v string) predicate.Session {
	return predicate.Session(sql.FieldLT(FieldData, v))
}

// DataLTE applies the LTE predicate on the "data" field.
func DataLTE(v string) predicate.Session {
	return predicate.Session(sql.FieldLTE(FieldData, v))
}

// DataContains applies the Contains predicate on the "data" field.
func DataContains(v string) predicate.Session {
	return predicate.Session(sql.FieldContains(FieldData, v))
}

// DataHasPrefix applies the HasPrefix predicate on the "data" field.
func DataHasPrefix(v string) predicate.Session {
	return predicate.Session(sql.FieldHasPrefix(FieldData, v))
}

// DataHasSuffix applies the HasSuffix predicate on the "data" field.
func DataHasSuffix(v string) predicate.Session {
	return predicate.Session(sql.FieldHasSuffix(FieldData, v))
}

// DataEqualFold applies the EqualFold predicate on the "data" field.
func DataEqualFold(v string) predicate.Session {
	return predicate.Session(sql.FieldEqualFold(FieldData, v))
}

// DataContainsFold applies the ContainsFold predicate on the "data" field.
func DataContainsFold(v string) predicate.Session {
	return predicate.Session(sql.FieldContainsFold(FieldData, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Session {
	return predicate.Session(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Session {
	return predicate.Session(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Session {
	return predicate.Session(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Session {
	return predicate.Session(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Session {
	return predicate.Session(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Session {
	return predicate.Session(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Session {
	return predicate.Session(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Session {
	return predicate.Session(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Session {
	return predicate.Session(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Session {
	return predicate.Session(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Session {
	return predicate.Session(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Session {
	return predicate.Session(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Session {
	return predicate.Session(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Session {
	return predicate.Session(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Session {
	return predicate.Session(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Session {
	return predicate.Session(sql.FieldLTE(FieldUpdatedAt, v))
}

// ExpiresAtEQ applies the EQ predicate on the "expires_at" field.
func ExpiresAtEQ(v time.Time) predicate.Session {
	return predicate.Session(sql.FieldEQ(FieldExpiresAt, v))
}

// ExpiresAtNEQ applies the NEQ predicate on the "expires_at" field.
func ExpiresAtNEQ(v time.Time) predicate.Session {
	return predicate.Session(sql.FieldNEQ(FieldExpiresAt, v))
}

// ExpiresAtIn applies the In predicate on the "expires_at" field.
func ExpiresAtIn(vs ...time.Time) predicate.Session {
	return predicate.Session(sql.FieldIn(FieldExpiresAt, vs...))
}

// ExpiresAtNotIn applies the NotIn predicate on the "expires_at" field.
func ExpiresAtNotIn(vs ...time.Time) predicate.Session {
	return predicate.Session(sql.FieldNotIn(FieldExpiresAt, vs...))
}

// ExpiresAtGT applies the GT predicate on the "expires_at" field.
func ExpiresAtGT(v time.Time) predicate.Session {
	return predicate.Session(sql.FieldGT(FieldExpiresAt, v))
}

// ExpiresAtGTE applies the GTE predicate on the "expires_at" field.
func ExpiresAtGTE(v time.Time) predicate.Session {
	return predicate.Session(sql.FieldGTE(FieldExpiresAt, v))
}

// ExpiresAtLT applies the LT predicate on the "expires_at" field.
func ExpiresAtLT(v time.Time) predicate.Session {
	return predicate.Session(sql.FieldLT(FieldExpiresAt, v))
}

// ExpiresAtLTE applies the LTE predicate on the "expires_at" field.
func ExpiresAtLTE(v time.Time) predicate.Session {
	return predicate.Session(sql.FieldLTE(FieldExpiresAt, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Session) predicate.Session {
	return predicate.Session(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Session) predicate.Session {
	return predicate.Session(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Session) predicate.Session {
	return predicate.Session(sql.NotPredicates(p))
}
