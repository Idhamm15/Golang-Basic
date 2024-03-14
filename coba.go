package main

import "fmt"

func main() {
	// x := 10
	// if x > 5 {
	// 	fmt.Println("x is greater than 5")
	// }else if x > 2{
	// 	fmt.Println("x is greater than 2")
	// }else {
	// 	fmt.Println("x is not greater than 2")

	// s1 := []int{1, 2 ,3}
	// s2 := []int{4,5,6}
	// s1 = append(s1, s2...)
	// s2[0] = 0
	// fmt.Println(s1)

	// a := [3]int{1,2,3}
	// b := &a[0]
	// c := &a[1]
	// fmt.Printf("%v, %p, %p\n",a,b,c)
	// b = c
	// *b = 4
	// fmt.Printf("%v, %p, %p\n",a,b,c)

	s:= make([]int,0)
	s= append(s,1)
	s= append(s,2,3,4)
	s= append(s,5)
	fmt.Println(s[2:4])
}
