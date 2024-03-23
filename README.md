# Go REST API

This is a REST API application written in Go that uses the following technologies:

- **Go**: An open-source programming language that makes it easy to build efficient and reliable applications.
- **gORM**: An Object-Relational Mapping (ORM) library for Go that makes it easy to manipulate data in relational databases.
- **Echo**: A lightweight and high-performance framework for building web applications and APIs in Go.

## Endpoints

The API offers the following endpoints:

### Users

- **GET /users**: Get the list of all users.
- **GET /users/:id**: Get a specific user by their ID.
- **POST /users**: Create a new user.
- **PUT /users/:id**: Update an existing user.
- **DELETE /users/:id**: Delete an existing user.

### Tasks

- **GET /tasks**: Get the list of all tasks.
- **GET /tasks/:id**: Get a specific task by its ID.
- **POST /tasks**: Create a new task.
- **PUT /tasks/:id**: Update an existing task.
- **DELETE /tasks/:id**: Delete an existing task.

Each endpoint accepts and returns data in JSON format.

## Installation

1. Clone the repository
2. Run go application:
    - linux
    ```
    go build -o app && ./app
    ```
    - windows
    ```
    go build -o app.exe; ./app.exe
    ```

3. The application will be available at `http://localhost:5000`.

Enjoy using the API!

