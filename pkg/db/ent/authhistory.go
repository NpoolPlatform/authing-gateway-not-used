// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/authing-gateway/pkg/db/ent/authhistory"
	"github.com/google/uuid"
)

// AuthHistory is the model entity for the AuthHistory schema.
type AuthHistory struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// AppID holds the value of the "app_id" field.
	AppID uuid.UUID `json:"app_id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID uuid.UUID `json:"user_id,omitempty"`
	// Resource holds the value of the "resource" field.
	Resource string `json:"resource,omitempty"`
	// Method holds the value of the "method" field.
	Method string `json:"method,omitempty"`
	// Allowed holds the value of the "allowed" field.
	Allowed bool `json:"allowed,omitempty"`
	// CreateAt holds the value of the "create_at" field.
	CreateAt uint32 `json:"create_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AuthHistory) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case authhistory.FieldAllowed:
			values[i] = new(sql.NullBool)
		case authhistory.FieldCreateAt:
			values[i] = new(sql.NullInt64)
		case authhistory.FieldResource, authhistory.FieldMethod:
			values[i] = new(sql.NullString)
		case authhistory.FieldID, authhistory.FieldAppID, authhistory.FieldUserID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type AuthHistory", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AuthHistory fields.
func (ah *AuthHistory) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case authhistory.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				ah.ID = *value
			}
		case authhistory.FieldAppID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field app_id", values[i])
			} else if value != nil {
				ah.AppID = *value
			}
		case authhistory.FieldUserID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value != nil {
				ah.UserID = *value
			}
		case authhistory.FieldResource:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field resource", values[i])
			} else if value.Valid {
				ah.Resource = value.String
			}
		case authhistory.FieldMethod:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field method", values[i])
			} else if value.Valid {
				ah.Method = value.String
			}
		case authhistory.FieldAllowed:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field allowed", values[i])
			} else if value.Valid {
				ah.Allowed = value.Bool
			}
		case authhistory.FieldCreateAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field create_at", values[i])
			} else if value.Valid {
				ah.CreateAt = uint32(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this AuthHistory.
// Note that you need to call AuthHistory.Unwrap() before calling this method if this AuthHistory
// was returned from a transaction, and the transaction was committed or rolled back.
func (ah *AuthHistory) Update() *AuthHistoryUpdateOne {
	return (&AuthHistoryClient{config: ah.config}).UpdateOne(ah)
}

// Unwrap unwraps the AuthHistory entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ah *AuthHistory) Unwrap() *AuthHistory {
	tx, ok := ah.config.driver.(*txDriver)
	if !ok {
		panic("ent: AuthHistory is not a transactional entity")
	}
	ah.config.driver = tx.drv
	return ah
}

// String implements the fmt.Stringer.
func (ah *AuthHistory) String() string {
	var builder strings.Builder
	builder.WriteString("AuthHistory(")
	builder.WriteString(fmt.Sprintf("id=%v", ah.ID))
	builder.WriteString(", app_id=")
	builder.WriteString(fmt.Sprintf("%v", ah.AppID))
	builder.WriteString(", user_id=")
	builder.WriteString(fmt.Sprintf("%v", ah.UserID))
	builder.WriteString(", resource=")
	builder.WriteString(ah.Resource)
	builder.WriteString(", method=")
	builder.WriteString(ah.Method)
	builder.WriteString(", allowed=")
	builder.WriteString(fmt.Sprintf("%v", ah.Allowed))
	builder.WriteString(", create_at=")
	builder.WriteString(fmt.Sprintf("%v", ah.CreateAt))
	builder.WriteByte(')')
	return builder.String()
}

// AuthHistories is a parsable slice of AuthHistory.
type AuthHistories []*AuthHistory

func (ah AuthHistories) config(cfg config) {
	for _i := range ah {
		ah[_i].config = cfg
	}
}
