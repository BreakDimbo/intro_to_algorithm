package main

import (
	"fmt"
	"time"

	"github.com/BreakDimbo/intro_to_algorithm/pkg/fib"
	"github.com/BreakDimbo/intro_to_algorithm/pkg/util"
)

func main() {
	defer util.ConsumeTime(time.Now())
	// r := fib.Fv(45) // fib result 1134903170 time consumed is 7.794661s
	// r := fib.Sv(1000) // fib result 1134903170 time consumed is 0.000042s
	r := fib.Tv(1000)
	fmt.Printf("fib result %d\n", r)
}
