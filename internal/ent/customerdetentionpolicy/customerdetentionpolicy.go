// Code generated by entc, DO NOT EDIT.

package customerdetentionpolicy

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the customerdetentionpolicy type in the database.
	Label = "customer_detention_policy"
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
	// FieldCustomerID holds the string denoting the customer_id field in the database.
	FieldCustomerID = "customer_id"
	// FieldCommodityID holds the string denoting the commodity_id field in the database.
	FieldCommodityID = "commodity_id"
	// FieldRevenueCodeID holds the string denoting the revenue_code_id field in the database.
	FieldRevenueCodeID = "revenue_code_id"
	// FieldShipmentTypeID holds the string denoting the shipment_type_id field in the database.
	FieldShipmentTypeID = "shipment_type_id"
	// FieldApplicationScope holds the string denoting the application_scope field in the database.
	FieldApplicationScope = "application_scope"
	// FieldChargeFreeTime holds the string denoting the charge_free_time field in the database.
	FieldChargeFreeTime = "charge_free_time"
	// FieldPaymentFreeTime holds the string denoting the payment_free_time field in the database.
	FieldPaymentFreeTime = "payment_free_time"
	// FieldLateArrivalPolicy holds the string denoting the late_arrival_policy field in the database.
	FieldLateArrivalPolicy = "late_arrival_policy"
	// FieldGracePeriod holds the string denoting the grace_period field in the database.
	FieldGracePeriod = "grace_period"
	// FieldAccessorialChargeID holds the string denoting the accessorial_charge_id field in the database.
	FieldAccessorialChargeID = "accessorial_charge_id"
	// FieldUnits holds the string denoting the units field in the database.
	FieldUnits = "units"
	// FieldAmount holds the string denoting the amount field in the database.
	FieldAmount = "amount"
	// FieldNotes holds the string denoting the notes field in the database.
	FieldNotes = "notes"
	// FieldEffectiveDate holds the string denoting the effective_date field in the database.
	FieldEffectiveDate = "effective_date"
	// FieldExpirationDate holds the string denoting the expiration_date field in the database.
	FieldExpirationDate = "expiration_date"
	// EdgeBusinessUnit holds the string denoting the business_unit edge name in mutations.
	EdgeBusinessUnit = "business_unit"
	// EdgeOrganization holds the string denoting the organization edge name in mutations.
	EdgeOrganization = "organization"
	// EdgeCustomer holds the string denoting the customer edge name in mutations.
	EdgeCustomer = "customer"
	// EdgeCommodity holds the string denoting the commodity edge name in mutations.
	EdgeCommodity = "commodity"
	// EdgeRevenueCode holds the string denoting the revenue_code edge name in mutations.
	EdgeRevenueCode = "revenue_code"
	// EdgeShipmentType holds the string denoting the shipment_type edge name in mutations.
	EdgeShipmentType = "shipment_type"
	// EdgeAccessorialCharge holds the string denoting the accessorial_charge edge name in mutations.
	EdgeAccessorialCharge = "accessorial_charge"
	// Table holds the table name of the customerdetentionpolicy in the database.
	Table = "customer_detention_policies"
	// BusinessUnitTable is the table that holds the business_unit relation/edge.
	BusinessUnitTable = "customer_detention_policies"
	// BusinessUnitInverseTable is the table name for the BusinessUnit entity.
	// It exists in this package in order to avoid circular dependency with the "businessunit" package.
	BusinessUnitInverseTable = "business_units"
	// BusinessUnitColumn is the table column denoting the business_unit relation/edge.
	BusinessUnitColumn = "business_unit_id"
	// OrganizationTable is the table that holds the organization relation/edge.
	OrganizationTable = "customer_detention_policies"
	// OrganizationInverseTable is the table name for the Organization entity.
	// It exists in this package in order to avoid circular dependency with the "organization" package.
	OrganizationInverseTable = "organizations"
	// OrganizationColumn is the table column denoting the organization relation/edge.
	OrganizationColumn = "organization_id"
	// CustomerTable is the table that holds the customer relation/edge.
	CustomerTable = "customer_detention_policies"
	// CustomerInverseTable is the table name for the Customer entity.
	// It exists in this package in order to avoid circular dependency with the "customer" package.
	CustomerInverseTable = "customers"
	// CustomerColumn is the table column denoting the customer relation/edge.
	CustomerColumn = "customer_id"
	// CommodityTable is the table that holds the commodity relation/edge.
	CommodityTable = "customer_detention_policies"
	// CommodityInverseTable is the table name for the Commodity entity.
	// It exists in this package in order to avoid circular dependency with the "commodity" package.
	CommodityInverseTable = "commodities"
	// CommodityColumn is the table column denoting the commodity relation/edge.
	CommodityColumn = "commodity_id"
	// RevenueCodeTable is the table that holds the revenue_code relation/edge.
	RevenueCodeTable = "customer_detention_policies"
	// RevenueCodeInverseTable is the table name for the RevenueCode entity.
	// It exists in this package in order to avoid circular dependency with the "revenuecode" package.
	RevenueCodeInverseTable = "revenue_codes"
	// RevenueCodeColumn is the table column denoting the revenue_code relation/edge.
	RevenueCodeColumn = "revenue_code_id"
	// ShipmentTypeTable is the table that holds the shipment_type relation/edge.
	ShipmentTypeTable = "customer_detention_policies"
	// ShipmentTypeInverseTable is the table name for the ShipmentType entity.
	// It exists in this package in order to avoid circular dependency with the "shipmenttype" package.
	ShipmentTypeInverseTable = "shipment_types"
	// ShipmentTypeColumn is the table column denoting the shipment_type relation/edge.
	ShipmentTypeColumn = "shipment_type_id"
	// AccessorialChargeTable is the table that holds the accessorial_charge relation/edge.
	AccessorialChargeTable = "customer_detention_policies"
	// AccessorialChargeInverseTable is the table name for the AccessorialCharge entity.
	// It exists in this package in order to avoid circular dependency with the "accessorialcharge" package.
	AccessorialChargeInverseTable = "accessorial_charges"
	// AccessorialChargeColumn is the table column denoting the accessorial_charge relation/edge.
	AccessorialChargeColumn = "accessorial_charge_id"
)

