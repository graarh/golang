package expert

import "github.com/graarh/golang/expert/data"

//Calculator structure initializes calculator
type Calculator struct {
	Rules []Rule
}

type ruleResult struct {
	ruleName string
	weight   data.Weight
	affected bool
}

func (c *Calculator) calculate(initial data.Weight, params data.Parameters) chan ruleResult {
	results := make(chan ruleResult, len(c.Rules))
	for _, rule := range c.Rules {
		// TODO: make additional checks that rule passed correctly
		// maybe use channel to pass rule and initial
		go func(rule Rule, results chan ruleResult) {
			result, affected := rule.Calculate(initial, params)
			results <- ruleResult{rule.Name(), result, affected}
		}(rule, results)
	}

	return results
}

//Optimal finds maximum weight
func (c *Calculator) Max(initial data.Weight, params data.Parameters) (data.Weight, string) {
	optimal := ruleResult{"", initial, true}

	results := c.calculate(initial, params)

	for _ = range c.Rules {
		calculatedWeight := <-results

		if calculatedWeight.affected && optimal.weight.Less(calculatedWeight.weight) {
			optimal = calculatedWeight
		}
	}

	return optimal.weight, optimal.ruleName
}

//Collect collects weights to initial
func (c *Calculator) Sum(initial data.Weight, params data.Parameters) []string {
	names := make([]string, 0)
	results := c.calculate(initial, params)

	for _ = range c.Rules {
		calculatedWeight := <-results

		if calculatedWeight.affected {
			names = append(names, calculatedWeight.ruleName)
			initial.Add(calculatedWeight.weight)
		}
	}

	return names
}
