package validator

func (v *Validator) BuildImports() {
	names := map[string]string{
		"github.com/foo/a": "a",
		"github.com/foo/b": "b",
	}
	v.f.ImportNames(names)
}
