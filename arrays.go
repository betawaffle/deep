package deep

// Append ...
type Append []interface{}

func (items Append) Apply(to interface{}) interface{} {
	var s []interface{}
	if to != nil {
		if v, ok := to.([]interface{}); ok {
			s = v
		}
	}
	return append(s, items...)
}

// Each ...
type Each struct {
	Applicable
}

func (e1 Each) Append(e2 Each) Each {
	e1.Applicable = appendApplicable(e1.Applicable, e2.Applicable)
	return e1
}

func (e Each) Apply(to interface{}) interface{} {
	if to == nil {
		return nil
	}
	if s, ok := to.([]interface{}); ok {
		for i, v := range s {
			s[i] = e.Applicable.Apply(v)
		}
		return s
	}
	// Is this the right thing to do?
	// We should probably panic or something...
	return to
}

// Prepend ...
type Prepend []interface{}

func (items Prepend) Apply(to interface{}) interface{} {
	s := append([]interface{}(nil), items...)
	if to != nil {
		if v, ok := to.([]interface{}); ok {
			return append(s, v...)
		}
	}
	return s
}
