package main

import (
	"fmt"
	"os"
)

func main() {
	var s Person = &Student{name: "halfrost"}
	s.sayHello("everyone")
	fmt.Printf("os:%v", os.Args[0])
	p, _ := os.Executable()
	fmt.Printf("os:%v", p)
}

type Person interface {
	sayHello(name string) string
	sayGoodbye(name string) string
}

type Student struct {
	name string
}

//go:noinline
func (s *Student) sayHello(name string) string {
	return fmt.Sprintf("%v: Hello %v, nice to meet you.\n", s.name, name)
}

//go:noinline
func (s *Student) sayGoodbye(name string) string {
	return fmt.Sprintf("%v: Hi %v, see you next time.\n", s.name, name)
}
