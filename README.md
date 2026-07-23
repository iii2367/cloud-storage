# Cloud Storage

Cloud Storage is a full-stack file storage application written in Go.

The project focuses on backend architecture, authentication, and hierarchical file storage while providing a minimal web interface for interacting with the system.

## Appearance

<table align="center">
<tr>
<td align="center">
<img src="docs/login.png" width="420"><br>
<b>Login</b>
</td>
</tr>
<tr>
<td align="center">
<img src="docs/storage.png" width="900"><br>
<b>Storage</b>
</td>
</tr>
<tr>
<td align="center">
<img src="docs/upload_file.png" width="250"><br>
<b>Upload</b>
</td>
<td align="center">
<img src="docs/account.png" width="250"><br>
<b>Account</b>
</td>
<td align="center">
<img src="docs/file_info.png" width="250"><br>
<b>File Information</b>
</td>
</tr>
</table>

## Features

**Authentication**
- User registration
- Login / Logout
- JWT authentication
- Access & Refresh tokens
- Refresh token rotation
- User sessions

**Storage**
- Hierarchical folder structure
- File upload
- File download
- File deletion
- Folder creation
- Tree navigation

**Infrastructure**
- PostgreSQL persistence
- Docker support
- Configuration via environment variables

## Architecture

The application follows a modular architecture with dependency injection and a simplified layering inspired by Clean Architecture principles.

### Modules

The codebase is split into independent modules under `internal/`, each owning its own entities, repositories, services, and HTTP handlers:

| Module     | Responsibility                                                                |
| ---------- | ------------------------------------------------------------------------------ |
| `auth`     | User registration, login, sessions, JWT issuing/rotation, account management  |
| `storage`  | Hierarchical file/folder tree, upload, download, deletion                     |
| `jwt`      | Access/refresh token generation and validation                                |
| `config`   | Environment-based configuration loading                                      |
| `database` | PostgreSQL connection pool setup                                              |
| `web`      | Server-rendered HTML pages (login, signup, storage UI)                        |

### Layer responsibilities

Each business module (`auth`, `storage`) is internally split into the same set of layers:

- **entity** — plain domain models (`User`, `Session`, `TreeNode`)
- **repository** — persistence interfaces, with a `postgres` implementation
- **service** — business logic, independent of HTTP and storage details
- **transport/http** — Gin handlers, routing, and middleware that adapt HTTP requests to service calls
- **dto** — request/response payloads used at the HTTP boundary

### Dependency graph

Dependencies flow inward, toward the service layer, so business logic never depends on transport or persistence details:

```
transport/http  --->  service  --->  repository (interface)
                                          ^
                                          |
                              repository/postgres (implementation)
```

`cmd/app/main.go` wires everything together: it builds the repositories, injects them into the services, injects the services into the HTTP handlers, and registers routes on a single Gin engine.

### Authentication

Authentication is based on JWT using Access and Refresh tokens.
Refresh tokens are stored as sessions in PostgreSQL and rotated on every refresh request.
Access and Refresh tokens use different signing secrets and configurable expiration times.

A reused (already-rotated) refresh token is treated as a compromise signal: all sessions belonging to that user are revoked, forcing re-authentication.

### Storage Design

Files are stored directly on the server's file system.
PostgreSQL stores only metadata describing the storage tree.
Each record represents either a file or a folder and belongs to a specific user.
This separates physical file storage from the logical representation of the directory tree.

Every user has a dedicated root node, and every other node (file or folder) references its parent, forming a tree that can be queried and rendered incrementally.

## Technologies

| Technology          | Purpose          |
| ------------------- | ---------------- |
| Go                  | Backend          |
| Gin                 | HTTP Router      |
| PostgreSQL          | Metadata storage |
| JWT                 | Authentication   |
| Docker              | Containerization |
| HTML/CSS/JavaScript | Web interface    |

## Running locally

### Directory Layout

