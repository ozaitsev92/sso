![Repository Top Language](https://img.shields.io/github/languages/top/ozaitsev92/sso)
![Github Repository Size](https://img.shields.io/github/repo-size/ozaitsev92/sso)
![Github Open Issues](https://img.shields.io/github/issues/ozaitsev92/sso)
![License](https://img.shields.io/badge/license-MIT-green)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/ozaitsev92/sso)
![GitHub last commit](https://img.shields.io/github/last-commit/ozaitsev92/sso)
![GitHub contributors](https://img.shields.io/github/contributors/ozaitsev92/sso)

---

# gRPC SSO Service

A lightweight Single Sign-On (SSO) microservice implemented in Go using gRPC.

---

## Project Structure

```

├── cmd                # Entrypoints
│   ├── migrator       # Runs DB migrations
│   └── sso            # Starts the gRPC server
├── config             # YAML configuration files
├── internal           # Application logic
│   ├── app            # App startup and gRPC middleware
│   ├── config         # Config parser
│   ├── domain         # Domain models
│   ├── grpc           # gRPC server implementation
│   ├── lib/jwt        # JWT logic
│   ├── services/auth  # Auth business logic
│   └── storage        # SQLite implementation
├── migrations         # SQL migration files
├── tests              # Functional test suite

```

---

## Getting Started

### Prerequisites

- [Docker](https://www.docker.com/)
- [Task](https://taskfile.dev) (optional, for automation)

---

### 1. Clone the Repository

```bash
git clone https://github.com/ozaitsev92/sso.git
cd sso
```

### 2. Run with Docker

```bash
task run
```

This builds and starts the container, then runs database migrations automatically.

---

## gRPC API Reference

The `Auth` service provides the following methods:

### 📌 `Register`

**RPC:** `Register(RegisterRequest) returns (RegisterResponse)`

Registers a new user.

**Request:**

```protobuf
message RegisterRequest {
  string email = 1;
  string password = 2;
}
```

**Response:**

```protobuf
message RegisterResponse {
  int64 user_id = 1;
}
```

---

### 🔐 `Login`

**RPC:** `Login(LoginRequest) returns (LoginResponse)`

Authenticates a user and returns a JWT token.

**Request:**

```protobuf
message LoginRequest {
  string email = 1;
  string password = 2;
  int64 app_id = 3;
}
```

**Response:**

```protobuf
message LoginResponse {
  string token = 1;
}
```
