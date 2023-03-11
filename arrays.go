package arrays

import "fmt"

func array(){
	fmt.Println("Welcome to arrays")
	//This elemnts are contigous in memory
	//This is a literal sysntax this alloctes enough space to hold the array
	grades := [...]int {87,34,90,45,23}
	a := [...]int{1,2,3,4,5,6,7,8,9}
	fmt.Println("a =",a[8])
	var students [3]string
	students[0] = "Lista"
	students[1] = "Arnold"
	fmt.Print(len(students))
	fmt.Print(grades)
	fmt.Print("students: %v\n",students[1])

}