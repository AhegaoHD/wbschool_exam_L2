package pattern

import "fmt"

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

type SubsystemOne struct{}

func (s *SubsystemOne) OperationOne() string {
	return "SubsystemOne: Ready!\n"
}

func (s *SubsystemOne) OperationN() string {
	return "SubsystemOne: Go!\n"
}

// Подсистема 2
type SubsystemTwo struct{}

func (s *SubsystemTwo) OperationOne() string {
	return "SubsystemTwo: Get Ready!\n"
}

func (s *SubsystemTwo) OperationZ() string {
	return "SubsystemTwo: Fire!\n"
}

// Фасад
type Facade struct {
	one *SubsystemOne
	two *SubsystemTwo
}

func NewFacade() *Facade {
	return &Facade{
		one: &SubsystemOne{},
		two: &SubsystemTwo{},
	}
}

func (f *Facade) Operation() string {
	result := "Facade initializes subsystems:\n"
	result += f.one.OperationOne()
	result += f.two.OperationOne()
	result += "Facade orders subsystems to perform the action:\n"
	result += f.one.OperationN()
	result += f.two.OperationZ()
	return result
}

func main() {
	facade := NewFacade()
	fmt.Print(facade.Operation())
}

//Применимость паттерна «Фасад»:
//
//Когда нужно предоставить простой или урезанный интерфейс к сложной системе.
//Когда нужно разложить подсистему на отдельные слои.

//Плюсы:
//
//Изолирует код от сложности подсистем.
//Упрощает взаимодействие со сложными подсистемами.

//Минусы:
//
//Фасад может стать "божественным объектом", слишком многим занимающимся в системе.
//Фасады могут скрывать возможности подсистем от клиентов.

//Примеры использования на практике:
//
//Библиотеки для работы с файловыми системами, где фасад предоставляет простые методы для чтения и записи файлов, скрывая сложность реальных операций с файловой системой.
//Веб-разработка: фреймворки часто используют фасады для упрощения работы с запросами и ответами HTTP.
//Приложения для работы с базами данных могут использовать фасад для предоставления простого интерфейса для выполнения запросов и обработки результатов, скрывая сложность SQL и работы с соединениями.
//Фасад упрощает взаимодействие с системой, но его следует использовать осторожно, чтобы не потерять гибкость и не скрыть полезные возможности подсистемы.
