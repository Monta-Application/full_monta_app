// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/emoss08/trenova/internal/ent/accessorialcharge"
	"github.com/emoss08/trenova/internal/ent/businessunit"
	"github.com/emoss08/trenova/internal/ent/commodity"
	"github.com/emoss08/trenova/internal/ent/customer"
	"github.com/emoss08/trenova/internal/ent/customerdetentionpolicy"
	"github.com/emoss08/trenova/internal/ent/organization"
	"github.com/emoss08/trenova/internal/ent/revenuecode"
	"github.com/emoss08/trenova/internal/ent/shipmenttype"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

// CustomerDetentionPolicy is the model entity for the CustomerDetentionPolicy schema.
type CustomerDetentionPolicy struct {
	config `json:"-" validate:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// BusinessUnitID holds the value of the "business_unit_id" field.
	BusinessUnitID uuid.UUID `json:"businessUnitId"`
	// OrganizationID holds the value of the "organization_id" field.
	OrganizationID uuid.UUID `json:"organizationId"`
	// The time that this entity was created.
	CreatedAt time.Time `json:"createdAt" validate:"omitempty"`
	// The last time that this entity was updated.
	UpdatedAt time.Time `json:"updatedAt" validate:"omitempty"`
	// The current version of this entity.
	Version int `json:"version" validate:"omitempty"`
	// Status holds the value of the "status" field.
	Status customerdetentionpolicy.Status `json:"status,omitempty"`
	// CustomerID holds the value of the "customer_id" field.
	CustomerID uuid.UUID `json:"customer_id,omitempty"`
	// The type of commodity to which the detention policy applies. This helps in customizing policies for different commodities.
	CommodityID *uuid.UUID `json:"commodityId" validate:"omitempty"`
	// A unique code associated with the revenue generated from detention charges.
	RevenueCodeID *uuid.UUID `json:"revenueCodeId" validate:"omitempty"`
	// Type of shipment (e.g., Standard, Expedited) to which the detention policy is applicable.
	ShipmentTypeID *uuid.UUID `json:"shipmentTypeId" validate:"omitempty"`
	// Specifies whether the policy applies to pickups, deliveries, or both.
	ApplicationScope customerdetentionpolicy.ApplicationScope `json:"applicationScope" validate:"required,oneof=PICKUP DELIVERY BOTH"`
	// The threshold time (in minutes) for the start of detention charges. This represents the allowed free time before charges apply.
	ChargeFreeTime int `json:"chargeFreeTime" validate:"omitempty,gt=0"`
	// The time (in minutes) considered for calculating detention payments. This can differ from charge_free_time in certain scenarios.
	PaymentFreeTime int `json:"paymentFreeTime" validate:"omitempty,gt=0"`
	// Indicates whether the policy applies to late arrivals. True if detention charges apply to late arrivals.
	LateArrivalPolicy bool `json:"lateArrivalPolicy" validate:"omitempty"`
	// An additional time buffer (in minutes) provided before detention charges kick in, often used to accommodate slight delays.
	GracePeriod int `json:"gracePeriod" validate:"omitempty,gt=0"`
	// The unique identifier for the accessorial charge associated with the detention policy.
	AccessorialChargeID *uuid.UUID `json:"accessorialChargeId" validate:"omitempty"`
	// The number of units (e.g., pallets, containers) considered for detention charges.
	Units int `json:"units" validate:"omitempty,gt=0"`
	// Amount holds the value of the "amount" field.
	Amount float64 `json:"amount" validate:"required,gt=0"`
	// Additional notes or comments about the detention policy.
	Notes string `json:"notes" validate:"omitempty"`
	// The date when the detention policy becomes effective.
	EffectiveDate *pgtype.Date `json:"effectiveDate"`
	// The date when the detention policy expires.
	ExpirationDate *pgtype.Date `json:"expirationDate"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CustomerDetentionPolicyQuery when eager-loading is set.
	Edges        CustomerDetentionPolicyEdges `json:"edges"`
	selectValues sql.SelectValues
}

