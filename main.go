package main

import (
	"fmt"
)

type person struct{
	Name string
	Email string
	Age int
}

func (p *person) SayHello() {
	fmt.Println("Hello", p.Name)

}


func main(){
	var guy = new(person)
	guy.Name = "Alvin"
	guy.SayHello()

}