package main

import (
	"fmt"
	"strconv"
)

var fib = func() func(int) uint64 {
	memo := make(map[int]uint64)
	var realfib func(int) uint64
	realfib = func(a int) uint64 {
		if a < 0 {
			return 0
		}
		if a == 0 {
			return 1
		}
		if a == 1 {
			return 1
		}
		r, found := memo[a]
		if found == false {
			memo[a] = realfib(a-1) + realfib(a-2)
			r = memo[a]
		} else {
			fmt.Println("Found:", r, "at index:", a)
		}
		return r
	}
	return realfib
}()

func main() {
	for i := 0; i < 100; i += 1 {
		fmt.Println("fib("+strconv.Itoa(i)+") == ", fib(i))
	}
}
