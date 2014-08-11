package cellphone

import (
	"fmt"
	"github.com/graarh/golang/expert"
	"github.com/graarh/golang/expert/condition"
	"github.com/graarh/golang/expert/data"
	"github.com/graarh/golang/expert/modifier"
)

func CreateExpertRules(rules []Rule) []expert.Rule {
	result := make([]expert.Rule, 0)

	for ruleNum, rule := range rules {
		//create static modifier, that will return predefined value
		mod := modifier.Static{
			//the value is common single param weight (price)
			&data.SingleWeight{float32(rule.Price)},
		}
		conditions := make([]expert.Condition, 0)
		for _, cond := range rule.Conditions {
			//make map condition
			mapCond := &condition.Map{
				Name:    cond.Name,
				Values:  condition.SliceToMap(cond.Values),
				Exclude: cond.Exclude,
			}
			conditions = append(conditions, mapCond)
		}
		//create rule record
		result = append(result, &expert.RuleRecord{
			fmt.Sprintf("Rule N%v", ruleNum+1),
			conditions,
			[]expert.Modifier{&mod},
		})
	}
	return result
}
