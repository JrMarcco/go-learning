// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// LaptopsColumns holds the columns for the "laptops" table.
	LaptopsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "uid", Type: field.TypeString, Unique: true},
		{Name: "brand", Type: field.TypeString},
		{Name: "laptop_name", Type: field.TypeString},
		{Name: "weight", Type: field.TypeFloat64, Default: 0},
		{Name: "price_rmb", Type: field.TypeUint32, Default: 0},
		{Name: "release_year", Type: field.TypeUint32, Default: 0},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// LaptopsTable holds the schema information for the "laptops" table.
	LaptopsTable = &schema.Table{
		Name:       "laptops",
		Columns:    LaptopsColumns,
		PrimaryKey: []*schema.Column{LaptopsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		LaptopsTable,
	}
)

func init() {
}
