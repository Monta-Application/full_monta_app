// Code generated by entc, DO NOT EDIT.

package revenuecode

import (
	"fmt"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the revenuecode type in the database.
	Label = "revenue_code"
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
	// FieldCode holds the string denoting the code field in the database.
	FieldCode = "code"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldExpenseAccountID holds the string denoting the expense_account_id field in the database.
	FieldExpenseAccountID = "expense_account_id"
	// FieldRevenueAccountID holds the string denoting the revenue_account_id field in the database.
	FieldRevenueAccountID = "revenue_account_id"
	// EdgeBusinessUnit holds the string denoting the business_unit edge name in mutations.
	EdgeBusinessUnit = "business_unit"
	// EdgeOrganization holds the string denoting the organization edge name in mutations.
	EdgeOrganization = "organization"
	// EdgeExpenseAccount holds the string denoting the expense_account edge name in mutations.
	EdgeExpenseAccount = "expense_account"
	// EdgeRevenueAccount holds the string denoting the revenue_account edge name in mutations.
	EdgeRevenueAccount = "revenue_account"
	// Table holds the table name of the revenuecode in the database.
	Table = "revenue_codes"
	// BusinessUnitTable is the table that holds the business_unit relation/edge.
	BusinessUnitTable = "revenue_codes"
	// BusinessUnitInverseTable is the table name for the BusinessUnit entity.
	// It exists in this package in order to avoid circular dependency with the "businessunit" package.
	BusinessUnitInverseTable = "business_units"
	// BusinessUnitColumn is the table column denoting the business_unit relation/edge.
	BusinessUnitColumn = "business_unit_id"
	// OrganizationTable is the table that holds the organization relation/edge.
	OrganizationTable = "revenue_codes"
	// OrganizationInverseTable is the table name for the Organization entity.
	// It exists in this package in order to avoid circular dependency with the "organization" package.
	OrganizationInverseTable = "organizations"
	// OrganizationColumn is the table column denoting the organization relation/edge.
	OrganizationColumn = "organization_id"
	// ExpenseAccountTable is the table that holds the expense_account relation/edge.
	ExpenseAccountTable = "revenue_codes"
	// ExpenseAccountInverseTable is the table name for the GeneralLedgerAccount entity.
	// It exists in this package in order to avoid circular dependency with the "generalledgeraccount" package.
	ExpenseAccountInverseTable = "general_ledger_accounts"
	// ExpenseAccountColumn is the table column denoting the expense_account relation/edge.
	ExpenseAccountColumn = "expense_account_id"
	// RevenueAccountTable is the table that holds the revenue_account relation/edge.
	RevenueAccountTable = "revenue_codes"
	// RevenueAccountInverseTable is the table name for the GeneralLedgerAccount entity.
	// It exists in this package in order to avoid circular dependency with the "generalledgeraccount" package.
	RevenueAccountInverseTable = "general_ledger_accounts"
	// RevenueAccountColumn is the table column denoting the revenue_account relation/edge.
	RevenueAccountColumn = "revenue_account_id"
)

// Columns holds all SQL columns for revenuecode fields.
var Columns = []string{
	FieldID,
	FieldBusinessUnitID,
	FieldOrganizationID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldVersion,
	FieldStatus,
	FieldCode,
	FieldDescription,
	FieldExpenseAccountID,
	FieldRevenueAccountID,
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

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/emoss08/trenova/ent/runtime"
var (
	Hooks [2]ent.Hook
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultVersion holds the default value on creation for the "version" field.
	DefaultVersion int
	// CodeValidator is a validator for the "code" field. It is called by the builders before save.
	CodeValidator func(string) error
	// DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	DescriptionValidator func(string) error
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
		return fmt.Errorf("revenuecode: invalid enum value for status field: %q", s)
	}
}

// OrderOption defines the ordering options for the RevenueCode queries.
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

// ByCode orders the results by the code field.
func ByCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCode, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByExpenseAccountID orders the results by the expense_account_id field.
func ByExpenseAccountID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExpenseAccountID, opts...).ToFunc()
}

// ByRevenueAccountID orders the results by the revenue_account_id field.
func ByRevenueAccountID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRevenueAccountID, opts...).ToFunc()
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

// ByExpenseAccountField orders the results by expense_account field.
func ByExpenseAccountField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newExpenseAccountStep(), sql.OrderByField(field, opts...))
	}
}

// ByRevenueAccountField orders the results by revenue_account field.
func ByRevenueAccountField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRevenueAccountStep(), sql.OrderByField(field, opts...))
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
func newExpenseAccountStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ExpenseAccountInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, ExpenseAccountTable, ExpenseAccountColumn),
	)
}
func newRevenueAccountStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RevenueAccountInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, RevenueAccountTable, RevenueAccountColumn),
	)
}
