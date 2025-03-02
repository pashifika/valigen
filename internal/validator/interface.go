package validator

import (
	"github.com/dave/jennifer/jen"

	"github.com/pashifika/valigen"
)

type Builder interface {
	JsonUnmarshal(sf valigen.StructField) *jen.Statement
}

//
