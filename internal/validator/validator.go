package validator

import (
	"bytes"

	"github.com/dave/jennifer/jen"

	"github.com/pashifika/valigen/internal/validator/builders"
)

type Validator struct {
	f       *jen.File
	buf     *bytes.Buffer
	methods []MethodType
	builder map[string]Builder

	pkgName string
}

func NewValidator(pkgName string, mts ...MethodType) *Validator {
	builder := map[string]Builder{
		"number": &builders.Number{},
	}
	if len(mts) == 0 {
		panic("no method type")
	}

	return &Validator{
		f:       jen.NewFile(pkgName),
		buf:     new(bytes.Buffer),
		methods: mts,
		builder: builder,
		pkgName: pkgName,
	}
}

func (v *Validator) WritePackage() error {
	return nil
}

func (v *Validator) WriteImport() error {
	return nil
}
