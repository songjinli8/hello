package hello

import (
	"rsc.io/quote"
)

//Hello say hello
func Hello() string {
	return "this is for test"
}

func HelloQuote() string {
	return quote.Hello()
}
