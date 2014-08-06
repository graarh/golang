package condition

import (
	"github.com/graarh/golang/expert/data"
)

// Map is the condition that searches params value with given
// name (aka key) in the predefined set of values
// and returns true if the value is found in set
type Map struct {
	Name string
	/*
		Map cannot work with any type, perhaps data.Parameter should
		implement something like <Hash() string> function later on,
		so map[string]struct{} can be used
	*/
	Values  map[data.Parameter]struct{}
	Exclude bool //true to pass if value is NOT in Values
}

// SliceToMap transforms Parameters slice into map
func SliceToMap(slice []data.Parameter) map[data.Parameter]struct{} {
	result := make(map[data.Parameter]struct{})
	for _, value := range slice {
		result[value] = struct{}{}
	}
	return result
}

// Check is the Condition interface implementation for Map
func (m *Map) Check(params data.Parameters) bool {
	if value, error := params.Get(m.Name); error == nil {
		if _, check := m.Values[value]; check {
			return !m.Exclude
		}
	}
	return m.Exclude
}
