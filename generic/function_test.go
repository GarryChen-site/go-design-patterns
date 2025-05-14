package generic

import (
	"fmt"
	"testing"
)

type Employee struct {
	Name     string
	Age      int
	Vacation int
	Salary   float32
}

var employee = []Employee{
	{"John", 30, 10, 50000},
	{"Jane", 25, 15, 60000},
	{"Bob", 35, 20, 70000},
	{"Alice", 28, 12, 55000},
	{"Charlie", 40, 25, 80000},
	{"Dave", 32, 18, 65000},
}

func Test_function(t *testing.T) {
	totalPay := gReduce(employee, 0.0, func(total float32, emp Employee) float32 {
		return total + emp.Salary
	})

	fmt.Printf("Total Pay: %.2f\n", totalPay)

	old := gCountIf(employee, func(emp Employee) bool {
		return emp.Age > 30
	})
	fmt.Printf("Number of employees older than 30: %d\n", old)

	highSalary := gCountIf(employee, func(emp Employee) bool {
		return emp.Salary > 60000
	})
	fmt.Printf("Employees with salary > 60000: %v\n", highSalary)

	youngerPay := gSum(employee, func(emp Employee) float32 {
		if emp.Age < 30 {
			return emp.Salary
		}
		return 0
	})
	fmt.Printf("Total pay for employees younger than 30: %.2f\n", youngerPay)

	totalVacation := gSum(employee, func(emp Employee) int {
		return emp.Vacation
	})
	fmt.Printf("Total vacation days: %d\n", totalVacation)
}
