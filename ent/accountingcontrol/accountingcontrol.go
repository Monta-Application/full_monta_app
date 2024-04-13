// Code generated by entc, DO NOT EDIT.

package accountingcontrol

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the accountingcontrol type in the database.
	Label = "accounting_control"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldRecThreshold holds the string denoting the rec_threshold field in the database.
	FieldRecThreshold = "rec_threshold"
	// FieldRecThresholdAction holds the string denoting the rec_threshold_action field in the database.
	FieldRecThresholdAction = "rec_threshold_action"
	// FieldAutoCreateJournalEntries holds the string denoting the auto_create_journal_entries field in the database.
	FieldAutoCreateJournalEntries = "auto_create_journal_entries"
	// FieldJournalEntryCriteria holds the string denoting the journal_entry_criteria field in the database.
	FieldJournalEntryCriteria = "journal_entry_criteria"
	// FieldRestrictManualJournalEntries holds the string denoting the restrict_manual_journal_entries field in the database.
	FieldRestrictManualJournalEntries = "restrict_manual_journal_entries"
	// FieldRequireJournalEntryApproval holds the string denoting the require_journal_entry_approval field in the database.
	FieldRequireJournalEntryApproval = "require_journal_entry_approval"
	// FieldEnableRecNotifications holds the string denoting the enable_rec_notifications field in the database.
	FieldEnableRecNotifications = "enable_rec_notifications"
	// FieldHaltOnPendingRec holds the string denoting the halt_on_pending_rec field in the database.
	FieldHaltOnPendingRec = "halt_on_pending_rec"
	// FieldCriticalProcesses holds the string denoting the critical_processes field in the database.
	FieldCriticalProcesses = "critical_processes"
	// FieldDefaultRevAccountID holds the string denoting the default_rev_account_id field in the database.
	FieldDefaultRevAccountID = "default_rev_account_id"
	// FieldDefaultExpAccountID holds the string denoting the default_exp_account_id field in the database.
	FieldDefaultExpAccountID = "default_exp_account_id"
	// EdgeOrganization holds the string denoting the organization edge name in mutations.
	EdgeOrganization = "organization"
	// EdgeBusinessUnit holds the string denoting the business_unit edge name in mutations.
	EdgeBusinessUnit = "business_unit"
	// EdgeDefaultRevAccount holds the string denoting the default_rev_account edge name in mutations.
	EdgeDefaultRevAccount = "default_rev_account"
	// EdgeDefaultExpAccount holds the string denoting the default_exp_account edge name in mutations.
	EdgeDefaultExpAccount = "default_exp_account"
	// Table holds the table name of the accountingcontrol in the database.
	Table = "accounting_controls"
	// OrganizationTable is the table that holds the organization relation/edge.
	OrganizationTable = "accounting_controls"
	// OrganizationInverseTable is the table name for the Organization entity.
	// It exists in this package in order to avoid circular dependency with the "organization" package.
	OrganizationInverseTable = "organizations"
	// OrganizationColumn is the table column denoting the organization relation/edge.
	OrganizationColumn = "organization_id"
	// BusinessUnitTable is the table that holds the business_unit relation/edge.
	BusinessUnitTable = "accounting_controls"
	// BusinessUnitInverseTable is the table name for the BusinessUnit entity.
	// It exists in this package in order to avoid circular dependency with the "businessunit" package.
	BusinessUnitInverseTable = "business_units"
	// BusinessUnitColumn is the table column denoting the business_unit relation/edge.
	BusinessUnitColumn = "business_unit_id"
	// DefaultRevAccountTable is the table that holds the default_rev_account relation/edge.
	DefaultRevAccountTable = "accounting_controls"
	// DefaultRevAccountInverseTable is the table name for the GeneralLedgerAccount entity.
	// It exists in this package in order to avoid circular dependency with the "generalledgeraccount" package.
	DefaultRevAccountInverseTable = "general_ledger_accounts"
	// DefaultRevAccountColumn is the table column denoting the default_rev_account relation/edge.
	DefaultRevAccountColumn = "default_rev_account_id"
	// DefaultExpAccountTable is the table that holds the default_exp_account relation/edge.
	DefaultExpAccountTable = "accounting_controls"
	// DefaultExpAccountInverseTable is the table name for the GeneralLedgerAccount entity.
	// It exists in this package in order to avoid circular dependency with the "generalledgeraccount" package.
	DefaultExpAccountInverseTable = "general_ledger_accounts"
	// DefaultExpAccountColumn is the table column denoting the default_exp_account relation/edge.
	DefaultExpAccountColumn = "default_exp_account_id"
)

// Columns holds all SQL columns for accountingcontrol fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldRecThreshold,
	FieldRecThresholdAction,
	FieldAutoCreateJournalEntries,
	FieldJournalEntryCriteria,
	FieldRestrictManualJournalEntries,
	FieldRequireJournalEntryApproval,
	FieldEnableRecNotifications,
	FieldHaltOnPendingRec,
	FieldCriticalProcesses,
	FieldDefaultRevAccountID,
	FieldDefaultExpAccountID,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "accounting_controls"
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
	// DefaultRecThreshold holds the default value on creation for the "rec_threshold" field.
	DefaultRecThreshold int8
	// RecThresholdValidator is a validator for the "rec_threshold" field. It is called by the builders before save.
	RecThresholdValidator func(int8) error
	// DefaultAutoCreateJournalEntries holds the default value on creation for the "auto_create_journal_entries" field.
	DefaultAutoCreateJournalEntries bool
	// DefaultRestrictManualJournalEntries holds the default value on creation for the "restrict_manual_journal_entries" field.
	DefaultRestrictManualJournalEntries bool
	// DefaultRequireJournalEntryApproval holds the default value on creation for the "require_journal_entry_approval" field.
	DefaultRequireJournalEntryApproval bool
	// DefaultEnableRecNotifications holds the default value on creation for the "enable_rec_notifications" field.
	DefaultEnableRecNotifications bool
	// DefaultHaltOnPendingRec holds the default value on creation for the "halt_on_pending_rec" field.
	DefaultHaltOnPendingRec bool
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// RecThresholdAction defines the type for the "rec_threshold_action" enum field.
type RecThresholdAction string