```
cloud-storage/
├── cmd/app/            # application entry point
├── internal/
│   ├── auth/           # authentication module (entity/repository/service/transport)
│   ├── storage/        # file storage module (entity/repository/service/transport)
│   ├── jwt/            # JWT manager
│   ├── config/         # environment configuration
│   ├── database/       # PostgreSQL connection
│   └── web/            # HTML page handlers
├── migrations/         # SQL migrations, applied automatically by the postgres container
├── web/
│   ├── static/         # CSS and JS assets
│   └── templates/      # HTML templates
├── docker/             # Dockerfile, docker-compose.yml, .env
└── docs/                # README screenshots
```

### Requirements

- Go 1.26 or newer
- PostgreSQL 17 (or compatible)
- Docker and Docker Compose (for the containerized setup)

### Environment variables

Configuration is loaded from environment variables prefixed with `MAIN_`. An example file is provided at `docker/.env`:

| Variable                  | Description                              |
| -------------------------- | ------------------------------------------ |
| `MAIN_SERVER_HOST`         | Host the HTTP server binds to             |
| `MAIN_SERVER_PORT`         | Port the HTTP server binds to             |
| `MAIN_DATABASE_HOST`       | PostgreSQL host                           |
| `MAIN_DATABASE_PORT`       | PostgreSQL port                           |
| `MAIN_DATABASE_USER`       | PostgreSQL user                           |
| `MAIN_DATABASE_PASSWORD`   | PostgreSQL password                       |
| `MAIN_DATABASE_DBNAME`     | PostgreSQL database name                  |
| `MAIN_JWT_ACCESS_SECRET`   | Signing secret for access tokens          |
| `MAIN_JWT_REFRESH_SECRET`  | Signing secret for refresh tokens         |
| `MAIN_JWT_ACCESSTTL`       | Access token lifetime (e.g. `15m`)        |
| `MAIN_JWT_REFRESHTTL`      | Refresh token lifetime (e.g. `2h`)        |
| `MAIN_JWT_ISSUER`          | JWT issuer claim                          |

> ⚠️ The values committed in `docker/.env` are examples for local development only. Replace the secrets and database credentials before deploying anywhere outside your machine.

### Docker

The easiest way to run the project is with Docker Compose, which starts both PostgreSQL and the application and applies migrations automatically:

```bash
cd docker
docker compose up --build
```

The application will be available at `http://localhost:8080`.

### Manual

1. Start a PostgreSQL instance and apply the SQL files in `migrations/` in order.
2. Export the environment variables listed above (or source `docker/.env`).
3. Run the application from the project root, so relative paths to `web/` resolve correctly:

```bash
go run ./cmd/app
```

## API Overview

### Auth

| Method     | Endpoint             | Description                                          |
| ---------- | --------------------- | ------------------------------------------------------ |
| **POST**   | `/api/auth/signup`   | Register a new user                                    |
| **POST**   | `/api/auth/login`    | Authenticate and issue an access/refresh token pair    |
| **POST**   | `/api/auth/refresh`  | Rotate the refresh token and issue a new token pair    |
| **POST**   | `/api/auth/logout`   | Revoke the current session                             |
| **GET**    | `/api/users/me`      | Get the authenticated user's profile                    |
| **DELETE** | `/api/users/me`      | Delete the authenticated user's account                 |

### Storage

| Method     | Endpoint                          | Description                                       |
| ---------- | ------------------------------------ | ---------------------------------------------------- |
| **GET**    | `/api/storage/tree`               | Get the authenticated user's root folder tree     |
| **GET**    | `/api/storage/tree/:id`           | Get the subtree of a specific folder               |
| **POST**   | `/api/storage/folders`            | Create a new folder                                |
| **POST**   | `/api/storage/root`               | Create the root node for the authenticated user     |
| **POST**   | `/api/storage/files`              | Upload a file                                      |
| **DELETE** | `/api/storage/nodes/:id`          | Delete a file or folder (and its descendants)       |
| **GET**    | `/api/storage/files/:id/download` | Download a file                                    |

All `auth` and `storage` endpoints (except signup/login/refresh/logout) require a valid access token.

## Design Decisions

- Physical files are stored separately from metadata.
- Authentication is isolated into its own module.
- Business logic is independent from HTTP handlers.
- Repository interfaces decouple persistence from services.
- Each user has an isolated storage tree rooted at a dedicated root node.
