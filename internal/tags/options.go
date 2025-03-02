package tags

type Options []Option

func (o Options) Has(s string) (*Option, bool) {
	for _, v := range o {
		if v.Name == s {
			return &v, true
		}
	}

	return nil, false
}

type Option struct {
	Name  string
	Value string
}
