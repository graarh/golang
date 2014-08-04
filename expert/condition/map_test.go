package condition

import (
	"github.com/graarh/golang/expert/data"
	"strconv"
	"testing"
)

var testmap = Map{
	Name: "testmap",
	Values: map[data.Parameter]struct{}{
		"one":   struct{}{},
		"two":   struct{}{},
		"three": struct{}{},
	},
}

var paramsMap = map[string]data.Parameter{
	"aaa":  "aaa",
	"bbb":  "bbb",
	"Time": int64(1),
	"Type": uint(0),
}

func TestKeyAbsent(t *testing.T) {
	nokey := testmap.Check(data.CreateParameters(paramsMap))
	if nokey {
		t.Error("Key 'testmap' is not present in params, check should return false")
	}
}

func TestKeyNotInMap(t *testing.T) {
	paramsMap["testmap"] = "not in map"
	keynotin := testmap.Check(data.CreateParameters(paramsMap))
	if keynotin {
		t.Error("Key 'testmap' has value, that is not exists in map, check should return false")
	}
}

func TestKeyInMap(t *testing.T) {
	paramsMap["testmap"] = "two"
	keyin := testmap.Check(data.CreateParameters(paramsMap))
	if !keyin {
		t.Error("Key 'testmap' has value, that present in map, check should return true")
	}
}

func BenchmarkMap(b *testing.B) {
	paramsMap["testmap"] = "two"
	for i := int64(0); i < 500; i++ {
		paramsMap[strconv.FormatInt(i, 10)+"qqq"] = strconv.FormatInt(i, 16)
		testmap.Values[strconv.FormatInt(i, 10)+"qwe"] = struct{}{}
	}
	params := data.CreateParameters(paramsMap)

	for i := 0; i < b.N; i++ {
		testmap.Check(params)
	}
}
