package Animal

import "fmt"

type Dog struct {
	name string
}

func NewDog(name string) *Dog {
	return &Dog{
		name: name,
	}
}

func (d Dog) Speak() {
	fmt.Println("I am a Dog")
}

func (d Dog) Run() {
	fmt.Println("Dog is running")
}
