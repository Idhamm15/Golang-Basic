package main

import "fmt"

func main()  {
	var name string
	
	name = "Eko Kurniawan"
	fmt.Println(name)

	name = "Eko Khannedy"
	fmt.Println(name)

	var friend = "Budi"
	fmt.Println(friend)

	var age = 38
	fmt.Println(age)

	// membuat variable tanpa menggunakan var
	country := "Indonesia"
	fmt.Println(country)

	country = "Amerika"

	// menggunakan var tetapi dengan model package
	var (
		firstname = "Eko"
		lastname = "Kurniawan"
	)

	fmt.Println(firstname)
	fmt.Println(lastname)


	
}