package main

import "fmt"
import "io/ioutil"

func Recover() {
	if r := recover(); r != nil {
		fmt.Println("Recovered:", r)
	}
}

func main() {
	defer Recover() // catch
	buffer, err := ioutil.ReadFile("file.txt")
	if (err != nil) {
		panic(err) // throw
	}

	fmt.Println(string(buffer))
}