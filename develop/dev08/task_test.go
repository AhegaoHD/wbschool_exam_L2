package main

import (
	"os"
	"testing"
)

func TestChangeDirectory(t *testing.T) {
	// Сохраняем текущий рабочий каталог для восстановления после теста
	originalWd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Cannot get current working directory: %s", err)
	}
	defer os.Chdir(originalWd)

	// Тестируем изменение директории
	testDir := "/"
	if err := os.Chdir(testDir); err != nil {
		t.Errorf("cd failed: %s", err)
	}

	// Проверяем, что директория действительно изменилась
	newWd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Cannot get new working directory: %s", err)
	}

	if newWd != testDir {
		t.Errorf("cd failed: expected new directory to be %s, got %s", testDir, newWd)
	}
}
