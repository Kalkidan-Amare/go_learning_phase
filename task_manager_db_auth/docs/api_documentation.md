# Task Management API Documentation

This API provides CRUD operations for managing tasks, now with persistent data storage using MongoDB.

## Postman Documentation

For detailed information and testing endpoints, please refer to the [Postman documentation](https://documenter.getpostman.com/view/37352369/2sA3s1nWpq).

## Endpoints

### 1. GET /tasks
- **Description**: Retrieve a list of all tasks.
- **Response**:
  - **Status**: 200 OK
  - **Body**: List of tasks

### 2. GET /tasks/:id
- **Description**: Retrieve the details of a specific task.
- **Parameters**:
  - `id`: The ID of the task
- **Response**:
  - **Status**: 200 OK
  - **Body**: Task details
  - **Status**: 404 Not Found (if task not found)

### 3. POST /tasks
- **Description**: Create a new task.
- **Request**:
  - **Body**: JSON object containing title, description, due date, and status
- **Response**:
  - **Status**: 201 Created
  - **Body**: Created task details

### 4. PUT /tasks/:id
- **Description**: Update a specific task.
- **Parameters**:
  - `id`: The ID of the task
- **Request**:
  - **Body**: JSON object containing new task details
- **Response**:
  - **Status**: 200 OK
  - **Body**: Updated task details
  - **Status**: 404 Not Found (if task not found)

### 5. DELETE /tasks/:id
- **Description**: Delete a specific task.
- **Parameters**:
  - `id`: The ID of the task
- **Response**:
  - **Status**: 200 OK
  - **Body**: Deletion confirmation
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
