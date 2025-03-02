package valigen

import "github.com/pashifika/valigen/internal/tags"

// StructField represents a struct field with its name and associated tag value.
type StructField struct {
	// Name represents the name of the struct field.
	Name string

	// Type represents the types.ExprString of the struct field.
	Type string

	// Value represents the parsed tag associated with the struct field, including its key, name, and any options.
	Value tags.Tag
}
