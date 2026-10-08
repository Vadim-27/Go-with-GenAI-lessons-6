// Package todo надає функції для збереження простого списку справ на диск
// у форматі JSON.
//
// Домашнє завдання — Завдання 1 (Урок 6: Робота з файлами, JSON та тестування):
// Реалізуйте SaveTodos і LoadTodos нижче так, щоб усі тести в
// todo_test.go проходили, і досягніть щонайменше 80% покриття операторів
// для цього пакета (перевіряється автоматично в CI — див. README
// репозиторію, як запустити це локально).
package todo

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// Todo представляє один елемент списку справ.
type Todo struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Done      bool      `json:"done"`
	CreatedAt time.Time `json:"created_at"`
}

// SaveTodos записує передані справи у файл за шляхом path у форматі JSON,
// створюючи файл, якщо його немає, і перезаписуючи, якщо він існує.
//
// TODO: реалізуйте цю функцію.
//   - Серіалізуйте todos у JSON (json.Marshal або json.MarshalIndent).
//   - Запишіть результат у path (найпростіший варіант — os.WriteFile).
//   - Обгортайте будь-яку помилку контекстом через fmt.Errorf("...: %w", err).
func SaveTodos(path string, todos []Todo) error {
	jsonBytes, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		return fmt.Errorf("SaveTodos: failed to marshal todos: %w", err)
	}
	err = os.WriteFile(path, jsonBytes, 0644)
	if err != nil {
		return fmt.Errorf("SaveTodos: failed to write file: %w", err)
	}
	return nil
}

// LoadTodos читає та розбирає список справ, збережений за шляхом path.
//
// TODO: реалізуйте цю функцію.
//   - Прочитайте файл за path (os.ReadFile).
//   - Десеріалізуйте JSON-байти у []Todo.
//   - Обгортайте помилки так, щоб викликач міг відрізнити «файл відсутній»
//     від «некоректний JSON» під час аналізу помилки (наприклад, через
//     errors.Is з os.ErrNotExist або перевіркою на *json.SyntaxError).
func LoadTodos(path string) ([]Todo, error) {
	jsonBytes, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("LoadTodos: file not found: %w", err)
		}
		return nil, fmt.Errorf("LoadTodos: failed to read file: %w", err)
	}

	var todos []Todo
	err = json.Unmarshal(jsonBytes, &todos)
	if err != nil {
		return nil, fmt.Errorf("LoadTodos: failed to unmarshal JSON: %w", err)
	}

	return todos, nil
}
