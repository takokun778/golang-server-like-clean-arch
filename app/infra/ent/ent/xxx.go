// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"
	"xxx/app/infra/ent/ent/xxx"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// Xxx is the model entity for the Xxx schema.
type Xxx struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Number holds the value of the "number" field.
	Number int `json:"number,omitempty"`
	// CreatedAt holds the value of the "createdAt" field.
	CreatedAt time.Time `json:"createdAt,omitempty"`
	// UpdatedAt holds the value of the "updatedAt" field.
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Xxx) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case xxx.FieldNumber:
			values[i] = new(sql.NullInt64)
		case xxx.FieldName:
			values[i] = new(sql.NullString)
		case xxx.FieldCreatedAt, xxx.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case xxx.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Xxx", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Xxx fields.
func (x *Xxx) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case xxx.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				x.ID = *value
			}
		case xxx.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				x.Name = value.String
			}
		case xxx.FieldNumber:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field number", values[i])
			} else if value.Valid {
				x.Number = int(value.Int64)
			}
		case xxx.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field createdAt", values[i])
			} else if value.Valid {
				x.CreatedAt = value.Time
			}
		case xxx.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updatedAt", values[i])
			} else if value.Valid {
				x.UpdatedAt = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Xxx.
// Note that you need to call Xxx.Unwrap() before calling this method if this Xxx
// was returned from a transaction, and the transaction was committed or rolled back.
func (x *Xxx) Update() *XxxUpdateOne {
	return (&XxxClient{config: x.config}).UpdateOne(x)
}

// Unwrap unwraps the Xxx entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (x *Xxx) Unwrap() *Xxx {
	tx, ok := x.config.driver.(*txDriver)
	if !ok {
		panic("ent: Xxx is not a transactional entity")
	}
	x.config.driver = tx.drv
	return x
}

// String implements the fmt.Stringer.
func (x *Xxx) String() string {
	var builder strings.Builder
	builder.WriteString("Xxx(")
	builder.WriteString(fmt.Sprintf("id=%v", x.ID))
	builder.WriteString(", name=")
	builder.WriteString(x.Name)
	builder.WriteString(", number=")
	builder.WriteString(fmt.Sprintf("%v", x.Number))
	builder.WriteString(", createdAt=")
	builder.WriteString(x.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updatedAt=")
	builder.WriteString(x.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Xxxes is a parsable slice of Xxx.
type Xxxes []*Xxx

func (x Xxxes) config(cfg config) {
	for _i := range x {
		x[_i].config = cfg
	}
}