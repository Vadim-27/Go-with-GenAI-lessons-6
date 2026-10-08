# Task 1 — Todo-list persistence (JSON)

Implement `SaveTodos` and `LoadTodos` in `todo.go`.

**Requirements**

- `SaveTodos(path string, todos []Todo) error` writes the list to `path`
  as JSON, creating or overwriting the file.
- `LoadTodos(path string) ([]Todo, error)` reads and parses the file at
  `path`, returning a clear error for a missing file or malformed JSON.
- At least **80% statement coverage** for this package
  (`go test -cover ./todo/...`).

**Do not edit `todo_test.go`** — it's the specification for this task.
Run it locally with:

```bash
go test -v ./todo/...
go test -cover ./todo/...
```


# Завдання 1 — збереження списку справ (JSON)

Реалізуйте `SaveTodos` та `LoadTodos` у файлі `todo.go`.

**Вимоги**

- `SaveTodos(path string, todos []Todo) error` записує список у `path`
  у форматі JSON, створюючи файл або перезаписуючи наявний.
- `LoadTodos(path string) ([]Todo, error)` читає та розбирає файл за
  шляхом `path`, повертаючи зрозумілу помилку, якщо файлу немає або JSON
  некоректний.
- Щонайменше **80% покриття операторів** для цього пакета
  (`go test -cover ./todo/...`).

**Не редагуйте `todo_test.go`** — це специфікація цього завдання.
Запускайте його локально так:

```bash
go test -v ./todo/...
go test -cover ./todo/...
```
