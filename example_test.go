package pasta_test

import (
	"fmt"

	"github.com/ardnew/pasta"
)

func Example() {
	pasta.WriteAll("Привет, мир!")
	text, _ := pasta.ReadAll()
	fmt.Println(text)

	// Output:
	// Привет, мир!
}
