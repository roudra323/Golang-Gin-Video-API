# Gin Simple API

A clean and simple REST API for managing videos, built with Go using the Gin web framework and GORM ORM. This project demonstrates a well-structured, layered architecture following best practices for building scalable web APIs.

## 🚀 Features

- **RESTful API**: Clean REST endpoints for video management
- **Layered Architecture**: Separation of concerns with handler, service, and repository layers
- **GORM Integration**: Database operations with SQLite (easily switchable to other databases)
- **Auto Migration**: Automatic database schema migration
- **JSON Validation**: Request validation with Gin's binding features
- **Error Handling**: Proper HTTP status codes and error responses

## 🛠 Tech Stack

- **Language**: Go 1.24.3
- **Web Framework**: [Gin](https://gin-gonic.com/)
- **ORM**: [GORM](https://gorm.io/)
- **Database**: SQLite (development), easily configurable for PostgreSQL/MySQL
- **Architecture**: Clean Architecture with dependency injection

## 📁 Project Structure

```
gin-simple-api/
├── main.go                 # Application entry point
├── go.mod                  # Go module dependencies
├── go.sum                  # Dependency checksums
├── test.db                 # SQLite database file
├── api/
│   └── video_handler.go    # HTTP handlers (controllers)
├── service/
│   └── video_service.go    # Business logic layer
├── repository/
│   └── video_repository.go # Data access layer
└── domain/
    └── video.go           # Domain models/entities
```

## 🏗 Architecture

This project follows a **layered architecture** pattern:

1. **Handler Layer** (`api/`): HTTP request/response handling
2. **Service Layer** (`service/`): Business logic and validation
3. **Repository Layer** (`repository/`): Data access and persistence
4. **Domain Layer** (`domain/`): Core business entities

## 🚀 Getting Started

### Prerequisites

- Go 1.24+ installed on your machine
- Git (optional, for cloning)

### Installation

1. **Clone the repository**
   ```bash
   git clone https://github.com/roudra323/gin-simple-api.git
   cd gin-simple-api
   ```

2. **Install dependencies**
   ```bash
   go mod tidy
   ```

3. **Run the application**
   ```bash
   go run main.go
   ```

The server will start on `http://localhost:8080`

## 📚 API Endpoints

### Video Management

| Method | Endpoint      | Description           | Request Body |
|--------|---------------|----------------------|-------------|
| GET    | `/videos`     | Get all videos       | -           |
| POST   | `/videos`     | Create a new video   | Video JSON  |
| GET    | `/videos/:id` | Get video by ID      | -           |

### Video Model

```json
{
  "id": 1,
  "title": "Sample Video",
  "description": "This is a sample video description",
  "url": "https://example.com/video.mp4",
  "author": "John Doe",
  "created_at": "2025-01-01T00:00:00Z",
  "updated_at": "2025-01-01T00:00:00Z"
}
```

## 🔥 Usage Examples

### Create a Video

```bash
curl -X POST http://localhost:8080/videos \
  -H "Content-Type: application/json" \
  -d '{
    "title": "My Awesome Video",
    "description": "This is an awesome video",
    "url": "https://example.com/awesome-video.mp4",
    "author": "Jane Doe"
  }'
```

### Get All Videos

```bash
curl http://localhost:8080/videos
```

### Get Video by ID

```bash
curl http://localhost:8080/videos/1
```

## 🔧 Configuration

### Database Configuration

The application uses SQLite by default. To switch to PostgreSQL or MySQL:

1. Update the database driver import in `main.go`
2. Modify the connection string in the `gorm.Open()` call

**Example for PostgreSQL:**
```go
import "gorm.io/driver/postgres"

// In main.go
dsn := "host=localhost user=username password=password dbname=videos port=5432 sslmode=disable"
db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
```

### Environment Variables

You can configure the application using environment variables:

- `PORT`: Server port (default: 8080)
- `DB_HOST`: Database host
- `DB_USER`: Database username
- `DB_PASSWORD`: Database password
- `DB_NAME`: Database name

## 🧪 Testing

### Manual Testing

Use the provided curl examples above, or import the API into Postman/Insomnia.

### Running Tests

```bash
go test ./...
```

## 🐳 Docker Support

Create a `Dockerfile`:

```dockerfile
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]
```

Build and run:
```bash
docker build -t gin-simple-api .
docker run -p 8080:8080 gin-simple-api
```

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 👨‍💻 Author

**Roudra** - [GitHub Profile](https://github.com/roudra323)

## 🙏 Acknowledgments

- [Gin Web Framework](https://gin-gonic.com/) - HTTP web framework
- [GORM](https://gorm.io/) - The fantastic ORM library for Golang
- Go community for excellent tooling and libraries
