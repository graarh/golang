package main

import (
	"flag"
	"github.com/graarh/golang/expert-examples/cellphone"
	"log"
)

const defaultCalls = "../calls/mixed.csv"
const defaultPlans = "../config/callplans.yml"

var calls, plans string

func main() {
	flag.StringVar(&calls, "calls", defaultCalls, "csv file with calls")
	flag.StringVar(&plans, "plans", defaultPlans, "billing plans configuration file")
	flag.Parse()

	calls, err := cellphone.LoadCalls(calls)
	if err != nil {
		panic(err)
	}

	cb, err := cellphone.CreateCallBiller(plans)
	if err != nil {
		panic(err)
	}

	var total map[string]float32 = make(map[string]float32)

	for _, call := range calls {
		log.Print("-------------------")
		log.Print("Call: ", call.CountryCode, " ", call.Prefix)

		params := cellphone.CreateParameters(call)

		prices := cb.Calculate(params)
		for operator, billing := range prices {
			log.Printf("%10s: %.02f (%s)", operator, billing.Price, billing.Rule)
			total[operator] += billing.Price
		}

		log.Print("-------------------")
		log.Println()
	}

	log.Print("-------Total:------")
	for operator, value := range total {
		log.Printf("%10s: %.2f", operator, value)
	}
}
