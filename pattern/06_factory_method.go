package pattern

import "fmt"

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

// Product определяет интерфейс продукта
type Product interface {
	Use() string
}

// Creator определяет интерфейс создателя, который возвращает продукт
type Creator interface {
	CreateProduct() Product
}

// ConcreteProductA конкретная реализация продукта A
type ConcreteProductA struct{}

func (p *ConcreteProductA) Use() string {
	return "Product A is used"
}

// ConcreteCreatorA конкретная реализация создателя для продукта A
type ConcreteCreatorA struct{}

func (c *ConcreteCreatorA) CreateProduct() Product {
	return &ConcreteProductA{}
}

// ConcreteProductB конкретная реализация продукта B
type ConcreteProductB struct{}

func (p *ConcreteProductB) Use() string {
	return "Product B is used"
}

// ConcreteCreatorB конкретная реализация создателя для продукта B
type ConcreteCreatorB struct{}

func (c *ConcreteCreatorB) CreateProduct() Product {
	return &ConcreteProductB{}
}

// Пример использования паттерна "Фабричный метод"
func main() {
	var creator Creator

	// Создаем продукт A
	creator = &ConcreteCreatorA{}
	productA := creator.CreateProduct()
	fmt.Println(productA.Use())

	// Создаем продукт B
	creator = &ConcreteCreatorB{}
	productB := creator.CreateProduct()
	fmt.Println(productB.Use())
}

//Применимость паттерна «Фабричный метод»:
//
//Когда класс не может предвидеть класс объектов, которые он должен создавать.
//Когда классы делегируют ответственность одной из нескольких вспомогательных подклассов, и вы хотите локализовать знание о том, какой вспомогательный подкласс делегат.
//Плюсы:
//
//Упрощает добавление новых продуктов в программу.
//Изолирует код создания продуктов от их использования.
//Реализует принцип открытости/закрытости.
//Минусы:
//
//Может привести к созданию большого числа маленьких классов, так как для каждого продукта нужен свой создатель.
//Примеры использования на практике:
//
//Фреймворки и библиотеки: когда фреймворк должен стандартизировать основной процесс, но детали реализации оставить подклассам.
//API для работы с различными типами ресурсов: например, классы, отвечающие за создание соединений с базами данных различных типов — MySQL, PostgreSQL, SQLite — могут использовать фабричный метод для инкапсуляции логики создания соединения.
