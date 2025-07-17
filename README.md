# Go Todo API

A sophisticated and professional Todo REST API built with Go, featuring clean architecture, comprehensive documentation, and production-ready features.

## 🚀 Features

- ✅ **Clean Architecture** - Layered architecture with clear separation of concerns
- ✅ **RESTful API** - Standard HTTP methods and status codes
- ✅ **Database Support** - SQLite (development) and PostgreSQL (production)
- ✅ **GORM ORM** - Type-safe database operations with auto-migration
- ✅ **UUID Primary Keys** - Secure and globally unique identifiers
- ✅ **Input Validation** - Comprehensive request validation
- ✅ **Structured Logging** - JSON-formatted logs with zerolog
- ✅ **CORS Support** - Cross-origin resource sharing configuration
- ✅ **Request ID Tracking** - Request tracing and debugging
- ✅ **Health Checks** - Application health monitoring
- ✅ **Prometheus Metrics** - Performance and business metrics
- ✅ **Swagger Documentation** - Auto-generated API documentation
- ✅ **Graceful Shutdown** - Proper server termination
- ✅ **Environment Configuration** - Flexible configuration management
- ✅ **Pagination** - Efficient data retrieval with metadata
- ✅ **Soft Deletes** - Data preservation with logical deletion

<!-- ## 🏗️ Architecture

```
├── cmd/api/           # Application entry point
├── internal/          # Private application code
│   ├── config/        # Configuration management
│   ├── handlers/      # HTTP request handlers
│   ├── middleware/    # HTTP middleware
│   ├── models/        # Data models and DTOs
│   ├── repository/    # Data access layer
│   └── services/      # Business logic layer
├── pkg/               # Public packages
│   ├── logger/        # Structured logging
│   ├── response/      # Standardized API responses
│   └── validator/     # Request validation
└── docs/              # Documentation
``` -->

## 🛠️ Technology Stack

