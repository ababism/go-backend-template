# GO Backend template
### Описание программы:

Стек:
- Gin + graceful shutdown, чистая архитектура
- Multistage докер сборка приложения
- Docker Compose с запуском БД и Healthcheck
- Файлы миграций
- Postman API коллекция для импорта (API.Postman_collection.json)
- PostgreSQL

### Строение репозитория программы:
Задание сделано по лэйауту привычному для чистой архитектуры у монолита:
```
├── cmd/app   // Приложение и его настройка
├── cmd/server   // точка входа
├── internal // директория не предполгающая "экспорт"
│  ├── config   // все что касается конфигураций
│  ├── transport/http   // транспортный слой: payload модели, хэдлеры, routes и middleware
│  ├── service   // бизнес логика
│  ├── service/domain   // общие модели и ошибки  
│  ├── repository/postgres   // модели для обращения к БД

```
### Строение репозитория программы:
Запуск:
- убедится что docker-compose.уml файл соответствует ожидаемым параметрам запуска 
- убедится что config/dockerconfig.уml файл на месте
- Запустить docker-compose
```shell
docker compose up --build
```
- Прогнать миграции
```shell
migrate -path ./schema -database 'postgres://postgres:password@localhost:5437/postgres?sslmode=disable' up
```
- БД на 5437 и сервер на 8080

