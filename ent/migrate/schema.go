// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// TokensColumns holds the columns for the "tokens" table.
	TokensColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "value", Type: field.TypeString, Unique: true},
	}
	// TokensTable holds the schema information for the "tokens" table.
	TokensTable = &schema.Table{
		Name:       "tokens",
		Columns:    TokensColumns,
		PrimaryKey: []*schema.Column{TokensColumns[0]},
	}
	// UrlsColumns holds the columns for the "urls" table.
	UrlsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "url", Type: field.TypeString, Unique: true},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"success", "error", "pending"}},
		{Name: "image_url", Type: field.TypeString},
	}
	// UrlsTable holds the schema information for the "urls" table.
	UrlsTable = &schema.Table{
		Name:       "urls",
		Columns:    UrlsColumns,
		PrimaryKey: []*schema.Column{UrlsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		TokensTable,
		UrlsTable,
	}
)

func init() {
}
