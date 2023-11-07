package pattern

import "fmt"

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

// Product определяет объект, который будет создаваться
type Product struct {
	partA string
	partB string
	partC string
}

// Builder предоставляет интерфейс для создания частей объекта Product
type Builder interface {
	BuildPartA()
	BuildPartB()
	BuildPartC()
	GetResult() Product
}

// ConcreteBuilder реализует интерфейс Builder и строит части продукта
type ConcreteBuilder struct {
	product Product
}

func (b *ConcreteBuilder) BuildPartA() {
	b.product.partA = "PartA"
}

func (b *ConcreteBuilder) BuildPartB() {
	b.product.partB = "PartB"
}

func (b *ConcreteBuilder) BuildPartC() {
	b.product.partC = "PartC"
}

func (b *ConcreteBuilder) GetResult() Product {
	return b.product
}

// Director определяет порядок вызова шагов построения
type Director struct {
	builder Builder
}

func NewDirector(b Builder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) Construct() {
	d.builder.BuildPartA()
	d.builder.BuildPartB()
	d.builder.BuildPartC()
}

func main() {
	builder := &ConcreteBuilder{}
	director := NewDirector(builder)
	director.Construct()
	product := builder.GetResult()

	fmt.Printf("Product Parts: %+v\n", product)
}

//Применимость паттерна «Строитель»:
//
//Когда процесс создания объекта не должен зависеть от его составляющих частей и их сборки.
//Когда нужно создавать объекты с различными представлениями.

//Плюсы:
//
//Позволяет изменять внутреннее представление продукта.
//Изолирует код построения продукта от его бизнес-логики.
//Дает более тонкий контроль над процессом конструирования, чем другие паттерны создания объектов.

//Минусы:
//
//Может усложнить код из-за введения дополнительных классов.
//Клиент должен знать различные строители, если ему нужны разные представления объекта.

//Примеры использования на практике:
//
//Компоновка сложных объектов в играх (например, создание различных уровней, персонажей).
//SQL запросы с помощью ORM, где строитель позволяет пошагово добавлять условия, поля, таблицы и так далее.
//Построение сложных объектов в GUI-библиотеках, например, сложных меню с множеством подменю и элементов.
//Строитель обеспечивает ясность и гибкость в процессе создания сложных объектов, позволяя построить различные представления объекта с той же построительной логикой.
