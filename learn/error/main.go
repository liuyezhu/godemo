// Golang program to illustrate the usage of
// fmt.Errorf() function

// Including the main package
package main

// Importing fmt
import (
	"fmt"
)

// Calling main
func main() {

	// Declaring some constant variables
	const name, dept = "GeeksforGeeks", "CS"

	// Calling the Errorf() function with verb
	// %q which is used for a single-quoted character
	err := fmt.Errorf("%q is a %q Portal.", name, dept)

	// Printing the error message
	fmt.Println(err.Error())
}
