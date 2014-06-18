package deep

// Delete ...
type Delete struct{}

func (Delete) Apply(_ interface{}) interface{} {
	return nil
}

// Set ...
type Set struct {
	Value interface{}
}

func (s Set) Apply(_ interface{}) interface{} {
	return s.Value
}
