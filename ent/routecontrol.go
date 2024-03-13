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
	"github.com/emoss08/trenova/ent/routecontrol"
	"github.com/google/uuid"
)

// RouteControl is the model entity for the RouteControl schema.
type RouteControl struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"createdAt"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updatedAt"`
	// DistanceMethod holds the value of the "distance_method" field.
	DistanceMethod routecontrol.DistanceMethod `json:"distanceMethod"`
	// MileageUnit holds the value of the "mileage_unit" field.
	MileageUnit routecontrol.MileageUnit `json:"mileageUnit"`
	// GenerateRoutes holds the value of the "generate_routes" field.
	GenerateRoutes bool `json:"generateRoutes"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RouteControlQuery when eager-loading is set.
	Edges            RouteControlEdges `json:"edges"`
	organization_id  *uuid.UUID
	business_unit_id *uuid.UUID
	selectValues     sql.SelectValues
}

// RouteControlEdges holds the relations/edges for other nodes in the graph.
type RouteControlEdges struct {
	// Organization holds the value of the organization edge.
	Organization *Organization `json:"organization,omitempty"`
	// BusinessUnit holds the value of the business_unit edge.
	BusinessUnit *BusinessUnit `json:"business_unit,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int
}

// OrganizationOrErr returns the Organization value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RouteControlEdges) OrganizationOrErr() (*Organization, error) {
	if e.Organization != nil {
		return e.Organization, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: organization.Label}
	}
	return nil, &NotLoadedError{edge: "organization"}
}

// BusinessUnitOrErr returns the BusinessUnit value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RouteControlEdges) BusinessUnitOrErr() (*BusinessUnit, error) {
	if e.BusinessUnit != nil {
		return e.BusinessUnit, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: businessunit.Label}
	}
	return nil, &NotLoadedError{edge: "business_unit"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*RouteControl) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case routecontrol.FieldGenerateRoutes:
			values[i] = new(sql.NullBool)
		case routecontrol.FieldDistanceMethod, routecontrol.FieldMileageUnit:
			values[i] = new(sql.NullString)
		case routecontrol.FieldCreatedAt, routecontrol.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case routecontrol.FieldID:
			values[i] = new(uuid.UUID)
		case routecontrol.ForeignKeys[0]: // organization_id
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case routecontrol.ForeignKeys[1]: // business_unit_id
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the RouteControl fields.
func (rc *RouteControl) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case routecontrol.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				rc.ID = *value
			}
		case routecontrol.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				rc.CreatedAt = value.Time
			}
		case routecontrol.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				rc.UpdatedAt = value.Time
			}
		case routecontrol.FieldDistanceMethod:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field distance_method", values[i])
			} else if value.Valid {
				rc.DistanceMethod = routecontrol.DistanceMethod(value.String)
			}
		case routecontrol.FieldMileageUnit:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mileage_unit", values[i])
			} else if value.Valid {
				rc.MileageUnit = routecontrol.MileageUnit(value.String)
			}
		case routecontrol.FieldGenerateRoutes:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field generate_routes", values[i])
			} else if value.Valid {
				rc.GenerateRoutes = value.Bool
			}
		case routecontrol.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field organization_id", values[i])
			} else if value.Valid {
				rc.organization_id = new(uuid.UUID)
				*rc.organization_id = *value.S.(*uuid.UUID)
			}
		case routecontrol.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field business_unit_id", values[i])
			} else if value.Valid {
				rc.business_unit_id = new(uuid.UUID)
				*rc.business_unit_id = *value.S.(*uuid.UUID)
			}
		default:
			rc.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the RouteControl.
// This includes values selected through modifiers, order, etc.
func (rc *RouteControl) Value(name string) (ent.Value, error) {
	return rc.selectValues.Get(name)
}

// QueryOrganization queries the "organization" edge of the RouteControl entity.
func (rc *RouteControl) QueryOrganization() *OrganizationQuery {
	return NewRouteControlClient(rc.config).QueryOrganization(rc)
}

// QueryBusinessUnit queries the "business_unit" edge of the RouteControl entity.
func (rc *RouteControl) QueryBusinessUnit() *BusinessUnitQuery {
	return NewRouteControlClient(rc.config).QueryBusinessUnit(rc)
}

// Update returns a builder for updating this RouteControl.
// Note that you need to call RouteControl.Unwrap() before calling this method if this RouteControl
// was returned from a transaction, and the transaction was committed or rolled back.
func (rc *RouteControl) Update() *RouteControlUpdateOne {
	return NewRouteControlClient(rc.config).UpdateOne(rc)
}

// Unwrap unwraps the RouteControl entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (rc *RouteControl) Unwrap() *RouteControl {
	_tx, ok := rc.config.driver.(*txDriver)
	if !ok {
		panic("ent: RouteControl is not a transactional entity")
	}
	rc.config.driver = _tx.drv
	return rc
}

// String implements the fmt.Stringer.
func (rc *RouteControl) String() string {
	var builder strings.Builder
	builder.WriteString("RouteControl(")
	builder.WriteString(fmt.Sprintf("id=%v, ", rc.ID))
	builder.WriteString("created_at=")
	builder.WriteString(rc.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(rc.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("distance_method=")
	builder.WriteString(fmt.Sprintf("%v", rc.DistanceMethod))
	builder.WriteString(", ")
	builder.WriteString("mileage_unit=")
	builder.WriteString(fmt.Sprintf("%v", rc.MileageUnit))
	builder.WriteString(", ")
	builder.WriteString("generate_routes=")
	builder.WriteString(fmt.Sprintf("%v", rc.GenerateRoutes))
	builder.WriteByte(')')
	return builder.String()
}

// RouteControls is a parsable slice of RouteControl.
type RouteControls []*RouteControl
