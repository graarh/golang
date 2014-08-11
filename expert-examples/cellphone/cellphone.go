package cellphone

import (
	"github.com/graarh/golang/expert"
	"github.com/graarh/golang/expert-examples/cellphone/config"
	"github.com/graarh/golang/expert/data"
)

type CallBill struct {
	calculators   map[string]expert.Calculator
	initialWeight data.Weight
}

type BillingResult struct {
	Price float32
	Rule  string
}

func CreateCallBiller(plansConfig string) (*CallBill, error) {
	plans, err := LoadPlans(plansConfig)
	if err != nil {
		return nil, err
	}

	cb := CallBill{make(map[string]expert.Calculator), &data.SingleWeight{float32(0)}}

	for name, plan := range plans {
		cb.calculators[name] = expert.Calculator{CreateExpertRules(plan)}
	}

	return &cb, nil
}

func (c *CallBill) Calculate(params data.Parameters) map[string]BillingResult {
	//preprocess
	params = config.Preprocess(params)

	result := make(map[string]BillingResult)
	for name, calculator := range c.calculators {
		max, rule := calculator.Max(c.initialWeight, params)
		swMax := max.(*data.SingleWeight)
		duration, _ := params.Get("duration")
		result[name] = BillingResult{
			swMax.Value.(float32) * float32(duration.(uint)),
			rule,
		}
	}

	return result
}
