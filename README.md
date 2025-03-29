# Test Case NDI

## Table of Contents

- [Installation](#installation)
- [Running the Application](#running-the-application)
- [API Endpoints](#api-endpoints)
- [Authentication](#authentication)
- [Error Handling](#error-handling)

## Installation

1. Clone the repository:
```bash
git clone https://github.com/fathxn/test-case-ndi.git
cd bank-app
```

2. Install dependencies:
```bash
make deps
```

## Running the Application

To build and run the application:
```bash
make run
```

## API Endpoints

The API provides the following endpoints:

### Public Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/login`   | Login and get authentication token |
| GET  | `/user/:id`| Get user information by ID (only returns username) |

### Protected Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
|  GET   |`/balance`| Get the authenticated user's balance |

## Authentication

The API uses JWT (JSON Web Token) for authentication. To access protected endpoints, you need to:

1. First, log in using the `/login` endpoint with valid credentials.
2. Include the received token in the `Authorization` header of your requests to protected endpoints.

Example:
```
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```

## Endpoint Details

### `POST /login`

Authenticate a user and get a JWT token.

**Request Body:**
```json
{
  "username": "orangpertama",
  "password": "password123"
}
```

**Response Example:**
```json
{
  "status": "success",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "user": {
    "id": 1,
    "username": "orangpertama"
  }
}
```

**Error Responses:**
- 400 Bad Request - If username or password is missing
- 401 Unauthorized - If credentials are invalid

### `GET /user/:id`

Get user information by ID. This endpoint only returns the username for security reasons.

**Response Example:**
```json
{
  "username": "orangpertama"
}
```

**Error Responses:**
- 400 Bad Request - If ID is not a valid number
- 404 Not Found - If user with the provided ID doesn't exist

### `GET /balance`

Get the authenticated user's balance. Requires authentication.

**Response Example:**
```json
{
  "status": "success",
  "user": "orangpertama",
  "balance": 1000.50
}
```

**Error Responses:**
- 401 Unauthorized - If token is missing or invalid
- 404 Not Found - If user does not exist

## Error Handling

The API returns appropriate HTTP status codes along with error messages in the response body:

```json
{
  "error": "Error message description"
}
```

## Test Users

For development and testing purposes, the application comes with two predefined users:

1. **User 1**
   - Username: orangpertama
   - Password: password123
   - Balance: 1000.50

2. **User 2**
   - Username: orangkedua
   - Password: password123
   - Balance: 2500.75
---