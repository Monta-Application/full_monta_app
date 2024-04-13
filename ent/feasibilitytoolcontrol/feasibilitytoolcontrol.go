// Code generated by entc, DO NOT EDIT.

package feasibilitytoolcontrol

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the feasibilitytoolcontrol type in the database.
	Label = "feasibility_tool_control"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldOtpOperator holds the string denoting the otp_operator field in the database.
	FieldOtpOperator = "otp_operator"
	// FieldOtpValue holds the string denoting the otp_value field in the database.
	FieldOtpValue = "otp_value"
	// FieldMpwOperator holds the string denoting the mpw_operator field in the database.
	FieldMpwOperator = "mpw_operator"
	// FieldMpwValue holds the string denoting the mpw_value field in the database.
	FieldMpwValue = "mpw_value"
	// FieldMpdOperator holds the string denoting the mpd_operator field in the database.
	FieldMpdOperator = "mpd_operator"
	// FieldMpdValue holds the string denoting the mpd_value field in the database.
	FieldMpdValue = "mpd_value"
	// FieldMpgOperator holds the string denoting the mpg_operator field in the database.
	FieldMpgOperator = "mpg_operator"
	// FieldMpgValue holds the string denoting the mpg_value field in the database.
	FieldMpgValue = "mpg_value"
	// EdgeOrganization holds the string denoting the organization edge name in mutations.
	EdgeOrganization = "organization"
	// EdgeBusinessUnit holds the string denoting the business_unit edge name in mutations.
	EdgeBusinessUnit = "business_unit"
	// Table holds the table name of the feasibilitytoolcontrol in the database.
	Table = "feasibility_tool_controls"
	// OrganizationTable is the table that holds the organization relation/edge.
	OrganizationTable = "feasibility_tool_controls"
	// OrganizationInverseTable is the table name for the Organization entity.
	// It exists in this package in order to avoid circular dependency with the "organization" package.
	OrganizationInverseTable = "organizations"
	// OrganizationColumn is the table column denoting the organization relation/edge.
	OrganizationColumn = "organization_id"
	// BusinessUnitTable is the table that holds the business_unit relation/edge.
	BusinessUnitTable = "feasibility_tool_controls"
	// BusinessUnitInverseTable is the table name for the BusinessUnit entity.
	// It exists in this package in order to avoid circular dependency with the "businessunit" package.
	BusinessUnitInverseTable = "business_units"
	// BusinessUnitColumn is the table column denoting the business_unit relation/edge.
	BusinessUnitColumn = "business_unit_id"
)

// Columns holds all SQL columns for feasibilitytoolcontrol fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldOtpOperator,
	FieldOtpValue,
	FieldMpwOperator,
	FieldMpwValue,
	FieldMpdOperator,
	FieldMpdValue,
	FieldMpgOperator,
	FieldMpgValue,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "feasibility_tool_controls"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"business_unit_id",
	"organization_id",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
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
	// DefaultOtpValue holds the default value on creation for the "otp_value" field.
	DefaultOtpValue float64
	// DefaultMpwValue holds the default value on creation for the "mpw_value" field.
	DefaultMpwValue float64
	// DefaultMpdValue holds the default value on creation for the "mpd_value" field.
	DefaultMpdValue float64
	// DefaultMpgValue holds the default value on creation for the "mpg_value" field.
	DefaultMpgValue float64
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OtpOperator defines the type for the "otp_operator" enum field.
type OtpOperator string

// OtpOperatorEq is the default value of the OtpOperator enum.
const DefaultOtpOperator = OtpOperatorEq

// OtpOperator values.
const (
	OtpOperatorEq  OtpOperator = "Eq"
	OtpOperatorNe  OtpOperator = "Ne"
	OtpOperatorGt  OtpOperator = "Gt"
	OtpOperatorGte OtpOperator = "Gte"
	OtpOperatorLt  OtpOperator = "Lt"
	OtpOperatorLte OtpOperator = "Lte"
)

func (oo OtpOperator) String() string {
	return string(oo)
}

// OtpOperatorValidator is a validator for the "otp_operator" field enum values. It is called by the builders before save.
func OtpOperatorValidator(oo OtpOperator) error {
	switch oo {
	case OtpOperatorEq, OtpOperatorNe, OtpOperatorGt, OtpOperatorGte, OtpOperatorLt, OtpOperatorLte:
		return nil
	default:
		return fmt.Errorf("feasibilitytoolcontrol: invalid enum value for otp_operator field: %q", oo)
	}
}

// MpwOperator defines the type for the "mpw_operator" enum field.
type MpwOperator string

// MpwOperatorEq is the default value of the MpwOperator enum.
const DefaultMpwOperator = MpwOperatorEq

