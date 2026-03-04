# Books API

REST API для управления книгами, реализованный на **Go + Fiber + MySQL**.

Проект демонстрирует:

* Clean Architecture
* работу с **Docker**
* **автоматические миграции**
* **Swagger документацию**
* **middleware валидацию**
* **централизованные ошибки**
* **структурированное логирование**

---

# 🚀 Технологии

* **Go**
* **Fiber**
* **MySQL**
* **Docker / Docker Compose**
* **golang-migrate**
* **Swagger (OpenAPI)**
* **Clean Architecture**
* **Centralized logging**

---

# 📁 Структура проекта

```
BooksAPI
│
├── cmd/app                # точка входа приложения
│
├── internal
│   ├── config             # загрузка конфигурации
│   ├── database           # подключение к БД + миграции
│   ├── domain             # доменные модели
│   ├── errors             # централизованные ошибки
│   ├── http               # handlers / routes / middleware
│   │   └── dto            # DTO для HTTP слоя
│   ├── repository         # работа с БД
│   └── service            # бизнес логика
│
├── migrations             # SQL миграции
├── docs                   # swagger документация
│
├── pkg
│   └── logger             # общий логгер приложения
│
├── curl_examples.txt      # примеры тестирования API
├── docker-compose.yml
├── Dockerfile
├── Makefile
├── .env-example
└── README.md
```

---

# 🧱 Архитектура

Проект реализован по принципу **Clean Architecture**.

```
HTTP Layer (handlers)
        ↓
Service Layer (business logic)
        ↓
Repository Layer (database)
        ↓
MySQL
```

### HTTP слой

* обработка запросов
* DTO
* middleware
* маршрутизация

### Service слой

* бизнес логика
* валидация данных

### Repository слой

* работа с БД
* SQL запросы
* обработка ошибок БД

---

# ⚙️ Конфигурация

Создайте `.env` на основе `.env-example`.

Пример:

```
APP_PORT=3000

DB_HOST=mysql
DB_PORT=3306
DB_USER=root
DB_PASSWORD=root
DB_NAME=books_db
```

---

# 🐳 Запуск через Docker

### Запуск приложения

```
make start
```

### Остановка

```
make down
```

### Логи

```
make logs
```

### Перезапуск

```
make restart
```

---

# 📚 Swagger документация

Сгенерировать swagger:

```
make swagger
```

Swagger UI:

```
http://localhost:3000/swagger/index.html
```

---

# 📡 API Endpoints

Base URL:

```
http://localhost:3000/api/v1
```

---

## Create book

```
POST /books
```

Request

```json
{
  "title": "Clean Code",
  "author": "Robert C. Martin",
  "isbn": "9780132350884"
}
```

Response **201**

```json
{
  "id": 1,
  "title": "Clean Code",
  "author": "Robert C. Martin",
  "isbn": "9780132350884"
}
```

Response **400**

```json
{
  "error": "validation",
  "fields": {
    "title": "required"
  }
}
```

---

## Get book

```
GET /books/{id}
```

Response **200**

```json
{
  "id": 1,
  "title": "Clean Code",
  "author": "Robert C. Martin",
  "isbn": "9780132350884"
}
```

Response **404**

```json
{
  "error": "not_found"
}
```

---

## Update book

```
PUT /books/{id}
```

Request

```json
{
  "title": "Clean Code (2nd Edition)",
  "author": "Robert C. Martin",
  "isbn": "9780132350884"
}
```

Response **200**

```json
{
  "id": 1,
  "title": "Clean Code (2nd Edition)",
  "author": "Robert C. Martin",
  "isbn": "9780132350884"
}
```

Response **400**

```json
{
  "error": "validation"
}
```

Response **404**

```json
{
  "error": "not_found"
}
```

---

## Delete book

```
DELETE /books/{id}
```

Response **200**

```json
{
  "status": "ok"
}
```

Response **404**

```json
{
  "error": "not_found"
}
```

---

# 🧪 Тестирование API

В проекте есть файл:

```
curl_examples.txt
```

Он содержит последовательность **curl-запросов**, демонстрирующих работу API:

* Create
* Validation error
* Get
* Update
* Delete

---

# 📊 Логирование (Logging)

В проекте используется единый логгер (`pkg/logger`), который применяется на всех уровнях приложения.

Логирование реализовано на нескольких уровнях приложения:

- **HTTP уровень**
    - логирование всех входящих запросов
    - используется middleware `fiber/middleware/logger`

- **Application уровень**
    - запуск сервера
    - загрузка конфигурации
    - подключение к базе данных
    - запуск миграций

- **Service уровень**
    - бизнес операции (создание, обновление, удаление книг)

