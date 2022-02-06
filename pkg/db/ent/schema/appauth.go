package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/google/uuid"
)

// AppAuth holds the schema definition for the AppAuth entity.
type AppAuth struct {
	ent.Schema
}

// Fields of the AppAuth.
func (AppAuth) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.UUID("app_id", uuid.UUID{}),
		field.String("resource"),
		field.String("method"),
		field.Uint32("create_at").
			DefaultFunc(func() uint32 {
				return uint32(time.Now().Unix())
			}),
		field.Uint32("update_at").
			DefaultFunc(func() uint32 {
				return uint32(time.Now().Unix())
			}).
			UpdateDefault(func() uint32 {
				return uint32(time.Now().Unix())
			}),
		field.Uint32("delete_at").
			DefaultFunc(func() uint32 {
				return 0
			}),
	}
}

// Edges of the AppAuth.
func (AppAuth) Edges() []ent.Edge {
	return nil
}

func (AppAuth) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("app_id", "resource", "method").Unique(),
	}
}
