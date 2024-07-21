## Create migration

```go
goose -dir ./internal/db/migrations create create_topics_table sql
```

## Goose up
Migrate the DB to the most recent version available
```go
goose -dir ./internal/db/migrations postgres "user=youruser dbname=news_topic_api sslmode=disable" up
```

## Goose down
Roll back the version by 1
```go
goose -dir ./internal/db/migrations postgres "user=youruser dbname=news_topic_api sslmode=disable" down
```

## Goose reset
Roll back all migrations

## Swagger init

```go
swag init -g cmd/main.go
```