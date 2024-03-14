// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/emoss08/trenova/ent/businessunit"
	"github.com/google/uuid"
)

// BusinessUnit is the model entity for the BusinessUnit schema.
type BusinessUnit struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"createdAt"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updatedAt"`
	// Status holds the value of the "status" field.
	Status businessunit.Status `json:"status,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// EntityKey holds the value of the "entity_key" field.
	EntityKey string `json:"entity_key,omitempty"`
	// PhoneNumber holds the value of the "phone_number" field.
	PhoneNumber string `json:"phoneNumber"`
	// Address holds the value of the "address" field.
	Address string `json:"address,omitempty"`
	// City holds the value of the "city" field.
	City *string `json:"city,omitempty"`
	// State holds the value of the "state" field.
	State *string `json:"state,omitempty"`
	// Country holds the value of the "country" field.
	Country *string `json:"country,omitempty"`
	// PostalCode holds the value of the "postal_code" field.
	PostalCode *string `json:"postalCode"`
	// TaxID holds the value of the "tax_id" field.
	TaxID *string `json:"taxId"`
	// SubscriptionPlan holds the value of the "subscription_plan" field.
	SubscriptionPlan *string `json:"subscriptionPlan"`
	// Description holds the value of the "description" field.
	Description *string `json:"description,omitempty"`
	// LegalName holds the value of the "legal_name" field.
	LegalName *string `json:"legalName"`
	// ContactName holds the value of the "contact_name" field.
	ContactName *string `json:"contactName"`
	// ContactEmail holds the value of the "contact_email" field.
	ContactEmail *string `json:"contactEmail"`
	// PaidUntil holds the value of the "paid_until" field.
	PaidUntil *time.Time `json:"-"`
	// Settings holds the value of the "settings" field.
	Settings map[string]interface{} `json:"settings,omitempty"`
	// FreeTrial holds the value of the "free_trial" field.
	FreeTrial bool `json:"freeTrial"`
	// ParentID holds the value of the "parent_id" field.
	ParentID *uuid.UUID `json:"parentId"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BusinessUnitQuery when eager-loading is set.
	Edges        BusinessUnitEdges `json:"edges"`
	selectValues sql.SelectValues
}

