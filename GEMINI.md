# Gemini Code Assistant Guidance for `isbetmf`

This document provides guidance for Gemini Code Assistant to understand the `isbetmf` project, its architecture, and coding conventions.

## Project Overview

`isbetmf` is a [TM Forum (TMF) Open API](https://www.tmforum.org/oda/open-apis/directory) server written in Go. It is designed to be highly flexible and can operate in two distinct modes:

1.  **Standalone Server:** Implements the TMF Open API specifications and manages entities in its own database (SQLite).
2.  **Proxy Server:** Acts as a Policy Enforcement Point (PEP) and Policy Decision Point (PDP) in front of a remote TM Forum API server. In this mode, it forwards requests to the upstream server while enforcing authentication, authorization, and other business rules.

The primary goal is to provide a production-ready, compliant, and extensible TMF API implementation.

## Architecture

The project follows a clean, layered architecture designed for modularity and testability. The core idea is to separate the HTTP transport layer from the business logic.

### 1. Entrypoint (`main.go`)

*   The `main` function in the root `main.go` is the application entrypoint.
*   It handles command-line flag parsing, configuration loading, and initializes all major components.
*   It uses the `tableflip` library to manage graceful, zero-downtime server restarts and upgrades.

### 2. Handler Layer (`tmfserver/handler/fiber`)

*   This is the thinnest layer, responsible for handling incoming HTTP requests.
*   It uses the [Fiber](https://gofiber.io/) web framework.
*   **Key Responsibility:** To translate the Fiber-specific request context (`fiber.Ctx`) into a generic, framework-agnostic `service.Request` struct. It then passes this request to the Service Layer.
*   After the service layer processes the request, the handler converts the returned `service.Response` back into an HTTP response.
*   **Generic Routes:** The routing, defined in `routes.go`, is generic. It uses URL parameters like `:apiFamily`, `:resourceName`, and `:id` to handle all TMF resources with a small set of common handlers (e.g., `CreateGenericObject`, `GetGenericObject`). This is a fundamental concept of the project.
*   **TMF API Path Structure:** The request path for a given TMF resource generally follows these patterns:
    ```
    /tmf-api/{apiFamily}/{apiVersion}/{resourceName}/{resourceID}?{queryParams}
    ```
    or
    ```
    /tmf-api/{apiFamily}/{apiVersion}/{resourceName}?{queryParams}
    ```
    -   `apiFamily`: The group of APIs (e.g., `productCatalogManagement`).
    -   `apiVersion`: The version of the API (e.g., `v4`, `v5`).
    -   `resourceName`: The specific resource being managed (e.g., `productOffering`).
    -   `resourceID`: The unique identifier of the resource (optional for POST requests).
    -   `queryParams`: Additional parameters for filtering, attribute selection, or pagination.
*   **Swagger UI:** The server also serves the Swagger UI for interactive API documentation and experimentation, available at paths like `/oapiv4/index.html` and `/oapiv5/index.html`.

### 3. Service Layer (`tmfserver/service`)

*   This is the core of the application, containing all business logic. It is completely decoupled from the Fiber framework.
*   The central component is the `Service` struct, which orchestrates all operations.
*   **Key Responsibilities:**
    *   Receives a `service.Request` and returns a `service.Response`.
    *   Performs authentication by processing JWT access tokens.
    *   Executes authorization logic by consulting the Policy Decision Point (PDP).
    *   Interacts with the data layer (either the local database or the remote TMF proxy).
    *   Manages the notification hub.
*   **Generic Request Processing:** For all TMF API requests (GET, POST, PATCH, DELETE), the service layer typically performs the following steps:
    1.  **Authentication:** Validates the provided access token (if any) and extracts user information. Read APIs can be used without authentication, but if a token is provided, it must be valid.
    2.  **Data Retrieval/Storage:**
        *   For read operations (GET, LIST), it retrieves the object(s) from the local database or forwards the request to the remote TMF server if in proxy mode.
        *   For write operations (POST, PATCH, DELETE), it interacts with the local database.
    3.  **Authorization Policies:** Calls the PDP to check if the authenticated user is authorized to perform the requested action on the specific resource.
    4.  **Object Filtering:** For read operations, it filters the returned object(s) based on `queryParams` (e.g., `fields` attribute), always including `id`, `href`, `lastUpdate`, `version`, `lifecycleStatus`.

### 4. Repository Layer (`tmfserver/repository`)

*   This layer abstracts all database interactions.
*   It is responsible for CRUD (Create, Read, Update, Delete) operations on TMF objects in the SQLite database.
*   It uses the `sqlx` library for database access.

### 5. Policy Engine (`pdp`)

*   The Policy Decision Point (PDP) is responsible for fine-grained authorization.
*   It uses the [Starlark](https://github.com/google/starlark-go) language, allowing operators to define custom authorization rules in `.star` files (see `auth_policies.star`).
*   The service layer calls the PDP to decide whether a given user/token has permission to perform an action on a resource.

## Coding Style and Conventions

*   **Modularity:** Keep layers decoupled. The service layer should never have knowledge of the `fiber.Ctx` or any other HTTP-specific constructs.
*   **Error Handling:** Use the custom `errl` package for annotating and wrapping errors to provide context.
*   **Configuration:** Configuration is managed in the `config` package. It uses a profile-based approach (`-run` flag) rather than numerous individual environment variables.
*   **Generics:** Embrace the "generic" approach. When adding new functionality, consider if it can be implemented in the generic handlers and service methods before creating specific ones.
*   **Dependencies:**
    *   **Web Framework:** `github.com/gofiber/fiber/v2`
    *   **Database:** `github.com/jmoiron/sqlx` with `github.com/mattn/go-sqlite3` driver.
    *   **Graceful Restarts:** `github.com/cloudflare/tableflip`
    *   **Policy Rules:** `go.starlark.net/starlark`

## How to Run

The application is started via `go run main.go`. The `-run` flag is important for selecting the correct configuration profile. For example:

```bash
go run main.go -run mycredential
```
