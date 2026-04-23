# Blog API Go

A production-grade RESTful blog API built with Go, featuring JWT authentication, role-based access control, and PostgreSQL persistence.

## Features

- **Authentication & Authorization**
  - JWT-based authentication
  - Role-based access control (user, admin, owner)
  - Secure password hashing with bcrypt

- **Blog Management**
  - Create, read, update, and delete posts
  - Author-based post ownership
  - Public read access, authenticated write access

- **User Management**
  - User registration and login
  - Admin user management endpoints
  - Role assignment (owner only)

- **Security**
  - CORS protection
  - Input validation
  - SQL injection prevention
  - Non-root Docker container execution

## Tech Stack

- **Go 1.25+** - Backend language
- **Gin** - HTTP web framework
- **PostgreSQL** - Database
- **pgx/v5** - PostgreSQL driver
- **JWT** - Token-based authentication
- **Docker** - Containerization
- **Goose** - Database migrations

## Quick Start

### Using Docker (Recommended)

1. Clone the repository:
```bash
git clone <repository-url>
cd blog-api-go
```

2. Create environment file:
```bash
cp example.env .env.prod
# Edit .env.prod and change POSTGRES_PASSWORD and JWT_SECRET
```

3. Start services:
```bash
docker-compose up -d
```

4. Apply database migrations:
```bash
GOOSE_DRIVER=postgres GOOSE_DBSTRING="postgres://postgres:your_password@localhost:5433/blog-go" goose -dir ./migrations up
```

5. Test the API:
```bash
curl http://localhost:8080/ping
# Response: {"message":"pong"}
```

### Local Development

1. Install dependencies:
```bash
go mod download
```

2. Set up PostgreSQL database and create `.env` file:
```bash
DB_URL=postgres://postgres:postgres@localhost:5432/blog-go
JWT_SECRET=your-secret-key-here
GIN_MODE=debug
```

3. Run migrations:
```bash
goose -dir ./migrations up
```

4. Start the server:
```bash
go run cmd/server/main.go
```

## API Endpoints

### Public Routes
- `GET /ping` - Health check
- `POST /api/v1/auth/register` - Register new user
- `POST /api/v1/auth/login` - Login and get JWT token
- `GET /api/v1/posts/` - List all posts
- `GET /api/v1/posts/:id` - Get single post

### Authenticated Routes (JWT required)
- `POST /api/v1/auth/logout` - Logout
- `POST /api/v1/posts/` - Create new post
- `PUT /api/v1/posts/:id` - Update post (author only)
- `DELETE /api/v1/posts/:id` - Delete post (author only)

### Admin Routes (admin/owner roles)
- `GET /api/v1/users/` - List all users
- `GET /api/v1/users/:id` - Get user details
- `PUT /api/v1/users/:id` - Update user
- `DELETE /api/v1/users/:id` - Delete user

### Owner Routes (owner role only)
- `PUT /api/v1/admin/users/:id/role` - Change user role

## Example Usage

### Register a new user
```bash
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{"username":"john","email":"john@example.com","password":"secret123"}'
```

### Login
```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"john@example.com","password":"secret123"}'
```

### Create a post (with JWT token)
```bash
curl -X POST http://localhost:8080/api/v1/posts/ \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{"title":"My First Post","content":"Hello World!"}'
```

### Get all posts
```bash
curl http://localhost:8080/api/v1/posts/
```

## Docker Commands

```bash
# Start services
docker-compose up -d

# View logs
docker-compose logs -f api

# Stop services
docker-compose down

# Rebuild after code changes
docker-compose build --no-cache

# Access PostgreSQL
docker exec -it blog-api-postgres psql -U postgres -d blog-go
```

## Development

### Run tests
```bash
go test ./...
go test ./... -race  # with race detector
```

### Lint code
```bash
go vet ./...
golangci-lint run
```

### Database migrations
```bash
# Create new migration
goose create migration_name sql

# Apply migrations
goose -dir ./migrations up

# Rollback migration
goose -dir ./migrations down

# Check status
goose -dir ./migrations status
```

## Project Structure

```
.
├── cmd/server/          # Application entry point
├── internal/
│   ├── database/        # Database connection pooling
│   ├── entities/        # Domain models
│   ├── handlers/        # HTTP handlers
│   ├── middleware/      # Auth & RBAC middleware
│   ├── repositories/    # Data access layer
│   ├── services/        # Business logic layer
│   └── libs/            # Utilities (config, JWT, hashing)
├── migrations/          # Database migrations
├── Dockerfile           # Multi-stage Docker build
├── docker-compose.yml   # Docker services configuration
└── example.env          # Environment variables template
```

## License

MIT
