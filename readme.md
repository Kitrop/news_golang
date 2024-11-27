# 📰 News-Go: A News Aggregator Application  
![News-Go Logo](newsgo_logo.png) <!-- Replace with actual logo URL -->

[![Go Version](https://img.shields.io/badge/Go-1.20+-blue)](https://golang.org)
[![Gin Framework](https://img.shields.io/badge/Framework-Gin-blue)](https://gin-gonic.com/)
[![PostgreSQL](https://img.shields.io/badge/Database-PostgreSQL-blue)](https://www.postgresql.org/)
[![License](https://img.shields.io/badge/License-MIT-green)](LICENSE)
[![Contributions Welcome](https://img.shields.io/badge/Contributions-Welcome-brightgreen)](#contributing)

News-Go is a robust and scalable news aggregator application built with Go and the Gin framework. It provides a secure, user-friendly platform for news management, user authentication, and admin controls.

---

## ✨ Features

- **🔒 User Authentication**  
  - Secure user registration and login with JWT (JSON Web Tokens).  
  - Password hashing with `bcrypt` for added security.  
  - Includes logout functionality.

- **📰 News Management**  
  - Create and view news posts.  
  - Admin panel for advanced controls.

- **👨‍💼 Admin Panel**  
  - Dedicated interface for managing users and news posts.  
  - Delete news posts and ban users with ease.

- **✅ Robust Validation**  
  - Input validation for user registration and password updates.  
  - Email validation and strong password checks.

- **💾 Database Integration**  
  - Powered by PostgreSQL with GORM as the ORM.

- **🛡️ Middleware**  
  - Custom middleware for authentication and authorization.
  - Middleware for collect user metadata 
  - Middleware for perfomance monitor

---

## 🏛️ Architecture

The application follows a clean, layered architecture:

1. **Controllers**: Handle HTTP requests and responses.  
2. **Services**: Contain business logic and interact with repositories.  
3. **Repositories**: Handle database operations, abstracting implementation details.  
4. **Models**: Define the core data structures (`User`, `News`).  
5. **Middleware**: Manage authentication (e.g., `Autification`, `UnAuthorized`) and authorization (e.g., `CheckIsAdmin`).  
6. **Utils**: Helper functions for JWT, password hashing, and validation.  
7. **Config**: Manages environment variables and database connections.

---

## 📌 Endpoints

### 👤 User Endpoints

| Method | Endpoint                | Description                   | Authentication |
|--------|--------------------------|-------------------------------|----------------|
| POST   | `/users/create`          | Creates a new user.           | Not required   |
| GET    | `/users/all`             | Retrieves all users.          | Required       |
| POST   | `/users/login`           | Logs in an existing user.     | Not required   |
| GET    | `/users/logout`          | Logs out the current user.    | Required       |
| POST   | `/users/change-password` | Changes the user's password.  | Required       |

### 📰 News Endpoints

| Method | Endpoint        | Description            | Authentication |
|--------|------------------|------------------------|----------------|
| POST   | `/news/new`      | Creates a new news post. | Required       |
| GET    | `/news/all`      | Retrieves all news posts. | Required       |

### 👨‍💼 Admin Endpoints

| Method | Endpoint             | Description             | Authentication  |
|--------|-----------------------|-------------------------|-----------------|
| POST   | `/admin/deleteNews`  | Deletes a news post.    | Required, Admin |
| GET    | `/admin/banUser`     | Bans a user.            | Required, Admin |

---

## 🛠️ Technology Stack

- **Go**: High-performance programming language.  
- **Gin**: Fast and flexible web framework.  
- **PostgreSQL**: Reliable and powerful database.  
- **GORM**: Object-Relational Mapper for database integration.  
- **JWT**: Secure authentication using JSON Web Tokens.  
- **bcrypt**: Industry-standard password hashing.

---

## 🚀 Setup and Run

1. **Clone the repository**:  
  ```bash
   git clone https://github.com/Kitrop/news_golang.git
  ```
2. **Install Dependencies**:  
  ```bash
   go mod tidy
  ```
3. **Run the application**:  
  ```bash
   go run main.go
  ```
