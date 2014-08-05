package modifier

import "github.com/graarh/golang/expert/data"

// Static modifier returns predefined value
type Static struct {
	Value data.Weight
}

//Modify is a Modifier interface implementation
func (s *Static) Modify(data.Weight, data.Parameters) data.Weight {
	return s.Value
}
