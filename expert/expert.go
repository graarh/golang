package expert

import "github.com/graarh/golang/expert/data"

// Calculate finds optimal weight using initial one, parameters and rule set
func Calculate(initial data.Weight, rules []Rule, params data.Parameters) (data.Weight, string) {
	type weightWithRuleName struct {
		ruleName string
		weight   data.Weight
	}

	optimal := weightWithRuleName{"", initial}

	results := make(chan weightWithRuleName, len(rules))
	for _, rule := range rules {
		// TODO: make additional checks that rule passed correctly
		// maybe use channel to pass rule and initial
		go func(rule Rule, results chan weightWithRuleName) {
			result := rule.Calculate(initial, params)
			results <- weightWithRuleName{rule.Name(), result}
		}(rule, results)
	}

	for _ = range rules {
		calculatedWeight := <-results

		if optimal.weight.Compare(calculatedWeight.weight) == -1 { // less than
			optimal = calculatedWeight
		}
	}

	return optimal.weight, optimal.ruleName
}
