package main

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
  - "a4bc2d5e" => "aaaabccddddde"
  - "abcd" => "abcd"
  - "45" => "" (некорректная строка)
  - "" => ""

Дополнительное задание: поддержка escape - последовательностей
  - qwe\4\5 => qwe45 (*)
  - qwe\45 => qwe44444 (*)
  - qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/
func main() {}

func unpack(s string) (string, error) {
	var builder strings.Builder
	var prevRune rune
	escaped := false
	multiplier := ""

	for _, r := range s {
		// Если предыдущий символ был слешем, обрабатываем текущий как литерал
		if escaped {
			builder.WriteRune(r)
			prevRune = r // Сбрасываем предыдущий символ
			escaped = false
			continue
		}

		// Если текущий символ слеш, включаем режим экранирования
		if r == '\\' {
			if escaped {
				// Двойной слеш превращаем в один
				builder.WriteRune(r)
				prevRune = 0 // Сбрасываем предыдущий символ
			}
			escaped = true
			if multiplier != "" {
				count, err := strconv.Atoi(multiplier)

				if err != nil {
					return string(r), errors.New("некорректная строка")
				}
				// Добавляем предыдущий символ нужное количество раз
				builder.WriteString(strings.Repeat(string(prevRune), count-1))
				multiplier = ""
			}
			continue
		}

		// Если текущий символ - цифра
		if unicode.IsDigit(r) {
			if prevRune == 0 { // Если цифра идет первым символом, возвращаем ошибку
				return "", errors.New("некорректная строка")
			}
			multiplier += string(r)
		} else {
			if multiplier != "" {
				count, err := strconv.Atoi(multiplier)

				if err != nil {
					return string(r), errors.New("некорректная строка")
				}
				// Добавляем предыдущий символ нужное количество раз
				builder.WriteString(strings.Repeat(string(prevRune), count-1))
				multiplier = ""
			}
			builder.WriteRune(r)
			prevRune = r
		}
	}

	// Если строка заканчивается экранированием, возвращаем ошибку
	if escaped {
		return "", errors.New("некорректная строка")
	}

	// Обработка множителя для последнего символа
	if multiplier != "" {
		count, err := strconv.Atoi(multiplier)
		if err != nil {
			return multiplier, errors.New("некорректная строка")
		}
		builder.WriteString(strings.Repeat(string(prevRune), count-1))
	}

	return builder.String(), nil
}
