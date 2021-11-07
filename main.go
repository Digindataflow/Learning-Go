package main

import (
	"fmt"

	"github.com/Digindataflow/Learning-Go/Account"
)

func addTwo(a int, b int) int {
	return a + b
}

func main() {
	add := Account.AddAdaptor(addTwo)
	aDepartment := Account.NewDepartment(add, 1)
	aDepartment.PrintCal(3, 4)

	ch, cancel := Account.CountTo(10)
	for i := range ch {
		if i > 5 {
			break
		}
		fmt.Println(i)
	}
	cancel()
}
