package main

import (
	"taskValihan/model"
)

func main() {

	//Задача 2: Инициализация структуры
	//Инициализируйте несколько экземпляров структуры Person и выведите их на экран.

	aaa := model.Person{
		Name:    "AAA",
		Age:     20,
		Address: "3-33-35",
	}

	bbb := model.Person{Name: "BBB", Age: 21, Address: "5-55-555"}

	model.PrintPerson(aaa)
	model.PrintPerson(bbb)

	//Задача 3: Изменение полей структуры
	//Измените значение поля Address у одного из экземпляров структуры Person
	//и выведите обновленный экземпляр на экран.

	bbb.Name = "ABC"
	model.PrintPerson(bbb)

	aaa.Age = 30
	model.PrintPerson(aaa)

	//Задача 4: Методы для структур
	//Добавьте метод Greet для структуры Person, который будет выводить приветствие
	//с именем человека.

	model.Greet(bbb)

	//Задача 6: Инициализация вложенной структуры
	//Инициализируйте структуру Company и выведите информацию о компании и её сотрудниках.

	emp1 := model.Person{
		Name:    "RTY",
		Age:     20,
		Address: "4-7-8",
	}

	company := model.Company{
		Name:      "RED",
		Employees: []model.Person{emp1},
	}

	model.PrintCompany(company)

}
