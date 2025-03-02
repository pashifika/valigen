package builders

import (
	. "github.com/dave/jennifer/jen"
)

var (
	variables []Code

	variablesCheckDuplicate = make(map[string]struct{})
)

func addGlobalVariable(nc namedCode) {
	if _, ok := variablesCheckDuplicate[nc.Name]; ok {
		return
	}
	variables = append(variables, nc.Code)
}
