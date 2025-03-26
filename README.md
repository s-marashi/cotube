# modular-monolith

This is a modular monolith application built in Go. The project is structured to facilitate scalability and maintainability by organizing code into distinct modules.

## Project Structure

```
modular-monolith
├── cmd
│   └── server
│       └── main.go          # Entry point of the application
├── internal
│   ├── core
│   │   ├── domain
│   │   │   └── models.go    # Domain models
│   │   └── ports
│   │       └── interfaces.go # Core application ports interfaces
│   ├── modules
│   │   ├── module1
│   │   │   ├── handlers      # HTTP handlers for module1
│   │   │   ├── repository    # Data access layer for module1
│   │   │   └── service       # Business logic for module1
│   │   └── module2
│   │       ├── handlers      # HTTP handlers for module2
│   │       ├── repository    # Data access layer for module2
│   │       └── service       # Business logic for module2
│   └── pkg
│       ├── config
│       │   └── config.go     # Application configuration
│       └── database
│           └── connection.go  # Database connection management
├── pkg
│   └── common
│       └── utils.go          # Utility functions
├── go.mod                    # Module dependencies
├── go.sum                    # Dependency checksums
└── README.md                 # Project documentation
```

## Setup Instructions

1. **Clone the repository:**
   ```bash
   git clone <repository-url>
   cd modular-monolith
   ```

2. **Install dependencies:**
   ```bash
   go mod tidy
   ```

3. **Run the application:**
   ```bash
   go run cmd/server/main.go
   ```

## Usage

- The application exposes various endpoints defined in the HTTP handlers of each module.
- Refer to the specific module documentation for details on available routes and their usage.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any enhancements or bug fixes.