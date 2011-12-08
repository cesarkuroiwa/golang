package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5} // Array with 5 elements
	
	s := a[:3] // s points to the first 3 elements of a
	
	fmt.Println("s: ", s)
	fmt.Println("a: ", a)
	s[0] = 9 // Change the first element of s2, s1 and a
	fmt.Println("s: ", s)
	fmt.Println("a: ", a)
}