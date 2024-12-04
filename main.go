package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Текстовый объект
type TextEntity struct {
	Type    string
	Text    string
	X, Y, Z float64
}

func main() {
	// Путь к DXF файлу
	filePath := "ex.dxf"

	// Открываем файл
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Failed to open file: ", err)
	}
	defer file.Close()

	// Читаем файл построчно
	var textObjects []TextEntity
	var currentEntity string
	var currentText string
	var x, y, z float64

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Ищем начало новой сущности
		if line == "0" {
			// Сохраняем текстовый объект, если он был найден
			if currentEntity == "TEXT" || currentEntity == "MTEXT" {
				textObjects = append(textObjects, TextEntity{
					Type: currentEntity,
					Text: currentText,
					X:    x,
					Y:    y,
					Z:    z,
				})
			}
			// Сбрасываем временные переменные
			currentEntity = ""
			currentText = ""
			x, y, z = 0, 0, 0
			continue
		}

		// Устанавливаем текущую сущность
		if currentEntity == "" {
			currentEntity = line
			continue
		}

		// Обрабатываем параметры текста
		switch currentEntity {
		case "TEXT", "MTEXT":
			if line == "1" { // Код группы 1 — текстовое содержимое
				if scanner.Scan() {
					currentText = scanner.Text()
				}
			} else if line == "10" { // Код группы 10 — X координата
				if scanner.Scan() {
					fmt.Sscanf(scanner.Text(), "%f", &x)
				}
			} else if line == "20" { // Код группы 20 — Y координата
				if scanner.Scan() {
					fmt.Sscanf(scanner.Text(), "%f", &y)
				}
			} else if line == "30" { // Код группы 30 — Z координата
				if scanner.Scan() {
					fmt.Sscanf(scanner.Text(), "%f", &z)
				}
			}
		}
	}

	// Проверяем на ошибки при чтении
	if err := scanner.Err(); err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	// Выводим найденные текстовые объекты
	for i, obj := range textObjects {
		fmt.Printf("Type: %s, Text: %s, X: %f, Y: %f, Z: %f\n",
			obj.Type, obj.Text, obj.X, obj.Y, obj.Z)
	}
}
