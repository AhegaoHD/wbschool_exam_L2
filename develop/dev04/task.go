package main

import (
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// Функция для сортировки строки
func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

// Функция поиска всех множеств анаграмм по словарю
func findAnagrams(dict []string) *map[string][]string {
	anagrams := make(map[string][]string)
	visited := make(map[string]bool)

	for _, word := range dict {
		// Приведение слова к нижнему регистру
		word = strings.ToLower(word)
		sortedWord := sortString(word)

		if visited[word] {
			continue
		}
		visited[word] = true

		// Если ключ для отсортированного слова существует, добавляем слово в множество
		if _, exists := anagrams[sortedWord]; exists {
			anagrams[sortedWord] = append(anagrams[sortedWord], word)
		} else {
			anagrams[sortedWord] = []string{word}
		}
	}

	// Удаление множеств, содержащих только одно слово
	for key, group := range anagrams {
		if len(group) < 2 {
			delete(anagrams, key)
		} else {
			// Сортировка слов в множестве по возрастанию
			sort.Strings(anagrams[key])
		}
	}

	// Формирование итоговой мапы с первым словом в качестве ключа
	result := make(map[string][]string)
	for _, group := range anagrams {
		sort.Strings(group)
		key := group[0]
		result[key] = group
	}

	return &result
}
