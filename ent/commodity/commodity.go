// Code generated by entc, DO NOT EDIT.

package commodity

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the commodity type in the database.
	Label = "commodity"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldBusinessUnitID holds the string denoting the business_unit_id field in the database.
	FieldBusinessUnitID = "business_unit_id"
	// FieldOrganizationID holds the string denoting the organization_id field in the database.
	FieldOrganizationID = "organization_id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldVersion holds the string denoting the version field in the database.
	FieldVersion = "version"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldIsHazmat holds the string denoting the is_hazmat field in the database.
	FieldIsHazmat = "is_hazmat"
	// FieldUnitOfMeasure holds the string denoting the unit_of_measure field in the database.
	FieldUnitOfMeasure = "unit_of_measure"
	// FieldMinTemp holds the string denoting the min_temp field in the database.
	FieldMinTemp = "min_temp"
	// FieldMaxTemp holds the string denoting the max_temp field in the database.
	FieldMaxTemp = "max_temp"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldHazardousMaterialID holds the string denoting the hazardous_material_id field in the database.
	FieldHazardousMaterialID = "hazardous_material_id"
	// EdgeBusinessUnit holds the string denoting the business_unit edge name in mutations.
	EdgeBusinessUnit = "business_unit"
	// EdgeOrganization holds the string denoting the organization edge name in mutations.
	EdgeOrganization = "organization"
	// EdgeHazardousMaterial holds the string denoting the hazardous_material edge name in mutations.
	EdgeHazardousMaterial = "hazardous_material"
	// Table holds the table name of the commodity in the database.
	Table = "commodities"
	// BusinessUnitTable is the table that holds the business_unit relation/edge.
	BusinessUnitTable = "commodities"
	// BusinessUnitInverseTable is the table name for the BusinessUnit entity.
	// It exists in this package in order to avoid circular dependency with the "businessunit" package.
	BusinessUnitInverseTable = "business_units"
	// BusinessUnitColumn is the table column denoting the business_unit relation/edge.
	BusinessUnitColumn = "business_unit_id"
	// OrganizationTable is the table that holds the organization relation/edge.
	OrganizationTable = "commodities"
	// OrganizationInverseTable is the table name for the Organization entity.
	// It exists in this package in order to avoid circular dependency with the "organization" package.
	OrganizationInverseTable = "organizations"
	// OrganizationColumn is the table column denoting the organization relation/edge.
	OrganizationColumn = "organization_id"
	// HazardousMaterialTable is the table that holds the hazardous_material relation/edge.
	HazardousMaterialTable = "commodities"
	// HazardousMaterialInverseTable is the table name for the HazardousMaterial entity.
	// It exists in this package in order to avoid circular dependency with the "hazardousmaterial" package.
	HazardousMaterialInverseTable = "hazardous_materials"
	// HazardousMaterialColumn is the table column denoting the hazardous_material relation/edge.
	HazardousMaterialColumn = "hazardous_material_id"
)

// Columns holds all SQL columns for commodity fields.
var Columns = []string{
	FieldID,
	FieldBusinessUnitID,
	FieldOrganizationID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldVersion,
	FieldStatus,
	FieldName,
	FieldIsHazmat,
	FieldUnitOfMeasure,
	FieldMinTemp,
	FieldMaxTemp,
	FieldDescription,
	FieldHazardousMaterialID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultVersion holds the default value on creation for the "version" field.
	DefaultVersion int
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultIsHazmat holds the default value on creation for the "is_hazmat" field.
	DefaultIsHazmat bool
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// Status defines the type for the "status" enum field.
type Status string

// StatusA is the default value of the Status enum.
const DefaultStatus = StatusA

// Status values.
const (
	StatusA Status = "A"
	StatusI Status = "I"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusA, StatusI:
		return nil
	default:
		return fmt.Errorf("commodity: invalid enum value for status field: %q", s)
	}
}

// OrderOption defines the ordering options for the Commodity queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByBusinessUnitID orders the results by the business_unit_id field.
func ByBusinessUnitID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBusinessUnitID, opts...).ToFunc()
}

// ByOrganizationID orders the results by the organization_id field.
func ByOrganizationID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOrganizationID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByVersion orders the results by the version field.
func ByVersion(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVersion, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByIsHazmat orders the results by the is_hazmat field.
func ByIsHazmat(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsHazmat, opts...).ToFunc()
}

// ByUnitOfMeasure orders the results by the unit_of_measure field.
func ByUnitOfMeasure(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUnitOfMeasure, opts...).ToFunc()
}

// ByMinTemp orders the results by the min_temp field.
func ByMinTemp(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMinTemp, opts...).ToFunc()
}

// ByMaxTemp orders the results by the max_temp field.
func ByMaxTemp(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMaxTemp, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByHazardousMaterialID orders the results by the hazardous_material_id field.
func ByHazardousMaterialID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHazardousMaterialID, opts...).ToFunc()
}

// ByBusinessUnitField orders the results by business_unit field.
func ByBusinessUnitField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newBusinessUnitStep(), sql.OrderByField(field, opts...))
	}
}

// ByOrganizationField orders the results by organization field.
func ByOrganizationField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOrganizationStep(), sql.OrderByField(field, opts...))
	}
}

// ByHazardousMaterialField orders the results by hazardous_material field.
func ByHazardousMaterialField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newHazardousMaterialStep(), sql.OrderByField(field, opts...))
	}
}
func newBusinessUnitStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(BusinessUnitInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, BusinessUnitTable, BusinessUnitColumn),
	)
}
func newOrganizationStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OrganizationInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, OrganizationTable, OrganizationColumn),
	)
}
func newHazardousMaterialStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(HazardousMaterialInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, HazardousMaterialTable, HazardousMaterialColumn),
	)
}
