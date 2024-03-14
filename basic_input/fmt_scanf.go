package main

import "fmt"

func main() {
    var name, address string

    fmt.Print("Enter your name and address : ")
    fmt.Scanf("%s %s", &name, &address) // menerima 2 input string yang dipisahkan oleh spasi, yang pertama akan diassign ke variable name dan yang kedua akan diassign ke variable address

    fmt.Println("Hello", name)
    fmt.Println("Your address", address)
}