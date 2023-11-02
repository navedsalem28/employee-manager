
- `/cmd`: Contains the entry point of the application, which sets up the HTTP server.
- `/internal`: Contains the core project logic, including handlers, models, and repositories.
- `/configs`: Stores the configuration file (e.g., database configuration) in `config.yaml`.
- `/scripts`: Contains the SQL script for initial database setup in `setup.sql`.
- `Dockerfile`: Defines the Docker image for the application.
- `go.mod` and `go.sum`: Define project dependencies.

## Implementation Details

### Database Setup

- GORM is used to define the Employee model and auto-migrate the database schema.

### Employee Management

- CRUD operations for Employee entities, including Create, Read, Update, and Delete.
- Echo Framework is used to set up routes for these operations.

### Validation

- Input validation is implemented for Employee data to ensure data integrity.

### Error Handling

- Proper error responses are provided for various failure cases (e.g., not found, bad request).

### Logging and Middleware

- Basic request and response logging are implemented.

### Configuration

- Database configuration and other settings are loaded from `config.yaml`.

## Running the Project

To run the project, follow these steps:

1. Initialize the project and install dependencies:

   ```bash
   go mod tidy
   go run cmd/main.go
   ./main

## API Endpoints
### Employee CRUD:
1. **GET** /employees<span style="white-space: pre">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</span> Get all employees
2. **GET** /employees/{id}<span style="white-space: pre">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</span> Get a employee by ID
3. **POST** /employees<span style="white-space: pre">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</span> Create a new employee
4. **PUT** /employees/{id}<span style="white-space: pre">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</span> Update an existing employee
5. **DELETE** /employees/{id}<span style="white-space: pre">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</span> Delete a employees by ID


To run the project, follow these steps:
#####    Initialize the project with Docker Composer:
   ```bash
        go test ./test/...
    ./service.sh