// Columns holds all SQL columns for customerdetentionpolicy fields.
var Columns = []string{
	FieldID,
	FieldBusinessUnitID,
	FieldOrganizationID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldVersion,
	FieldStatus,
	FieldCustomerID,
	FieldCommodityID,
	FieldRevenueCodeID,
	FieldShipmentTypeID,
	FieldApplicationScope,
	FieldChargeFreeTime,
	FieldPaymentFreeTime,
	FieldLateArrivalPolicy,
	FieldGracePeriod,
	FieldAccessorialChargeID,
	FieldUnits,
	FieldAmount,
	FieldNotes,
	FieldEffectiveDate,
	FieldExpirationDate,
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
	// ChargeFreeTimeValidator is a validator for the "charge_free_time" field. It is called by the builders before save.
	ChargeFreeTimeValidator func(int) error
	// PaymentFreeTimeValidator is a validator for the "payment_free_time" field. It is called by the builders before save.
	PaymentFreeTimeValidator func(int) error
	// DefaultLateArrivalPolicy holds the default value on creation for the "late_arrival_policy" field.
	DefaultLateArrivalPolicy bool
	// GracePeriodValidator is a validator for the "grace_period" field. It is called by the builders before save.
	GracePeriodValidator func(int) error
	// UnitsValidator is a validator for the "units" field. It is called by the builders before save.
	UnitsValidator func(int) error
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
		return fmt.Errorf("customerdetentionpolicy: invalid enum value for status field: %q", s)
	}
}

// ApplicationScope defines the type for the "application_scope" enum field.
type ApplicationScope string

// ApplicationScopePICKUP is the default value of the ApplicationScope enum.
const DefaultApplicationScope = ApplicationScopePICKUP

// ApplicationScope values.
const (
	ApplicationScopePICKUP   ApplicationScope = "PICKUP"
	ApplicationScopeDELIVERY ApplicationScope = "DELIVERY"
	ApplicationScopeBOTH     ApplicationScope = "BOTH"
)

