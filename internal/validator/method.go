package validator

import (
	"errors"

	. "github.com/dave/jennifer/jen"

	"github.com/pashifika/valigen"
)

// BuildMethods generates methods for the provided struct data by iterating over each methods
// list and invoking the appropriate builder function.
func (v *Validator) BuildMethods(structData map[string][]valigen.StructField) error {
	for name, structs := range structData {
		for _, method := range v.methods {
			blocks := make([]Code, 0)
			for _, sf := range structs {
				if bi, ok := v.builder[sf.Type]; ok {
					builder, err := getBuilder(bi, method)
					if err != nil {
						return err
					}
					blocks = append(blocks, builder(sf))
				} else {
					return errors.New("unknown builder type")
				}
			}

			fn, err := getFuncName(method)
			if err != nil {
				return err
			}
			v.f.Func().
				Params(Id("v").Op("*").Id(name)).
				Id(fn).Params(Id("data").Index().Byte()).
				Error().
				Block(blocks...)
		}
	}

	return nil
}

func getBuilder(b Builder, m MethodType) (func(sf valigen.StructField) *Statement, error) {
	switch m {
	case MethodTypeJSONUnmarshal:
		return b.JsonUnmarshal, nil
	default:
		return nil, errors.New("getBuilder: unknown method type")
	}
}

func getFuncName(m MethodType) (string, error) {
	switch m {
	case MethodTypeJSONUnmarshal:
		return "UnmarshalJSON", nil
	default:
		return "", errors.New("getFuncName: unknown method type")
	}
}
