package model

//Задача 7: Композиция структур
//Создайте структуру Address с полями Street, City и PostalCode.
//Измените структуру Person, чтобы она включала Address вместо строки Address.

type Address struct {
	City       string
	Street     string
	PostalCode int
}