- **Repository уровень**
    - SQL операции
    - ошибки базы данных

### Поток логирования

```
HTTP Request
↓
Fiber Logger Middleware
↓
Handler
↓
Service (business logs)
↓
Repository (database logs)
```

Пример логов при запуске и работе API:

```
books-api    | 2026/03/04 03:10:54 main.go:26: starting Books API
books-api    | 2026/03/04 03:10:54 config.go:17: loading configuration
books-api    | 2026/03/04 03:10:54 config.go:21: .env file not found, using environment variables
books-api    | 2026/03/04 03:10:54 config.go:33: configuration loaded (port=3000)
books-api    | 2026/03/04 03:10:54 main.go:29: config loaded
books-api    | 2026/03/04 03:10:54 mysql.go:12: initializing MySQL connection
books-api    | 2026/03/04 03:10:54 mysql.go:19: mysql ping error: dial tcp 172.27.0.2:3306: connect: connection refused
books-api    | 2026/03/04 03:10:55 main.go:26: starting Books API
books-api    | 2026/03/04 03:10:55 config.go:17: loading configuration
books-api    | 2026/03/04 03:10:55 config.go:21: .env file not found, using environment variables
books-api    | 2026/03/04 03:10:55 config.go:33: configuration loaded (port=3000)
books-api    | 2026/03/04 03:10:55 main.go:29: config loaded
books-api    | 2026/03/04 03:10:55 mysql.go:12: initializing MySQL connection
books-api    | 2026/03/04 03:10:55 mysql.go:19: mysql ping error: dial tcp 172.27.0.2:3306: connect: connection refused
books-api    | 2026/03/04 03:10:56 main.go:26: starting Books API
books-api    | 2026/03/04 03:10:56 config.go:17: loading configuration
books-api    | 2026/03/04 03:10:56 config.go:21: .env file not found, using environment variables
books-api    | 2026/03/04 03:10:56 config.go:33: configuration loaded (port=3000)
books-api    | 2026/03/04 03:10:56 main.go:29: config loaded
books-api    | 2026/03/04 03:10:56 mysql.go:12: initializing MySQL connection
books-api    | 2026/03/04 03:10:56 mysql.go:21: MySQL connection established
books-api    | 2026/03/04 03:10:56 main.go:32: database connected
books-api    | 2026/03/04 03:10:56 migrate.go:16: running database migrations
books-api    | 2026/03/04 03:10:56 migrate.go:36: no new migrations
books-api    | 2026/03/04 03:10:56 main.go:50: server started on port 3000
books-api    | 
books-api    |  ┌───────────────────────────────────────────────────┐ 
books-api    |  │                  Fiber v2.52.12                   │ 
books-api    |  │               http://127.0.0.1:3000               │ 
books-api    |  │       (bound on host 0.0.0.0 and port 3000)       │ 
books-api    |  │                                                   │ 
books-api    |  │ Handlers ............ 12  Processes ........... 1 │ 
books-api    |  │ Prefork ....... Disabled  PID ................. 1 │ 
books-api    |  └───────────────────────────────────────────────────┘ 
books-api    | 
books-api    | 03:12:36 | 200 |     790.226µs | 172.27.0.1 | GET | index.html | -
books-api    | 03:12:36 | 200 |     184.928µs | 172.27.0.1 | GET | swagger-ui.css | -
books-api    | 03:12:36 | 200 |      74.716µs | 172.27.0.1 | GET | swagger-ui-bundle.js | -
books-api    | 03:12:36 | 200 |     114.806µs | 172.27.0.1 | GET | swagger-ui-standalone-preset.js | -
books-api    | 03:12:36 | 200 |     357.179µs | 172.27.0.1 | GET | doc.json | -
books-api    | 03:12:36 | 200 |      25.018µs | 172.27.0.1 | GET | favicon-32x32.png | -
books-api    | 03:12:39 | 204 |     120.114µs | 172.27.0.1 | OPTIONS | /api/v1/books | -
books-api    | 2026/03/04 03:12:39 book_service_impl.go:21: creating book: title=Clean Code author=Robert C. Martin
books-api    | 2026/03/04 03:12:39 book_repository_impl.go:22: repository: creating book title=Clean Code author=Robert C. Martin
books-api    | 2026/03/04 03:12:39 book_repository_impl.go:44: repository: book created id=3
books-api    | 03:12:39 | 201 |    26.16982ms | 172.27.0.1 | POST | /api/v1/books | -
books-api    | 03:12:55 | 204 |      26.139µs | 172.27.0.1 | OPTIONS | /api/v1/books | -
books-api    | 03:12:55 | 400 |     120.979µs | 172.27.0.1 | POST | /api/v1/books | -
books-api    | 2026/03/04 03:13:10 book_service_impl.go:35: getting book id=1
books-api    | 2026/03/04 03:13:10 book_repository_impl.go:51: repository: getting book id=1
books-api    | 2026/03/04 03:13:10 book_repository_impl.go:68: repository: book not found id=1
books-api    | 03:13:10 | 404 |    1.101446ms | 172.27.0.1 | GET | /api/v1/books/1 | -
books-api    | 2026/03/04 03:13:15 book_service_impl.go:35: getting book id=2
books-api    | 2026/03/04 03:13:15 book_repository_impl.go:51: repository: getting book id=2
books-api    | 03:13:15 | 200 |     892.266µs | 172.27.0.1 | GET | /api/v1/books/2 | -
books-api    | 2026/03/04 03:13:23 book_service_impl.go:35: getting book id=3
books-api    | 2026/03/04 03:13:23 book_repository_impl.go:51: repository: getting book id=3
books-api    | 03:13:23 | 200 |    2.030569ms | 172.27.0.1 | GET | /api/v1/books/3 | -
books-api    | 2026/03/04 03:13:25 book_service_impl.go:35: getting book id=787
books-api    | 2026/03/04 03:13:25 book_repository_impl.go:51: repository: getting book id=787
books-api    | 2026/03/04 03:13:25 book_repository_impl.go:68: repository: book not found id=787
books-api    | 03:13:25 | 404 |    1.119364ms | 172.27.0.1 | GET | /api/v1/books/787 | -
books-api    | 2026/03/04 03:13:27 book_service_impl.go:35: getting book id=0
books-api    | 2026/03/04 03:13:27 book_repository_impl.go:51: repository: getting book id=0
books-api    | 2026/03/04 03:13:27 book_repository_impl.go:68: repository: book not found id=0
books-api    | 03:13:27 | 404 |    1.047917ms | 172.27.0.1 | GET | /api/v1/books/0 | -
books-api    | 2026/03/04 03:13:32 book_service_impl.go:35: getting book id=-1
books-api    | 2026/03/04 03:13:32 book_repository_impl.go:51: repository: getting book id=-1
books-api    | 2026/03/04 03:13:32 book_repository_impl.go:68: repository: book not found id=-1
books-api    | 03:13:32 | 404 |    2.237574ms | 172.27.0.1 | GET | /api/v1/books/-1 | -
books-api    | 03:13:38 | 204 |      12.966µs | 172.27.0.1 | OPTIONS | /api/v1/books/2 | -
books-api    | 2026/03/04 03:13:38 book_service_impl.go:39: updating book id=2
books-api    | 2026/03/04 03:13:38 book_repository_impl.go:82: repository: updating book id=2
books-api    | 2026/03/04 03:13:38 book_repository_impl.go:108: repository: book updated id=2
books-api    | 03:13:38 | 200 |   10.467228ms | 172.27.0.1 | PUT | /api/v1/books/2 | -
books-api    | 03:13:42 | 204 |       12.82µs | 172.27.0.1 | OPTIONS | /api/v1/books/3 | -
books-api    | 2026/03/04 03:13:42 book_service_impl.go:39: updating book id=3
books-api    | 2026/03/04 03:13:42 book_repository_impl.go:82: repository: updating book id=3
books-api    | 2026/03/04 03:13:42 book_repository_impl.go:108: repository: book updated id=3
books-api    | 03:13:42 | 200 |   11.737154ms | 172.27.0.1 | PUT | /api/v1/books/3 | -
books-api    | 03:13:46 | 204 |       10.81µs | 172.27.0.1 | OPTIONS | /api/v1/books/4 | -
books-api    | 2026/03/04 03:13:46 book_service_impl.go:52: deleting book id=4
books-api    | 2026/03/04 03:13:46 book_repository_impl.go:115: repository: deleting book id=4
books-api    | 2026/03/04 03:13:46 book_repository_impl.go:134: repository: delete failed book not found id=4
books-api    | 03:13:46 | 404 |    2.270787ms | 172.27.0.1 | DELETE | /api/v1/books/4 | -
books-api    | 03:13:48 | 204 |      16.016µs | 172.27.0.1 | OPTIONS | /api/v1/books/3 | -
books-api    | 2026/03/04 03:13:48 book_service_impl.go:52: deleting book id=3
books-api    | 2026/03/04 03:13:48 book_repository_impl.go:115: repository: deleting book id=3
books-api    | 2026/03/04 03:13:48 book_repository_impl.go:138: repository: book deleted id=3
books-api    | 03:13:48 | 200 |   11.610399ms | 172.27.0.1 | DELETE | /api/v1/books/3 | -
```

---

# 📌 Автор

AlexTrav
