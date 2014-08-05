package modifier

import (
	"github.com/graarh/golang/expert/data"
	"testing"
)

func TestSingleModifier(t *testing.T) {
	f := func(value data.Comparable) data.Comparable {
		return value.(int) + 5
	}

	iw := data.SingleWeight{7}
	iwm := SingleModifier{f}
	params := data.CreateParameters(map[string]data.Parameter{"Type": uint(0), "Time": int64(0)})

	result := iwm.Modify(&iw, params)
	iwmResult := result.(*data.SingleWeight)

	if iwmResult.Value != 12 {
		t.Error("Int value modifier error, result should be 12, but given ", iwmResult.Value)
	}
}
