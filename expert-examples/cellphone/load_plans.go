package cellphone

import (
	"github.com/graarh/golang/expert/data"
	"gopkg.in/yaml.v1"
	"io/ioutil"
)

type Rule struct {
	Conditions []Condition
	Price      float64
}

type Condition struct {
	Name    string
	Values  []data.Parameter
	Exclude bool
}

type Plans map[string][]Rule

func LoadPlans(fileName string) (Plans, error) {
	var plans []byte
	var err error

	if plans, err = ioutil.ReadFile(fileName); err != nil {
		return nil, err
	}

	var data Plans

	if err = yaml.Unmarshal(plans, &data); err != nil {
		return nil, err
	}

	return data, nil
}