// CustomerDetentionPolicyEdges holds the relations/edges for other nodes in the graph.
type CustomerDetentionPolicyEdges struct {
	// BusinessUnit holds the value of the business_unit edge.
	BusinessUnit *BusinessUnit `json:"business_unit,omitempty"`
	// Organization holds the value of the organization edge.
	Organization *Organization `json:"organization,omitempty"`
	// Customer holds the value of the customer edge.
	Customer *Customer `json:"customer,omitempty"`
	// Commodity holds the value of the commodity edge.
	Commodity *Commodity `json:"commodity"`
	// RevenueCode holds the value of the revenue_code edge.
	RevenueCode *RevenueCode `json:"revenueCode"`
	// ShipmentType holds the value of the shipment_type edge.
	ShipmentType *ShipmentType `json:"shipmentType"`
	// AccessorialCharge holds the value of the accessorial_charge edge.
	AccessorialCharge *AccessorialCharge `json:"accessorialCharge"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [7]bool
}

// BusinessUnitOrErr returns the BusinessUnit value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CustomerDetentionPolicyEdges) BusinessUnitOrErr() (*BusinessUnit, error) {
	if e.BusinessUnit != nil {
		return e.BusinessUnit, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: businessunit.Label}
	}
	return nil, &NotLoadedError{edge: "business_unit"}
}

// OrganizationOrErr returns the Organization value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CustomerDetentionPolicyEdges) OrganizationOrErr() (*Organization, error) {
	if e.Organization != nil {
		return e.Organization, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: organization.Label}
	}
	return nil, &NotLoadedError{edge: "organization"}
}

// CustomerOrErr returns the Customer value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CustomerDetentionPolicyEdges) CustomerOrErr() (*Customer, error) {
	if e.Customer != nil {
		return e.Customer, nil
	} else if e.loadedTypes[2] {
		return nil, &NotFoundError{label: customer.Label}
	}
	return nil, &NotLoadedError{edge: "customer"}
}

// CommodityOrErr returns the Commodity value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CustomerDetentionPolicyEdges) CommodityOrErr() (*Commodity, error) {
	if e.Commodity != nil {
		return e.Commodity, nil
	} else if e.loadedTypes[3] {
		return nil, &NotFoundError{label: commodity.Label}
	}
	return nil, &NotLoadedError{edge: "commodity"}
}

// RevenueCodeOrErr returns the RevenueCode value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CustomerDetentionPolicyEdges) RevenueCodeOrErr() (*RevenueCode, error) {
	if e.RevenueCode != nil {
		return e.RevenueCode, nil
	} else if e.loadedTypes[4] {
		return nil, &NotFoundError{label: revenuecode.Label}
	}
	return nil, &NotLoadedError{edge: "revenue_code"}
}

// ShipmentTypeOrErr returns the ShipmentType value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CustomerDetentionPolicyEdges) ShipmentTypeOrErr() (*ShipmentType, error) {
	if e.ShipmentType != nil {
		return e.ShipmentType, nil
	} else if e.loadedTypes[5] {
		return nil, &NotFoundError{label: shipmenttype.Label}
	}
	return nil, &NotLoadedError{edge: "shipment_type"}
}

// AccessorialChargeOrErr returns the AccessorialCharge value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CustomerDetentionPolicyEdges) AccessorialChargeOrErr() (*AccessorialCharge, error) {
	if e.AccessorialCharge != nil {
		return e.AccessorialCharge, nil
	} else if e.loadedTypes[6] {
		return nil, &NotFoundError{label: accessorialcharge.Label}
	}
	return nil, &NotLoadedError{edge: "accessorial_charge"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*CustomerDetentionPolicy) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case customerdetentionpolicy.FieldEffectiveDate, customerdetentionpolicy.FieldExpirationDate:
			values[i] = &sql.NullScanner{S: new(pgtype.Date)}
		case customerdetentionpolicy.FieldCommodityID, customerdetentionpolicy.FieldRevenueCodeID, customerdetentionpolicy.FieldShipmentTypeID, customerdetentionpolicy.FieldAccessorialChargeID:
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case customerdetentionpolicy.FieldLateArrivalPolicy:
			values[i] = new(sql.NullBool)
		case customerdetentionpolicy.FieldAmount:
			values[i] = new(sql.NullFloat64)
		case customerdetentionpolicy.FieldVersion, customerdetentionpolicy.FieldChargeFreeTime, customerdetentionpolicy.FieldPaymentFreeTime, customerdetentionpolicy.FieldGracePeriod, customerdetentionpolicy.FieldUnits:
			values[i] = new(sql.NullInt64)
		case customerdetentionpolicy.FieldStatus, customerdetentionpolicy.FieldApplicationScope, customerdetentionpolicy.FieldNotes:
			values[i] = new(sql.NullString)
		case customerdetentionpolicy.FieldCreatedAt, customerdetentionpolicy.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case customerdetentionpolicy.FieldID, customerdetentionpolicy.FieldBusinessUnitID, customerdetentionpolicy.FieldOrganizationID, customerdetentionpolicy.FieldCustomerID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CustomerDetentionPolicy fields.
func (cdp *CustomerDetentionPolicy) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case customerdetentionpolicy.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				cdp.ID = *value
			}
		case customerdetentionpolicy.FieldBusinessUnitID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field business_unit_id", values[i])
			} else if value != nil {
				cdp.BusinessUnitID = *value
			}
		case customerdetentionpolicy.FieldOrganizationID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field organization_id", values[i])
			} else if value != nil {
				cdp.OrganizationID = *value
			}
		case customerdetentionpolicy.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				cdp.CreatedAt = value.Time
			}
		case customerdetentionpolicy.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				cdp.UpdatedAt = value.Time
			}
		case customerdetentionpolicy.FieldVersion:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field version", values[i])
			} else if value.Valid {
				cdp.Version = int(value.Int64)
			}
		case customerdetentionpolicy.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				cdp.Status = customerdetentionpolicy.Status(value.String)
			}
		case customerdetentionpolicy.FieldCustomerID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field customer_id", values[i])
			} else if value != nil {
				cdp.CustomerID = *value
			}
		case customerdetentionpolicy.FieldCommodityID:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field commodity_id", values[i])
			} else if value.Valid {
				cdp.CommodityID = new(uuid.UUID)
				*cdp.CommodityID = *value.S.(*uuid.UUID)
			}
		case customerdetentionpolicy.FieldRevenueCodeID:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field revenue_code_id", values[i])
			} else if value.Valid {
				cdp.RevenueCodeID = new(uuid.UUID)
				*cdp.RevenueCodeID = *value.S.(*uuid.UUID)
			}
		case customerdetentionpolicy.FieldShipmentTypeID:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field shipment_type_id", values[i])
			} else if value.Valid {
				cdp.ShipmentTypeID = new(uuid.UUID)
				*cdp.ShipmentTypeID = *value.S.(*uuid.UUID)
			}
		case customerdetentionpolicy.FieldApplicationScope:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field application_scope", values[i])
			} else if value.Valid {
				cdp.ApplicationScope = customerdetentionpolicy.ApplicationScope(value.String)
			}
		case customerdetentionpolicy.FieldChargeFreeTime:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field charge_free_time", values[i])
			} else if value.Valid {
				cdp.ChargeFreeTime = int(value.Int64)
			}
		case customerdetentionpolicy.FieldPaymentFreeTime:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field payment_free_time", values[i])
			} else if value.Valid {
				cdp.PaymentFreeTime = int(value.Int64)
			}
		case customerdetentionpolicy.FieldLateArrivalPolicy:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field late_arrival_policy", values[i])
			} else if value.Valid {
				cdp.LateArrivalPolicy = value.Bool
			}
		case customerdetentionpolicy.FieldGracePeriod:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field grace_period", values[i])
			} else if value.Valid {
				cdp.GracePeriod = int(value.Int64)
			}
		case customerdetentionpolicy.FieldAccessorialChargeID:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field accessorial_charge_id", values[i])
			} else if value.Valid {
				cdp.AccessorialChargeID = new(uuid.UUID)
				*cdp.AccessorialChargeID = *value.S.(*uuid.UUID)
			}
		case customerdetentionpolicy.FieldUnits:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field units", values[i])
			} else if value.Valid {
				cdp.Units = int(value.Int64)
			}
		case customerdetentionpolicy.FieldAmount:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field amount", values[i])
			} else if value.Valid {
				cdp.Amount = value.Float64
			}
		case customerdetentionpolicy.FieldNotes:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field notes", values[i])
			} else if value.Valid {
				cdp.Notes = value.String
			}
		case customerdetentionpolicy.FieldEffectiveDate:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field effective_date", values[i])
			} else if value.Valid {
				cdp.EffectiveDate = value.S.(*pgtype.Date)
			}
		case customerdetentionpolicy.FieldExpirationDate:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field expiration_date", values[i])
			} else if value.Valid {
				cdp.ExpirationDate = value.S.(*pgtype.Date)
			}
		default:
			cdp.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the CustomerDetentionPolicy.
// This includes values selected through modifiers, order, etc.
func (cdp *CustomerDetentionPolicy) Value(name string) (ent.Value, error) {
	return cdp.selectValues.Get(name)
}

// QueryBusinessUnit queries the "business_unit" edge of the CustomerDetentionPolicy entity.
func (cdp *CustomerDetentionPolicy) QueryBusinessUnit() *BusinessUnitQuery {
	return NewCustomerDetentionPolicyClient(cdp.config).QueryBusinessUnit(cdp)
}

// QueryOrganization queries the "organization" edge of the CustomerDetentionPolicy entity.
func (cdp *CustomerDetentionPolicy) QueryOrganization() *OrganizationQuery {
	return NewCustomerDetentionPolicyClient(cdp.config).QueryOrganization(cdp)
}

// QueryCustomer queries the "customer" edge of the CustomerDetentionPolicy entity.
func (cdp *CustomerDetentionPolicy) QueryCustomer() *CustomerQuery {
	return NewCustomerDetentionPolicyClient(cdp.config).QueryCustomer(cdp)
}

// QueryCommodity queries the "commodity" edge of the CustomerDetentionPolicy entity.
func (cdp *CustomerDetentionPolicy) QueryCommodity() *CommodityQuery {
	return NewCustomerDetentionPolicyClient(cdp.config).QueryCommodity(cdp)
}

// QueryRevenueCode queries the "revenue_code" edge of the CustomerDetentionPolicy entity.
func (cdp *CustomerDetentionPolicy) QueryRevenueCode() *RevenueCodeQuery {
	return NewCustomerDetentionPolicyClient(cdp.config).QueryRevenueCode(cdp)
}

// QueryShipmentType queries the "shipment_type" edge of the CustomerDetentionPolicy entity.
func (cdp *CustomerDetentionPolicy) QueryShipmentType() *ShipmentTypeQuery {
	return NewCustomerDetentionPolicyClient(cdp.config).QueryShipmentType(cdp)
}

// QueryAccessorialCharge queries the "accessorial_charge" edge of the CustomerDetentionPolicy entity.
func (cdp *CustomerDetentionPolicy) QueryAccessorialCharge() *AccessorialChargeQuery {
	return NewCustomerDetentionPolicyClient(cdp.config).QueryAccessorialCharge(cdp)
}

// Update returns a builder for updating this CustomerDetentionPolicy.
// Note that you need to call CustomerDetentionPolicy.Unwrap() before calling this method if this CustomerDetentionPolicy
// was returned from a transaction, and the transaction was committed or rolled back.
func (cdp *CustomerDetentionPolicy) Update() *CustomerDetentionPolicyUpdateOne {
	return NewCustomerDetentionPolicyClient(cdp.config).UpdateOne(cdp)
}

// Unwrap unwraps the CustomerDetentionPolicy entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (cdp *CustomerDetentionPolicy) Unwrap() *CustomerDetentionPolicy {
	_tx, ok := cdp.config.driver.(*txDriver)
	if !ok {
		panic("ent: CustomerDetentionPolicy is not a transactional entity")
	}
	cdp.config.driver = _tx.drv
	return cdp
}

// String implements the fmt.Stringer.
func (cdp *CustomerDetentionPolicy) String() string {
	var builder strings.Builder
	builder.WriteString("CustomerDetentionPolicy(")
	builder.WriteString(fmt.Sprintf("id=%v, ", cdp.ID))
	builder.WriteString("business_unit_id=")
	builder.WriteString(fmt.Sprintf("%v", cdp.BusinessUnitID))
	builder.WriteString(", ")
	builder.WriteString("organization_id=")
	builder.WriteString(fmt.Sprintf("%v", cdp.OrganizationID))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(cdp.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(cdp.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("version=")
	builder.WriteString(fmt.Sprintf("%v", cdp.Version))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", cdp.Status))
	builder.WriteString(", ")
	builder.WriteString("customer_id=")
	builder.WriteString(fmt.Sprintf("%v", cdp.CustomerID))
	builder.WriteString(", ")
	if v := cdp.CommodityID; v != nil {
		builder.WriteString("commodity_id=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := cdp.RevenueCodeID; v != nil {
		builder.WriteString("revenue_code_id=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := cdp.ShipmentTypeID; v != nil {
		builder.WriteString("shipment_type_id=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	builder.WriteString("application_scope=")
	builder.WriteString(fmt.Sprintf("%v", cdp.ApplicationScope))
	builder.WriteString(", ")
	builder.WriteString("charge_free_time=")
	builder.WriteString(fmt.Sprintf("%v", cdp.ChargeFreeTime))
	builder.WriteString(", ")
	builder.WriteString("payment_free_time=")
	builder.WriteString(fmt.Sprintf("%v", cdp.PaymentFreeTime))
	builder.WriteString(", ")
	builder.WriteString("late_arrival_policy=")
	builder.WriteString(fmt.Sprintf("%v", cdp.LateArrivalPolicy))
	builder.WriteString(", ")
	builder.WriteString("grace_period=")
	builder.WriteString(fmt.Sprintf("%v", cdp.GracePeriod))
	builder.WriteString(", ")
	if v := cdp.AccessorialChargeID; v != nil {
		builder.WriteString("accessorial_charge_id=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	builder.WriteString("units=")
	builder.WriteString(fmt.Sprintf("%v", cdp.Units))
	builder.WriteString(", ")
	builder.WriteString("amount=")
	builder.WriteString(fmt.Sprintf("%v", cdp.Amount))
	builder.WriteString(", ")
	builder.WriteString("notes=")
	builder.WriteString(cdp.Notes)
	builder.WriteString(", ")
	if v := cdp.EffectiveDate; v != nil {
		builder.WriteString("effective_date=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := cdp.ExpirationDate; v != nil {
		builder.WriteString("expiration_date=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteByte(')')
	return builder.String()
}

// CustomerDetentionPolicies is a parsable slice of CustomerDetentionPolicy.
type CustomerDetentionPolicies []*CustomerDetentionPolicy
