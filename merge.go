package deep

import (
	"strings"
)

// Merge ...
type Merge map[string]Applicable

func (m *Merge) Add(path string, value Applicable) {
	m.add(strings.Split(path, "."), value)
}

func (changes Merge) Apply(to interface{}) interface{} {
	o := makeObject(to)
	for k, v := range changes {
		switch change := v.(type) {
		case Delete:
			delete(o, k)
		case Applicable:
			o[k] = change.Apply(o[k])
		default:
			o[k] = v
		}
	}
	return o
}

func (m *Merge) Append(path string, items ...interface{}) {
	m.Add(path, Append(items))
}

func (m *Merge) Delete(path string) {
	m.Add(path, Delete{})
}

func (m *Merge) Each(path string, applicable Applicable) {
	m.Add(path, Each{applicable})
}

func (m *Merge) EachFunc(path string, f func(interface{}) interface{}) {
	m.Add(path, Each{ApplicableFunc(f)})
}

func (m *Merge) Prepend(path string, items ...interface{}) {
	m.Add(path, Prepend(items))
}

func (m *Merge) Set(path string, value interface{}) {
	m.Add(path, Set{value})
}

func (mPtr *Merge) add(path []string, newValue Applicable) {
	m := *mPtr // I want this to fail if m is a nil pointer.
	if m == nil {
		m = make(Merge)
	}
	k := path[0] // I want this to fail if path is empty.
	if len(path) > 1 {
		merge := makeMerge(m[k])
		merge.add(path[1:], newValue)
		m[k] = merge
	} else {
		m[k] = appendApplicable(m[k], newValue)
	}
	*mPtr = m
}

func makeMerge(i Applicable) Merge {
	if i != nil {
		if v, ok := i.(Merge); ok {
			return v
		}
	}
	return make(Merge)
}

func makeObject(i interface{}) map[string]interface{} {
	if i != nil {
		if v, ok := i.(map[string]interface{}); ok {
			return v
		}
	}
	return make(map[string]interface{})
}
