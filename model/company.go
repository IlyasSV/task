package model

import "fmt"

//Задача 5: Вложенные структуры
//Создайте структуру Company с полями Name и Employees, где Employees — это срез структур Person.

type Company struct {
	Name      string
	Employees []Person
}

func (c Company) PrintCompany() {
	fmt.Printf("Name company: %s\nInfo employee:%+v\n", c.Name, c.Employees)
}
