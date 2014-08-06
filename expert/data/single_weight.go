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
func (sw1 *SingleWeight) Less(w Weight) bool {
	sw2 := w.(*SingleWeight)

	if sw1.Value == sw2.Value {
		return false
	}

	switch sw1.Value.(type) {
	case string:
		if sw1.Value.(string) < sw2.Value.(string) {
			return true
		}
	case uint, uint8, uint16, uint32, uint64:
		if reflect.ValueOf(sw1.Value).Uint() < reflect.ValueOf(sw2.Value).Uint() {
			return true
		}
	case int, int8, int16, int32, int64:
		if reflect.ValueOf(sw1.Value).Int() < reflect.ValueOf(sw2.Value).Int() {
			return true
		}
	case float32, float64:
		if reflect.ValueOf(sw1.Value).Float() < reflect.ValueOf(sw2.Value).Float() {
			return true
		}
	default:
		trace.Log("Not comparable type in SingleWeight.Value", sw1.Value)
	}
	return false
}

func (sw1 *SingleWeight) Add(w Weight) {
	sw2 := w.(*SingleWeight)

	switch sw1.Value.(type) {
	case string:
		sw1.Value = sw1.Value.(string) + sw2.Value.(string)
	case uint, uint8, uint16, uint32, uint64:
		sw1.Value = reflect.ValueOf(sw1.Value).Uint() + reflect.ValueOf(sw2.Value).Uint()
	case int, int8, int16, int32, int64:
		sw1.Value = reflect.ValueOf(sw1.Value).Int() + reflect.ValueOf(sw2.Value).Int()
	case float32, float64:
		sw1.Value = reflect.ValueOf(sw1.Value).Float() + reflect.ValueOf(sw2.Value).Float()
	default:
		trace.Log("Not comparable type in SingleWeight.Value", sw1.Value)
	}
}