func (as ApplicationScope) String() string {
	return string(as)
}

// ApplicationScopeValidator is a validator for the "application_scope" field enum values. It is called by the builders before save.
func ApplicationScopeValidator(as ApplicationScope) error {
	switch as {
	case ApplicationScopePICKUP, ApplicationScopeDELIVERY, ApplicationScopeBOTH:
		return nil
	default:
		return fmt.Errorf("customerdetentionpolicy: invalid enum value for application_scope field: %q", as)
	}
}

// OrderOption defines the ordering options for the CustomerDetentionPolicy queries.
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

// ByCustomerID orders the results by the customer_id field.
func ByCustomerID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCustomerID, opts...).ToFunc()
}

// ByCommodityID orders the results by the commodity_id field.
func ByCommodityID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCommodityID, opts...).ToFunc()
}

// ByRevenueCodeID orders the results by the revenue_code_id field.
func ByRevenueCodeID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRevenueCodeID, opts...).ToFunc()
}

// ByShipmentTypeID orders the results by the shipment_type_id field.
func ByShipmentTypeID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldShipmentTypeID, opts...).ToFunc()
}

// ByApplicationScope orders the results by the application_scope field.
func ByApplicationScope(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldApplicationScope, opts...).ToFunc()
}

// ByChargeFreeTime orders the results by the charge_free_time field.
func ByChargeFreeTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldChargeFreeTime, opts...).ToFunc()
}

// ByPaymentFreeTime orders the results by the payment_free_time field.
func ByPaymentFreeTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPaymentFreeTime, opts...).ToFunc()
}

// ByLateArrivalPolicy orders the results by the late_arrival_policy field.
func ByLateArrivalPolicy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLateArrivalPolicy, opts...).ToFunc()
}

// ByGracePeriod orders the results by the grace_period field.
func ByGracePeriod(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGracePeriod, opts...).ToFunc()
}

// ByAccessorialChargeID orders the results by the accessorial_charge_id field.
func ByAccessorialChargeID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAccessorialChargeID, opts...).ToFunc()
}

// ByUnits orders the results by the units field.
func ByUnits(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUnits, opts...).ToFunc()
}

// ByAmount orders the results by the amount field.
func ByAmount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAmount, opts...).ToFunc()
}

// ByNotes orders the results by the notes field.
func ByNotes(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNotes, opts...).ToFunc()
}

// ByEffectiveDate orders the results by the effective_date field.
func ByEffectiveDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEffectiveDate, opts...).ToFunc()
}

// ByExpirationDate orders the results by the expiration_date field.
func ByExpirationDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExpirationDate, opts...).ToFunc()
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

// ByCustomerField orders the results by customer field.
func ByCustomerField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCustomerStep(), sql.OrderByField(field, opts...))
	}
}

// ByCommodityField orders the results by commodity field.
func ByCommodityField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCommodityStep(), sql.OrderByField(field, opts...))
	}
}

// ByRevenueCodeField orders the results by revenue_code field.
func ByRevenueCodeField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRevenueCodeStep(), sql.OrderByField(field, opts...))
	}
}

// ByShipmentTypeField orders the results by shipment_type field.
func ByShipmentTypeField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newShipmentTypeStep(), sql.OrderByField(field, opts...))
	}
}

// ByAccessorialChargeField orders the results by accessorial_charge field.
func ByAccessorialChargeField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAccessorialChargeStep(), sql.OrderByField(field, opts...))
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
func newCustomerStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CustomerInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, CustomerTable, CustomerColumn),
	)
}
func newCommodityStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CommodityInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, CommodityTable, CommodityColumn),
	)
}
func newRevenueCodeStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RevenueCodeInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, RevenueCodeTable, RevenueCodeColumn),
	)
}
func newShipmentTypeStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ShipmentTypeInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, ShipmentTypeTable, ShipmentTypeColumn),
	)
}
func newAccessorialChargeStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AccessorialChargeInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, AccessorialChargeTable, AccessorialChargeColumn),
	)
}
