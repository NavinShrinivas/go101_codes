package examplemodules

import(
  "fmt"
  "strconv"
)

func Fizzbuzz_stdout_sequential(start, end int) {
	for i := start; i <= end; i++ {
		if i%3 == 0 {
			fmt.Print("fizz")
		}
		if i%5 == 0 {
			fmt.Print("buzz")
		}
		if i%3 != 0 && i%5 != 0 {
			fmt.Print(strconv.Itoa(i))
		}
		fmt.Println("")
	}
}

