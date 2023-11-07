package pattern

import "fmt"

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

// Handler представляет интерфейс обработчика в цепочке
type Handler interface {
	HandleRequest(request string)
	SetNext(handler Handler)
}

// BaseHandler базовая структура обработчика
type BaseHandler struct {
	next Handler
}

// SetNext устанавливает следующий обработчик в цепочке
func (h *BaseHandler) SetNext(handler Handler) {
	h.next = handler
}

// HandleRequest по умолчанию передает запрос дальше по цепочке
func (h *BaseHandler) HandleRequest(request string) {
	if h.next != nil {
		h.next.HandleRequest(request)
	}
}

// ConcreteHandlerA конкретный обработчик A
type ConcreteHandlerA struct {
	BaseHandler
}

// HandleRequest реализация обработчика A
func (h *ConcreteHandlerA) HandleRequest(request string) {
	if request == "A" {
		fmt.Println("ConcreteHandlerA handled the request")
	} else {
		fmt.Println("ConcreteHandlerA passed the request")
		h.BaseHandler.HandleRequest(request)
	}
}

// ConcreteHandlerB конкретный обработчик B
type ConcreteHandlerB struct {
	BaseHandler
}

// HandleRequest реализация обработчика B
func (h *ConcreteHandlerB) HandleRequest(request string) {
	if request == "B" {
		fmt.Println("ConcreteHandlerB handled the request")
	} else {
		fmt.Println("ConcreteHandlerB passed the request")
		h.BaseHandler.HandleRequest(request)
	}
}

// Пример использования паттерна "Цепочка обязанностей"
func main() {
	handlerA := &ConcreteHandlerA{}
	handlerB := &ConcreteHandlerB{}

	handlerA.SetNext(handlerB) // Установка цепочки

	// Отправка запросов
	handlerA.HandleRequest("A")
	handlerA.HandleRequest("B")
	handlerA.HandleRequest("C")
}

//Применимость паттерна «Цепочка обязанностей»:
//
//Когда имеется более одного объекта, способного обработать запрос, и конкретный обработчик заранее неизвестен.
//Когда набор объектов, способных обработать запрос, должен динамически определяться.
//Плюсы:
//
//Уменьшает зависимость между клиентом и обработчиками.
//Реализует принцип единственной обязанности.
//Позволяет динамически добавлять или изменять обработчики в цепочке.
//Минусы:
//
//Запрос может остаться никем не обработанным.
//Усложняет архитектуру программы из-за множества дополнительных классов.
//Примеры использования на практике:
//
//Обработка событий в графическом пользовательском интерфейсе: событие передается от элемента к его родителю, пока не будет обработано.
//Middleware в веб-фреймворках: запросы обрабатываются последовательностью middleware функций до тех пор, пока одна из них не завершит обработку запроса.
//Проверка доступа: запрос на выполнение действия проходит через цепочку проверок прав, пока не будет одобрен или отклонен.
