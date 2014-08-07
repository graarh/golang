package data

import "testing"

func TestAddParamsWeight(t *testing.T) {
	pw1 := &ParamsWeight{CreateParameters(map[string]Parameter{
		"Time":int64(1),
		"Type":uint(2),
	}), nil}
	pw2 := &ParamsWeight{CreateParameters(map[string]Parameter{
		"Fld1":1,
		"Fld2":2,
	}), nil}
	pw3 := &ParamsWeight{CreateParameters(map[string]Parameter{
		"Fld3":3,
	}), nil}

	pw1.Add(pw2)

	result, error := pw1.Params.Get("Fld2")
	if error != nil {
		t.Error("Fld2 param should be added to initial parameters")
	}
	if result != 2 {
		t.Error("Fld2 value is 2, but given ", result)
	}

	pw1.Add(pw3)

	result, error = pw1.Params.Get("Fld3")
	if error != nil {
		t.Error("Fld3 param should be added to initial parameters")
	}
	if result != 3 {
		t.Error("Fld3 value is 3, but given ", result)
	}
}

func BenchmarkParametersAdd(b *testing.B) {
	pw2 := &ParamsWeight{CreateParameters(map[string]Parameter{
		"Fld1":1,
		"Fld2":2,
	}), nil}
	for i := 0; i < b.N; i++ {
		pw1 := &ParamsWeight{CreateParameters(map[string]Parameter{
			"Time":int64(1),
			"Type":uint(2),
		}), nil}
		pw1.Add(pw2)
	}
}
