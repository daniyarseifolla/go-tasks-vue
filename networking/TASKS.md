# Задача
Создать сервис управления пользователями с одновременной поддержкой gRPC и REST API для изучения различий между этими подходами.
## Требования к API

### REST Endpoints:

```text
POST   /users     - создание пользователя
GET    /users/{id} - получение пользователя
PUT    /users/{id} - обновление пользователя
DELETE /users/{id} - удаление пользователя
```
### gRPC Methods:

```proto
rpc CreateUser(CreateUserRequest) returns (CreateUserResponse)
rpc GetUser(GetUserRequest) returns (GetUserResponse)
rpc UpdateUser(UpdateUserRequest) returns (UpdateUserResponse)
rpc DeleteUser(DeleteUserRequest) returns (DeleteUserResponse)
```
### Шаг 1: Определение модели данных

Создай структуру User в пакете models:

- ID (string)
- Name (string)
- Email (string)
- Age (int32)
- CreatedAt (time.Time)
- UpdatedAt (time.Time)

### Шаг 2: Создание Protocol Buffers

1. Создай файл `user.proto` в api/proto/
2. Определи сервис UserService с 4 методами:
    - CreateUser
    - GetUser
    - UpdateUser
    - DeleteUser
3. Определи сообщения для запросов и ответов
4. Сгенерируй Go код из .proto файла

### Шаг 4: Реализация репозитория (хранилища)

Создай in-memory хранилище в пакете repository:

- Используй map для хранения пользователей
- Реализуй мьютекс для потокобезопасности
- Реализуй методы: Create, GetByID, Update, Delete
- Обработай случай, когда пользователь не найден

### Шаг 5: Реализация сервисного слоя

Создай бизнес-логику в пакете service:

- Принимай зависимости через конструктор
- Реализуй те же 4 метода, что и в репозитории
- Добавь валидацию данных (например, проверку email)

### Шаг 6: Реализация gRPC обработчика

Создай gRPC handler в internal/handler/grpc:

- Реализуй интерфейс из сгенерированного protobuf кода
- Конвертируй между protobuf сообщениями и моделями
- Обработай ошибки и верни соответствующие gRPC статусы

### Шаг 7: Реализация REST обработчика

Создай REST handler в internal/handler/rest:

- Реализуй HTTP endpoints для тех же операций
- Используй JSON для запросов и ответов
- Настрой правильные HTTP методы и статусы
- Добавь обработку path parameters для ID

### Шаг 8: Настройка сервера

В main.go:

- Инициализируй все зависимости (репозиторий → сервис → обработчики)
- Запусти gRPC сервер на порту 50051 в отдельной goroutine
- Запусти REST сервер на порту 8080
- Настрой graceful shutdown
