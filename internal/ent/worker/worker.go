// Code generated by entc, DO NOT EDIT.

package worker

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the worker type in the database.
	Label = "worker"
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
	// FieldProfilePictureURL holds the string denoting the profile_picture_url field in the database.
	FieldProfilePictureURL = "profile_picture_url"
	// FieldWorkerType holds the string denoting the worker_type field in the database.
	FieldWorkerType = "worker_type"
	// FieldFirstName holds the string denoting the first_name field in the database.
	FieldFirstName = "first_name"
	// FieldLastName holds the string denoting the last_name field in the database.
	FieldLastName = "last_name"
	// FieldAddressLine1 holds the string denoting the address_line_1 field in the database.
	FieldAddressLine1 = "address_line_1"
	// FieldAddressLine2 holds the string denoting the address_line_2 field in the database.
	FieldAddressLine2 = "address_line_2"
	// FieldCity holds the string denoting the city field in the database.
	FieldCity = "city"
	// FieldPostalCode holds the string denoting the postal_code field in the database.
	FieldPostalCode = "postal_code"
	// FieldStateID holds the string denoting the state_id field in the database.
	FieldStateID = "state_id"
	// FieldFleetCodeID holds the string denoting the fleet_code_id field in the database.
	FieldFleetCodeID = "fleet_code_id"
	// FieldManagerID holds the string denoting the manager_id field in the database.
	FieldManagerID = "manager_id"
	// FieldExternalID holds the string denoting the external_id field in the database.
	FieldExternalID = "external_id"
	// EdgeBusinessUnit holds the string denoting the business_unit edge name in mutations.
	EdgeBusinessUnit = "business_unit"
	// EdgeOrganization holds the string denoting the organization edge name in mutations.
	EdgeOrganization = "organization"
	// EdgeState holds the string denoting the state edge name in mutations.
	EdgeState = "state"
	// EdgeFleetCode holds the string denoting the fleet_code edge name in mutations.
	EdgeFleetCode = "fleet_code"
	// EdgeManager holds the string denoting the manager edge name in mutations.
	EdgeManager = "manager"
	// EdgePrimaryTractor holds the string denoting the primary_tractor edge name in mutations.
	EdgePrimaryTractor = "primary_tractor"
	// EdgeSecondaryTractor holds the string denoting the secondary_tractor edge name in mutations.
	EdgeSecondaryTractor = "secondary_tractor"
	// EdgeWorkerProfile holds the string denoting the worker_profile edge name in mutations.
	EdgeWorkerProfile = "worker_profile"
	// EdgeWorkerComments holds the string denoting the worker_comments edge name in mutations.
	EdgeWorkerComments = "worker_comments"
	// EdgeWorkerContacts holds the string denoting the worker_contacts edge name in mutations.
	EdgeWorkerContacts = "worker_contacts"
	// Table holds the table name of the worker in the database.
	Table = "workers"
	// BusinessUnitTable is the table that holds the business_unit relation/edge.
	BusinessUnitTable = "workers"
	// BusinessUnitInverseTable is the table name for the BusinessUnit entity.
	// It exists in this package in order to avoid circular dependency with the "businessunit" package.
	BusinessUnitInverseTable = "business_units"
	// BusinessUnitColumn is the table column denoting the business_unit relation/edge.
	BusinessUnitColumn = "business_unit_id"
	// OrganizationTable is the table that holds the organization relation/edge.
	OrganizationTable = "workers"
	// OrganizationInverseTable is the table name for the Organization entity.
	// It exists in this package in order to avoid circular dependency with the "organization" package.
	OrganizationInverseTable = "organizations"
	// OrganizationColumn is the table column denoting the organization relation/edge.
	OrganizationColumn = "organization_id"
	// StateTable is the table that holds the state relation/edge.
	StateTable = "workers"
	// StateInverseTable is the table name for the UsState entity.
	// It exists in this package in order to avoid circular dependency with the "usstate" package.
	StateInverseTable = "us_states"
	// StateColumn is the table column denoting the state relation/edge.
	StateColumn = "state_id"
	// FleetCodeTable is the table that holds the fleet_code relation/edge.
	FleetCodeTable = "workers"
	// FleetCodeInverseTable is the table name for the FleetCode entity.
	// It exists in this package in order to avoid circular dependency with the "fleetcode" package.
	FleetCodeInverseTable = "fleet_codes"
	// FleetCodeColumn is the table column denoting the fleet_code relation/edge.
	FleetCodeColumn = "fleet_code_id"
	// ManagerTable is the table that holds the manager relation/edge.
	ManagerTable = "workers"
	// ManagerInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	ManagerInverseTable = "users"
	// ManagerColumn is the table column denoting the manager relation/edge.
	ManagerColumn = "manager_id"
	// PrimaryTractorTable is the table that holds the primary_tractor relation/edge.
	PrimaryTractorTable = "tractors"
	// PrimaryTractorInverseTable is the table name for the Tractor entity.
	// It exists in this package in order to avoid circular dependency with the "tractor" package.
	PrimaryTractorInverseTable = "tractors"
	// PrimaryTractorColumn is the table column denoting the primary_tractor relation/edge.
	PrimaryTractorColumn = "primary_worker_id"
	// SecondaryTractorTable is the table that holds the secondary_tractor relation/edge.
	SecondaryTractorTable = "tractors"
	// SecondaryTractorInverseTable is the table name for the Tractor entity.
	// It exists in this package in order to avoid circular dependency with the "tractor" package.
	SecondaryTractorInverseTable = "tractors"
	// SecondaryTractorColumn is the table column denoting the secondary_tractor relation/edge.
	SecondaryTractorColumn = "secondary_worker_id"
	// WorkerProfileTable is the table that holds the worker_profile relation/edge.
	WorkerProfileTable = "worker_profiles"
	// WorkerProfileInverseTable is the table name for the WorkerProfile entity.
	// It exists in this package in order to avoid circular dependency with the "workerprofile" package.
	WorkerProfileInverseTable = "worker_profiles"
	// WorkerProfileColumn is the table column denoting the worker_profile relation/edge.
	WorkerProfileColumn = "worker_id"
	// WorkerCommentsTable is the table that holds the worker_comments relation/edge.
	WorkerCommentsTable = "worker_comments"
	// WorkerCommentsInverseTable is the table name for the WorkerComment entity.
	// It exists in this package in order to avoid circular dependency with the "workercomment" package.
	WorkerCommentsInverseTable = "worker_comments"
	// WorkerCommentsColumn is the table column denoting the worker_comments relation/edge.
	WorkerCommentsColumn = "worker_id"
	// WorkerContactsTable is the table that holds the worker_contacts relation/edge.
	WorkerContactsTable = "worker_contacts"
	// WorkerContactsInverseTable is the table name for the WorkerContact entity.
	// It exists in this package in order to avoid circular dependency with the "workercontact" package.
	WorkerContactsInverseTable = "worker_contacts"
	// WorkerContactsColumn is the table column denoting the worker_contacts relation/edge.
	WorkerContactsColumn = "worker_id"
)

