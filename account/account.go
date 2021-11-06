package main

import (
	"fmt"
)

type calculator interface {
	to(a int, b int) int
}

type addAdaptor func(a int, b int) int

func (f addAdaptor) to(a int, b int) int {
	return f(a, b)
}

func addTwo(a int, b int) int {
	return a + b
}

type department struct {
	c  calculator
	Id int
}

func (d department) printCal(a, b int) {
	res := d.c.to(a, b)
	fmt.Println(res)
}

func newDepartment(c calculator, Id int) department {
	return department{
		c:  c,
		Id: Id,
	}
}

func main() {
	add := addAdaptor(addTwo)
	aDepartment := newDepartment(add, 1)
	aDepartment.printCal(3, 4)
}