// RecThresholdActionHalt is the default value of the RecThresholdAction enum.
const DefaultRecThresholdAction = RecThresholdActionHalt

// RecThresholdAction values.
const (
	RecThresholdActionHalt RecThresholdAction = "Halt"
	RecThresholdActionWarn RecThresholdAction = "Warn"
)

func (rta RecThresholdAction) String() string {
	return string(rta)
}

// RecThresholdActionValidator is a validator for the "rec_threshold_action" field enum values. It is called by the builders before save.
func RecThresholdActionValidator(rta RecThresholdAction) error {
	switch rta {
	case RecThresholdActionHalt, RecThresholdActionWarn:
		return nil
	default:
		return fmt.Errorf("accountingcontrol: invalid enum value for rec_threshold_action field: %q", rta)
	}
}

// JournalEntryCriteria defines the type for the "journal_entry_criteria" enum field.
type JournalEntryCriteria string

// JournalEntryCriteriaOnShipmentBill is the default value of the JournalEntryCriteria enum.
const DefaultJournalEntryCriteria = JournalEntryCriteriaOnShipmentBill

// JournalEntryCriteria values.
const (
	JournalEntryCriteriaOnShipmentBill       JournalEntryCriteria = "OnShipmentBill"
	JournalEntryCriteriaOnReceiptOfPayment   JournalEntryCriteria = "OnReceiptOfPayment"
	JournalEntryCriteriaOnExpenseRecognition JournalEntryCriteria = "OnExpenseRecognition"
)

func (jec JournalEntryCriteria) String() string {
	return string(jec)
}

// JournalEntryCriteriaValidator is a validator for the "journal_entry_criteria" field enum values. It is called by the builders before save.
func JournalEntryCriteriaValidator(jec JournalEntryCriteria) error {
	switch jec {
	case JournalEntryCriteriaOnShipmentBill, JournalEntryCriteriaOnReceiptOfPayment, JournalEntryCriteriaOnExpenseRecognition:
		return nil
	default:
		return fmt.Errorf("accountingcontrol: invalid enum value for journal_entry_criteria field: %q", jec)
	}
}

// OrderOption defines the ordering options for the AccountingControl queries.
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

// ByRecThreshold orders the results by the rec_threshold field.
func ByRecThreshold(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRecThreshold, opts...).ToFunc()
}

// ByRecThresholdAction orders the results by the rec_threshold_action field.
func ByRecThresholdAction(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRecThresholdAction, opts...).ToFunc()
}

// ByAutoCreateJournalEntries orders the results by the auto_create_journal_entries field.
func ByAutoCreateJournalEntries(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAutoCreateJournalEntries, opts...).ToFunc()
}

// ByJournalEntryCriteria orders the results by the journal_entry_criteria field.
func ByJournalEntryCriteria(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldJournalEntryCriteria, opts...).ToFunc()
}

// ByRestrictManualJournalEntries orders the results by the restrict_manual_journal_entries field.
func ByRestrictManualJournalEntries(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRestrictManualJournalEntries, opts...).ToFunc()
}

// ByRequireJournalEntryApproval orders the results by the require_journal_entry_approval field.
func ByRequireJournalEntryApproval(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRequireJournalEntryApproval, opts...).ToFunc()
}

// ByEnableRecNotifications orders the results by the enable_rec_notifications field.
func ByEnableRecNotifications(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEnableRecNotifications, opts...).ToFunc()
}

// ByHaltOnPendingRec orders the results by the halt_on_pending_rec field.
func ByHaltOnPendingRec(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHaltOnPendingRec, opts...).ToFunc()
}

// ByCriticalProcesses orders the results by the critical_processes field.
func ByCriticalProcesses(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCriticalProcesses, opts...).ToFunc()
}

// ByDefaultRevAccountID orders the results by the default_rev_account_id field.
func ByDefaultRevAccountID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDefaultRevAccountID, opts...).ToFunc()
}

// ByDefaultExpAccountID orders the results by the default_exp_account_id field.
func ByDefaultExpAccountID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDefaultExpAccountID, opts...).ToFunc()
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

// ByDefaultRevAccountField orders the results by default_rev_account field.
func ByDefaultRevAccountField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDefaultRevAccountStep(), sql.OrderByField(field, opts...))
	}
}

// ByDefaultExpAccountField orders the results by default_exp_account field.
func ByDefaultExpAccountField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDefaultExpAccountStep(), sql.OrderByField(field, opts...))
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
func newDefaultRevAccountStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DefaultRevAccountInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, DefaultRevAccountTable, DefaultRevAccountColumn),
	)
}
func newDefaultExpAccountStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DefaultExpAccountInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, DefaultExpAccountTable, DefaultExpAccountColumn),
	)
}
