# Task Management System API

This is a Task Management System API built using the Echo Framework in Go and MongoDB as the database. The API allows users to perform CRUD operations on tasks, including creating, reading, updating, and deleting tasks.

## Features

- **Create a Task**: Add a new task with a title, description, and status.
- **Get All Tasks**: Retrieve a list of all tasks.
- **Get Task by ID**: Fetch details of a specific task by its ID.
- **Update a Task**: Modify the title, description, or status of a task.
- **Delete a Task**: Remove a task by its ID.

## Technologies Used

- **Go**: Programming language used for building the API.
- **Echo Framework**: Web framework for handling HTTP requests and routing.
- **MongoDB**: NoSQL database used for data storage.
- **Remote MongoDB**: A remote MongoDB cluster was used for this project. MongoDB Compass was utilized for managing the database.

## Installation and Setup

### Prerequisites
- Go (version 1.19 or above)
- MongoDB Compass for database management
- Git

### Steps to Run the Application
1. Clone the repository:
    ```bash
    git clone <repository_url>
    cd <repository_name>
    ```

2. Install dependencies:
    ```bash
    go mod tidy
    ```

3. Update the MongoDB connection string in the `connectMongo` function with your remote MongoDB URI:
    ```go
    options.Client().ApplyURI("mongodb://<your_mongodb_uri>")
    ```

4. Run the application:
    ```bash
    go run main.go
    ```

5. The API will be available at `http://localhost:8080`.

## API Endpoints and Sample Responses

### Create a Task
**POST** `/tasks`

**Request Body**:
```json
{
  "title": "Complete Project",
  "description": "Finalize all modules and submit by deadline",
  "status": "In Progress"
}
```

**Response**:
```json
{
  "id": "64c2fba9357e9f7c09a823db",
  "title": "Complete Project",
  "description": "Finalize all modules and submit by deadline",
  "status": "In Progress",
  "created_at": "2024-12-19T12:00:00Z",
  "updated_at": "2024-12-19T12:00:00Z"
}
```

**GET ALL TASK** `/tasks`

**Response**:
```json
{
    "id": "task_id",
    "title": "Task Title",
    "description": "Optional Task Description",
    "status": "Pending",
    "created_at": "2024-12-19T10:00:00Z",
    "updated_at": "2024-12-19T10:00:00Z"
  }
```

**GET TASK BY ID** `/tasks`

**Response**:
```json

  {
  "id": "task_id",
  "title": "Task Title",
  "description": "Optional Task Description",
  "status": "Pending",
  "created_at": "2024-12-19T10:00:00Z",
  "updated_at": "2024-12-19T10:00:00Z"
} 
```


**PUT** `/tasks`

**Request Body**:
```json
{
  "title": "Updated Task Title",
  "description": "Updated Task Description",
  "status": "Completed"
}

```

**Response**:
```json

{
  "id": "task_id",
  "title": "Updated Task Title",
  "description": "Updated Task Description",
  "status": "Completed",
  "created_at": "2024-12-19T10:00:00Z",
  "updated_at": "2024-12-19T12:00:00Z"
}

```

**DELETE** `/tasks`

**Response**:
```json

{
  "message": "Task deleted successfully"
}

```

