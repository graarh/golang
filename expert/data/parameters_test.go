package data

import "testing"

func TestCreate(t *testing.T) {
	p := CreateParameters(
		map[string]Parameter{
			"one":  "one",
			"two":  "two",
			"Time": int64(123),
			"Type": uint(456),
		},
	)
	if res, _ := p.Get("one"); res != "one" {
		t.Error("Wrong map created, Get function returns incorrect values")
	}
	if p.Time() != 123 {
		t.Error("Time() was set to 123, but result is ", p.Time())
	}
	if p.Type() != 456 {
		t.Error("Type() was set to 456 but result is", p.Type())
	}
}

func TestGet(t *testing.T) {
	p := parametersRecord{make(map[string]Parameter), 0, 0}
	p.parameters["one"] = "1"

	v, e := p.Get("one")
	if v != "1" {
		t.Error("Get should return value from the Parameters map")
	}
	if e != nil {
		t.Error("Get should not produce error on existing name")
	}

	_, e1 := p.Get("two")

	if e1 == nil {
		t.Error("Get should produce error if no existing name found")
	}
}
