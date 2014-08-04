expert
======

Expert system library

### Installation

    go get github.com/graarh/golang/expert

### Usage

Complete examples will be available later.

Look for rule_test.go and expert_test.go
for usage examples

### Quick schema

1. Input parameters
   - request parameters <map[string]Parameter> <type Parameter interface{}>
   - request initial weight <Weight> <type Weight interface{}>
   - rule set <[]Rule>
2. Output result
   - weight <Weight>, ruleName <string>
3. Construction elements
   - <Rule: []conditions, []modifiers>
   - <condition.Check(params) bool>
   - <modifier.apply(Weight, params) Weight>
