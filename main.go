package main

import (
	"fmt"
)

func FizzBuzz(n int) {
	for i := 1; i <= n; i++ {
		var res string

		if i%3 == 0 {
			res += "Fizz"
		}

		if i%5 == 0 {
			res += "Buzz"
		}

		if res != "" {
			fmt.Println(res)
		} else {
			fmt.Println(i)
		}
	}
}

func FizzBuzz_v2(n int) {
	for i := 0; i <= n; i++ {
		if i%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

func main() {
	fmt.Println("Nothing much here")
}
