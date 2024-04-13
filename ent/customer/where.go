// Code generated by entc, DO NOT EDIT.

package customer

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/emoss08/trenova/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Customer {
	return predicate.Customer(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Customer {
	return predicate.Customer(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Customer {
	return predicate.Customer(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Customer {
	return predicate.Customer(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Customer {
	return predicate.Customer(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Customer {
	return predicate.Customer(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Customer {
	return predicate.Customer(sql.FieldLTE(FieldID, id))
}

// BusinessUnitID applies equality check predicate on the "business_unit_id" field. It's identical to BusinessUnitIDEQ.
func BusinessUnitID(v uuid.UUID) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldBusinessUnitID, v))
}

// OrganizationID applies equality check predicate on the "organization_id" field. It's identical to OrganizationIDEQ.
func OrganizationID(v uuid.UUID) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldOrganizationID, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldUpdatedAt, v))
}

// Version applies equality check predicate on the "version" field. It's identical to VersionEQ.
func Version(v int) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldVersion, v))
}

// Code applies equality check predicate on the "code" field. It's identical to CodeEQ.
func Code(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldCode, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldName, v))
}

// AddressLine1 applies equality check predicate on the "address_line_1" field. It's identical to AddressLine1EQ.
func AddressLine1(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldAddressLine1, v))
}

// AddressLine2 applies equality check predicate on the "address_line_2" field. It's identical to AddressLine2EQ.
func AddressLine2(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldAddressLine2, v))
}

// City applies equality check predicate on the "city" field. It's identical to CityEQ.
func City(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldCity, v))
}

// StateID applies equality check predicate on the "state_id" field. It's identical to StateIDEQ.
func StateID(v uuid.UUID) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldStateID, v))
}

// PostalCode applies equality check predicate on the "postal_code" field. It's identical to PostalCodeEQ.
func PostalCode(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldPostalCode, v))
}

// HasCustomerPortal applies equality check predicate on the "has_customer_portal" field. It's identical to HasCustomerPortalEQ.
func HasCustomerPortal(v bool) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldHasCustomerPortal, v))
}

// AutoMarkReadyToBill applies equality check predicate on the "auto_mark_ready_to_bill" field. It's identical to AutoMarkReadyToBillEQ.
func AutoMarkReadyToBill(v bool) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldAutoMarkReadyToBill, v))
}

// BusinessUnitIDEQ applies the EQ predicate on the "business_unit_id" field.
func BusinessUnitIDEQ(v uuid.UUID) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldBusinessUnitID, v))
}

// BusinessUnitIDNEQ applies the NEQ predicate on the "business_unit_id" field.
func BusinessUnitIDNEQ(v uuid.UUID) predicate.Customer {
	return predicate.Customer(sql.FieldNEQ(FieldBusinessUnitID, v))
}

// BusinessUnitIDIn applies the In predicate on the "business_unit_id" field.
func BusinessUnitIDIn(vs ...uuid.UUID) predicate.Customer {
	return predicate.Customer(sql.FieldIn(FieldBusinessUnitID, vs...))
}

// BusinessUnitIDNotIn applies the NotIn predicate on the "business_unit_id" field.
func BusinessUnitIDNotIn(vs ...uuid.UUID) predicate.Customer {
	return predicate.Customer(sql.FieldNotIn(FieldBusinessUnitID, vs...))
}

// OrganizationIDEQ applies the EQ predicate on the "organization_id" field.
func OrganizationIDEQ(v uuid.UUID) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldOrganizationID, v))
}

// OrganizationIDNEQ applies the NEQ predicate on the "organization_id" field.
func OrganizationIDNEQ(v uuid.UUID) predicate.Customer {
	return predicate.Customer(sql.FieldNEQ(FieldOrganizationID, v))
}

// OrganizationIDIn applies the In predicate on the "organization_id" field.
func OrganizationIDIn(vs ...uuid.UUID) predicate.Customer {
	return predicate.Customer(sql.FieldIn(FieldOrganizationID, vs...))
}

