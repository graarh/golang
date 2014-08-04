package expert

import "github.com/graarh/golang/expert/data"

// Condition is the interface that defines one single check
// using input parameters
type Condition interface {
	Check(params data.Parameters) bool
}
