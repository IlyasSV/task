package model

import "fmt"

//Задача 1: Создание структуры
//Создайте структуру Person с полями Name, Age и Address.

type Person struct {
	Name    string
	Age     int
	Address Address
}

func (p Person) Print() {
	fmt.Printf("Name: %s\nAge: %d\nAddress: %+v\n", p.Name, p.Age, p.Address)
}

func (p Person) PrintName() {
	fmt.Printf("Name: %s\n", p.Name)
}

func (p Person) PrintAge() {
	fmt.Printf("Age: %d\n", p.Age)
}

func (p Person) PrintAddress() {
	fmt.Printf("Address: %s\n", p.Address)
}
