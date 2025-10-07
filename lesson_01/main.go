package main

import "fmt"

type employee struct {
	name   string
	sex    string
	age    int
	salary int
}

func newEmployee(name string, sex string, age int, salary int) employee {
	return employee{
		name:   name,
		sex:    sex,
		age:    age,
		salary: salary,
	}
}

// pointer receiver
func (e *employee) getEmployeeInfo() {
	fmt.Printf("New employee with name %s and age %d\n", e.name, e.age)
}

// value receiver
func (e employee) setEmployeeName(name string) {
	e.name = name
}

type age int

func (a age) isAdult() bool {
	return a >= 20
}

func main() {
	employee1 := newEmployee("Gosha", "male", 30, 2500)
	employee1.setEmployeeName("Dosha")
	employee1.getEmployeeInfo()

	myAge := age(20)
	fmt.Printf("Are am adult? %t\n", myAge.isAdult())
}
