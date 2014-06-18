package deep

// Applicable ...
type Applicable interface {
	Apply(interface{}) interface{}
}

// ApplicableFunc ...
type ApplicableFunc func(interface{}) interface{}

func (f ApplicableFunc) Apply(to interface{}) interface{} {
	return f(to)
}

// Applicables ...
type Applicables []Applicable

func (items Applicables) Append(item Applicable) Applicables {
	if e2, ok := item.(Each); ok {
		last := len(items) - 1
		if e1, ok := items[last].(Each); ok {
			items[last] = e1.Append(e2)
			return items
		}
	}
	return append(items, item)
}

func (applicables Applicables) Apply(to interface{}) interface{} {
	for _, applicable := range applicables {
		to = applicable.Apply(to)
	}
	return to
}

func appendApplicable(to, item Applicable) Applicable {
	if to == nil {
		return item
	}
	switch item.(type) {
	case Delete, Set:
		return item
	}
	switch v1 := to.(type) {
	case Applicables:
		return v1.Append(item)
	case Each:
		if v2, ok := item.(Each); ok {
			return v1.Append(v2)
		}
	}
	return Applicables{to, item}
}
