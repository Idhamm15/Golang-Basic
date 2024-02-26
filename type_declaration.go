package main

import "fmt"

func main() {
	type noKTP string
	type Married bool

	var noKTPEko noKTP = "321946129746218794621"
	var MarriedStatus Married = true
	
	fmt.Println(noKTPEko)
	fmt.Println(MarriedStatus)
}