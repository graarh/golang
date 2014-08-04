package condition

import (
	"github.com/graarh/golang/expert/data"
	"testing"
)

var intervalParams = data.CreateParameters(map[string]data.Parameter{
	"Time": int64(10),
	"Type": uint(0),
})

var Any = Point{0, false}

func getPoint(value int64) Point {
	if value == 0 {
		return Any
	}
	return Point{value, true}
}

func createInterval(from, to int64) Interval {
	return Interval{getPoint(from), getPoint(to), TimeValue}
}

func TestAlwaysPass(t *testing.T) {
	interval := createInterval(0, 0)
	if interval.Check(intervalParams) != true {
		t.Error("0,0 interval should always pass")
	}
}

func TestLeftPass(t *testing.T) {
	interval := createInterval(0, 10)
	if interval.Check(intervalParams) != true {
		t.Error("0, 10 interval should pass 10 time")
	}

	interval = createInterval(0, 9)
	if interval.Check(intervalParams) != false {
		t.Error("0, 9 interval should NOT pass 10 time")
	}
}

func TestRightPass(t *testing.T) {
	interval := createInterval(10, 0)
	if interval.Check(intervalParams) != true {
		t.Error("10, 0 interval should pass 10 time")
	}

	interval = createInterval(11, 0)
	if interval.Check(intervalParams) != false {
		t.Error("11, 0 interval should NOT pass 10 time")
	}
}

func TestBothPass(t *testing.T) {
	interval := createInterval(9, 11)
	if interval.Check(intervalParams) != true {
		t.Error("9, 11 interval should pass 10 time")
	}

	interval = createInterval(11, 12)
	if interval.Check(intervalParams) != false {
		t.Error("11, 12 interval should NOT pass 10 time")
	}

	interval = createInterval(8, 9)
	if interval.Check(intervalParams) != false {
		t.Error("8, 9 interval should NOT pass 10 time")
	}
}

func BenchmarkInterval(b *testing.B) {
	interval := createInterval(9, 11)
	for i := 0; i < b.N; i++ {
		interval.Check(intervalParams)
	}
}
