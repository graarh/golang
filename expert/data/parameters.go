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
	All() map[string]Parameter          // Get all parameters as map
	Type() uint                         // Get type of the request
	Time() int64                        // Get time of the request
}

// CreateParameters generates Parameters interface
// from the common key-value map
func CreateParameters(data map[string]Parameter) Parameters {
	params := parametersRecord{data, 0, 0}

	if _, ok := data["Type"]; ok {
		if requestType, ok := data["Type"].(uint); ok {
			params.requestType = requestType
		} else {
			trace.Log("Parameters 'Type' value type must be 'uint'")
		}
	}

	if _, ok := data["Time"]; ok {
		if requestTime, ok := data["Time"].(int64); ok {
			params.requestTime = int64(requestTime)
		} else {
			trace.Log("Parameters 'Time' value type must be 'int64'")
		}
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

func (p *parametersRecord) All() map[string]Parameter {
	return p.parameters
}

func (p *parametersRecord) Type() uint {
	return p.requestType
}

func (p *parametersRecord) Time() int64 {
	return p.requestTime
}
