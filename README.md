# News Topic API
---
This project is a RESTful API for managing news and topics using Golang with GORM, Go-Chi router, and Postgres. It includes API documentation with Swagger.

## Project Structure

Here's an overview of the project's folder structure:

```
News Project
├── cmd
│   └── main.go                        # Entry point for the application.
├── common
│   ├── base.entity.go                 # Base entity definitions and common methods.
│   └── pagination.go                  # Pagination utility functions.
├── docs
│   ├── docs.go                        # Swagger documentation setup.
│   ├── swagger.json                   # Swagger JSON file for API documentation.
│   └── swagger.yaml                   # Swagger YAML file for API documentation.
├── internal
│   ├── db
│   │   ├── migrations                 # Database migration files.
│   │   │   ├── 20240720141608_create_topics_table.sql # Migration for topics table.
│   │   │   ├── 20240720180753_create_statuses_table.sql # Migration for statuses table.
│   │   │   ├── 20240720180805_create_news_table.sql # Migration for news table.
│   │   │   └── 20240720180835_create_news_topics_table.sql # Migration for news_topics table.
│   │   └── postgres.go                # Postgres database connection setup.
│   ├── delivery
│   │   ├── data
│   │   │   ├── dtos                   # Data Transfer Objects (DTOs) for requests and responses.
│   │   │   │   ├── news.dto.go        # DTOs related to news.
│   │   │   │   └── topic.dto.go       # DTOs related to topics.
│   │   │   └── responses              # Response structures for API responses.
│   │   │       ├── news.response.go   # Response structure for news.
│   │   │       ├── response.go        # General response structure.
│   │   │       └── topic.response.go  # Response structure for topics.
│   │   └── handlers                   # HTTP handlers for different routes.
│   │       ├── news.handler.go        # Handlers for news-related requests.
│   │       └── topic.handler.go       # Handlers for topic-related requests.
│   ├── entities                       # Database entity definitions.
│   │   ├── news.entity.go             # News entity definition.
│   │   └── topics.entity.go           # Topic entity definition.
│   ├── repositories                   # Repository interfaces and implementations.
│   │   ├── news_interface.repository.go # Interface for news repository.
│   │   └── news.repository.go         # Implementation of news repository.
│   │   ├── topic_interface.repository.go # Interface for topic repository.
│   │   └── topic.repository.go        # Implementation of topic repository.
│   ├── routes                        # Route definitions.
│   │   ├── news.router.go             # Routes for news endpoints.
│   │   ├── routes.go                 # Main route configuration.
│   │   └── topic.router.go            # Routes for topic endpoints.
│   └── usecase                       # Use cases for business logic.
│       ├── news_interface.usecase.go  # Interface for news use cases.
│       └── news.usecase.go            # Implementation of news use cases.
│       ├── topic_interface.usecase.go # Interface for topic use cases.
│       └── topic.usecase.go           # Implementation of topic use cases.
├── .env                              # Environment configuration file.
├── .env.example                       # Example environment configuration file.
├── .gitignore                         # Git ignore file.
├── go.mod                             # Go module file.
├── go.sum                             # Go module checksum file.
├── main                               # Compiled application binary (may vary).
└── README.md                          # Project documentation.
```

## Setup and Installation

### Prerequisites

1. **Go**: Ensure you have Go installed. You can download it from [golang.org](https://golang.org/dl/).
2. **Postgres**: Make sure PostgreSQL is installed and running on your local machine.
3. **Swagger**: Install Swagger if you want to access the API documentation.

### 1. Clone the Repository

Clone the repository to your local machine:

```bash
git clone https://github.com/awahids/go-news-topic.git
cd go-news-topic
```

### 2. Set Up the Environment

Create a `.env` file in the root directory based on `.env.example`. Configure the environment variables as needed:

```bash
cp .env.example .env
```

### 3. Install Dependencies

Install Go dependencies:

```bash
go mod tidy
```

### 4. Database Migrations

You can manage migrations using GORM or Goose. Choose one of the following methods:

#### Using GORM Migrations

Run GORM migrations to set up the database schema:
- open postgres.go
- uncomment line 32

```bash
// db.AutoMigrate(&entities.News{}, &entities.Topic{})
```

#### Using Goose Migrations

1. **Migrate Up:**

   Migrate the database to the most recent version:

   ```bash
   goose -dir ./internal/db/migrations postgres "user=youruser dbname=news_topic_api sslmode=disable" up
   ```

2. **Migrate Down:**

   Roll back the last migration:

   ```bash
   goose -dir ./internal/db/migrations postgres "user=youruser dbname=news_topic_api sslmode=disable" down
   ```

3. **Reset Migrations:**

   Roll back all migrations:

   ```bash
   goose -dir ./internal/db/migrations postgres "user=youruser dbname=news_topic_api sslmode=disable" reset
   ```

### 5. Start the Application

Use CompileDaemon to automatically rebuild and restart the application when files change:

```bash
CompileDaemon --build="go build cmd/main.go" --command=./main
```

### 6. Access Swagger Documentation

Once the server is running, you can access the Swagger documentation at:

```
http://localhost:9000/api/v1/swagger/index.html
```

## Troubleshooting

- **Postgres Connection**: Ensure your `.env` file has the correct Postgres connection details.
- **Migration Issues**: Check the migration files in `internal/db/migrations` if you encounter schema issues.
- **CompileDaemon Issues**: Ensure `CompileDaemon` is installed. You can install it via `go get`:

    ```bash
    go get github.com/githubnemo/CompileDaemon
    ```

## Additional Notes

- **Swagger Documentation**: The Swagger UI provides interactive documentation for your API endpoints.
- **Configuration**: Adjust configurations as needed in the `.env` file or directly in the source code.

---