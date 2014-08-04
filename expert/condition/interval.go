package condition

import (
	"github.com/graarh/golang/expert/data"
)

// Interval condition is the condition to check that value
// is in the given interval, zero values is reserved and mean
// "no limit for this side of interval"
// One of the usages is to check the time
// Avoid of using value functions that converts value from
// string to int using name (aka key)
// it can affect performance significantly
// convert it once and return converted result
type Interval struct {
	From, To Point
	Value    func(data.Parameters) int64
}

// Point structure represents point of interval
// Point can be undefined, so if exist == false
// value means nothing
type Point struct {
	Value int64
	Exist bool
}

//TimeValue is the common interval value function
func TimeValue(params data.Parameters) int64 {
	return params.Time()
}

//TypeValue is the common interval value function
func TypeValue(params data.Parameters) int64 {
	return int64(params.Type())
}

// Check is the interface implementation function
func (i *Interval) Check(params data.Parameters) bool {
	currentTime := i.Value(params)
	if (!i.From.Exist || i.From.Value <= currentTime) &&
		(!i.To.Exist || i.To.Value >= currentTime) {
		return true
	}
	return false
}