// Columns holds all SQL columns for worker fields.
var Columns = []string{
	FieldID,
	FieldBusinessUnitID,
	FieldOrganizationID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldVersion,
	FieldStatus,
	FieldCode,
	FieldProfilePictureURL,
	FieldWorkerType,
	FieldFirstName,
	FieldLastName,
	FieldAddressLine1,
	FieldAddressLine2,
	FieldCity,
	FieldPostalCode,
	FieldStateID,
	FieldFleetCodeID,
	FieldManagerID,
	FieldExternalID,
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
	// CodeValidator is a validator for the "code" field. It is called by the builders before save.
	CodeValidator func(string) error
	// FirstNameValidator is a validator for the "first_name" field. It is called by the builders before save.
	FirstNameValidator func(string) error
	// LastNameValidator is a validator for the "last_name" field. It is called by the builders before save.
	LastNameValidator func(string) error
	// AddressLine1Validator is a validator for the "address_line_1" field. It is called by the builders before save.
	AddressLine1Validator func(string) error
	// AddressLine2Validator is a validator for the "address_line_2" field. It is called by the builders before save.
	AddressLine2Validator func(string) error
	// CityValidator is a validator for the "city" field. It is called by the builders before save.
	CityValidator func(string) error
	// PostalCodeValidator is a validator for the "postal_code" field. It is called by the builders before save.
	PostalCodeValidator func(string) error
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
		return fmt.Errorf("worker: invalid enum value for status field: %q", s)
	}
}

// WorkerType defines the type for the "worker_type" enum field.
type WorkerType string

// WorkerTypeEmployee is the default value of the WorkerType enum.
const DefaultWorkerType = WorkerTypeEmployee

// WorkerType values.
const (
	WorkerTypeEmployee   WorkerType = "Employee"
	WorkerTypeContractor WorkerType = "Contractor"
)

