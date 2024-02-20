package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// SampleTable holds the schema definition for the SampleTable entity.
type SampleTable struct {
	ent.Schema
}

func (SampleTable) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "sample_table"},
	}
}

// Fields of the SampleTable.
func (SampleTable) Fields() []ent.Field {
	return []ent.Field{
		field.String("varchar_column").
			NotEmpty().
			MaxLen(255),
		field.Int("integer_column"),
		field.Int64("bigint_column"),
		field.Float("numeric_column"),
		field.Bool("boolean_column"),
		field.String("date_column"),
		field.Time("timestamp_column").
			Default(time.Now),
		field.String("text_column").
			NotEmpty(),
	}
}

// Edges of the SampleTable.
func (SampleTable) Edges() []ent.Edge {
	return nil
}
