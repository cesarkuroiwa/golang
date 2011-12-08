package main

import "container/list"

type MyStruct struct {}

func main() {
	for i := 0; i < 100000; i++ {
		list := list.New()
		for j := 0; j < 10000; j++ {
			myStruct := new(MyStruct)
			list.PushBack(myStruct)
		}
	}
}