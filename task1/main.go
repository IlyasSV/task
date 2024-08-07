package main

import (
	"fmt"
	"taskValihan/model"
)

func main() {

	////Задача 2: Инициализация структуры
	////Инициализируйте несколько экземпляров структуры Person и выведите их на экран.
	//
	//aaa := model.Person{
	//	Name: "AAA",
	//	Age:  20,
	//	Address: model.Address{
	//		City:       "Иннополис",
	//		Street:     "ул. Спортивная, 132",
	//		PostalCode: 123456,
	//	},
	//}
	//
	//bbb := model.Person{Name: "BBB", Age: 21, Address: model.Address{City: "Казань",
	//	Street:     "ул. Спортивная, 132",
	//	PostalCode: 123456}}
	//
	//aaa.Print()
	//bbb.Print()
	//
	////Задача 3: Изменение полей структуры
	////Измените значение поля Address у одного из экземпляров структуры Person
	////и выведите обновленный экземпляр на экран.
	//
	//bbb.Name = "ABC"
	//bbb.Print()
	//
	//aaa.Age = 30
	//aaa.Print()
	//
	////Задача 4: Методы для структур
	////Добавьте метод Greet для структуры Person, который будет выводить приветствие
	////с именем человека.
	//
	//model.Greet(bbb)
	//
	////Задача 6: Инициализация вложенной структуры
	////Инициализируйте структуру Company и выведите информацию о компании и её сотрудниках.
	//
	//emp1 := model.Person{
	//	Name: "RTY",
	//	Age:  20,
	//	Address: model.Address{
	//		City:       "Москва",
	//		Street:     "ул. Спортивная, 132",
	//		PostalCode: 123456},
	//}
	//
	//emp1.Print()
	//
	//company := model.Company{
	//	Name:      "RED",
	//	Employees: []model.Person{emp1},
	//}
	//
	//company.PrintCompany()

	//Задача 10: Использование анонимных полей
	//Инициализируйте несколько экземпляров структуры Employee и выведите их на экран.

	employee1 := model.Employee{
		Person: model.Person{
			Name: "Jon",
			Age:  20,
			Address: model.Address{
				City:       "Japan",
				Street:     "ASDsd",
				PostalCode: 1234,
			},
		},
		Position: "1",
	}
	employee1.PrintEmployee()

	employees := []model.Employee{
		{
			Person: model.Person{
				Name: "Deni",
				Age:  24,
				Address: model.Address{
					City:       "Moscow",
					Street:     "Arbat",
					PostalCode: 1234,
				},
			},
			Position: "2",
		},
		{
			Person: model.Person{
				Name: "Nino",
				Age:  19,
				Address: model.Address{
					City:       "Bishkek",
					Street:     "Micro",
					PostalCode: 1234,
				},
			},
			Position: "3",
		},
		{
			Person: model.Person{
				Name: "Roni",
				Age:  20,
				Address: model.Address{
					City:       "Lisobon",
					Street:     "Pravda",
					PostalCode: 1234,
				},
			},
			Position: "4",
		},
	}

	employees = append(employees, employee1)
	printEmployees(employees)

	emp2 := []model.Employee{
		{
			Person: model.Person{
				Name: "Titi",
				Age:  27,
				Address: model.Address{
					City:       "Praga",
					Street:     "Rjaka",
					PostalCode: 10292,
				},
			},
			Position: "5",
		},
		{
			Person: model.Person{
				Name: "Yoyo",
				Age:  24,
				Address: model.Address{
					City:       "Riga",
					Street:     "Kievskaya",
					PostalCode: 808090,
				},
			},
			Position: "6",
		},
	}

	employees = append(employees, emp2...)

	printEmployees(employees)

}

func printEmployees(employees []model.Employee) {
	for _, emp := range employees {
		fmt.Println("Информация о сотруднике:")
		fmt.Printf("Имя: %s, Возврат: %d\nДомашний адрес:\nГород: %s\nУлица: %s\nИндекс:"+
			" %d\nПозиция: %s\n", emp.Name, emp.Age, emp.Address.City,
			emp.Address.Street, emp.Address.PostalCode, emp.Position)
	}
}
