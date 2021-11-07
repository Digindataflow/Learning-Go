package Account

import (
	"fmt"
)

type calculator interface {
	to(a int, b int) int
}

type AddAdaptor func(a int, b int) int

func (f AddAdaptor) to(a int, b int) int {
	return f(a, b)
}

type department struct {
	c  calculator
	Id int
}

func (d department) PrintCal(a, b int) {
	res := d.c.to(a, b)
	fmt.Println(res)
}

func NewDepartment(c calculator, Id int) department {
	return department{
		c:  c,
		Id: Id,
	}
}
