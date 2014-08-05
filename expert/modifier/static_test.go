package modifier

import (
	"github.com/graarh/golang/expert/data"
	"testing"
)

func TestStaticModifier(t *testing.T) {
	s := Static{&data.SingleWeight{5}}
	params := data.CreateParameters(map[string]data.Parameter{"Type": uint(0), "Time": int64(0)})
	s1 := data.SingleWeight{6}

	result := s.Modify(&s1, params)

	if result.(*data.SingleWeight).Value != 5 {
		t.Error("Static modifier should return weight, stored in value, it is 5, not ",
			result.(*data.SingleWeight).Value)
	}
}
