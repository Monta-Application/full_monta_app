// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/emoss08/trenova/ent/businessunit"
	"github.com/emoss08/trenova/ent/hazardousmaterialsegregation"
	"github.com/emoss08/trenova/ent/organization"
	"github.com/google/uuid"
)

// HazardousMaterialSegregation is the model entity for the HazardousMaterialSegregation schema.
type HazardousMaterialSegregation struct {
	config `json:"-" validate:"-"`
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
	// Version holds the value of the "version" field.
	Version int `json:"version" validate:"omitempty"`
	// ClassA holds the value of the "class_a" field.
	ClassA hazardousmaterialsegregation.ClassA `json:"classA" validate:"required"`
	// ClassB holds the value of the "class_b" field.
	ClassB hazardousmaterialsegregation.ClassB `json:"classB" validate:"required"`
	// SegregationType holds the value of the "segregation_type" field.
	SegregationType hazardousmaterialsegregation.SegregationType `json:"segregationType" validate:"required"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the HazardousMaterialSegregationQuery when eager-loading is set.
	Edges        HazardousMaterialSegregationEdges `json:"edges"`
	selectValues sql.SelectValues
}

// HazardousMaterialSegregationEdges holds the relations/edges for other nodes in the graph.
type HazardousMaterialSegregationEdges struct {
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
func (e HazardousMaterialSegregationEdges) BusinessUnitOrErr() (*BusinessUnit, error) {
	if e.BusinessUnit != nil {
		return e.BusinessUnit, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: businessunit.Label}
	}
	return nil, &NotLoadedError{edge: "business_unit"}
}

// OrganizationOrErr returns the Organization value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e HazardousMaterialSegregationEdges) OrganizationOrErr() (*Organization, error) {
	if e.Organization != nil {
		return e.Organization, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: organization.Label}
	}
	return nil, &NotLoadedError{edge: "organization"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*HazardousMaterialSegregation) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case hazardousmaterialsegregation.FieldVersion:
			values[i] = new(sql.NullInt64)
		case hazardousmaterialsegregation.FieldClassA, hazardousmaterialsegregation.FieldClassB, hazardousmaterialsegregation.FieldSegregationType:
			values[i] = new(sql.NullString)
		case hazardousmaterialsegregation.FieldCreatedAt, hazardousmaterialsegregation.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case hazardousmaterialsegregation.FieldID, hazardousmaterialsegregation.FieldBusinessUnitID, hazardousmaterialsegregation.FieldOrganizationID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the HazardousMaterialSegregation fields.
func (hms *HazardousMaterialSegregation) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case hazardousmaterialsegregation.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				hms.ID = *value
			}
		case hazardousmaterialsegregation.FieldBusinessUnitID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field business_unit_id", values[i])
			} else if value != nil {
				hms.BusinessUnitID = *value
			}
		case hazardousmaterialsegregation.FieldOrganizationID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field organization_id", values[i])
			} else if value != nil {
				hms.OrganizationID = *value
			}
		case hazardousmaterialsegregation.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				hms.CreatedAt = value.Time
			}
		case hazardousmaterialsegregation.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				hms.UpdatedAt = value.Time
			}
		case hazardousmaterialsegregation.FieldVersion:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field version", values[i])
			} else if value.Valid {
				hms.Version = int(value.Int64)
			}
		case hazardousmaterialsegregation.FieldClassA:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field class_a", values[i])
			} else if value.Valid {
				hms.ClassA = hazardousmaterialsegregation.ClassA(value.String)
			}
		case hazardousmaterialsegregation.FieldClassB:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field class_b", values[i])
			} else if value.Valid {
				hms.ClassB = hazardousmaterialsegregation.ClassB(value.String)
			}
		case hazardousmaterialsegregation.FieldSegregationType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field segregation_type", values[i])
			} else if value.Valid {
				hms.SegregationType = hazardousmaterialsegregation.SegregationType(value.String)
			}
		default:
			hms.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the HazardousMaterialSegregation.
// This includes values selected through modifiers, order, etc.
func (hms *HazardousMaterialSegregation) Value(name string) (ent.Value, error) {
	return hms.selectValues.Get(name)
}

// QueryBusinessUnit queries the "business_unit" edge of the HazardousMaterialSegregation entity.
func (hms *HazardousMaterialSegregation) QueryBusinessUnit() *BusinessUnitQuery {
	return NewHazardousMaterialSegregationClient(hms.config).QueryBusinessUnit(hms)
}

// QueryOrganization queries the "organization" edge of the HazardousMaterialSegregation entity.
func (hms *HazardousMaterialSegregation) QueryOrganization() *OrganizationQuery {
	return NewHazardousMaterialSegregationClient(hms.config).QueryOrganization(hms)
}

// Update returns a builder for updating this HazardousMaterialSegregation.
// Note that you need to call HazardousMaterialSegregation.Unwrap() before calling this method if this HazardousMaterialSegregation
// was returned from a transaction, and the transaction was committed or rolled back.
func (hms *HazardousMaterialSegregation) Update() *HazardousMaterialSegregationUpdateOne {
	return NewHazardousMaterialSegregationClient(hms.config).UpdateOne(hms)
}

// Unwrap unwraps the HazardousMaterialSegregation entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (hms *HazardousMaterialSegregation) Unwrap() *HazardousMaterialSegregation {
	_tx, ok := hms.config.driver.(*txDriver)
	if !ok {
		panic("ent: HazardousMaterialSegregation is not a transactional entity")
	}
	hms.config.driver = _tx.drv
	return hms
}

// String implements the fmt.Stringer.
func (hms *HazardousMaterialSegregation) String() string {
	var builder strings.Builder
	builder.WriteString("HazardousMaterialSegregation(")
	builder.WriteString(fmt.Sprintf("id=%v, ", hms.ID))
	builder.WriteString("business_unit_id=")
	builder.WriteString(fmt.Sprintf("%v", hms.BusinessUnitID))
	builder.WriteString(", ")
	builder.WriteString("organization_id=")
	builder.WriteString(fmt.Sprintf("%v", hms.OrganizationID))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(hms.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(hms.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("version=")
	builder.WriteString(fmt.Sprintf("%v", hms.Version))
	builder.WriteString(", ")
	builder.WriteString("class_a=")
	builder.WriteString(fmt.Sprintf("%v", hms.ClassA))
	builder.WriteString(", ")
	builder.WriteString("class_b=")
	builder.WriteString(fmt.Sprintf("%v", hms.ClassB))
	builder.WriteString(", ")
	builder.WriteString("segregation_type=")
	builder.WriteString(fmt.Sprintf("%v", hms.SegregationType))
	builder.WriteByte(')')
	return builder.String()
}

// HazardousMaterialSegregations is a parsable slice of HazardousMaterialSegregation.
type HazardousMaterialSegregations []*HazardousMaterialSegregation
