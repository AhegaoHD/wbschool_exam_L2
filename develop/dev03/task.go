package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func sortLines(lines []string, columnIndex int, numericSort bool, reverse bool, unique bool) []string {
	sortFunc := func(i, j int) bool {
		// Разделение строк на колонки
		splitI := strings.Fields(lines[i])
		splitJ := strings.Fields(lines[j])

		// Обработка ключа -k
		var strI, strJ string
		if columnIndex < len(splitI) && columnIndex < len(splitJ) {
			strI = splitI[columnIndex]
			strJ = splitJ[columnIndex]
		} else {
			return false
		}

		// Обработка ключа -n
		if numericSort {
			intI, errI := strconv.Atoi(strI)
			intJ, errJ := strconv.Atoi(strJ)
			if errI == nil && errJ == nil {
				return intI < intJ
			}
		}
		// Сортировка как строки по умолчанию
		return strI < strJ
	}

	// Выполняем сортировку
	if reverse {
		sort.SliceStable(lines, func(i, j int) bool {
			return sortFunc(j, i)
		})
	} else {
		sort.SliceStable(lines, sortFunc)
	}

	// Обработка ключа -u
	if unique {
		uniqueLines := make([]string, 0, len(lines))
		seen := make(map[string]bool)
		for _, line := range lines {
			if !seen[line] {
				seen[line] = true
				uniqueLines = append(uniqueLines, line)
			}
		}
		return uniqueLines
	}

	return lines
}

func main() {
	// Парсинг аргументов командной строки
	column := flag.Int("k", 0, "column index for sorting")
	numeric := flag.Bool("n", false, "numeric sort")
	reverse := flag.Bool("r", false, "reverse sort")
	unique := flag.Bool("u", false, "unique sort")
	flag.Parse()

	// Чтение файла
	fileName := flag.Arg(0)
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка открытия файла: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var lines []string

	// Чтение строк файла
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "Ошибка чтения файла: %v\n", err)
			os.Exit(1)
		}
		lines = append(lines, strings.TrimRight(line, "\n"))
	}

	// Сортировка
	sortedLines := sortLines(lines, *column, *numeric, *reverse, *unique)

	// Вывод отсортированных строк
	for _, line := range sortedLines {
		fmt.Println(line)
	}
}
