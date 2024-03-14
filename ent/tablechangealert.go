// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/emoss08/trenova/ent/businessunit"
	"github.com/emoss08/trenova/ent/organization"
	"github.com/emoss08/trenova/ent/tablechangealert"
	"github.com/google/uuid"
)

// TableChangeAlert is the model entity for the TableChangeAlert schema.
type TableChangeAlert struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// BusinessUnitID holds the value of the "business_unit_id" field.
	BusinessUnitID uuid.UUID `json:"businessUnitId"`
	// OrganizationID holds the value of the "organization_id" field.
	OrganizationID uuid.UUID `json:"organizationId"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"createdAt"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updatedAt"`
	// Status holds the value of the "status" field.
	Status tablechangealert.Status `json:"status"`
	// Name holds the value of the "name" field.
	Name string `json:"name"`
	// DatabaseAction holds the value of the "database_action" field.
	DatabaseAction tablechangealert.DatabaseAction `json:"databaseAction"`
	// Source holds the value of the "source" field.
	Source tablechangealert.Source `json:"source"`
	// TableName holds the value of the "table_name" field.
	TableName *string `json:"tableName"`
	// Topic holds the value of the "topic" field.
	Topic *string `json:"topic"`
	// Description holds the value of the "description" field.
	Description *string `json:"description"`
	// CustomSubject holds the value of the "custom_subject" field.
	CustomSubject *string `json:"customSubject"`
	// FunctionName holds the value of the "function_name" field.
	FunctionName *string `json:"functionName"`
	// TriggerName holds the value of the "trigger_name" field.
	TriggerName *string `json:"triggerName"`
	// ListenerName holds the value of the "listener_name" field.
	ListenerName *string `json:"listenerName"`
	// EmailRecipients holds the value of the "email_recipients" field.
	EmailRecipients *string `json:"emailRecipients"`
	// EffectiveDate holds the value of the "effective_date" field.
	EffectiveDate *time.Time `json:"effectiveDate"`
	// ExpirationDate holds the value of the "expiration_date" field.
	ExpirationDate *time.Time `json:"expirationDate"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TableChangeAlertQuery when eager-loading is set.
	Edges        TableChangeAlertEdges `json:"edges"`
	selectValues sql.SelectValues
}

// TableChangeAlertEdges holds the relations/edges for other nodes in the graph.
type TableChangeAlertEdges struct {
	// BusinessUnit holds the value of the business_unit edge.
	BusinessUnit *BusinessUnit `json:"business_unit,omitempty"`
	// Organization holds the value of the organization edge.
	Organization *Organization `json:"organization,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// BusinessUnitOrErr returns the BusinessUnit value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TableChangeAlertEdges) BusinessUnitOrErr() (*BusinessUnit, error) {
	if e.BusinessUnit != nil {
		return e.BusinessUnit, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: businessunit.Label}
	}
	return nil, &NotLoadedError{edge: "business_unit"}
}

// OrganizationOrErr returns the Organization value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TableChangeAlertEdges) OrganizationOrErr() (*Organization, error) {
	if e.Organization != nil {
		return e.Organization, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: organization.Label}
	}
	return nil, &NotLoadedError{edge: "organization"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*TableChangeAlert) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case tablechangealert.FieldStatus, tablechangealert.FieldName, tablechangealert.FieldDatabaseAction, tablechangealert.FieldSource, tablechangealert.FieldTableName, tablechangealert.FieldTopic, tablechangealert.FieldDescription, tablechangealert.FieldCustomSubject, tablechangealert.FieldFunctionName, tablechangealert.FieldTriggerName, tablechangealert.FieldListenerName, tablechangealert.FieldEmailRecipients:
			values[i] = new(sql.NullString)
		case tablechangealert.FieldCreatedAt, tablechangealert.FieldUpdatedAt, tablechangealert.FieldEffectiveDate, tablechangealert.FieldExpirationDate:
			values[i] = new(sql.NullTime)
		case tablechangealert.FieldID, tablechangealert.FieldBusinessUnitID, tablechangealert.FieldOrganizationID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the TableChangeAlert fields.
