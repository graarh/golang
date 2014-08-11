package cellphone

import (
	"github.com/graarh/golang/expert/data"
	"time"
)

func CreateParameters(call Call) data.Parameters {
	return data.CreateParameters(map[string]data.Parameter{
		"Time": time.Now().Unix(),

		"callTime": call.CallTime.Format(time.RFC822),
		"code":     call.CountryCode,
		"prefix":   call.Prefix,
		"number":   call.Number,
		"duration": uint(call.Duration),
	})
}
