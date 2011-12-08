package main

import "mypackage"

func main() {
	var m mypackage.MyClass
	m.SetX(123)
	mypackage.PrintX(&m)
}
