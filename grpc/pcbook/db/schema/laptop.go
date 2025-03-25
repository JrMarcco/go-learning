package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// Laptop holds the schema definition for the Laptop entity.
type Laptop struct {
	ent.Schema
}

// Fields of the Laptop.
func (Laptop) Fields() []ent.Field {
	return []ent.Field{
		field.String("uid").NotEmpty().Unique(),
		field.String("brand").NotEmpty(),
		field.String("laptop_name").NotEmpty(),
		field.Float("weight").Default(0.00),
		field.Uint32("price_rmb").Default(0),
		field.Uint32("release_year").Default(0),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Laptop.
func (Laptop) Edges() []ent.Edge {
	return nil
}
