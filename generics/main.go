package main

//Generics help reduce code repetition
import (
	"fmt"
	"golang.org/x/exp/constraints"
)

func addNums(num1 int, num2 int) int{
	return num1 + num2
}


type Num interface{
	int | int8 | int16 | int64 | float64
}

func addInt[T Num](a T, b T)T{
	return a + b
}


func retNums[C constraints.Ordered](a C, b C) C {
	return a + b
}

func main(){
	res := addInt(4.2,4.5)
	resv := retNums(5,4)
	result := addNums(3,4)
	fmt.Printf("The sum is :%v\n", result)
	fmt.Printf("The sum is :%v\n", resv)
	fmt.Printf("The sum is :%v", res)
}