func (tca *TableChangeAlert) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case tablechangealert.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				tca.ID = *value
			}
		case tablechangealert.FieldBusinessUnitID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field business_unit_id", values[i])
			} else if value != nil {
				tca.BusinessUnitID = *value
			}
		case tablechangealert.FieldOrganizationID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field organization_id", values[i])
			} else if value != nil {
				tca.OrganizationID = *value
			}
		case tablechangealert.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				tca.CreatedAt = value.Time
			}
		case tablechangealert.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				tca.UpdatedAt = value.Time
			}
		case tablechangealert.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				tca.Status = tablechangealert.Status(value.String)
			}
		case tablechangealert.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				tca.Name = value.String
			}
		case tablechangealert.FieldDatabaseAction:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field database_action", values[i])
			} else if value.Valid {
				tca.DatabaseAction = tablechangealert.DatabaseAction(value.String)
			}
		case tablechangealert.FieldSource:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field source", values[i])
			} else if value.Valid {
				tca.Source = tablechangealert.Source(value.String)
			}
		case tablechangealert.FieldTableName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field table_name", values[i])
			} else if value.Valid {
				tca.TableName = new(string)
				*tca.TableName = value.String
			}
		case tablechangealert.FieldTopic:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field topic", values[i])
			} else if value.Valid {
				tca.Topic = new(string)
				*tca.Topic = value.String
			}
		case tablechangealert.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				tca.Description = new(string)
				*tca.Description = value.String
			}
		case tablechangealert.FieldCustomSubject:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field custom_subject", values[i])
			} else if value.Valid {
				tca.CustomSubject = new(string)
				*tca.CustomSubject = value.String
			}
		case tablechangealert.FieldFunctionName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field function_name", values[i])
			} else if value.Valid {
				tca.FunctionName = new(string)
				*tca.FunctionName = value.String
			}
		case tablechangealert.FieldTriggerName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field trigger_name", values[i])
			} else if value.Valid {
				tca.TriggerName = new(string)
				*tca.TriggerName = value.String
			}
		case tablechangealert.FieldListenerName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field listener_name", values[i])
			} else if value.Valid {
				tca.ListenerName = new(string)
				*tca.ListenerName = value.String
			}
		case tablechangealert.FieldEmailRecipients:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email_recipients", values[i])
			} else if value.Valid {
				tca.EmailRecipients = new(string)
				*tca.EmailRecipients = value.String
			}
		case tablechangealert.FieldEffectiveDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field effective_date", values[i])
			} else if value.Valid {
				tca.EffectiveDate = new(time.Time)
				*tca.EffectiveDate = value.Time
			}
		case tablechangealert.FieldExpirationDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field expiration_date", values[i])
			} else if value.Valid {
				tca.ExpirationDate = new(time.Time)
				*tca.ExpirationDate = value.Time
			}
		default:
			tca.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the TableChangeAlert.
// This includes values selected through modifiers, order, etc.
func (tca *TableChangeAlert) Value(name string) (ent.Value, error) {
	return tca.selectValues.Get(name)
}

// QueryBusinessUnit queries the "business_unit" edge of the TableChangeAlert entity.
func (tca *TableChangeAlert) QueryBusinessUnit() *BusinessUnitQuery {
	return NewTableChangeAlertClient(tca.config).QueryBusinessUnit(tca)
}

// QueryOrganization queries the "organization" edge of the TableChangeAlert entity.
func (tca *TableChangeAlert) QueryOrganization() *OrganizationQuery {
	return NewTableChangeAlertClient(tca.config).QueryOrganization(tca)
}

// Update returns a builder for updating this TableChangeAlert.
// Note that you need to call TableChangeAlert.Unwrap() before calling this method if this TableChangeAlert
// was returned from a transaction, and the transaction was committed or rolled back.
func (tca *TableChangeAlert) Update() *TableChangeAlertUpdateOne {
	return NewTableChangeAlertClient(tca.config).UpdateOne(tca)
}

// Unwrap unwraps the TableChangeAlert entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (tca *TableChangeAlert) Unwrap() *TableChangeAlert {
	_tx, ok := tca.config.driver.(*txDriver)
	if !ok {
		panic("ent: TableChangeAlert is not a transactional entity")
	}
	tca.config.driver = _tx.drv
	return tca
}

// String implements the fmt.Stringer.
func (tca *TableChangeAlert) String() string {
	var builder strings.Builder
	builder.WriteString("TableChangeAlert(")
	builder.WriteString(fmt.Sprintf("id=%v, ", tca.ID))
	builder.WriteString("business_unit_id=")
	builder.WriteString(fmt.Sprintf("%v", tca.BusinessUnitID))
	builder.WriteString(", ")
	builder.WriteString("organization_id=")
	builder.WriteString(fmt.Sprintf("%v", tca.OrganizationID))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(tca.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(tca.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", tca.Status))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(tca.Name)
	builder.WriteString(", ")
	builder.WriteString("database_action=")
	builder.WriteString(fmt.Sprintf("%v", tca.DatabaseAction))
	builder.WriteString(", ")
	builder.WriteString("source=")
	builder.WriteString(fmt.Sprintf("%v", tca.Source))
	builder.WriteString(", ")
	if v := tca.TableName; v != nil {
		builder.WriteString("table_name=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := tca.Topic; v != nil {
		builder.WriteString("topic=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := tca.Description; v != nil {
		builder.WriteString("description=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := tca.CustomSubject; v != nil {
		builder.WriteString("custom_subject=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := tca.FunctionName; v != nil {
		builder.WriteString("function_name=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := tca.TriggerName; v != nil {
		builder.WriteString("trigger_name=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := tca.ListenerName; v != nil {
		builder.WriteString("listener_name=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := tca.EmailRecipients; v != nil {
		builder.WriteString("email_recipients=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := tca.EffectiveDate; v != nil {
		builder.WriteString("effective_date=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := tca.ExpirationDate; v != nil {
		builder.WriteString("expiration_date=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteByte(')')
	return builder.String()
}

// TableChangeAlerts is a parsable slice of TableChangeAlert.
type TableChangeAlerts []*TableChangeAlert
