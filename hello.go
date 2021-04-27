package main

import (
	"fmt"
	"github.com/mujl/gomod_test/logicOP"
	"rsc.io/quote"
)

func main() {

	fmt.Println(quote.Hello())
	fmt.Println("3+5=", logicOP.SUM(3, 5))
}
