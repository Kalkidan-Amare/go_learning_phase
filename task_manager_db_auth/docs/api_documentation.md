# Task Management API Documentation

This API provides CRUD operations for managing tasks, now with persistent data storage using MongoDB and JWT-based authentication and authorization.

## Postman Documentation

For detailed information and testing endpoints, please refer to the [Postman documentation](https://documenter.getpostman.com/view/37352369/2sA3s1nWpq).

## Authentication & Authorization

- **Registration**:
  - **Endpoint**: 
    - **Admin**: `POST /register-admin`
    - **User**: `POST /register-user`
  - **Description**: Registers a new admin or user account with the API. Admins can perform all CRUD operations on tasks, while users can only manage their own tasks and read others.
  - **Request**:
    - **Body** (for both admin and user): 
    ```json
    {
      "username": "exampleuser",
      "password": "securepassword"
    }
    ```
  - **Response**: 
    - **Status**: 201 Created
    - **Body**: Confirmation message or error

- **Login**:
  - **Endpoint**: `POST /login`
  - **Description**: Authenticates the user and returns a JWT token.
  - **Request**:
    - **Body**: 
    ```json
    {
      "username": "exampleuser",
      "password": "securepassword"
    }
    ```
  - **Response**:
    - **Status**: 200 OK
    - **Body**: JWT token

- **JWT Authentication**:
  - Include the JWT token in the `Authorization` header with the format `Bearer <token>` for all protected routes.

## Endpoints

### 1. GET /tasks
- **Description**: Retrieve a list of all tasks.
- **Authorization**: 
  - **Role**: `admin` or `user`
- **Response**:
  - **Status**: 200 OK
  - **Body**: List of tasks

### 2. GET /tasks/:id
- **Description**: Retrieve the details of a specific task.
- **Authorization**: 
  - **Role**: `admin` or `user`
- **Parameters**:
  - `id`: The ID of the task
- **Response**:
  - **Status**: 200 OK
  - **Body**: Task details
  - **Status**: 404 Not Found (if task not found)

### 3. POST /tasks
- **Description**: Create a new task.
- **Authorization**: 
  - **Role**: `admin` or `user`
- **Request**:
  - **Body**: JSON object containing title, description, due date, and status
    ```json
    {
      "title": "New Task",
      "description": "Task description",
      "due_date": "2024-08-15",
      "status": "pending"
    }
    ```
- **Response**:
  - **Status**: 201 Created
  - **Body**: Created task details

### 4. PUT /tasks/:id
- **Description**: Update a specific task.
- **Authorization**: 
  - **Role**: `admin` or `user` (Users can only update their own tasks)
- **Parameters**:
  - `id`: The ID of the task
- **Request**:
  - **Body**: JSON object containing new task details
    ```json
    {
      "title": "Updated Task Title",
      "description": "Updated task description",
      "due_date": "2024-08-20",
      "status": "completed"
    }
    ```
- **Response**:
  - **Status**: 200 OK
  - **Body**: Updated task details
  - **Status**: 403 Forbidden (if user tries to update another user's task)
  - **Status**: 404 Not Found (if task not found)

### 5. DELETE /tasks/:id
- **Description**: Delete a specific task.
- **Authorization**: 
  - **Role**: `admin` or `user` (Users can only delete their own tasks)
- **Parameters**:
  - `id`: The ID of the task
- **Response**:
  - **Status**: 200 OK
  - **Body**: Deletion confirmation
  - **Status**: 403 Forbidden (if user tries to delete another user's task)
  - **Status**: 404 Not Found (if task not found)

## Sample Data for Testing

### Sample Task 1

```json
{
  "title": "Finish project report",
  "description": "Complete the final report for the project and submit it by the due date.",
  "due_date": "2024-08-15",
  "status": "pending"
}