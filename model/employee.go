package model

import "fmt"

// Задача 9: Анонимные поля
// Создайте структуру Employee, которая будет включать анонимное поле типа Person, а также поле Position.
type Employee struct {
	Person
	Position string
}

func (e *Employee) PrintEmployee() {
	fmt.Printf("Информация о сотруднике:\nИмя: %s Возвраст: %d\nАдрес: %+v\nPosition: %s\n", e.Name, e.Age, e.Address, e.Position)
}
