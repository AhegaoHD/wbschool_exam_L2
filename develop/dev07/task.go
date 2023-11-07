package main

import "reflect"

/*
=== Or channel ===

Реализовать функцию, которая будет объединять один или более done каналов в single канал если один из его составляющих каналов закроется.
Одним из вариантов было бы очевидно написать выражение при помощи select, которое бы реализовывало эту связь,
однако иногда неизестно общее число done каналов, с которыми вы работаете в рантайме.
В этом случае удобнее использовать вызов единственной функции, которая, приняв на вход один или более or каналов, реализовывала весь функционал.

Определение функции:
var or func(channels ...<- chan interface{}) <- chan interface{}

Пример использования функции:
sig := func(after time.Duration) <- chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
}()
return c
}

start := time.Now()
<-or (
	sig(2*time.Hour),
	sig(5*time.Minute),
	sig(1*time.Second),
	sig(1*time.Hour),
	sig(1*time.Minute),
)

fmt.Printf(“fone after %v”, time.Since(start))
*/

// or объединяет один или более done каналов в single канал.
var or func(channels ...<-chan interface{}) <-chan interface{} = func(channels ...<-chan interface{}) <-chan interface{} {
	switch len(channels) {
	case 0:
		// Если нет каналов, возвращаем закрытый канал
		c := make(chan interface{})
		close(c)
		return c
	case 1:
		// Если канал один, просто возвращаем его
		return channels[0]
	}

	orDone := make(chan interface{})
	go func() {
		defer close(orDone)
		var cases []reflect.SelectCase
		for _, ch := range channels {
			cases = append(cases, reflect.SelectCase{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(ch),
			})
		}

		// Ждём сигнала из любого канала
		reflect.Select(cases)
		// Как только получаем сигнал, выходной канал будет закрыт
	}()

	return orDone
}
