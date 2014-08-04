package modifier

import "github.com/graarh/golang/expert/data"

// IwModifierFunc is a function definition to modify value
type IwModifierFunc func(int) int

// IwModifier is a common int value function modifier
// Mostly, for test purposes
type IwModifier struct {
	ModifierFunc IwModifierFunc
}

//Modify is a Modifier interface implementation
func (iwm *IwModifier) Modify(weight data.Weight, params data.Parameters) data.Weight {
	iw := weight.(*data.IntWeight)
	return &data.IntWeight{iwm.ModifierFunc(iw.Value)}
}
