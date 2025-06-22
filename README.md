# Task Management API

Простое REST API для управления задачами (tasks) с возможностью создания, получения и удаления.

## 🚀 Запуск проекта

### Сборка и запуск
```bash
# Запуск напрямую
go run cmd/main.go

# Или сборка в бинарный файл
go build -o "main" cmd/main.go
./main
```

### Ожидаемый вывод
Сервер запускается на `localhost:8080` (по умолчанию).  
Логи выводятся в консоль.

---

## 📌 Endpoints

### 1. Создание задачи (`POST`)
**URL:** `http://localhost:8080/api/v1/task`  
**Пример**
```bash
curl -X POST http://localhost:8080/api/v1/create_task \
  -H "Content-Type: application/json" \
  -d '{"Title": "my_task"}'
```
**Request:**
```json
{
    "Title": "some_title"
}
```
**Response:**
```json
{
    "ID": 1,
    "Title": "some_title",
    "Status": "completed",
    "CreatedAt": "11:11:11 0000-00-00",
    "RunTime": "12:00:00"
}
```

---

### 2. Получение задачи (`GET`)
**URL:** `http://localhost:8080/api/v1/task/{id}`  
**Пример:**
```bash
curl http://localhost:8080/api/v1/task/1
```
**Response:**
```json
{
    "ID": 1,
    "Title": "some_title",
    "Status": "completed",
    "CreatedAt": "11:11:11 0000-00-00",
    "RunTime": "12:00:00"
}
```

---

### 3. Удаление задачи (`DELETE`)
**URL:** `http://localhost:8080/api/v1/task/{id}`  
**Пример:**
```bash
curl -X DELETE http://localhost:8080/api/v1/task/1
```
**Response:**
```json
{
    "ID": 1,
    "Title": "some_title",
    "Status": "completed",
    "CreatedAt": "11:11:11 0000-00-00",
    "RunTime": "12:00:00"
}
```

---

## 📊 Структура задачи (Task)
| Поле        | Тип     | Описание                  | Пример значения                      |
|-------------|---------|---------------------------|--------------------------------------|
| `ID`        | `int`   | Уникальный ID задачи      | `1`                                  |
| `Title`     | `string`| Название задачи           | `"some_title"`                       |
| `Status`    | `string`| Статус (`pending` `in_progress``completed` и др.)| `"completed"` |
| `CreatedAt` | `string`| Время создания            | `"11:11:11 0000-00-00"`              |
| `RunTime`   | `string`| Время выполнения          | `"12:00:00"`                         |

---

## 🛠 Технологии
- **Язык:** `Go`
- **Роутинг:** `gorrila/mux`
- **Логирование:** `logrus`
- **Ститль архитектуры** `Чистая архитектура с доменной областью`
