# Task Management Service
A microservice for managing tasks using Go, Gin, and GORM.

## Setup
1. Setup DB - PostgreSQL
   1. DB Name: `alle-task-db`
   2. Username: `postgres`
   3. Password: `admin123`
2. Install `brew install golang-migrate` for golang db migration support.
3. Run migrations to setup the database schema.
   `migrate -path migrations -database "postgres://localhost:5432/alle-task-db?sslmode=disable" up`
4. Install all dependencies: `go mod tidy`
5. Run the service: `go run cmd/main.go`
   1. The service by default adds dummy entry to the DB on each startup for ease of testing.
6. Service running on http://localhost:8080

## API Endpoints
- **POST /tasks**: Create a new task.
  ```
    POST: http://localhost:8080/tasks
    BODY:
    { 
        "title": "Task 1",
        "description": "Description 1"
    }
    RESPONSE:
    {
        "id": 43,
        "title": "Task 1",
        "description": "Description 1",
        "status": "Pending",
        "created_at": "2025-02-15T02:44:20.57623-05:00",
        "updated_at": "2025-02-15T02:44:20.57623-05:00"
    }
  ```
- **GET /tasks**: Get all tasks with pagination.
  ```
    GET: localhost:8080/tasks?page=<page_no>
    RESPONSE:
    {
        "count": 5,
        "page": 1,
        "tasks": [
        {
            "id": 1,
            "title": "Task 1",
            "description": "Description 1",
            "status": "Pending",
            "created_at": "2025-02-14T02:09:54.361871Z",
            "updated_at": "2025-02-14T02:09:54.361871Z"
        },
        {
            "id": 2,
            "title": "Task 2",
            "description": "Description 2",
            "status": "In Progress",
            "created_at": "2025-02-14T02:09:54.365125Z",
            "updated_at": "2025-02-14T02:09:54.365125Z"
        },
        {
            "id": 3,
            "title": "Task 3",
            "description": "Description 3",
            "status": "Completed",
            "created_at": "2025-02-14T02:09:54.365524Z",
            "updated_at": "2025-02-14T02:09:54.365524Z"
        },
        {
            "id": 4,
            "title": "Task 1",
            "description": "Description 1",
            "status": "Pending",
            "created_at": "2025-02-14T02:10:15.331329Z",
            "updated_at": "2025-02-14T02:10:15.331329Z"
        },
        {
            "id": 5,
            "title": "Task 2",
            "description": "Description 2",
            "status": "In Progress",
            "created_at": "2025-02-14T02:10:15.333727Z",
            "updated_at": "2025-02-14T02:10:15.333727Z"
        }
    ]
    }
  ```
- **GET /tasks/filter**: Filter tasks by status.
    ```
    GET: localhost:8080/tasks?status=Pending
    RESPONSE:
    {
    "count": 5,
    "page": 1,
    "tasks": [
        {
            "id": 1,
            "title": "Task 1",
            "description": "Description 1",
            "status": "Pending",
            "created_at": "2025-02-14T02:09:54.361871Z",
            "updated_at": "2025-02-14T02:09:54.361871Z"
        },
        {
            "id": 4,
            "title": "Task 1",
            "description": "Description 1",
            "status": "Pending",
            "created_at": "2025-02-14T02:10:15.331329Z",
            "updated_at": "2025-02-14T02:10:15.331329Z"
        },
        {
            "id": 7,
            "title": "Task 1",
            "description": "Description 1",
            "status": "Pending",
            "created_at": "2025-02-14T02:11:12.214242Z",
            "updated_at": "2025-02-14T02:11:12.214242Z"
        },
        {
            "id": 10,
            "title": "Task 1",
            "description": "Description 1",
            "status": "Pending",
            "created_at": "2025-02-14T02:15:29.323029Z",
            "updated_at": "2025-02-14T02:15:29.323029Z"
        },
        {
            "id": 13,
            "title": "Task 1",
            "description": "Description 1",
            "status": "Pending",
            "created_at": "2025-02-14T02:44:48.632864Z",
            "updated_at": "2025-02-14T02:44:48.632864Z"
        }
    ]
    }
    
  ```
- **GET /tasks/:id**: Get a task by ID.
- **PUT /tasks/:id**: Update a task.
- **DELETE /tasks/:id**: Delete a task.


