package _Functions

import "rsc.io/quote"

const MyName = "yash" // exported, can be used outside

const age = 18 // not exported

// exported, can be used outside

func GetIp() string {
	return "10.20.40.227"
}

func GetQuote() string {
	return quote.Go()
}
