package main

import "fmt"

func main() {
    var name, address string
    fmt.Print("Enter your name : ")
    fmt.Scanln(&name)

    fmt.Print("Enter your address : ")
    fmt.Scanln(&address)

    fmt.Println("Hello", name)
    fmt.Println("Your address", address)
}