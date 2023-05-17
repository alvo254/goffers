package main

import (
	"fmt"

	"github.com/alvo254/goffers/dataTypes"
)

type person struct {
	Name  string
	Email string
	Age   int
}

func (p *person) SayHello() {
	fmt.Println("Hello", p.Name)

}

func main() {
	res := dataTypes.AddNums(8,2)
	fmt.Printf("Here %v\n", res)
	fmt.Println(dataTypes.Data())
	var guy = new(person)
	guy.Name = "Alvin"
	guy.SayHello()

}
