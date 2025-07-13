package main

import (
	"fmt"
	"github.com/land1725/go-daily-algorithms/problems/day11_object_oriented"
)

func main() {
	rectangle := day11_object_oriented.Rectangle{Width: 10, Height: 10}
	circle := day11_object_oriented.Circle{
		Radius: 5,
	}
	fmt.Println("rectangle area", rectangle.Area())
	fmt.Println("rectangle perimeter", rectangle.Perimeter())
	fmt.Println("circle area", circle.Area())
	fmt.Println("circle perimeter", circle.Perimeter())

	employee := day11_object_oriented.Employee{
		Person: day11_object_oriented.Person{
			Name: "tiansongye",
			Age:  40,
		},
		EmployeeNumber: 1111,
	}
	employee.GetEmployeeInfo()
}
