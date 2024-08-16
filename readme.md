# Concurrent Ticketing System

A ticket allocation system that allows users to purchase tickets concurrently. The system ensures that tickets are allocated correctly and that no purchase exceeds more than the available tickets.

Built with Go Fiber, GORM, and PostgreSQL

## Getting Started

### Prerequisites

- Docker
- Docker Compose

### Running the Project

To build and start the application, use:

```bash
docker compose up --build
```

This command starts both the application and PostgreSQL database.

### Running Tests

To execute the tests, use:

```bash
docker compose --profile test up test --build --force-recreate
```

## API Documentation

Access the API documentation (Swagger) at:

[http://localhost:3000/swagger](http://localhost:3000/swagger)

## Project Structure

- config/: Manages database connection & migrations. Also loads environment variables.
- dal/: Data Access Layer for database operations.
- docs/: OpenAPI (Swagger) documentation files.
- routes/: API route definitions.
- services/: Business logic for ticket and purchase operations.
- tests/: Unit tests with mocked database interactions.
- types/: Custom types for request/response validation.
- utils/: Helper functions for error handling and validation.
- Dockerfile: Docker configuration for building the application.
- docker-compose.yml: Docker Compose configuration for development and testing.
- main.go: Application entry point.

### Usage

- Create Ticket: POST /tickets - Add a new ticket.
- Get Ticket: GET /tickets/{id} - Retrieve ticket details.
- Create Purchase: POST /tickets/{id}/purchases - Purchase tickets.
