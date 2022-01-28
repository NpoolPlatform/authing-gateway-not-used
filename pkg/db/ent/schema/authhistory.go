package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"

	"github.com/google/uuid"
)

// AuthHistory holds the schema definition for the AuthHistory entity.
type AuthHistory struct {
	ent.Schema
}

// Fields of the AuthHistory.
func (AuthHistory) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.UUID("app_id", uuid.UUID{}),
		field.UUID("user_id", uuid.UUID{}).Optional(),
		field.String("resource"),
		field.String("method"),
		field.Uint32("create_at").
			DefaultFunc(func() uint32 {
				return uint32(time.Now().Unix())
			}),
	}
}

// Edges of the AuthHistory.
func (AuthHistory) Edges() []ent.Edge {
	return nil
}
