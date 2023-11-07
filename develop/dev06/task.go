package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	// Определение флагов
	fields := flag.String("f", "", "Choose fields (columns)")
	delimiter := flag.String("d", "\t", "Use different delimiter")
	onlySeparated := flag.Bool("s", false, "Only lines with delimiter")

	flag.Parse()

	// Вызов функции cutCommand для обработки входных данных
	err := cutCommand(os.Stdin, os.Stdout, *fields, *delimiter, *onlySeparated)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}
func cutCommand(input io.Reader, output io.Writer, fieldsArg string, delimiter string, onlySep bool) error {
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()

		// Если включен флаг -s и в строке нет разделителя, пропускаем строку
		if onlySep && !strings.Contains(line, delimiter) {
			continue
		}

		// Разделение строки по разделителю
		columns := strings.Split(line, delimiter)

		// Обработка флага -f, выборка колонок
		if fieldsArg != "" {
			fieldIndexes, err := parseFields(fieldsArg)
			if err != nil {
				return fmt.Errorf("invalid fields argument: %v", err)
			}

			selectedFields := selectFields(columns, fieldIndexes)
			fmt.Fprintln(output, strings.Join(selectedFields, delimiter))
		} else {
			fmt.Fprintln(output, line)
		}
	}

	return scanner.Err()
}

// parseFields парсит строку с индексами полей и возвращает их в виде среза int
func parseFields(fieldsArg string) ([]int, error) {
	if fieldsArg == "" {
		return []int{}, nil
	}
	fieldStrs := strings.Split(fieldsArg, ",")
	fields := make([]int, 0, len(fieldStrs))
	for _, f := range fieldStrs {
		var i int
		_, err := fmt.Sscanf(f, "%d", &i)
		if err != nil {
			return nil, fmt.Errorf("parse error: %v", err)
		}
		fields = append(fields, i-1) // в Unix-стиле колонки начинаются с 1, в Go с 0
	}
	return fields, nil
}

// selectFields возвращает только выбранные колонки
func selectFields(columns []string, fieldIndexes []int) []string {
	selected := make([]string, 0, len(fieldIndexes))
	for _, idx := range fieldIndexes {
		if idx >= 0 && idx < len(columns) {
			selected = append(selected, columns[idx])
		}
	}
	return selected
}
