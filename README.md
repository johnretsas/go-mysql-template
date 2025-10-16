# Go SQL Template with Docker

A simple template for Go applications with MySQL database integration using Docker.

## Features

- Go application with MySQL database connectivity
- Docker containerization for easy deployment
- Database initialization with custom SQL scripts
- Environment variable configuration
- Docker Compose for simplified orchestration

## Project Structure

```
go-sql-template/
├── cmd/
│   └── main.go          # Main Go application
├── docker-compose.yaml  # Docker Compose configuration
├── Dockerfile          # MySQL container configuration
├── go.mod             # Go module definition
├── init.sql           # Database initialization script
└── README.md          # This file
```

## Prerequisites

- [Go](https://golang.org/dl/) 1.23.2 or later
- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/install/)

## Quick Start

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd go-sql-template
   ```

2. **Set up environment variables**
   Create a `.env` file in the root directory:
   ```env
   CONTAINER_NAME=my-mysql-db
   MYSQL_ROOT_PASSWORD=rootpassword
   MYSQL_DATABASE=myapp
   MYSQL_USER=appuser
   MYSQL_PASSWORD=apppassword
   ```

3. **Start the database**
   ```bash
   docker-compose up -d
   ```

4. **Install Go dependencies**
   ```bash
   go mod tidy
   ```

5. **Set environment variables for the Go application**
   ```bash
   export DB_USER=appuser
   export DB_PASSWORD=apppassword
   export DB_NAME=myapp
   ```

6. **Run the Go application**
   ```bash
   go run cmd/main.go
   ```

## Configuration

### Environment Variables

The application uses the following environment variables:

| Variable | Description | Default |
|----------|-------------|---------|
| `DB_USER` | Database user | - |
| `DB_PASSWORD` | Database password | - |
| `DB_NAME` | Database name | - |

### Docker Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `CONTAINER_NAME` | MySQL container name | - |
| `MYSQL_ROOT_PASSWORD` | MySQL root password | - |
| `MYSQL_DATABASE` | Database to create | - |
| `MYSQL_USER` | MySQL user to create | - |
| `MYSQL_PASSWORD` | Password for MySQL user | - |

## Database Initialization

The `init.sql` file is automatically executed when the MySQL container starts for the first time. Add your table creation scripts and initial data here.

Example:
```sql
CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(100) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO users (username, email) VALUES 
('john_doe', 'john@example.com'),
('jane_smith', 'jane@example.com');
```

## Development

### Adding Dependencies

To add new Go dependencies:
```bash
go get <package-name>
go mod tidy
```

### Database Connection

The application connects to MySQL using the standard `database/sql` package. The connection string is built from environment variables:

```go
dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s", user, password, database)
db, err := sql.Open("mysql", dsn)
```

## Docker Commands

### Start services
```bash
docker-compose up -d
```

### Stop services
```bash
docker-compose down
```

### View logs
```bash
docker-compose logs database
```

### Rebuild containers
```bash
docker-compose up --build
```

## Troubleshooting

### Connection Issues

1. **Ensure MySQL container is running:**
   ```bash
   docker-compose ps
   ```

2. **Check MySQL logs:**
   ```bash
   docker-compose logs database
   ```

3. **Verify environment variables are set correctly**

### Database Access

Connect to MySQL directly:
```bash
docker exec -it <container-name> mysql -u <username> -p
```

## License

This project is open source and available under the [MIT License](LICENSE).