// OrganizationIDNotIn applies the NotIn predicate on the "organization_id" field.
func OrganizationIDNotIn(vs ...uuid.UUID) predicate.Customer {
	return predicate.Customer(sql.FieldNotIn(FieldOrganizationID, vs...))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Customer {
	return predicate.Customer(sql.FieldLTE(FieldUpdatedAt, v))
}

// VersionEQ applies the EQ predicate on the "version" field.
func VersionEQ(v int) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldVersion, v))
}

// VersionNEQ applies the NEQ predicate on the "version" field.
func VersionNEQ(v int) predicate.Customer {
	return predicate.Customer(sql.FieldNEQ(FieldVersion, v))
}

// VersionIn applies the In predicate on the "version" field.
func VersionIn(vs ...int) predicate.Customer {
	return predicate.Customer(sql.FieldIn(FieldVersion, vs...))
}

// VersionNotIn applies the NotIn predicate on the "version" field.
func VersionNotIn(vs ...int) predicate.Customer {
	return predicate.Customer(sql.FieldNotIn(FieldVersion, vs...))
}

// VersionGT applies the GT predicate on the "version" field.
func VersionGT(v int) predicate.Customer {
	return predicate.Customer(sql.FieldGT(FieldVersion, v))
}

// VersionGTE applies the GTE predicate on the "version" field.
func VersionGTE(v int) predicate.Customer {
	return predicate.Customer(sql.FieldGTE(FieldVersion, v))
}

// VersionLT applies the LT predicate on the "version" field.
func VersionLT(v int) predicate.Customer {
	return predicate.Customer(sql.FieldLT(FieldVersion, v))
}

