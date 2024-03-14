package main

import "fmt"

func main() {
    var name, address string
    fmt.Print("Enter your name and address : ")
    fmt.Scan(&name, &address) // set input to variable 'name' and 'address'

    fmt.Println("Hello", name)
    fmt.Println("Your address", address)
}