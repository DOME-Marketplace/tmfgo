# TM Forum API Server

[![Signed with Sigstore](https://img.shields.io/badge/Signed%20with-Sigstore-blue?logo=sigstore)](https://www.sigstore.dev)
[![Rekor Transparency Log](https://img.shields.io/badge/Provenance-Rekor%20Log-green?logo=sigstore)](https://rekor.sigstore.dev)
[![Release Workflow](https://github.com/hesusruiz/tmforum/actions/workflows/release.yml/badge.svg)](https://github.com/hesusruiz/tmforum/actions/workflows/release.yml)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/hesusruiz/tmforum.svg)](https://pkg.go.dev/github.com/hesusruiz/tmforum)
[![GHCR](https://img.shields.io/badge/GHCR-tmforum-blue?logo=github)](https://github.com/hesusruiz/tmforum/pkgs/container/tmforum)


`tmforum` is a [TM Forum (TMF) Open API](https://www.tmforum.org/oda/open-apis/directory) server written in Go, which can operate in two modes.

In **proxy** mode it runs in front of a remote TMF API server, adding **authentication** and **authorization**. The server includes both a [Policy Enforcement Point (PEP)](https://csrc.nist.gov/glossary/term/policy_enforcement_point) and a [Policy Decision Point (PDP)](https://csrc.nist.gov/glossary/term/policy_decision_point), enforcing authentication and fine-grained authorization using [Starlark](https://starlark-lang.org/) scripts.
In this mode, `tmforum` acts also as a smart **cache** in front of the remote server, reducing the load on it and improving the response times.

In **standalone** mode it is a full self-contained TM Forum server implemented in Go, as a single binary simple to install and operate. It uses few resources (compared to other implementations like ones based in Java), and it is also fast to start (less than 1 second). In this mode the authentication and authorization features provided by the PEP/PDP are also available.

The server passes the [TM Forum Conformance Test Kit](https://github.com/tmforum-randd/CTK) tests for TMF V4 APIs.

There are several ways to run the server. The simplest is to use the already built container image in
[GHCR](https://github.com/hesusruiz/tmforum/pkgs/container/tmforum), which is [**signed using Sigstore Cosign**](https://www.sigstore.dev/) to provide traceability and verifiability of the software supply chain.

Alternatively, the software can be built and run from the source code, integrating with your CI/CD system.


## Using the pre-built container image

The container image at its different versions is available at [GHCR](https://github.com/hesusruiz/tmforum/pkgs/container/tmforum) and can be run as follows:

```bash
docker run -p 9991:9991 ghcr.io/hesusruiz/tmforum:<version>
```

For example:

```bash
docker run -p 9991:9991 ghcr.io/hesusruiz/tmforum:v1.0.2
```

The container can be configured as described in [Configuring and running](#configuring-and-running).

### Verifying the Signed Container Image

This step is not strictly required to run the server, but it is recommended for security reasons.

All container images published for this project are [**signed using Sigstore Cosign**](https://www.sigstore.dev/) in keyless mode.  
This allows anyone to verify that an image was built by the official GitHub Actions workflow and is recorded in the public [Rekor transparency log](https://rekor.sigstore.dev/).

#### 1. Install Cosign

Cosign is a single binary and does **not** require Go:

**Direct download (Linux/macOS/Windows):**
```sh
curl -sSL -o cosign https://github.com/sigstore/cosign/releases/latest/download/cosign-linux-amd64
chmod +x cosign
sudo mv cosign /usr/local/bin/
```

(Replace `linux-amd64` with `darwin-amd64` or `windows-amd64.exe` as needed.)

**macOS (Homebrew):**
```sh
brew install cosign
```

**Linux (APT):**
```sh
sudo apt-get install cosign
```

#### 2. Retrieve the image digest

Cosign verifies **digests**, not tags.  
To get the digest for any version:

```sh
cosign triangulate ghcr.io/hesusruiz/tmforum:<version>
```
Where `<version>` is the tag (e.g., `v1.0.2`) of the image you want to verify.

Example:

```sh
cosign triangulate ghcr.io/hesusruiz/tmforum:v1.0.2
```

This prints something like:

```
ghcr.io/hesusruiz/tmforum@sha256:<digest>
```

Copy the digest for the next step.

#### 3. Verify the signature

Use the digest you obtained:

```sh
cosign verify \
  --certificate-oidc-issuer "https://token.actions.githubusercontent.com" \
  --certificate-identity-regexp "https://github.com/hesusruiz/tmforum/.github/workflows/.*" \
  ghcr.io/hesusruiz/tmforum@sha256:<digest>
```

If verification succeeds, Cosign will show:

- the Fulcio certificate  
- the GitHub workflow identity  
- the Rekor entry UUID  
- a “Verified OK” message  


#### 4. Verification Script (optional)

Users may run this script to verify any version automatically.  
Save it as `verify.sh`:

```sh
#!/usr/bin/env bash
set -euo pipefail

if [ $# -ne 1 ]; then
    echo "Usage: $0 <version>"
    echo "Example: $0 v1.0.2"
    exit 1
fi

VERSION="$1"
IMAGE="ghcr.io/hesusruiz/tmforum:${VERSION}"

# Check cosign availability
if ! command -v cosign >/dev/null 2>&1; then
    echo "Error: cosign is not installed."
    echo "Install it from https://github.com/sigstore/cosign/releases/latest"
    exit 1
fi

echo "Retrieving digest for ${IMAGE}..."
DIGEST=$(cosign triangulate "${IMAGE}")

echo "Digest: ${DIGEST}"
echo "Verifying signature..."

cosign verify \
  --certificate-oidc-issuer "https://token.actions.githubusercontent.com" \
  --certificate-identity-regexp "https://github.com/hesusruiz/tmforum/.github/workflows/.*" \
  "${DIGEST}"

echo "✔ Verification successful"
echo "The image is authentic, signed by the tmforum GitHub workflow, and recorded in the Rekor transparency log."
```

Run it like:

```sh
./verify.sh v1.0.2
```

### Security & Provenance

This project uses **Sigstore** to provide strong supply‑chain security for all published container images and module artifacts.

#### Keyless Signing (No Keys to Manage)

All releases are signed using **Sigstore Cosign** in keyless mode.  
Instead of long‑lived private keys, signatures are tied to:

- a GitHub Actions workflow identity  
- a short‑lived Fulcio certificate  
- an immutable Rekor transparency log entry  

This eliminates key management risks and ensures signatures cannot be forged.

#### Immutable Provenance

Each signed artifact includes:

- the Git commit  
- the GitHub workflow that built it  
- the tag that triggered the release  
- the container image digest  
- a timestamped Rekor entry  

This creates a tamper‑proof chain of custody from source -> build -> artifact.

#### Protection Against Supply‑Chain Attacks

By verifying signatures, users ensure:

- the image was built by the **official tmforum GitHub workflow** in this repository.
- the image has **not been tampered with**.  
- the image corresponds to the **source code** at the tagged version, enabling easy auditing of the real contents of the image.
- the signature is **publicly logged** and cannot be removed or altered.  

This protects against:

- compromised CI systems  
- malicious image replacement  
- tag hijacking  
- dependency poisoning  
- unauthorized rebuilds  

#### Public Transparency (Rekor)

All signatures are stored in the **Rekor transparency log**, a public, append‑only ledger.  
Anyone can independently audit:

- who signed the artifact  
- when it was signed  
- what digest was signed  
- the certificate used  

This provides verifiable trust without relying on private infrastructure.

#### Easy Verification

Users can verify any version using Cosign:

```sh
cosign triangulate ghcr.io/hesusruiz/tmforum:<version>
cosign verify \
  --certificate-oidc-issuer "https://token.actions.githubusercontent.com" \
  --certificate-identity-regexp "https://github.com/hesusruiz/tmforum/.github/workflows/.*" \
  ghcr.io/hesusruiz/tmforum@sha256:<digest>
```

Or simply run the included `verify.sh` script.

## Configuring and running <a id="configuring-and-running"></a>

### Docker Deployment

The application is designed to be containerized.

1.  **Build the container:**
    ```bash
    docker build -t tmforum .
    ```

2.  **Run the container:**
    The container requires the `TMF_ADMIN_TOKEN` environment variable to specify the admin token, wich callers will have to use to perform admin operations.

    ```bash
    docker run -e TMF_ADMIN_TOKEN=1234567890 -p 9991:9991 tmforum
    ```

    The container also acceps other environment variables which configure the behavior of the server. These are:

    * `TMF_RUN_ENVIRONMENT`: The environment to use. Possible values: `isbedev`, `isbepre`, `isbepro`, `domedev`, `domepre`, `domepro`, `local`.
    * `TMF_DEBUG`: Enable debug mode. Possible values: `true`, `false`.
    * `TMF_PROXY_ENABLED`: Enable proxy mode. Possible values: `true`, `false`.
    * `TMF_REMOTE_SERVER`: The URL of the remote TMForum API server when we act as proxy.
    * `TMF_VERIFIER`: The URL of the verifier server, which is used to verify the access tokens.

3.  **The database and backups:**
    The application uses SQLite as the database. The database is stored in the `/data` directory of the container. To persist the database in the host filesystem, you can use a volume mapped to the `/data` directory of the container.

    The system maintains backups of the database in the `/data/backups` directory of the container. It maintains a rotating set of 7 backups, one for each day of the week. The backups contain the number of the day of the week in the name (e.g., `backup_01.db`, `backup_02.db`, etc.). The backups are performed every 2 hours at even hours (e.g., 11:00, 13:00, etc.) in the backup file corresponding to the day of the week. Storing the backups in a different system (e.g., in a cloud storage service) is as simple as copying the `/data/backups` directory (or just the latest backup file) to the desired location. To ensure that there are no conflicts with the scheduled updates of the backups, it is recommended to make the copy on even hours (e.g., 10:00, 12:00, etc.).

    ```bash
    docker run -e TMF_ADMIN_TOKEN=1234567890 -p 9991:9991 -v /path/to/your/database:/data tmforum
    ```
    This will persist the database in the `/path/to/your/database` directory on the host and create backups in the same directory as the database. It will create a subdirectory `backups` to store the backups.

### Profiles

Configuration is managed via profiles defined in `config/config_data.go`. This approach reduces configuration errors by grouping settings into well-defined environments.

To add or modify a profile, edit `config/config_data.go`. You can specify the profile to use at runtime with the `-run` flag:
```bash
./bin/tmforum -run <profile_name>
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

### Prerequisites

* [Go](https://go.dev/dl/) 1.26 or higher
* [Docker](https://www.docker.com/) (for containerized deployment)

### Local Development

To run the server locally for development:

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/hesusruiz/tmforum.git
    cd tmforum
    ```

2.  **Run with default configuration:**
    ```bash
    go run main.go -run local
    ```
    This starts the server using the `local` profile (configured in `config/config_data.go`).

3.  **Build the binary:**
    ```bash
    go build -o bin/tmforum main.go
    ```

4.  **Run Tests:**
    ```bash
    make test
    ```

## License

This project is licensed under the Apache 2.0 License - see the [LICENSE](LICENSE) file for details.
