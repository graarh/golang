package expert

import (
	"github.com/graarh/golang/expert/data"
	"github.com/graarh/golang/expert/modifier"
	"testing"
)

// no rules -> initial weight
func TestZeroRules(t *testing.T) {
	rules := []Rule{}
	resultWeight, ruleName := Calculate(&weight, rules, params)
	iwWeight := resultWeight.(*data.IntWeight)

	if iwWeight.Value != weight.Value {
		t.Error("Calculate with no rules should return initial weight, not", iwWeight)
	}

	if ruleName != "" {
		t.Error("Calculate with no rules should return empty string as rule name, not", ruleName)
	}
}

// rules that does not provide optimal weight -> initial weight
func TestGoroutineRulesNotOptimal(t *testing.T) {
	m1 := []Modifier{
		&modifier.IwModifier{func(int) int { return 1 }},
		&modifier.IwModifier{func(int) int { return 2 }},
	}
	m2 := []Modifier{
		&modifier.IwModifier{func(int) int { return 2 }},
		&modifier.IwModifier{func(int) int { return 3 }},
	}
	m3 := []Modifier{
		&modifier.IwModifier{func(int) int { return 3 }},
		&modifier.IwModifier{func(int) int { return 4 }},
	}

	rules := []Rule{
		&RuleRecord{"rule1", conditions[0:2], m1},
		&RuleRecord{"rule2", conditions[0:3], m2},
		&RuleRecord{"rule3", conditions[1:3], m3},
	}
	resultWeight, ruleName := Calculate(&weight, rules, params)
	iwWeight := resultWeight.(*data.IntWeight)

	if iwWeight.Value != weight.Value {
		t.Error("Calculate should return optimal weight, initial one in this case, not ", iwWeight)
	}
	if ruleName != "" {
		t.Error("Calculate should return empty rule if initial weight is optimal, not ", ruleName)
	}
}

// rules that provide optimal weight -> optimal weight
func TestGoroutineRulesOptimalSelection(t *testing.T) {
	m1 := []Modifier{
		&modifier.IwModifier{func(int) int { return 1 }},
		&modifier.IwModifier{func(int) int { return 2 }},
	}
	m2 := []Modifier{
		&modifier.IwModifier{func(int) int { return 6 }},
		&modifier.IwModifier{func(int) int { return 7 }},
	}
	m3 := []Modifier{
		&modifier.IwModifier{func(int) int { return 3 }},
		&modifier.IwModifier{func(int) int { return 4 }},
	}
	rules := []Rule{
		&RuleRecord{"rule1", conditions[0:2], m1},
		&RuleRecord{"rule2", conditions[0:3], m2},
		&RuleRecord{"rule3", conditions[1:3], m3},
	}
	resultWeight, ruleName := Calculate(&weight, rules, params)
	iwWeight := resultWeight.(*data.IntWeight)

	if iwWeight.Value != 7 {
		t.Error("Calculate should return optimal weight, IntWeight.Value == 2 in this case, not", iwWeight)
	}
	if ruleName != "rule2" {
		t.Error("Calculate should return 'rule2' rule, that returns optimal weight, not ", ruleName)
	}
}

func BenchmarkCommonCalculation(b *testing.B) {
	//try to increase amount of rules and look for benchmark changes
	m1 := []Modifier{
		&modifier.IwModifier{func(int) int { return 9 }},
		&modifier.IwModifier{func(int) int { return 10 }},
	}
	m2 := []Modifier{
		&modifier.IwModifier{func(int) int { return 1 }},
		&modifier.IwModifier{func(int) int { return 2 }},
	}
	m3 := []Modifier{
		&modifier.IwModifier{func(int) int { return 13 }},
		&modifier.IwModifier{func(int) int { return 14 }},
	}
	rules := []Rule{
		&RuleRecord{"rule1", conditions[0:2], m1},
		&RuleRecord{"rule2", conditions[0:3], m2},
		&RuleRecord{"rule3", conditions[1:3], m3},
	}
	for i := 0; i < b.N; i++ {
		Calculate(&weight, rules, params)
	}
}
