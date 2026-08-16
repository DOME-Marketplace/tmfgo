# TM Forum API Server

[![Go Report Card](https://goreportcard.com/badge/github.com/hesusruiz/tmforum)](https://goreportcard.com/report/github.com/hesusruiz/tmforum)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](LICENSE)

`isbetmf` is a [TM Forum (TMF) Open API](https://www.tmforum.org/oda/open-apis/directory) server written in Go. It is designed to be highly flexible and compliant with TMF standards.

## Features

*   **Dual Operation Modes:**
    *   **Standalone Server:** Implements TMF Open API specifications using a local SQLite database.
    *   **Proxy Server:** Acts as a gateway in front of a remote TMF API server.
*   **Policy Enforcement:** Acts as a Policy Enforcement Point (PEP) and Policy Decision Point (PDP), enforcing authentication and fine-grained authorization using Starlark scripts.
*   **Reporting & Validation:** Includes a dedicated tool (`tmf-reporting`) for validating TMF API implementations and generating compliance reports.
*   **Replication:** Supports data replication capabilities (see `replicate` directory).
*   **Conformance:** Passes the official TMF Conformance Test Kit for TMF V4 and V5 APIs.
*   **Zero-Downtime Updates:** Uses `tableflip` for graceful restarts and upgrades.

## Getting Started

### Prerequisites

*   [Go](https://go.dev/dl/) 1.24 or higher
*   [Docker](https://www.docker.com/) (for containerized deployment)
*   [Make](https://www.gnu.org/software/make/) (optional, for build automation)

### Local Development

To run the server locally for development:

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/hesusruiz/tmforum.git
    cd isbetmf
    ```

2.  **Run with default configuration:**
    ```bash
    go run main.go -run isbedev
    ```
    Or using Make:
    ```bash
    make run
    ```
    This starts the server using the `isbedev` profile (configured in `config/config_data.go`), which typically uses a local SQLite database.

3.  **Build the binaries:**
    ```bash
    make build          # Builds the server (bin/isbetmf)
    make build-reporting # Builds the reporting tool (bin/tmf-reporting)
    ```

4.  **Run Tests:**
    ```bash
    make test
    ```

### Docker Deployment

The application is designed to be containerized.

1.  **Build the container:**
    ```bash
    docker build -t isbetmf .
    ```

2.  **Run the container:**
    The container requires the `ISBETMF_RUN_ENVIRONMENT` environment variable to select the configuration profile.
    ```bash
    docker run -e ISBETMF_RUN_ENVIRONMENT=isbedev -p 9991:9991 isbetmf
    ```
    Possible values for `ISBETMF_RUN_ENVIRONMENT`: `isbedev`, `isbepre`, `isbepro`, `domedev`, `domepre`, `domepro`.

3.  **The database and backups:**
    The application uses SQLite as the database. The database is stored in the `/data` directory of the container. To persist the database in the host filesystem, you can use a volume mapped to the `/data` directory of the container.

    The system maintains backups of the database in the `/data/backups` directory of the container. It maintains a rotating set of 7 backups, one for each day of the week. The backups contain the number of the day of the week in the name (e.g., `backup_01.db`, `backup_02.db`, etc.). The backups are performed every 2 hours at even hours (e.g., 11:00, 13:00, etc.) in the backup file corresponding to the day of the week. Storing the backups in a different system (e.g., in a cloud storage service) is as simple as copying the `/data/backups` directory (or just the latest backup file) to the desired location. To ensure that there are no conflicts with the scheduled updates of the backups, it is recommended to make the copy on even hours (e.g., 10:00, 12:00, etc.).

    ```bash
    docker run -e ISBETMF_RUN_ENVIRONMENT=isbedev -p 9991:9991 -v /path/to/your/database:/data isbetmf
    ```
    This will persist the database in the `/data` directory of the container.

## Configuration

### Profiles

Configuration is managed via profiles defined in `config/config_data.go`. This approach reduces configuration errors by grouping settings into well-defined environments.

To add or modify a profile, edit `config/config_data.go`. You can specify the profile to use at runtime with the `-run` flag:
```bash
./bin/isbetmf -run <profile_name>
```

### Policy Engine (PDP)

Authorization rules are defined in Starlark scripts (e.g., `auth_policies.star`). This allows for dynamic and complex permission logic without recompiling the server. The PDP evaluates these rules for every request to determine access.

## Architecture

The project follows a layered architecture:

1.  **Entrypoint (`main.go`):** Handles initialization, config loading, and graceful restarts.
2.  **Handler Layer (`tmfserver/handler/fiber`):** Manages HTTP requests using Fiber, translating them into transport-agnostic service requests. Handles generic TMF routing.
3.  **Service Layer (`tmfserver/service`):** Contains the core business logic, decoupled from HTTP. Handles authentication, authorization (via PDP), and orchestration.
4.  **Repository Layer (`tmfserver/repository`):** Abstracts database interactions (SQLite).
5.  **Policy Engine (`pdp`):** Executes Starlark rules for authorization.

## Contributing

Contributions are welcome! Please follow these steps:

1.  Fork the repository.
2.  Create a feature branch (`git checkout -b feature/amazing-feature`).
3.  Commit your changes (`git commit -m 'Add amazing feature'`).
4.  Push to the branch (`git push origin feature/amazing-feature`).
5.  Open a Pull Request.

Please ensure you run tests (`make test`) and lint your code (`make lint`) before submitting.

## License

This project is licensed under the Apache 2.0 License - see the [LICENSE](LICENSE) file for details.
