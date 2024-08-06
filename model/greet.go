package model

import "fmt"

//Задача 4: Методы для структур
//Добавьте метод Greet для структуры Person, который будет выводить приветствие с именем человека.

func Greet(name Person) {
	fmt.Printf("Hello %s!\n", name.Name)
}
