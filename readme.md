# Amar Dokan

`Amar Dokan` is a Go backend for a small e-commerce system built with `Gin`, `GORM`, and `PostgreSQL`. It provides APIs for user registration and login, OTP-based account verification, JWT authentication, product management, category management, add-to-cart operations, and Swagger API docs.

## Features

- User registration with email OTP verification
- JWT access token and refresh token flow
- Protected routes with authentication middleware
- Product create, list, update, and soft delete APIs
- Category create and list APIs
- Add-to-cart create, list, update, and delete APIs
- PostgreSQL integration with auto migration
- Swagger documentation at `/api/v1/docs/index.html`
- Basic health check endpoint at `/health`

## Tech Stack

- `Go 1.25.6`
- `Gin` for HTTP routing
- `GORM` for ORM/database access
- `PostgreSQL` as the primary database
- `swaggo/gin-swagger` for API docs
- `godotenv` for environment loading

## Project Structure

```text
amar_dokan/
|-- app_error/         # Application-level error mapping
|-- assets/            # Static assets and landing/docs page
|-- cmd/               # App bootstrap
|-- config/            # App and database configuration
|-- controllers/       # HTTP handlers
|-- docs/              # Generated Swagger files
|-- infra/db/          # DB connection and migrations
|-- middleware/        # Auth, CORS, rate limiting
|-- models/            # Database models and request DTOs
|-- repositories/      # Data access layer
|-- routes/            # Route registration
|-- services/          # Business logic
|-- utils/             # JWT, password, OTP, response helpers
|-- main.go            # Entrypoint
```

## Environment Variables

Create a `.env` file in the project root with the following keys:

```env
VERSION=1.0.0
SERVICE_NAME=Amar Dokan
PORT=8000

JWT_SECURE_KEY=your-secret-key
JWT_EXPIRY_DAYS=2
REFRESH_JWT_EXPIRY_DAYS=30

DB_HOST=127.0.0.1
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your-db-password
DB_NAME=amar_dokan
DB_SSLMODE=disable

AppPass=your-gmail-app-password
SenderMail=your-email@gmail.com
```

## Prerequisites

- Install `Go 1.25+`
- Install and run `PostgreSQL`
- Create a database named `amar_dokan`
- Use a valid Gmail app password if you want OTP emails to work

## Getting Started

1. Clone the repository.
2. Create and fill the `.env` file.
3. Install dependencies:

```bash
go mod tidy
```

4. Run the application:

```bash
go run main.go
```

The server starts on:

```text
http://localhost:8000
```

On startup the app will:

- Load environment configuration
- Connect to PostgreSQL
- Run auto migrations
- Start the Gin router

## API Base URL

```text
http://localhost:8000/api/v1
```

## Main Endpoints

### Public

- `GET /health` - service health check
- `POST /api/v1/users` - register a new user and send OTP
- `POST /api/v1/users/verification` - verify OTP and activate account
- `POST /api/v1/users/login` - login and get tokens
- `GET /api/v1/users/refresh-token` - generate a new access token from refresh token
- `GET /api/v1/products` - list products
- `GET /api/v1/category/` - list categories
- `GET /api/v1/docs/index.html` - Swagger UI

### Protected

These routes require:

```http
Authorization: Bearer <access-token>
```

- `GET /api/v1/users` - list users
- `GET /api/v1/users/profile` - current user profile
- `GET /api/v1/users/:id` - get user by ID
- `PUT /api/v1/users/:id` - update user
- `DELETE /api/v1/users/:id` - delete user
- `POST /api/v1/products` - create product
- `PUT /api/v1/products/:id` - update product
- `DELETE /api/v1/products/:id` - delete product
- `POST /api/v1/category/` - create category
- `POST /api/v1/add-to-cart/` - add product to cart
- `GET /api/v1/add-to-cart/` - get current user's cart
- `PUT /api/v1/add-to-cart/:id` - update cart item quantity
- `DELETE /api/v1/add-to-cart/:id` - remove cart item

## Authentication Flow

1. Register with `POST /api/v1/users`
2. Receive OTP by email
3. Verify the OTP with `POST /api/v1/users/verification`
4. Receive an access token after successful verification
5. Use `POST /api/v1/users/login` for later logins
6. Use the refresh token endpoint to issue a new access token when needed

## Data Model Overview

### User

- `id`
- `name`
- `email`
- `password`
- `is_owner`
- `created_at`
- `updated_at`

### Product

- `id`
- `uid`
- `image_url`
- `name`
- `description`
- `price`
- `is_delete`

### Category

- `id`
- `uid`
- `image_url`
- `name`
- `is_delete`

### Add To Cart

- `id`
- `product_id`
- `user_id`
- `quantity`
- `is_delete`

## Swagger Documentation

Swagger is already integrated into the project. After starting the server, open:

```text
http://localhost:8000/api/v1/docs/index.html
```

Generated Swagger files are stored in the [`docs/`](./docs) directory.

## Notes

- The project uses auto migration on startup through GORM.
- Products and categories use soft-delete style flags via `is_delete`.
- OTP validity is currently set to `2 minutes`.
- CORS middleware is enabled globally.
- A rate limiter middleware exists in the codebase but is currently commented out in router setup.

## Example Run

```bash
go run main.go
```

Expected startup flow:

```text
Config loaded
Database connected
Migrations applied
Server running at http://localhost:8000
```

## Future Improvements

- Add owner/admin authorization rules for product and category mutation
- Add tests for services and handlers
- Add pagination and filtering for product listing
- Add category update and delete routes to the router
- Add Docker support and deployment config
- Add stronger validation and centralized logging
