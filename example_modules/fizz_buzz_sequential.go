package examplemodules

import(
  "strconv"
)

func Fizzbuzz_sequential(start, end int) []string {

	// declare the array we're using to keep the strings in
	var fb []string

	for i := start; i <= end; i++ {

		// keep a variable to check if the number is divisible by 3 or 5
		divisible := 0

		// no else used here to ensure numbers divisbile by 3 and 5 show up as "fizzbuzz"
		if i%3 == 0 {
			fb = append(fb, "fizz")
			divisible = 1
		}
		if i%5 == 0 {
			fb = append(fb, "buzz")
			divisible = 1
		}

		// if we haven't appended anything, push the number itself
		if divisible == 0 {
			fb = append(fb, strconv.Itoa(i))
		}

		fb = append(fb, "\n")

	}
	return fb
}