// MpwOperator values.
const (
	MpwOperatorEq  MpwOperator = "Eq"
	MpwOperatorNe  MpwOperator = "Ne"
	MpwOperatorGt  MpwOperator = "Gt"
	MpwOperatorGte MpwOperator = "Gte"
	MpwOperatorLt  MpwOperator = "Lt"
	MpwOperatorLte MpwOperator = "Lte"
)

func (mo MpwOperator) String() string {
	return string(mo)
}

// MpwOperatorValidator is a validator for the "mpw_operator" field enum values. It is called by the builders before save.
func MpwOperatorValidator(mo MpwOperator) error {
	switch mo {
	case MpwOperatorEq, MpwOperatorNe, MpwOperatorGt, MpwOperatorGte, MpwOperatorLt, MpwOperatorLte:
		return nil
	default:
		return fmt.Errorf("feasibilitytoolcontrol: invalid enum value for mpw_operator field: %q", mo)
	}
}

// MpdOperator defines the type for the "mpd_operator" enum field.
type MpdOperator string

// MpdOperatorEq is the default value of the MpdOperator enum.
const DefaultMpdOperator = MpdOperatorEq

// MpdOperator values.
const (
	MpdOperatorEq  MpdOperator = "Eq"
	MpdOperatorNe  MpdOperator = "Ne"
	MpdOperatorGt  MpdOperator = "Gt"
	MpdOperatorGte MpdOperator = "Gte"
	MpdOperatorLt  MpdOperator = "Lt"
	MpdOperatorLte MpdOperator = "Lte"
)

func (mo MpdOperator) String() string {
	return string(mo)
}

// MpdOperatorValidator is a validator for the "mpd_operator" field enum values. It is called by the builders before save.
func MpdOperatorValidator(mo MpdOperator) error {
	switch mo {
	case MpdOperatorEq, MpdOperatorNe, MpdOperatorGt, MpdOperatorGte, MpdOperatorLt, MpdOperatorLte:
		return nil
	default:
		return fmt.Errorf("feasibilitytoolcontrol: invalid enum value for mpd_operator field: %q", mo)
	}
}

// MpgOperator defines the type for the "mpg_operator" enum field.
type MpgOperator string

// MpgOperatorEq is the default value of the MpgOperator enum.
const DefaultMpgOperator = MpgOperatorEq

// MpgOperator values.
const (
	MpgOperatorEq  MpgOperator = "Eq"
	MpgOperatorNe  MpgOperator = "Ne"
	MpgOperatorGt  MpgOperator = "Gt"
	MpgOperatorGte MpgOperator = "Gte"
	MpgOperatorLt  MpgOperator = "Lt"
	MpgOperatorLte MpgOperator = "Lte"
)

func (mo MpgOperator) String() string {
	return string(mo)
}

// MpgOperatorValidator is a validator for the "mpg_operator" field enum values. It is called by the builders before save.
func MpgOperatorValidator(mo MpgOperator) error {
	switch mo {
	case MpgOperatorEq, MpgOperatorNe, MpgOperatorGt, MpgOperatorGte, MpgOperatorLt, MpgOperatorLte:
		return nil
	default:
		return fmt.Errorf("feasibilitytoolcontrol: invalid enum value for mpg_operator field: %q", mo)
	}
}

// OrderOption defines the ordering options for the FeasibilityToolControl queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByOtpOperator orders the results by the otp_operator field.
func ByOtpOperator(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOtpOperator, opts...).ToFunc()
}

// ByOtpValue orders the results by the otp_value field.
func ByOtpValue(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOtpValue, opts...).ToFunc()
}

// ByMpwOperator orders the results by the mpw_operator field.
func ByMpwOperator(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMpwOperator, opts...).ToFunc()
}

// ByMpwValue orders the results by the mpw_value field.
func ByMpwValue(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMpwValue, opts...).ToFunc()
}

// ByMpdOperator orders the results by the mpd_operator field.
func ByMpdOperator(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMpdOperator, opts...).ToFunc()
}

// ByMpdValue orders the results by the mpd_value field.
func ByMpdValue(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMpdValue, opts...).ToFunc()
}

// ByMpgOperator orders the results by the mpg_operator field.
func ByMpgOperator(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMpgOperator, opts...).ToFunc()
}

// ByMpgValue orders the results by the mpg_value field.
func ByMpgValue(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMpgValue, opts...).ToFunc()
}

// ByOrganizationField orders the results by organization field.
func ByOrganizationField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOrganizationStep(), sql.OrderByField(field, opts...))
	}
}

// ByBusinessUnitField orders the results by business_unit field.
func ByBusinessUnitField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newBusinessUnitStep(), sql.OrderByField(field, opts...))
	}
}
func newOrganizationStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OrganizationInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, OrganizationTable, OrganizationColumn),
	)
}
func newBusinessUnitStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(BusinessUnitInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, BusinessUnitTable, BusinessUnitColumn),
	)
}
