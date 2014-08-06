package expert

import "github.com/graarh/golang/expert/data"

// Calculate finds optimal weight using initial one, parameters and rule set
func Calculate(initial data.Weight, rules []Rule, params data.Parameters) (data.Weight, string) {
	type weightWithRuleName struct {
		ruleName string
		weight   data.Weight
		affected bool
	}

	optimal := weightWithRuleName{"", initial, true}

	results := make(chan weightWithRuleName, len(rules))
	for _, rule := range rules {
		// TODO: make additional checks that rule passed correctly
		// maybe use channel to pass rule and initial
		go func(rule Rule, results chan weightWithRuleName) {
			result, affected := rule.Calculate(initial, params)
			results <- weightWithRuleName{rule.Name(), result, affected}
		}(rule, results)
	}

	for _ = range rules {
		calculatedWeight := <-results

		if calculatedWeight.affected && optimal.weight.Less(calculatedWeight.weight) {
			optimal = calculatedWeight
		}
	}

	return optimal.weight, optimal.ruleName
}