func (wt WorkerType) String() string {
	return string(wt)
}

// WorkerTypeValidator is a validator for the "worker_type" field enum values. It is called by the builders before save.
func WorkerTypeValidator(wt WorkerType) error {
	switch wt {
	case WorkerTypeEmployee, WorkerTypeContractor:
		return nil
	default:
		return fmt.Errorf("worker: invalid enum value for worker_type field: %q", wt)
	}
}

// OrderOption defines the ordering options for the Worker queries.
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

// ByProfilePictureURL orders the results by the profile_picture_url field.
func ByProfilePictureURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProfilePictureURL, opts...).ToFunc()
}

// ByWorkerType orders the results by the worker_type field.
func ByWorkerType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWorkerType, opts...).ToFunc()
}

// ByFirstName orders the results by the first_name field.
func ByFirstName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFirstName, opts...).ToFunc()
}

// ByLastName orders the results by the last_name field.
func ByLastName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastName, opts...).ToFunc()
}

// ByAddressLine1 orders the results by the address_line_1 field.
func ByAddressLine1(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAddressLine1, opts...).ToFunc()
}

// ByAddressLine2 orders the results by the address_line_2 field.
func ByAddressLine2(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAddressLine2, opts...).ToFunc()
}

// ByCity orders the results by the city field.
func ByCity(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCity, opts...).ToFunc()
}

// ByPostalCode orders the results by the postal_code field.
func ByPostalCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPostalCode, opts...).ToFunc()
}

// ByStateID orders the results by the state_id field.
func ByStateID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStateID, opts...).ToFunc()
}

// ByFleetCodeID orders the results by the fleet_code_id field.
func ByFleetCodeID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFleetCodeID, opts...).ToFunc()
}

// ByManagerID orders the results by the manager_id field.
func ByManagerID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldManagerID, opts...).ToFunc()
}

// ByExternalID orders the results by the external_id field.
func ByExternalID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExternalID, opts...).ToFunc()
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

// ByStateField orders the results by state field.
func ByStateField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newStateStep(), sql.OrderByField(field, opts...))
	}
}

// ByFleetCodeField orders the results by fleet_code field.
func ByFleetCodeField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newFleetCodeStep(), sql.OrderByField(field, opts...))
	}
}

// ByManagerField orders the results by manager field.
func ByManagerField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newManagerStep(), sql.OrderByField(field, opts...))
	}
}

// ByPrimaryTractorField orders the results by primary_tractor field.
func ByPrimaryTractorField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPrimaryTractorStep(), sql.OrderByField(field, opts...))
	}
}

// BySecondaryTractorField orders the results by secondary_tractor field.
func BySecondaryTractorField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSecondaryTractorStep(), sql.OrderByField(field, opts...))
	}
}

// ByWorkerProfileField orders the results by worker_profile field.
func ByWorkerProfileField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newWorkerProfileStep(), sql.OrderByField(field, opts...))
	}
}

// ByWorkerCommentsCount orders the results by worker_comments count.
func ByWorkerCommentsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newWorkerCommentsStep(), opts...)
	}
}

// ByWorkerComments orders the results by worker_comments terms.
func ByWorkerComments(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newWorkerCommentsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByWorkerContactsCount orders the results by worker_contacts count.
func ByWorkerContactsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newWorkerContactsStep(), opts...)
	}
}

// ByWorkerContacts orders the results by worker_contacts terms.
func ByWorkerContacts(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newWorkerContactsStep(), append([]sql.OrderTerm{term}, terms...)...)
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
func newStateStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(StateInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, StateTable, StateColumn),
	)
}
func newFleetCodeStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(FleetCodeInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, FleetCodeTable, FleetCodeColumn),
	)
}
func newManagerStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ManagerInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, ManagerTable, ManagerColumn),
	)
}
func newPrimaryTractorStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PrimaryTractorInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, PrimaryTractorTable, PrimaryTractorColumn),
	)
}
func newSecondaryTractorStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SecondaryTractorInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, SecondaryTractorTable, SecondaryTractorColumn),
	)
}
func newWorkerProfileStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(WorkerProfileInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, WorkerProfileTable, WorkerProfileColumn),
	)
}
func newWorkerCommentsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(WorkerCommentsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, WorkerCommentsTable, WorkerCommentsColumn),
	)
}
func newWorkerContactsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(WorkerContactsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, WorkerContactsTable, WorkerContactsColumn),
	)
}
