package expert

import "github.com/graarh/golang/expert/data"

// Rule is the interface of expert system that should
// affect weight then params pass conditions
type Rule interface {
	Name() string
	Conditions() []Condition
	Modifiers() []Modifier

	Calculate(initial data.Weight, params data.Parameters) data.Weight
}

// RuleRecord is the base implementation of the Rule interface
type RuleRecord struct {
	RuleName       string
	RuleConditions []Condition
	RuleModifiers  []Modifier
}

// Name is the Rule interface implementation
func (p *RuleRecord) Name() string {
	return p.RuleName
}

// Conditions is the Rule interface implementation
func (p *RuleRecord) Conditions() []Condition {
	return p.RuleConditions
}

// Modifiers is the Rule interface implementation
func (p *RuleRecord) Modifiers() []Modifier {
	return p.RuleModifiers
}

// Calculate applies modifiers to initial weight if conditions are passed with
// given params, implementation of the Rule interface
func (p *RuleRecord) Calculate(initial data.Weight, params data.Parameters) data.Weight {
	applyModifiers := true
	for _, condition := range p.Conditions() {
		check := condition.Check(params)
		if !check {
			applyModifiers = false
			break
		}
	}

	if applyModifiers {
		modWeight := initial
		for _, modifier := range p.Modifiers() {
			modWeight = modifier.Modify(modWeight, params)
		}
		return modWeight
	}

	return initial
}
