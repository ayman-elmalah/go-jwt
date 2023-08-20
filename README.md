# Go JWT User Authentication Starter Kit

Welcome to the Go JWT User Authentication Starter Kit! This project provides a strong foundation for building a web application with user authentication using JSON Web Tokens (JWT). It is built upon the [Go Starter Kit](https://github.com/ayman-elmalah/go-starter-kit) skeleton, offering a well-organized starting point for your Go web application development.

Developed by Ayman Elmalah. You can find more of Ayman's work on [GitHub](https://github.com/ayman-elmalah).

## Features

- **User Authentication:** Implement user registration and login functionality using JWT for secure authentication.

- **Clean Code Structure:** The project follows a structured code layout, which separates concerns and promotes easy management and extension of your application.

- **Dependency Management:** Utilizes Go Modules for efficient dependency management, allowing seamless integration of third-party libraries.

- **Robust Middleware:** Includes a middleware for JWT-based authentication, ensuring secure access to protected routes.

## Installation

To install and get started with the **go-starter-kit**, follow these steps:

1. Clone the repository:

```sh
git clone https://github.com/ayman-elmalah/go-jwt
```

Change to the project directory:

```sh
cd go-jwt
```

## Prerequisites

Before you start using the **go-jwt**, make sure you have the following prerequisites installed:

1. **CompileDaemon**: A tool to watch and rebuild your Go application on changes. Install it using the following command:

```sh
go install -mod=mod github.com/githubnemo/CompileDaemon
```

2. **golang-migrate**: A tool to manage database migrations in your Go project. Follow the installation guide [here](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate) to install it.

Then run the following command to tidy up your Go modules:

```sh
go mod tidy
```

## Configuration

Copy the example configuration file:

```sh
cp config/config.yaml.example config/config.yaml
```

Open the `config/config.yaml` file and provide your data and database credentials.

## Starting the Project

To start the project, use the following command:

```sh
CompileDaemon -build="go build -o main main.go" -command="./main serve"
```

## Database Migrations

Follow these steps to create and manage database migrations:

### Applying Migrations

To apply migrations, use the following commands:

- Migrate Up:

```sh
go run main.go migrate-up
```

- Migrate Down

```sh
go run main.go migrate-down
```

For more advanced migration features and options, refer to the [golang-migrate guide](https://github.com/golang-migrate/migrate).

# Using Postman Collections

To make it easier for developers to interact with the API, we've included a set of Postman collections in this project. You can find the collections in the `postman` directory, specifically `collection.json`. These collections provide pre-configured API requests that you can import into Postman for testing and exploring the endpoints.
