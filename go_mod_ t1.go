package main

import (
	"fmt"

	"github.com/sauravdeshpande/gomodules/hello"
	"rsc.io/quote"
)

func main() {

	x := hello.Hello()
	fmt.Println(quote.Hello())

	fmt.Println(x)
}
