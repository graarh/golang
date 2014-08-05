package data

import (
	"reflect"
	"github.com/graarh/golang/trace"
)

//Comparable is alias to any buildin type that can be compared
type Comparable interface{}

//SingleWeight is the common single number weight
type SingleWeight struct {
	Value Comparable
}

//Compare is the Weight interface function implementation
func (sw1 *SingleWeight) Compare(w Weight) int {
	sw2 := w.(*SingleWeight)

	if sw1.Value == sw2.Value {
		return 0
	}

	switch sw1.Value.(type) {
	case string:
		if sw1.Value.(string) < sw2.Value.(string) {
			return -1
		}
		return 1
	case uint, uint8, uint16, uint32, uint64:
		if reflect.ValueOf(sw1.Value).Uint() < reflect.ValueOf(sw2.Value).Uint() {
			return -1
		}
		return 1
	case int, int8, int16, int32, int64:
		if reflect.ValueOf(sw1.Value).Int() < reflect.ValueOf(sw2.Value).Int() {
			return -1
		}
		return 1
	case float32, float64:
		if reflect.ValueOf(sw1.Value).Float() < reflect.ValueOf(sw2.Value).Float() {
			return -1
		}
		return 1
	default:
		trace.Log("Not comparable type in SingleWeight.Value", sw1.Value)
	}
	return 1
}
