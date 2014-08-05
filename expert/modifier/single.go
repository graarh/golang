package modifier

import "github.com/graarh/golang/expert/data"

// SingleModifierFunc is a function definition to modify value
type SingleModifierFunc func(data.Comparable) data.Comparable

// IwModifier is a common int value function modifier
// Mostly, for test purposes
type SingleModifier struct {
	ModifierFunc SingleModifierFunc
}

//Modify is a Modifier interface implementation
func (swm *SingleModifier) Modify(weight data.Weight, params data.Parameters) data.Weight {
	iw := weight.(*data.SingleWeight)
	return &data.SingleWeight{swm.ModifierFunc(iw.Value)}
}
