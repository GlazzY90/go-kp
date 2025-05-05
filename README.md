# RestFul API with Go-Gin and JWT

This is the result of one week Go Project. The descriptions contained in this are as follow: 

- RESTful CRUD operations on a **Task** resource  
- **MySQL** integration via **GORM**  
- **User authentication** with **JWT**  
- Clean architecture using the **Repository Pattern**  
- Git-based workflow  

---
## Project Structure

```
go-api/
├── main.go
├── config/
│   └── config.go
├── controllers/
│   └── task_handler.go
│   └── user_handler.go
├── middlewares/
│   └── task.go
│   └── user.go
├── models/
│   └── task.go
│   └── user.go
├── repos/
│   └── task_repository.go
│   └── user_repository.go
├── services/
│   └── task_service.go
│   └── user_service.go
├── utils/
│   └── init.sql
├── go.mod
├── go.sum
└── README.md
```

---
## 1. Project Setup

1. **Clone the repository**  
  ```sh
  git clone https://github.com/yourusername/myapp.git
  cd go-api
  ```

2. **Download dependencies**  
  ```sh
  go mod download
  ```

3. **Create environment file**  
  ```sh
  DB_USER=root
  DB_PASSWORD=
  DB_HOST=127.0.0.1
  DB_PORT=3306
  DB_NAME=go_api
  JWT_SECRET=
  ```

## 2. Running the database migration + Running the Project
  ```sh
  go run main.go
  ```

## 3. How to test the API endpoints 

Using Postman:
1. Import a Postman collection or manually create requests.
2. Set headers like Content-Type: application/json
3. Test routes such as:
  ```sh
  GET    /api/users
  POST   /api/login
  POST   /api/register
  ```
Using curl:
  ```sh
  curl -X POST http://localhost:8080/api/register \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser","password":"123456"}'
  ```
## 4. Any assumptions or decisions made during development 

Clean Architecture:

cmd/api for the application entrypoint

internal/ for all business logic

GORM + MySQL for ORM and persistence

JWT for stateless authentication

Repository Pattern to separate data access from business logic

SQLYog can be used to inspect and manage the MySQL database