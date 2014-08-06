package expert

import (
	"github.com/graarh/golang/expert/data"
	"github.com/graarh/golang/expert/modifier"
	"testing"
)

// no rules -> initial weight
func TestZeroRules(t *testing.T) {
	calculator := Calculator{[]Rule{}}
	resultWeight, ruleName := calculator.Max(&weight, params)
	iwWeight := resultWeight.(*data.SingleWeight)

	if iwWeight.Value.(int) != weight.Value.(int) {
		t.Error("Calculate with no rules should return initial weight, not", iwWeight)
	}

	if ruleName != "" {
		t.Error("Calculate with no rules should return empty string as rule name, not", ruleName)
	}
}

// rules that does not provide optimal weight -> initial weight
func TestGoroutineRulesNotOptimal(t *testing.T) {
	m1 := []Modifier{
		&modifier.Static{&data.SingleWeight{1}},
		&modifier.Static{&data.SingleWeight{2}},
	}
	m2 := []Modifier{
		&modifier.Static{&data.SingleWeight{2}},
		&modifier.Static{&data.SingleWeight{3}},
	}
	m3 := []Modifier{
		&modifier.Static{&data.SingleWeight{3}},
		&modifier.Static{&data.SingleWeight{4}},
	}

	calculator := Calculator{[]Rule{
		&RuleRecord{"rule1", conditions[0:2], m1},
		&RuleRecord{"rule2", conditions[0:3], m2},
		&RuleRecord{"rule3", conditions[1:3], m3},
	}}
	resultWeight, ruleName := calculator.Max(&weight, params)
	iwWeight := resultWeight.(*data.SingleWeight)

	if iwWeight.Value.(int) != weight.Value.(int) {
		t.Error("Calculate should return optimal weight, initial one in this case, not ", iwWeight)
	}
	if ruleName != "" {
		t.Error("Calculate should return empty rule if initial weight is optimal, not ", ruleName)
	}
}

// rules that provide optimal weight -> optimal weight
func TestGoroutineRulesOptimalSelection(t *testing.T) {
	m1 := []Modifier{
		&modifier.Static{&data.SingleWeight{1}},
		&modifier.Static{&data.SingleWeight{2}},
	}
	m2 := []Modifier{
		&modifier.Static{&data.SingleWeight{6}},
		&modifier.Static{&data.SingleWeight{7}},
	}
	m3 := []Modifier{
		&modifier.Static{&data.SingleWeight{3}},
		&modifier.Static{&data.SingleWeight{4}},
	}
	calculator := Calculator{[]Rule{
		&RuleRecord{"rule1", conditions[0:2], m1},
		&RuleRecord{"rule2", conditions[0:3], m2},
		&RuleRecord{"rule3", conditions[1:3], m3},
	}}
	resultWeight, ruleName := calculator.Max(&weight, params)
	iwWeight := resultWeight.(*data.SingleWeight)

	if iwWeight.Value.(int) != 7 {
		t.Error("Calculate should return optimal weight, IntWeight.Value == 2 in this case, not", iwWeight)
	}
	if ruleName != "rule2" {
		t.Error("Calculate should return 'rule2' rule, that returns optimal weight, not ", ruleName)
	}
}

func BenchmarkCommonCalculation(b *testing.B) {
	//try to increase amount of rules and look for benchmark changes
	m1 := []Modifier{
		&modifier.Static{&data.SingleWeight{9}},
		&modifier.Static{&data.SingleWeight{10}},
	}
	m2 := []Modifier{
		&modifier.Static{&data.SingleWeight{1}},
		&modifier.Static{&data.SingleWeight{2}},
	}
	m3 := []Modifier{
		&modifier.Static{&data.SingleWeight{13}},
		&modifier.Static{&data.SingleWeight{14}},
	}
	calculator := Calculator{[]Rule{
		&RuleRecord{"rule1", conditions[0:2], m1},
		&RuleRecord{"rule2", conditions[0:3], m2},
		&RuleRecord{"rule3", conditions[1:3], m3},
	}}
	for i := 0; i < b.N; i++ {
		calculator.Max(&weight, params)
	}
}
