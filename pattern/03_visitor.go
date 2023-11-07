package pattern

import "fmt"

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

// Component интерфейс для элементов, которые могут быть посещены
type Component interface {
	Accept(v Visitor)
}

// ConcreteComponentA один из элементов
type ConcreteComponentA struct {
	// поля, специфичные для ConcreteComponentA
}

// Accept реализация метода Accept для ConcreteComponentA
func (c *ConcreteComponentA) Accept(v Visitor) {
	v.VisitConcreteComponentA(c)
}

// ExclusiveMethodOfConcreteComponentA метод, специфичный для ConcreteComponentA
func (c *ConcreteComponentA) ExclusiveMethodOfConcreteComponentA() string {
	return "A"
}

// ConcreteComponentB другой элемент
type ConcreteComponentB struct {
	// поля, специфичные для ConcreteComponentB
}

// Accept реализация метода Accept для ConcreteComponentB
func (c *ConcreteComponentB) Accept(v Visitor) {
	v.VisitConcreteComponentB(c)
}

// SpecialMethodOfConcreteComponentB метод, специфичный для ConcreteComponentB
func (c *ConcreteComponentB) SpecialMethodOfConcreteComponentB() string {
	return "B"
}

// Visitor интерфейс посетителя, объявляет набор методов посещения для каждого конкретного компонента
type Visitor interface {
	VisitConcreteComponentA(*ConcreteComponentA)
	VisitConcreteComponentB(*ConcreteComponentB)
}

// ConcreteVisitor1 конкретный посетитель, реализующий интерфейс Visitor
type ConcreteVisitor1 struct{}

func (v *ConcreteVisitor1) VisitConcreteComponentA(c *ConcreteComponentA) {
	fmt.Println("ConcreteVisitor1: ComponentA ->", c.ExclusiveMethodOfConcreteComponentA())
}

func (v *ConcreteVisitor1) VisitConcreteComponentB(c *ConcreteComponentB) {
	fmt.Println("ConcreteVisitor1: ComponentB ->", c.SpecialMethodOfConcreteComponentB())
}

// Пример использования паттерна "Посетитель"
func main() {
	components := []Component{
		&ConcreteComponentA{},
		&ConcreteComponentB{},
	}

	visitor := &ConcreteVisitor1{}
	for _, comp := range components {
		comp.Accept(visitor)
	}
}

//Применимость паттерна «Посетитель»:
//
//Когда нужно выполнить операцию на целом комплексе объектов с разнородными классами.
//Когда новые операции должны быть добавлены в библиотеку классов, и при этом нежелательно менять код самих классов.

//Плюсы:
//
//Упрощает добавление операций, работающих со сложными структурами объектов.
//Собирает родственные операции и разделяет несвязанные.
//Посетитель может накапливать состояние при обходе структуры элементов.

//Минусы:
//
//Может привести к нарушению инкапсуляции элементов, если посетитель должен иметь доступ к их внутреннему состоянию.
//Паттерн сложно применять к иерархиям классов, часто изменяющимся, так как любое добавление или удаление класса ведет к изменению всех посетителей.

//Примеры использования на практике:
//
//Обработка XML или JSON объектов с различной структурой: посетитель может использоваться для навигации по структуре и выполнения операций, таких как сериализация или валидация.
