package builders

import (
	. "github.com/dave/jennifer/jen"
)

var (
	regexAlpha = namedCode{
		Name: "regexAlpha",
		Code: Id("regexAlpha").Op(":=").
			Qual("regexp", "MustCompile").Call(Lit(`^[a-zA-Z]+$`)),
	}
	regexAlphaNum = namedCode{
		Name: "regexAlphaNum",
		Code: Id("regexAlphaNum").Op(":=").
			Qual("regexp", "MustCompile").Call(Lit(`^[a-zA-Z0-9]+$`)),
	}
	regexNumber = namedCode{
		Name: "regexNumber",
		Code: Id("regexNumber").Op(":=").
			Qual("regexp", "MustCompile").Call(Lit(`^[0-9]+$`)),
	}
	regexNumeric namedCode = namedCode{
		Name: "regexNumeric",
		Code: Id("regexNumeric").Op(":=").
			Qual("regexp", "MustCompile").Call(Lit(`^[-+]?[0-9]+(?:\\.[0-9]+)?$`)),
	}
)
