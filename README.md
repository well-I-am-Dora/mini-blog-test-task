### Инструкция к запуску
#### Для in memory хранилища
```
IN_MEMORY=true go run cmd/main.go
```
#### Для PostgreSQL хранилища.
```
DSN='<your-dsn>' go run cmd/main.go
```
Так же перед использованием надо применить миграции.
```
goose up
```
#### Запуск через docker
```
docker build . -t post-service
docker run --env IN_MEMORY=true -p=8080:8080 post-service
```

### Комментарий к решению
Если мы хотим получать список постов с комментариями, необходимо добавить dataloader чтобы решить проблему с n + 1 запросом. Так как в задаче нет условия что это необходимо, dataloader не был добавлен.