// BusinessUnitEdges holds the relations/edges for other nodes in the graph.
type BusinessUnitEdges struct {
	// Prev holds the value of the prev edge.
	Prev *BusinessUnit `json:"parent_id"`
	// Next holds the value of the next edge.
	Next *BusinessUnit `json:"next,omitempty"`
	// Organizations holds the value of the organizations edge.
	Organizations []*Organization `json:"organizations,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// PrevOrErr returns the Prev value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BusinessUnitEdges) PrevOrErr() (*BusinessUnit, error) {
	if e.Prev != nil {
		return e.Prev, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: businessunit.Label}
	}
	return nil, &NotLoadedError{edge: "prev"}
}

// NextOrErr returns the Next value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BusinessUnitEdges) NextOrErr() (*BusinessUnit, error) {
	if e.Next != nil {
		return e.Next, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: businessunit.Label}
	}
	return nil, &NotLoadedError{edge: "next"}
}

// OrganizationsOrErr returns the Organizations value or an error if the edge
// was not loaded in eager-loading.
func (e BusinessUnitEdges) OrganizationsOrErr() ([]*Organization, error) {
	if e.loadedTypes[2] {
		return e.Organizations, nil
	}
	return nil, &NotLoadedError{edge: "organizations"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*BusinessUnit) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case businessunit.FieldParentID:
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case businessunit.FieldSettings:
			values[i] = new([]byte)
		case businessunit.FieldFreeTrial:
			values[i] = new(sql.NullBool)
		case businessunit.FieldStatus, businessunit.FieldName, businessunit.FieldEntityKey, businessunit.FieldPhoneNumber, businessunit.FieldAddress, businessunit.FieldCity, businessunit.FieldState, businessunit.FieldCountry, businessunit.FieldPostalCode, businessunit.FieldTaxID, businessunit.FieldSubscriptionPlan, businessunit.FieldDescription, businessunit.FieldLegalName, businessunit.FieldContactName, businessunit.FieldContactEmail:
			values[i] = new(sql.NullString)
		case businessunit.FieldCreatedAt, businessunit.FieldUpdatedAt, businessunit.FieldPaidUntil:
			values[i] = new(sql.NullTime)
		case businessunit.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the BusinessUnit fields.
func (bu *BusinessUnit) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case businessunit.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				bu.ID = *value
			}
		case businessunit.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				bu.CreatedAt = value.Time
			}
		case businessunit.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				bu.UpdatedAt = value.Time
			}
		case businessunit.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				bu.Status = businessunit.Status(value.String)
			}
		case businessunit.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				bu.Name = value.String
			}
		case businessunit.FieldEntityKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field entity_key", values[i])
			} else if value.Valid {
				bu.EntityKey = value.String
			}
		case businessunit.FieldPhoneNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field phone_number", values[i])
			} else if value.Valid {
				bu.PhoneNumber = value.String
			}
		case businessunit.FieldAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address", values[i])
			} else if value.Valid {
				bu.Address = value.String
			}
		case businessunit.FieldCity:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field city", values[i])
			} else if value.Valid {
				bu.City = new(string)
				*bu.City = value.String
			}
		case businessunit.FieldState:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field state", values[i])
			} else if value.Valid {
				bu.State = new(string)
				*bu.State = value.String
			}
		case businessunit.FieldCountry:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field country", values[i])
			} else if value.Valid {
				bu.Country = new(string)
				*bu.Country = value.String
			}
		case businessunit.FieldPostalCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field postal_code", values[i])
			} else if value.Valid {
				bu.PostalCode = new(string)
				*bu.PostalCode = value.String
			}
		case businessunit.FieldTaxID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tax_id", values[i])
			} else if value.Valid {
				bu.TaxID = new(string)
				*bu.TaxID = value.String
			}
		case businessunit.FieldSubscriptionPlan:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field subscription_plan", values[i])
			} else if value.Valid {
				bu.SubscriptionPlan = new(string)
				*bu.SubscriptionPlan = value.String
			}
		case businessunit.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				bu.Description = new(string)
				*bu.Description = value.String
			}
		case businessunit.FieldLegalName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field legal_name", values[i])
			} else if value.Valid {
				bu.LegalName = new(string)
				*bu.LegalName = value.String
			}
		case businessunit.FieldContactName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field contact_name", values[i])
			} else if value.Valid {
				bu.ContactName = new(string)
				*bu.ContactName = value.String
			}
		case businessunit.FieldContactEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field contact_email", values[i])
			} else if value.Valid {
				bu.ContactEmail = new(string)
				*bu.ContactEmail = value.String
			}
		case businessunit.FieldPaidUntil:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field paid_until", values[i])
			} else if value.Valid {
				bu.PaidUntil = new(time.Time)
				*bu.PaidUntil = value.Time
			}
		case businessunit.FieldSettings:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field settings", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &bu.Settings); err != nil {
					return fmt.Errorf("unmarshal field settings: %w", err)
				}
			}
		case businessunit.FieldFreeTrial:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field free_trial", values[i])
			} else if value.Valid {
				bu.FreeTrial = value.Bool
			}
		case businessunit.FieldParentID:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field parent_id", values[i])
			} else if value.Valid {
				bu.ParentID = new(uuid.UUID)
				*bu.ParentID = *value.S.(*uuid.UUID)
			}
		default:
			bu.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the BusinessUnit.
// This includes values selected through modifiers, order, etc.
func (bu *BusinessUnit) Value(name string) (ent.Value, error) {
	return bu.selectValues.Get(name)
}

// QueryPrev queries the "prev" edge of the BusinessUnit entity.
func (bu *BusinessUnit) QueryPrev() *BusinessUnitQuery {
	return NewBusinessUnitClient(bu.config).QueryPrev(bu)
}

// QueryNext queries the "next" edge of the BusinessUnit entity.
func (bu *BusinessUnit) QueryNext() *BusinessUnitQuery {
	return NewBusinessUnitClient(bu.config).QueryNext(bu)
}

// QueryOrganizations queries the "organizations" edge of the BusinessUnit entity.
func (bu *BusinessUnit) QueryOrganizations() *OrganizationQuery {
	return NewBusinessUnitClient(bu.config).QueryOrganizations(bu)
}

// Update returns a builder for updating this BusinessUnit.
// Note that you need to call BusinessUnit.Unwrap() before calling this method if this BusinessUnit
// was returned from a transaction, and the transaction was committed or rolled back.
func (bu *BusinessUnit) Update() *BusinessUnitUpdateOne {
	return NewBusinessUnitClient(bu.config).UpdateOne(bu)
}

// Unwrap unwraps the BusinessUnit entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (bu *BusinessUnit) Unwrap() *BusinessUnit {
	_tx, ok := bu.config.driver.(*txDriver)
	if !ok {
		panic("ent: BusinessUnit is not a transactional entity")
	}
	bu.config.driver = _tx.drv
	return bu
}

// String implements the fmt.Stringer.
func (bu *BusinessUnit) String() string {
	var builder strings.Builder
	builder.WriteString("BusinessUnit(")
	builder.WriteString(fmt.Sprintf("id=%v, ", bu.ID))
	builder.WriteString("created_at=")
	builder.WriteString(bu.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(bu.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", bu.Status))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(bu.Name)
	builder.WriteString(", ")
	builder.WriteString("entity_key=")
	builder.WriteString(bu.EntityKey)
	builder.WriteString(", ")
	builder.WriteString("phone_number=")
	builder.WriteString(bu.PhoneNumber)
	builder.WriteString(", ")
	builder.WriteString("address=")
	builder.WriteString(bu.Address)
	builder.WriteString(", ")
	if v := bu.City; v != nil {
		builder.WriteString("city=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := bu.State; v != nil {
		builder.WriteString("state=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := bu.Country; v != nil {
		builder.WriteString("country=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := bu.PostalCode; v != nil {
		builder.WriteString("postal_code=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := bu.TaxID; v != nil {
		builder.WriteString("tax_id=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := bu.SubscriptionPlan; v != nil {
		builder.WriteString("subscription_plan=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := bu.Description; v != nil {
		builder.WriteString("description=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := bu.LegalName; v != nil {
		builder.WriteString("legal_name=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := bu.ContactName; v != nil {
		builder.WriteString("contact_name=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := bu.ContactEmail; v != nil {
		builder.WriteString("contact_email=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := bu.PaidUntil; v != nil {
		builder.WriteString("paid_until=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("settings=")
	builder.WriteString(fmt.Sprintf("%v", bu.Settings))
	builder.WriteString(", ")
	builder.WriteString("free_trial=")
	builder.WriteString(fmt.Sprintf("%v", bu.FreeTrial))
	builder.WriteString(", ")
	if v := bu.ParentID; v != nil {
		builder.WriteString("parent_id=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteByte(')')
	return builder.String()
}

// BusinessUnits is a parsable slice of BusinessUnit.
type BusinessUnits []*BusinessUnit
