package Functions

import "rsc.io/quote"

const MyName = "yash"

const age = 18

func GetIp() string {
	return "10.20.40.227"
}

func GetQuote() string {
	return quote.Go()
}
