// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AuthHistoriesColumns holds the columns for the "auth_histories" table.
	AuthHistoriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "user_id", Type: field.TypeUUID, Nullable: true},
		{Name: "resource", Type: field.TypeString},
		{Name: "method", Type: field.TypeString},
		{Name: "create_at", Type: field.TypeUint32},
	}
	// AuthHistoriesTable holds the schema information for the "auth_histories" table.
	AuthHistoriesTable = &schema.Table{
		Name:       "auth_histories",
		Columns:    AuthHistoriesColumns,
		PrimaryKey: []*schema.Column{AuthHistoriesColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AuthHistoriesTable,
	}
)

func init() {
}
