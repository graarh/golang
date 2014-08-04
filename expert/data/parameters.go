package data

import (
	"fmt"
	"github.com/graarh/golang/trace"
)

//Parameter is a one value of input data
type Parameter interface{}

// Parameters is the interface for
// expert system input parameters
type Parameters interface {
	Get(name string) (Parameter, error) // Get one of the parameters
	Type() uint                         // Get type of the request
	Time() int64                        // Get time of the request
}

// CreateParameters generates Parameters interface
// from the common key-value map
func CreateParameters(data map[string]Parameter) Parameters {
	params := parametersRecord{data, 0, 0}

	if requestType, ok := data["Type"].(uint); ok {
		params.requestType = requestType
	} else {
		trace.LogBacktrace("Parameters 'Type' value type must be 'uint'")
	}

	if requestTime, ok := data["Time"].(int64); ok {
		params.requestTime = int64(requestTime)
	} else {
		trace.LogBacktrace("Parameters 'Time' value type must be 'int64'")
	}

	return &params
}

type parametersRecord struct {
	parameters  map[string]Parameter
	requestTime int64
	requestType uint
}

func (p *parametersRecord) Get(name string) (Parameter, error) {
	if val, ok := p.parameters[name]; ok {
		return val, nil
	}
	return "", fmt.Errorf("Parameter with '%s' name not found", name)
}

func (p *parametersRecord) Type() uint {
	return p.requestType
}

func (p *parametersRecord) Time() int64 {
	return p.requestTime
}
