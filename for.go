package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// Go では if を短くすることが推奨されている
	// ネストを浅くすることが推奨されている

	// NG
	// ネストが深く、また if 文の中身も長い
	for i := 0; i <= 100; i++ {
		if i%3 == 0 {
			if i%5 == 0 {
				fmt.Println(i, "FizzBuzz")
			} else {
				fmt.Println(i, "Fizz")
			}
		} else if i%5 == 0 {
			fmt.Println(i, "Buzz")
		} else {
			fmt.Println(i)
		}
	}

	// OK (推奨)
	for i := 0; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println(i, "FizzBuzz")
			continue
		}
		if i%3 == 0 {
			fmt.Println(i, "Fizz")
			continue
		}
		if i%5 == 0 {
			fmt.Println(i, "Buzz")
			continue
		}
		fmt.Println(i)
	}

	// for-range
	evanVals := [6]int{2, 4, 6, 8, 10, 12}
	for _, v := range evanVals {
		fmt.Println(v)
	}
}
