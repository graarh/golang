package modifier

import (
	"github.com/graarh/golang/expert/data"
	"testing"
)

func TestIwModifier(t *testing.T) {
	f := func(value int) int {
		return value + 5
	}

	iw := data.IntWeight{7}
	iwm := IwModifier{f}
	params := data.CreateParameters(map[string]data.Parameter{"Type": uint(0), "Time": int64(0)})

	result := iwm.Modify(&iw, params)
	iwmResult := result.(*data.IntWeight)

	if iwmResult.Value != 12 {
		t.Error("Int value modifier error, result should be 12, but given ", iwmResult.Value)
	}
}
