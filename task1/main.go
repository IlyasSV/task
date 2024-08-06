package main

import (
	"taskValihan/model"
)

func main() {

	//Задача 2: Инициализация структуры
	//Инициализируйте несколько экземпляров структуры Person и выведите их на экран.

	aaa := model.Person{
		Name: "AAA",
		Age:  20,
		Address: model.Address{
			City:       "Иннополис",
			Street:     "ул. Спортивная, 132",
			PostalCode: 123456,
		},
	}

	bbb := model.Person{Name: "BBB", Age: 21, Address: model.Address{City: "Казань",
		Street:     "ул. Спортивная, 132",
		PostalCode: 123456}}

	aaa.Print()
	bbb.Print()

	//Задача 3: Изменение полей структуры
	//Измените значение поля Address у одного из экземпляров структуры Person
	//и выведите обновленный экземпляр на экран.

	bbb.Name = "ABC"
	bbb.Print()

	aaa.Age = 30
	aaa.Print()

	//Задача 4: Методы для структур
	//Добавьте метод Greet для структуры Person, который будет выводить приветствие
	//с именем человека.

	model.Greet(bbb)

	//Задача 6: Инициализация вложенной структуры
	//Инициализируйте структуру Company и выведите информацию о компании и её сотрудниках.

	emp1 := model.Person{
		Name: "RTY",
		Age:  20,
		Address: model.Address{
			City:       "Москва",
			Street:     "ул. Спортивная, 132",
			PostalCode: 123456},
	}

	emp1.Print()

	company := model.Company{
		Name:      "RED",
		Employees: []model.Person{emp1},
	}

	company.PrintCompany()

}
