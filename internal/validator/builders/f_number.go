package builders

import (
	. "github.com/dave/jennifer/jen"

	"github.com/pashifika/valigen"
)

// Number represents a numeric data type and provides methods for validator.Builder.
type Number struct{}

// JsonUnmarshal generates a validation statement for UnmarshalJSON into a specified numeric type for a struct field.
func (n *Number) JsonUnmarshal(sf valigen.StructField) *Statement {
	// add global variable
	addGlobalVariable(regexNumber)

	res := &Statement{}
	switch sf.Type {
	case "int":
		res.Add(
			List(Id("val"), Id("err")).Qual("strconv", "Atoi").Call(Id("value")),
		)
	}

	return res
}
