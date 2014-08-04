package expert

import "github.com/graarh/golang/expert/data"

// Modifier interface modifies weight using initial weight and parameters
type Modifier interface {
	Modify(weight data.Weight, params data.Parameters) data.Weight
}