- **Framework**: [Gin](https://github.com/gin-gonic/gin) - High-performance HTTP web framework
- **ORM**: [GORM](https://gorm.io/) - Object-relational mapping library
- **Logging**: [Zerolog](https://github.com/rs/zerolog) - Structured JSON logging
- **Validation**: [Validator](https://github.com/go-playground/validator) - Request validation
- **Documentation**: [Swagger](https://swagger.io/) - API documentation
- **Database**: SQLite (dev) / PostgreSQL (prod)
- **Monitoring**: Prometheus metrics

## 📋 Prerequisites

- Go 1.21 or higher
- Git
- SQLite (for development)
- PostgreSQL (for production, optional)

## 🚀 Quick Start

### 1. Clone the repository

```bash
git clone https://github.com/1cbyc/go-todo-api.git
cd go-todo-api
```

### 2. Install dependencies

```bash
go mod tidy
```

### 3. Set up environment variables

Create a `.env` file in the root directory:

```env
# Server Configuration
PORT=8080
GIN_MODE=debug
LOG_LEVEL=info

# Database Configuration
DB_DRIVER=sqlite
DB_NAME=todo_api

# For PostgreSQL (production)
# DB_DRIVER=postgres
# DB_HOST=localhost
# DB_PORT=5432
# DB_USER=postgres
# DB_PASSWORD=password
# DB_NAME=todo_api
# DB_SSLMODE=disable
```

### 4. Run the application

```bash
go run cmd/api/main.go
```

The API will be available at `http://localhost:8080`

## 📚 API Documentation

### Interactive Documentation

Once the server is running, visit:
- **Swagger UI**: http://localhost:8080/swagger/index.html
- **Health Check**: http://localhost:8080/health
- **Metrics**: http://localhost:8080/api/v1/metrics

### API Endpoints

#### Todo Management

| Method | Endpoint | Description |
|--------|----------|-------------|
| `GET` | `/api/v1/todos` | List all todos with pagination |
| `GET` | `/api/v1/todos/:id` | Get a specific todo |
| `POST` | `/api/v1/todos` | Create a new todo |
| `PUT` | `/api/v1/todos/:id` | Update a todo |
| `DELETE` | `/api/v1/todos/:id` | Delete a todo |
| `PATCH` | `/api/v1/todos/:id/toggle` | Toggle todo completion |

#### System Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| `GET` | `/health` | Health check |
| `GET` | `/api/v1/metrics` | Prometheus metrics |
| `GET` | `/swagger/*` | API documentation |

### Example Requests

#### Create a Todo

```bash
curl -X POST http://localhost:8080/api/v1/todos \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Learn Go",
    "description": "Study Go programming language",
    "priority": "high",
    "due_date": "2024-12-31T23:59:59Z"
  }'
```

#### Get All Todos

```bash
curl -X GET "http://localhost:8080/api/v1/todos?page=1&per_page=10"
```

#### Update a Todo

```bash
curl -X PUT http://localhost:8080/api/v1/todos/{id} \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Learn Go Programming",
    "completed": true
  }'
```

#### Toggle Todo Completion

```bash
curl -X PATCH http://localhost:8080/api/v1/todos/{id}/toggle
```

## 🗄️ Database

### SQLite (Development)

The application uses SQLite by default for development. The database file will be created automatically as `todo_api.db`.

### PostgreSQL (Production)

To use PostgreSQL in production:

1. Install PostgreSQL
2. Create a database
3. Update environment variables:

```env
DB_DRIVER=postgres
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=todo_api
DB_SSLMODE=disable
```

## 🧪 Testing

### Run Tests

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests with verbose output
go test -v ./...
```

### Run Specific Tests

```bash
# Run only unit tests
go test ./internal/services/...

# Run only integration tests
go test ./internal/repository/...

# Run benchmarks
go test -bench=. ./...
```

## 🐳 Docker

### Build Docker Image

```bash
docker build -t todo-api .
```

### Run with Docker

```bash
docker run -p 8080:8080 todo-api
```

### Docker Compose

Create a `docker-compose.yml` file:

```yaml
version: '3.8'
services:
  api:
    build: .
    ports:
      - "8080:8080"
    environment:
      - DB_DRIVER=sqlite
      - DB_NAME=todo_api
    volumes:
      - ./data:/app/data

  # For PostgreSQL
  postgres:
    image: postgres:15
    environment:
      POSTGRES_DB: todo_api
      POSTGRES_USER: postgres
      POSTGRES_PASSWORD: password
    ports:
      - "5432:5432"
    volumes:
      - postgres_data:/var/lib/postgresql/data

volumes:
  postgres_data:
```

Run with:

```bash
docker-compose up -d
```

## 📊 Monitoring

### Health Check

```bash
curl http://localhost:8080/health
```

Response:
```json
{
  "status": "ok",
  "message": "Todo API is running",
  "version": "1.0.0",
  "timestamp": "2024-01-01T00:00:00Z",
  "uptime": "1h2m3s",
  "memory": {
    "alloc": 1234567,
    "total_alloc": 2345678,
    "sys": 3456789,
    "num_gc": 5
  },
  "goroutines": 10
}
```

### Prometheus Metrics

```bash
curl http://localhost:8080/api/v1/metrics
```

## 🔧 Configuration

### Environment Variables

| Variable | Default | Description |
|----------|---------|-------------|
| `PORT` | `8080` | Server port |
| `GIN_MODE` | `debug` | Gin mode (debug/release) |
| `LOG_LEVEL` | `info` | Log level (debug/info/warn/error) |
| `DB_DRIVER` | `sqlite` | Database driver (sqlite/postgres) |
| `DB_HOST` | `localhost` | Database host |
| `DB_PORT` | `5432` | Database port |
| `DB_USER` | `postgres` | Database user |
| `DB_PASSWORD` | `` | Database password |
| `DB_NAME` | `todo_api` | Database name |
| `DB_SSLMODE` | `disable` | Database SSL mode |

## 🚀 Deployment

### Production Build

```bash
# Build for production
go build -o bin/api cmd/api/main.go

# Run production binary
./bin/api
```

### Environment Setup

For production, set:

```env
GIN_MODE=release
LOG_LEVEL=info
DB_DRIVER=postgres
# ... other production settings
```

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

### Development Guidelines

- Follow Go coding standards and conventions
- Write comprehensive tests for new features
- Update documentation for API changes
- Use conventional commit messages
- Ensure all tests pass before submitting PRs

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 📚 Documentation

- [Technical Explanation](docs/explanation.md) - Detailed architecture and design decisions
- [What's Next](docs/whats-next.md) - Development roadmap and future features
- [API Documentation](http://localhost:8080/swagger/index.html) - Interactive API docs

## 🆘 Support

If you encounter any issues or have questions:

1. Check the [documentation](docs/)
2. Search existing [issues](https://github.com/1cbyc/go-todo-api/issues)
3. Create a new issue with detailed information

## 🙏 Acknowledgments

- [Gin](https://github.com/gin-gonic/gin) - Web framework
- [GORM](https://gorm.io/) - ORM library
- [Zerolog](https://github.com/rs/zerolog) - Logging library
- [Swagger](https://swagger.io/) - API documentation

---

**Happy coding! 🎉**
