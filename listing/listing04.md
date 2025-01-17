Что выведет программа? Объяснить вывод программы.

```go
package main

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	for n := range ch {
		println(n)
	}
}
```

Ответ:
```
от 0 до 9

```
Программа в текущем виде вызовет deadlock (взаимную блокировку), потому что канал ch никогда не закрывается, и основная горутина будет бесконечно ждать новых значений из канала после того, как горутина, запущенная в go func(), завершит отправку 10 значений.

Каналы в Go предоставляют способ общения между горутинами. Операция чтения из канала блокируется, пока не будет доступно значение для чтения. Операция записи в канал блокируется, пока другая горутина не прочитает значение из канала. Ключевое слово range используется для чтения из канала в цикле до тех пор, пока канал не будет закрыт и пуст.

Чтобы программа корректно завершилась, нужно закрыть канал после завершения отправки значений в горутине. Вот как должен выглядеть исправленный код:
