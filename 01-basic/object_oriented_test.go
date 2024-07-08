package main

import (
	"testing"
)

func TestObjectOriented(t *testing.T) {
	var sayer Sayer

	sayer = &Dog{Animal{Name: "Yoda"}}
	sayer.Say("hello world") // Dog[Yoda] say: hello world
	sayer.SayHi()            // Animal[Yoda] say: Hi

	sayer = &Cat{Animal{Name: "Jerry"}}
	sayer.Say("hello world") // Cat[Jerry] say: hello world
	sayer.SayHi()            // Cat[Jerry] say: Hi
}
