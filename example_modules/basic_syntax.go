package examplemodules

import "fmt"

func Basic_syntax() {

	// variable declaration with var
	var greet string
	greet = "Hello, world!"

	// variable declaration with assignment
	greet_1 := "Hello, world! 1"

	// what is the type of `answer`?
	answer := 1 + 2

	// show both
	fmt.Println(greet)
	fmt.Println(greet_1)
	fmt.Println(answer)

	for i := 5; i > -1; i-- {
		fmt.Println("Counting Down ..", i)
	}
}

