package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
	"time"
)

/*
=== Базовая задача ===

Создать программу печатающую точное время с использованием NTP библиотеки.Инициализировать как go module.
Использовать библиотеку https://github.com/beevik/ntp.
Написать программу печатающую текущее время / точное время с использованием этой библиотеки.

Программа должна быть оформлена с использованием как go module.
Программа должна корректно обрабатывать ошибки библиотеки: распечатывать их в STDERR и возвращать ненулевой код выхода в OS.
Программа должна проходить проверки go vet и golint.
*/

func main() {
	// Получаем точное время с помощью NTP
	response, err := ntp.Query("pool.ntp.org")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка при получении времени:", err)
		os.Exit(1)
	}

	// Выравниваем время относительно локальных часов
	err = response.Validate()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка при валидации NTP ответа:", err)
		os.Exit(1)
	}

	// Время после синхронизации с NTP сервером
	exactTime := time.Now().Add(response.ClockOffset)

	fmt.Printf("Текущее время: %v\n", time.Now())
	fmt.Printf("Точное время: %v\n", exactTime)
}
