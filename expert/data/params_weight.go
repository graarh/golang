package data

import "github.com/graarh/golang/trace"

//ParametersComparator is a function definition that can compare
//two maps of parameters
type ParametersComparator func(p1 Parameters, p2 Parameters) bool

//ParamsWeight is weight implementation for map of parameters
type ParamsWeight struct {
	Params     Parameters
	Comparator ParametersComparator
}

//Less is the Weight interface implementation
func (pw1 *ParamsWeight) Less(w Weight) bool {
	pw2 := w.(*ParamsWeight)
	return pw1.Comparator(pw1.Params, pw2.Params)
}

//Add is the Weight interface implementation
func (pw1 *ParamsWeight) Add(w Weight) {
	pw2 := w.(*ParamsWeight)

	pw1Params := pw1.Params.All()

	for key, value := range pw2.Params.All() {
		if _, ok := pw1Params[key]; ok {
			trace.Log("Key ", key, " already exists into initial parameters set. ", pw1.Params, pw2.Params)
		}
		pw1Params[key] = value
	}

	pw1.Params = CreateParameters(pw1Params)
}
