package model

import "fmt"

//Задача 1: Создание структуры
//Создайте структуру Person с полями Name, Age и Address.

type Person struct {
	Name    string
	Age     int
	Address Address
}

type Address struct {
}

func PrintPerson(person Person) {
	fmt.Printf("Name: %s\nAge: %d\nAddress: %s\n", person.Name, person.Age, person.Address)
}
