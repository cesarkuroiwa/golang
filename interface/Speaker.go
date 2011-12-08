package main

import "fmt"

type SpeakerInterface interface { speak() }

type Speaker1 struct {}

func (p Speaker1) speak() {
	fmt.Println("Hello from Speaker1")
}

type Speaker2 struct {}

func (p Speaker2) speak() {
	fmt.Println("Hello from Speaker2")
}

func saySomething(s SpeakerInterface) {	s.speak() }

func main() {
	var s1 Speaker1
	var s2 Speaker2
	saySomething(s1)
	saySomething(s2)
}