// VersionLTE applies the LTE predicate on the "version" field.
func VersionLTE(v int) predicate.Customer {
	return predicate.Customer(sql.FieldLTE(FieldVersion, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.Customer {
	return predicate.Customer(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.Customer {
	return predicate.Customer(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.Customer {
	return predicate.Customer(sql.FieldNotIn(FieldStatus, vs...))
}

// CodeEQ applies the EQ predicate on the "code" field.
func CodeEQ(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldCode, v))
}

// CodeNEQ applies the NEQ predicate on the "code" field.
func CodeNEQ(v string) predicate.Customer {
	return predicate.Customer(sql.FieldNEQ(FieldCode, v))
}

// CodeIn applies the In predicate on the "code" field.
func CodeIn(vs ...string) predicate.Customer {
	return predicate.Customer(sql.FieldIn(FieldCode, vs...))
}

// CodeNotIn applies the NotIn predicate on the "code" field.
func CodeNotIn(vs ...string) predicate.Customer {
	return predicate.Customer(sql.FieldNotIn(FieldCode, vs...))
}

// CodeGT applies the GT predicate on the "code" field.
func CodeGT(v string) predicate.Customer {
	return predicate.Customer(sql.FieldGT(FieldCode, v))
}

// CodeGTE applies the GTE predicate on the "code" field.
func CodeGTE(v string) predicate.Customer {
	return predicate.Customer(sql.FieldGTE(FieldCode, v))
}

// CodeLT applies the LT predicate on the "code" field.
func CodeLT(v string) predicate.Customer {
	return predicate.Customer(sql.FieldLT(FieldCode, v))
}

// CodeLTE applies the LTE predicate on the "code" field.
func CodeLTE(v string) predicate.Customer {
	return predicate.Customer(sql.FieldLTE(FieldCode, v))
}

// CodeContains applies the Contains predicate on the "code" field.
func CodeContains(v string) predicate.Customer {
	return predicate.Customer(sql.FieldContains(FieldCode, v))
}

// CodeHasPrefix applies the HasPrefix predicate on the "code" field.
func CodeHasPrefix(v string) predicate.Customer {
	return predicate.Customer(sql.FieldHasPrefix(FieldCode, v))
}

// CodeHasSuffix applies the HasSuffix predicate on the "code" field.
func CodeHasSuffix(v string) predicate.Customer {
	return predicate.Customer(sql.FieldHasSuffix(FieldCode, v))
}

// CodeEqualFold applies the EqualFold predicate on the "code" field.
func CodeEqualFold(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEqualFold(FieldCode, v))
}

// CodeContainsFold applies the ContainsFold predicate on the "code" field.
func CodeContainsFold(v string) predicate.Customer {
	return predicate.Customer(sql.FieldContainsFold(FieldCode, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Customer {
	return predicate.Customer(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Customer {
	return predicate.Customer(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Customer {
	return predicate.Customer(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Customer {
	return predicate.Customer(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Customer {
	return predicate.Customer(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Customer {
	return predicate.Customer(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Customer {
	return predicate.Customer(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Customer {
	return predicate.Customer(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Customer {
	return predicate.Customer(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Customer {
	return predicate.Customer(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Customer {
	return predicate.Customer(sql.FieldContainsFold(FieldName, v))
}

// AddressLine1EQ applies the EQ predicate on the "address_line_1" field.
func AddressLine1EQ(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldAddressLine1, v))
}

// AddressLine1NEQ applies the NEQ predicate on the "address_line_1" field.
func AddressLine1NEQ(v string) predicate.Customer {
	return predicate.Customer(sql.FieldNEQ(FieldAddressLine1, v))
}

// AddressLine1In applies the In predicate on the "address_line_1" field.
func AddressLine1In(vs ...string) predicate.Customer {
	return predicate.Customer(sql.FieldIn(FieldAddressLine1, vs...))
}

// AddressLine1NotIn applies the NotIn predicate on the "address_line_1" field.
func AddressLine1NotIn(vs ...string) predicate.Customer {
	return predicate.Customer(sql.FieldNotIn(FieldAddressLine1, vs...))
}

// AddressLine1GT applies the GT predicate on the "address_line_1" field.
func AddressLine1GT(v string) predicate.Customer {
	return predicate.Customer(sql.FieldGT(FieldAddressLine1, v))
}

// AddressLine1GTE applies the GTE predicate on the "address_line_1" field.
func AddressLine1GTE(v string) predicate.Customer {
	return predicate.Customer(sql.FieldGTE(FieldAddressLine1, v))
}

// AddressLine1LT applies the LT predicate on the "address_line_1" field.
func AddressLine1LT(v string) predicate.Customer {
	return predicate.Customer(sql.FieldLT(FieldAddressLine1, v))
}

// AddressLine1LTE applies the LTE predicate on the "address_line_1" field.
func AddressLine1LTE(v string) predicate.Customer {
	return predicate.Customer(sql.FieldLTE(FieldAddressLine1, v))
}

// AddressLine1Contains applies the Contains predicate on the "address_line_1" field.
func AddressLine1Contains(v string) predicate.Customer {
	return predicate.Customer(sql.FieldContains(FieldAddressLine1, v))
}

// AddressLine1HasPrefix applies the HasPrefix predicate on the "address_line_1" field.
func AddressLine1HasPrefix(v string) predicate.Customer {
	return predicate.Customer(sql.FieldHasPrefix(FieldAddressLine1, v))
}

// AddressLine1HasSuffix applies the HasSuffix predicate on the "address_line_1" field.
func AddressLine1HasSuffix(v string) predicate.Customer {
	return predicate.Customer(sql.FieldHasSuffix(FieldAddressLine1, v))
}

// AddressLine1EqualFold applies the EqualFold predicate on the "address_line_1" field.
func AddressLine1EqualFold(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEqualFold(FieldAddressLine1, v))
}

// AddressLine1ContainsFold applies the ContainsFold predicate on the "address_line_1" field.
func AddressLine1ContainsFold(v string) predicate.Customer {
	return predicate.Customer(sql.FieldContainsFold(FieldAddressLine1, v))
}

// AddressLine2EQ applies the EQ predicate on the "address_line_2" field.
func AddressLine2EQ(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldAddressLine2, v))
}

// AddressLine2NEQ applies the NEQ predicate on the "address_line_2" field.
func AddressLine2NEQ(v string) predicate.Customer {
	return predicate.Customer(sql.FieldNEQ(FieldAddressLine2, v))
}

// AddressLine2In applies the In predicate on the "address_line_2" field.
func AddressLine2In(vs ...string) predicate.Customer {
	return predicate.Customer(sql.FieldIn(FieldAddressLine2, vs...))
}

// AddressLine2NotIn applies the NotIn predicate on the "address_line_2" field.
func AddressLine2NotIn(vs ...string) predicate.Customer {
	return predicate.Customer(sql.FieldNotIn(FieldAddressLine2, vs...))
}

// AddressLine2GT applies the GT predicate on the "address_line_2" field.
func AddressLine2GT(v string) predicate.Customer {
	return predicate.Customer(sql.FieldGT(FieldAddressLine2, v))
}

// AddressLine2GTE applies the GTE predicate on the "address_line_2" field.
func AddressLine2GTE(v string) predicate.Customer {
	return predicate.Customer(sql.FieldGTE(FieldAddressLine2, v))
}

// AddressLine2LT applies the LT predicate on the "address_line_2" field.
func AddressLine2LT(v string) predicate.Customer {
	return predicate.Customer(sql.FieldLT(FieldAddressLine2, v))
}

// AddressLine2LTE applies the LTE predicate on the "address_line_2" field.
func AddressLine2LTE(v string) predicate.Customer {
	return predicate.Customer(sql.FieldLTE(FieldAddressLine2, v))
}

// AddressLine2Contains applies the Contains predicate on the "address_line_2" field.
func AddressLine2Contains(v string) predicate.Customer {
	return predicate.Customer(sql.FieldContains(FieldAddressLine2, v))
}

// AddressLine2HasPrefix applies the HasPrefix predicate on the "address_line_2" field.
func AddressLine2HasPrefix(v string) predicate.Customer {
	return predicate.Customer(sql.FieldHasPrefix(FieldAddressLine2, v))
}

// AddressLine2HasSuffix applies the HasSuffix predicate on the "address_line_2" field.
func AddressLine2HasSuffix(v string) predicate.Customer {
	return predicate.Customer(sql.FieldHasSuffix(FieldAddressLine2, v))
}

// AddressLine2IsNil applies the IsNil predicate on the "address_line_2" field.
func AddressLine2IsNil() predicate.Customer {
	return predicate.Customer(sql.FieldIsNull(FieldAddressLine2))
}

// AddressLine2NotNil applies the NotNil predicate on the "address_line_2" field.
func AddressLine2NotNil() predicate.Customer {
	return predicate.Customer(sql.FieldNotNull(FieldAddressLine2))
}

// AddressLine2EqualFold applies the EqualFold predicate on the "address_line_2" field.
func AddressLine2EqualFold(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEqualFold(FieldAddressLine2, v))
}

// AddressLine2ContainsFold applies the ContainsFold predicate on the "address_line_2" field.
func AddressLine2ContainsFold(v string) predicate.Customer {
	return predicate.Customer(sql.FieldContainsFold(FieldAddressLine2, v))
}

// CityEQ applies the EQ predicate on the "city" field.
func CityEQ(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldCity, v))
}

// CityNEQ applies the NEQ predicate on the "city" field.
func CityNEQ(v string) predicate.Customer {
	return predicate.Customer(sql.FieldNEQ(FieldCity, v))
}

// CityIn applies the In predicate on the "city" field.
func CityIn(vs ...string) predicate.Customer {
	return predicate.Customer(sql.FieldIn(FieldCity, vs...))
}

// CityNotIn applies the NotIn predicate on the "city" field.
func CityNotIn(vs ...string) predicate.Customer {
	return predicate.Customer(sql.FieldNotIn(FieldCity, vs...))
}

// CityGT applies the GT predicate on the "city" field.
func CityGT(v string) predicate.Customer {
	return predicate.Customer(sql.FieldGT(FieldCity, v))
}

// CityGTE applies the GTE predicate on the "city" field.
func CityGTE(v string) predicate.Customer {
	return predicate.Customer(sql.FieldGTE(FieldCity, v))
}

// CityLT applies the LT predicate on the "city" field.
func CityLT(v string) predicate.Customer {
	return predicate.Customer(sql.FieldLT(FieldCity, v))
}

// CityLTE applies the LTE predicate on the "city" field.
func CityLTE(v string) predicate.Customer {
	return predicate.Customer(sql.FieldLTE(FieldCity, v))
}

// CityContains applies the Contains predicate on the "city" field.
func CityContains(v string) predicate.Customer {
	return predicate.Customer(sql.FieldContains(FieldCity, v))
}

// CityHasPrefix applies the HasPrefix predicate on the "city" field.
func CityHasPrefix(v string) predicate.Customer {
	return predicate.Customer(sql.FieldHasPrefix(FieldCity, v))
}

// CityHasSuffix applies the HasSuffix predicate on the "city" field.
func CityHasSuffix(v string) predicate.Customer {
	return predicate.Customer(sql.FieldHasSuffix(FieldCity, v))
}

// CityEqualFold applies the EqualFold predicate on the "city" field.
func CityEqualFold(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEqualFold(FieldCity, v))
}

// CityContainsFold applies the ContainsFold predicate on the "city" field.
func CityContainsFold(v string) predicate.Customer {
	return predicate.Customer(sql.FieldContainsFold(FieldCity, v))
}

// StateIDEQ applies the EQ predicate on the "state_id" field.
func StateIDEQ(v uuid.UUID) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldStateID, v))
}

// StateIDNEQ applies the NEQ predicate on the "state_id" field.
func StateIDNEQ(v uuid.UUID) predicate.Customer {
	return predicate.Customer(sql.FieldNEQ(FieldStateID, v))
}

// StateIDIn applies the In predicate on the "state_id" field.
func StateIDIn(vs ...uuid.UUID) predicate.Customer {
	return predicate.Customer(sql.FieldIn(FieldStateID, vs...))
}

// StateIDNotIn applies the NotIn predicate on the "state_id" field.
func StateIDNotIn(vs ...uuid.UUID) predicate.Customer {
	return predicate.Customer(sql.FieldNotIn(FieldStateID, vs...))
}

// PostalCodeEQ applies the EQ predicate on the "postal_code" field.
func PostalCodeEQ(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldPostalCode, v))
}

// PostalCodeNEQ applies the NEQ predicate on the "postal_code" field.
func PostalCodeNEQ(v string) predicate.Customer {
	return predicate.Customer(sql.FieldNEQ(FieldPostalCode, v))
}

// PostalCodeIn applies the In predicate on the "postal_code" field.
func PostalCodeIn(vs ...string) predicate.Customer {
	return predicate.Customer(sql.FieldIn(FieldPostalCode, vs...))
}

// PostalCodeNotIn applies the NotIn predicate on the "postal_code" field.
func PostalCodeNotIn(vs ...string) predicate.Customer {
	return predicate.Customer(sql.FieldNotIn(FieldPostalCode, vs...))
}

// PostalCodeGT applies the GT predicate on the "postal_code" field.
func PostalCodeGT(v string) predicate.Customer {
	return predicate.Customer(sql.FieldGT(FieldPostalCode, v))
}

// PostalCodeGTE applies the GTE predicate on the "postal_code" field.
func PostalCodeGTE(v string) predicate.Customer {
	return predicate.Customer(sql.FieldGTE(FieldPostalCode, v))
}

// PostalCodeLT applies the LT predicate on the "postal_code" field.
func PostalCodeLT(v string) predicate.Customer {
	return predicate.Customer(sql.FieldLT(FieldPostalCode, v))
}

// PostalCodeLTE applies the LTE predicate on the "postal_code" field.
func PostalCodeLTE(v string) predicate.Customer {
	return predicate.Customer(sql.FieldLTE(FieldPostalCode, v))
}

// PostalCodeContains applies the Contains predicate on the "postal_code" field.
func PostalCodeContains(v string) predicate.Customer {
	return predicate.Customer(sql.FieldContains(FieldPostalCode, v))
}

// PostalCodeHasPrefix applies the HasPrefix predicate on the "postal_code" field.
func PostalCodeHasPrefix(v string) predicate.Customer {
	return predicate.Customer(sql.FieldHasPrefix(FieldPostalCode, v))
}

// PostalCodeHasSuffix applies the HasSuffix predicate on the "postal_code" field.
func PostalCodeHasSuffix(v string) predicate.Customer {
	return predicate.Customer(sql.FieldHasSuffix(FieldPostalCode, v))
}

// PostalCodeEqualFold applies the EqualFold predicate on the "postal_code" field.
func PostalCodeEqualFold(v string) predicate.Customer {
	return predicate.Customer(sql.FieldEqualFold(FieldPostalCode, v))
}

// PostalCodeContainsFold applies the ContainsFold predicate on the "postal_code" field.
func PostalCodeContainsFold(v string) predicate.Customer {
	return predicate.Customer(sql.FieldContainsFold(FieldPostalCode, v))
}

// HasCustomerPortalEQ applies the EQ predicate on the "has_customer_portal" field.
func HasCustomerPortalEQ(v bool) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldHasCustomerPortal, v))
}

// HasCustomerPortalNEQ applies the NEQ predicate on the "has_customer_portal" field.
func HasCustomerPortalNEQ(v bool) predicate.Customer {
	return predicate.Customer(sql.FieldNEQ(FieldHasCustomerPortal, v))
}

// AutoMarkReadyToBillEQ applies the EQ predicate on the "auto_mark_ready_to_bill" field.
func AutoMarkReadyToBillEQ(v bool) predicate.Customer {
	return predicate.Customer(sql.FieldEQ(FieldAutoMarkReadyToBill, v))
}

// AutoMarkReadyToBillNEQ applies the NEQ predicate on the "auto_mark_ready_to_bill" field.
func AutoMarkReadyToBillNEQ(v bool) predicate.Customer {
	return predicate.Customer(sql.FieldNEQ(FieldAutoMarkReadyToBill, v))
}

// HasBusinessUnit applies the HasEdge predicate on the "business_unit" edge.
func HasBusinessUnit() predicate.Customer {
	return predicate.Customer(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, BusinessUnitTable, BusinessUnitColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBusinessUnitWith applies the HasEdge predicate on the "business_unit" edge with a given conditions (other predicates).
func HasBusinessUnitWith(preds ...predicate.BusinessUnit) predicate.Customer {
	return predicate.Customer(func(s *sql.Selector) {
		step := newBusinessUnitStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOrganization applies the HasEdge predicate on the "organization" edge.
func HasOrganization() predicate.Customer {
	return predicate.Customer(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, OrganizationTable, OrganizationColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOrganizationWith applies the HasEdge predicate on the "organization" edge with a given conditions (other predicates).
func HasOrganizationWith(preds ...predicate.Organization) predicate.Customer {
	return predicate.Customer(func(s *sql.Selector) {
		step := newOrganizationStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasState applies the HasEdge predicate on the "state" edge.
func HasState() predicate.Customer {
	return predicate.Customer(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, StateTable, StateColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasStateWith applies the HasEdge predicate on the "state" edge with a given conditions (other predicates).
func HasStateWith(preds ...predicate.UsState) predicate.Customer {
	return predicate.Customer(func(s *sql.Selector) {
		step := newStateStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasShipments applies the HasEdge predicate on the "shipments" edge.
func HasShipments() predicate.Customer {
	return predicate.Customer(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ShipmentsTable, ShipmentsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasShipmentsWith applies the HasEdge predicate on the "shipments" edge with a given conditions (other predicates).
func HasShipmentsWith(preds ...predicate.Shipment) predicate.Customer {
	return predicate.Customer(func(s *sql.Selector) {
		step := newShipmentsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Customer) predicate.Customer {
	return predicate.Customer(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Customer) predicate.Customer {
	return predicate.Customer(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Customer) predicate.Customer {
	return predicate.Customer(sql.NotPredicates(p))
}
