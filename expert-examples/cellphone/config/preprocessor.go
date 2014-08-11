package config

import (
	"github.com/graarh/golang/expert"
	"github.com/graarh/golang/expert/condition"
	"github.com/graarh/golang/expert/data"
	"github.com/graarh/golang/expert/modifier"
)

var tele2 = &expert.RuleRecord{
	"Tele2 prefixes",
	[]expert.Condition{
		&condition.Map{
			"code",
			condition.SliceToMap([]data.Parameter{
				"+7", "8",
			}),
			false,
		},
		&condition.Map{
			"prefix",
			condition.SliceToMap([]data.Parameter{
				"900", "902", "904", "908", "950", "951", "952", "953",
			}),
			false,
		},
	},
	[]expert.Modifier{
		&modifier.Static{
			&data.ParamsWeight{
				data.CreateParameters(map[string]data.Parameter{
					"operator": "Tele2",
				}),
				nil,
			},
		},
	},
}

var local = &expert.RuleRecord{
	"Tele2 locations",
	[]expert.Condition{
		&condition.Map{
			"code",
			condition.SliceToMap([]data.Parameter{
				"+7", "8",
			}),
			false,
		},
		&condition.Map{
			"prefix",
			condition.SliceToMap([]data.Parameter{
				"951", "953", "381",
			}),
			false,
		},
	},
	[]expert.Modifier{
		&modifier.Static{
			&data.ParamsWeight{
				data.CreateParameters(map[string]data.Parameter{
					"region": "Local",
				}),
				nil,
			},
		},
	},
}

var preprocessor = expert.Calculator{[]expert.Rule{tele2, local}}

func Preprocess(params data.Parameters) data.Parameters {
	weight := &data.ParamsWeight{params, nil}
	preprocessor.Sum(weight, params)
	return weight.Params
}
