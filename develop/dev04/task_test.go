package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	tests := []struct {
		name string
		dict []string
		want map[string][]string
	}{
		{
			name: "basic anagrams",
			dict: []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"},
			want: map[string][]string{
				"пятак":  {"пятак", "пятка", "тяпка"},
				"листок": {"листок", "слиток", "столик"},
			},
		},
		{
			name: "no anagrams",
			dict: []string{"пятак", "тест", "слово"},
			want: map[string][]string{},
		},
		{
			name: "case insensitivity",
			dict: []string{"пЯтак", "ПяТка", "тЯпка"},
			want: map[string][]string{
				"пятак": {"пятак", "пятка", "тяпка"},
			},
		},
		{
			name: "including single words",
			dict: []string{"пятак", "тест", "тяпка", "тесто"},
			want: map[string][]string{
				"пятак": {"пятак", "тяпка"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findAnagrams(tt.dict); !reflect.DeepEqual(got, &tt.want) {
				t.Errorf("findAnagrams() = %v, want %v", got, &tt.want)
			}
		})
	}
}

// Вспомогательная функция для сортировки мапы перед сравнением результатов
func sortAnagramMap(anagramMap map[string][]string) map[string][]string {
	for _, words := range anagramMap {
		sort.Strings(words)
	}
	return anagramMap
}
