# go-ogen-crud-example

HTTP-сервер для управления данными об автомобилях с in-memory хранилищем , сгенерированный с помощью
библиотеки [ogen](https://github.com/ogen-go/ogen).

## Запуск сервера

```bash
go run cmd/api-server/main.go
```

Сервер будет доступен по адресу `http://localhost:8080`.

## Примеры использования API

### Получение информации о всех автомобилях

```bash
curl -X GET "http://localhost:8080/cars" -H "accept: application/json"
```

### Получение информации об одном автомобиле

```bash
curl -X GET "http://localhost:8080/cars/0" -H "accept: application/json"
```

### Добавление нового автомобиля

```bash
curl -X POST "http://localhost:8080/cars" \
  -H "Content-Type: application/json" \
  -d '{"manufacturer":"Toyota","model":"Camry","year":2022,"color":"black"}'
```

### Обновление данных автомобиля

```bash
curl -X PUT "http://localhost:8080/cars/0" \
  -H "Content-Type: application/json" \
  -d '{"model":"Camry Hybrid","year":2023}'
```

### Удаление автомобиля

```bash
curl -X DELETE "http://localhost:8080/cars/0"
```

## Тестовый клиент

Для тестирования API можно использовать клиент:

```bash
go run cmd/api-client/main.go
```