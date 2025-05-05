# RestFul API with Go-Gin and JWT

This is the result of one week Go Project. The descriptions contained in this are as follow: 

- RESTful CRUD operations on a **Task** resource  
- **MySQL** integration via **GORM**  
- **User authentication** with **JWT**  
- Clean architecture using the **Repository Pattern**  
- Git-based workflow  

---
## Project Structure

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

## 2. Running the Project
  ```sh
  go run main.go
  ```