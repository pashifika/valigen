package valigen

import (
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"

	"github.com/pashifika/valigen/internal/tags"
)

// Phase represents an internal processing unit that includes a token.FileSet for managing Go source files.
type Phase struct {
	fset *token.FileSet
}

// NewPhase creates and returns a pointer to a new Phase instance with an initialized token.FileSet.
func NewPhase(fset *token.FileSet) *Phase {
	return &Phase{fset: fset}
}

// ParseStructFields extracts struct fields with a specific tag from a Go source file and maps them to their struct names.
func (p *Phase) ParseStructFields(src, tag string) (map[string][]StructField, error) {
	fAst, err := parser.ParseFile(p.fset, src, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	res := make(map[string][]StructField)
	for _, decl := range fAst.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok {
			continue
		}
		for _, spec := range genDecl.Specs {
			typeSpec, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}
			structType, ok := typeSpec.Type.(*ast.StructType)
			if !ok {
				continue
			}

			// find target fields
			var fields []StructField
			for _, field := range structType.Fields.List {
				if field.Tag == nil {
					continue
				}
				fieldTags, err := tags.ParseTag(field.Tag.Value)
				if err != nil {
					return nil, err
				}
				targetTag, ok := fieldTags.Get(tag)
				if !ok {
					continue
				}
				typeString := types.ExprString(field.Type)

				// multiple fields may share the same tag
				for _, fieldName := range field.Names {
					fields = append(fields, StructField{
						Name:  fieldName.Name,
						Type:  typeString,
						Value: *targetTag,
					})
				}
			}
			if len(fields) > 0 {
				res[typeSpec.Name.Name] = fields
			}
		}
	}

	return res, nil
}
