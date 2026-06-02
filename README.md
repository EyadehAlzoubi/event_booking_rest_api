# Event Booking REST API

A simple REST API built with Go, Gin, and SQLite for managing events with user authentication.

## Features
- User signup and login
- Password hashing with bcrypt
- JWT-based authentication
- Create, read, update, and delete events
- SQLite database with automatic table creation

## Tech Stack
- Go
- Gin (`github.com/gin-gonic/gin`)
- SQLite (`github.com/mattn/go-sqlite3`)
- JWT (`github.com/golang-jwt/jwt/v5`)

## Quick Start
```bash
go mod tidy
go run .
```

Server runs on: `http://localhost:8080`

## Main Endpoints
- `POST /signup` - Create a new user
- `POST /login` - Login and get JWT token
- `GET /events` - Get all events
- `GET /events/:id` - Get event by ID
- `POST /events` - Create event (requires `Authorization` token)
- `PUT /events/:id` - Update event
- `DELETE /events/:id` - Delete event

## Notes
- Database file is `api.db` (created automatically).
- Add your JWT token in the `Authorization` header for protected routes.
