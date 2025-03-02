package tags

import (
	"errors"
	"strings"
	"text/scanner"
)

// Tags represent a set of tags from a single struct field
type Tags []Tag

// Tag defines a single struct's string literal tag
type Tag struct {
	// Key is the tag key, such as json, xml, etc...
	// i.e: `json:"foo,omitempty"`. Here key is: "json"
	Key string

	// Name is a part of the value
	// i.e: `json:"foo,omitempty"`. Here name is: "foo"
	Name string

	// Options is a part of the value. It contains a slice of tag options
	// i.e: `json:"foo,omitempty"`. Here options is: ["omitempty"]
	Options []string
}

// Get retrieves a tag from the Tags slice with the specified key.
func (t Tags) Get(key string) (*Tag, bool) {
	for _, tag := range t {
		if tag.Key == key {
			return &tag, true
		}
	}

	return nil, false
}

// ParseTag parses a struct field's tag string into a slice of Tag objects.
func ParseTag(s string) (Tags, error) {
	var (
		scan   scanner.Scanner
		result Tags
	)
	scan.Init(strings.NewReader(s))
	step := _stepKey
	tag := Tag{}
	var prevRune rune
	for {
		r := scan.Next()
		switch r {
		case scanner.EOF:
			if prevRune != _quote || step < _stepName {
				return nil, ErrTagSyntax
			}
			goto end
		case _colon:
			if step > _stepName {
				// case: `name:"test,opt1:ab"`
				tag.Options[step] += string(r)
				break
			}
			if tag.Key == "" {
				return nil, ErrTagKeySyntax
			}
			if scan.Peek() == _quote {
				prevRune = r
				scan.Next()
				step++
			}
		case _comma: // case: `key1:"name,opt"`
			if step == _stepKey {
				// case: `n,ame:""`
				return nil, ErrTagSyntax
			}
			if scan.Peek() != _quote {
				tag.Options = append(tag.Options, "")
				step++
			}
		case _quote:
			switch scan.Peek() {
			case scanner.EOF:
				result = append(result, tag)
			case _space: // case: `key1:"name" key2:"name"`
				prevRune = r
				scan.Next()
				step = _stepKey // is next key, initialize step
				result = append(result, tag)
				tag = Tag{}
			default:
				return nil, ErrTagSyntax
			}
		default:
			if r == _bs && scan.Peek() == _comma {
				r = scan.Next()
			}
			switch step {
			case _stepKey:
				tag.Key += string(r)
			case _stepName:
				tag.Name += string(r)
			default:
				tag.Options[step] += string(r)
			}
		}
		prevRune = r
	}

end:
	return result, nil
}

var (
	ErrTagSyntax    = errors.New("bad syntax for struct tag pair")
	ErrTagKeySyntax = errors.New("bad syntax for struct tag key")
)
