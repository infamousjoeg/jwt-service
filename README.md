<div align="center">
<h1 align="center">
<img src="https://raw.githubusercontent.com/PKief/vscode-material-icon-theme/ec559a9f6bfd399b82bb44393651661b08aaf7ba/icons/folder-markdown-open.svg" width="100" />
<br>jwt-service
</h1>
<h3>◦ Unlocking seamless authentication with jwt-service.</h3>
<h3>◦ Developed with the software and tools listed below.</h3>

<p align="center">
<img src="https://img.shields.io/badge/Docker-2496ED.svg?style&logo=Docker&logoColor=white" alt="Docker" />
<img src="https://img.shields.io/badge/Go-00ADD8.svg?style&logo=Go&logoColor=white" alt="Go" />
<img src="https://img.shields.io/badge/Markdown-000000.svg?style&logo=Markdown&logoColor=white" alt="Markdown" />
</p>
<img src="https://img.shields.io/github/languages/top/infamousjoeg/jwt-service?style&color=5D6D7E" alt="GitHub top language" />
<img src="https://img.shields.io/github/languages/code-size/infamousjoeg/jwt-service?style&color=5D6D7E" alt="GitHub code size in bytes" />
<img src="https://img.shields.io/github/commit-activity/m/infamousjoeg/jwt-service?style&color=5D6D7E" alt="GitHub commit activity" />
<img src="https://img.shields.io/github/license/infamousjoeg/jwt-service?style&color=5D6D7E" alt="GitHub license" />
</div>

---

## 📒 Table of Contents
- [📒 Table of Contents](#-table-of-contents)
- [📍 Overview](#-overview)
- [⚙️ Features](#️-features)
- [📂 Project Structure](#-project-structure)
- [🧩 Modules](#-modules)
- [🚀 Getting Started](#-getting-started)
  - [✔️ Prerequisites](#️-prerequisites)
  - [📦 Installation](#-installation)
  - [🎮 Using jwt-service](#-using-jwt-service)
    - [GET /generate-jwt](#get-generate-jwt)
    - [GET /.well-known/jwks.json](#get-well-knownjwksjson)
  - [🧪 Running Tests](#-running-tests)
- [🤝 Contributing](#-contributing)
- [📄 License](#-license)

---


## 📍 Overview

The "jwt-service" project is a Golang application that serves an HTTP endpoint for generating JSON Web Tokens (JWTs). It leverages RSA key pairs for token signing and periodically rotates the keys based on a specified time-to-live (TTL). Additionally, it provides a JSON Web Key Set (JWKS) endpoint, allowing clients to efficiently access and verify public keys for token verification. This project offers a secure and scalable solution for managing and generating JWTs, making it highly valuable for applications requiring authentication and access control.

---

## ⚙️ Features

| Feature                | Description                           |
| ---------------------- | ------------------------------------- |
| **⚙️ Architecture**     | The codebase follows a simple modular design using a traditional web service architecture with a Go server serving HTTP endpoints for token generation and a JWKS endpoint for public key information.    |
| **📖 Documentation**   | The project lacks comprehensive documentation. The README provides a brief overview and setup instructions, but more documentation is needed to understand the implementation details and usage.    |
| **🔗 Dependencies**    | The code relies on the jwt-go package (v3.2.0+) from the dgrijalva/jwt-go repository for JWT functionalities. There don't appear to be any direct dependencies on other external systems.    |
| **🧩 Modularity**      | The codebase is well-organized into separate files and packages, with clear responsibilities for generating JWTs and handling the JWKS endpoint. The modular structure allows for easy extensibility and maintainability.    |
| **✔️ Testing**          | The project includes test functions that verify the correctness of the JWKS endpoint and JWT generation. The tests cover basic scenarios but could be expanded to include more cases and handle error scenarios. No automated test runner/tool is mentioned.    |
| **⚡️ Performance**      | Performance analysis requires further examination, but using Go as the implementation language should result in good performance. However, actual performance characteristics can be influenced by factors such as the JWT payload size and the underlying infrastructure.    |
| **🔐 Security**        | The codebase utilizes RSA key pairs for token signing, providing a secure approach for JWT generation. The periodic key rotation based on the provided TTL ensures a higher level of security. The code follows standard security practices for JWT usage.    |
| **🔀 Version Control** | The project utilizes Git for version control. The codebase is hosted on GitHub, allowing for collaboration, version management, and issue tracking. Further details on specific version control strategies and tools are not mentioned.    |
| **🔌 Integrations**    | As a standalone JWT service, the codebase does not have explicit integrations with other systems. However, it can be easily integrated into larger systems where JWT authentication/authorization is required.    |
| **📶 Scalability**     | The system's scalability depends on the underlying infrastructure. Since the codebase is written in Go, it can handle a considerable amount of traffic. It supports horizontal scalability through load balancing.    |

---


## 📂 Project Structure




---

## 🧩 Modules

<details closed><summary>Root</summary>

| File                                                                               | Summary                                                                                                                                                                                                                                                                                                                                                                        |
| ---                                                                                | ---                                                                                                                                                                                                                                                                                                                                                                            |
| [main.go](https://github.com/infamousjoeg/jwt-service/blob/main/main.go)           | This code is a Golang application that serves an HTTP endpoint to generate JSON Web Tokens (JWTs). It uses RSA key pairs for token signing and periodically rotates the keys based on the provided time-to-live (TTL). The code also exposes a JSON Web Key Set (JWKS) endpoint that provides information about the available public keys for token verification.              |
| [Dockerfile](https://github.com/infamousjoeg/jwt-service/blob/main/Dockerfile)     | This code builds a Docker image for a Go application called "jwt-service". It first creates a build artifact using the official Golang image and then copies the necessary files and dependencies. The final Docker image is built from scratch and includes the pre-built binary of the application. It sets a non-root user and specifies the command to run the executable. |
| [go.mod](https://github.com/infamousjoeg/jwt-service/blob/main/go.mod)             | The code is a module called jwt-service, which primarily focuses on implementing JWT (JSON Web Token) functionalities. It requires the jwt-go package version v3.2.0+incompatible from the dgrijalva/jwt-go GitHub repository.                                                                                                                                                 |
| [main_test.go](https://github.com/infamousjoeg/jwt-service/blob/main/main_test.go) | This code contains test functions for handling JWKS endpoint and generating JWT. It verifies if the handlers return the expected status codes and checks the validity of the returned key and generated JWT. It can be expanded to include more tests and handle error scenarios.                                                                                              |

</details>

---

## 🚀 Getting Started

### ✔️ Prerequisites

Before you begin, ensure that you have the following prerequisites installed:
> - `ℹ️ Docker`

### 📦 Installation

1. Clone the jwt-service repository:
```sh
git clone https://github.com/infamousjoeg/jwt-service
```

2. Change to the project directory:
```sh
cd jwt-service
```

3. Build the containers:
```sh
docker compose build
```

4. Update the environment variables in the [docker-compose.yml]() file.

```yaml
      - VIRTUAL_HOST=yourdomain.com
      - VIRTUAL_PORT=8080
      - LETSENCRYPT_HOST=yourdomain.com
      - LETSENCRYPT_EMAIL=youremail@example.com
      - JWT_ISSUER=yourdomain.com
      - JWT_SUBJECT=host/workload/id
      - JWT_AUDIENCE=example.secretsmgr.cyberark.cloud
      - JWT_TTL=5
      - JWKS_KEY_TTL=6
```

|Variable|Description|
|---|---|
|`VIRTUAL_HOST`|The host name for the service for the NGINX reverse proxy.|
|`VIRTUAL_PORT`|The port number for the service for the NGINX reverse proxy.|
|`LETSENCRYPT_HOST`|The domain name for the service for the LetsEncrypt SSL Certificate.|
|`LETSENCRYPT_EMAIL`|The email address for the service for the LetsEncrypt SSL Certificate.|
|`JWT_ISSUER`|The issuer for the JWT.|
|`JWT_SUBJECT`|The subject for the JWT.|
|`JWT_AUDIENCE`|The audience for the JWT.|
|`JWT_TTL`|The time-to-live for the JSON Web Token (JWT) in minutes (Default: 60)|
|`JWKS_KEY_TTL`|The time-to-live for the JSON Web Key Set (JWKS) signing key in minutes. (Default: 60)|

5. Run the containers:
```sh
docker compose up -d
```

### 🎮 Using jwt-service

#### GET /generate-jwt

Responds with a JSON Web Token (JWT) signed with the current signing key.

#### GET /.well-known/jwks.json

Responds with a JSON Web Key Set (JWKS) containing the current signing key's public certificate.

### 🧪 Running Tests
```sh
go test
```

---

## 🤝 Contributing

Contributions are always welcome! Please follow these steps:
1. Fork the project repository. This creates a copy of the project on your account that you can modify without affecting the original project.
2. Clone the forked repository to your local machine using a Git client like Git or GitHub Desktop.
3. Create a new branch with a descriptive name (e.g., `new-feature-branch` or `bugfix-issue-123`).
```sh
git checkout -b new-feature-branch
```
4. Make changes to the project's codebase.
5. Commit your changes to your local branch with a clear commit message that explains the changes you've made.
```sh
git commit -m 'Implemented new feature.'
```
6. Push your changes to your forked repository on GitHub using the following command
```sh
git push origin new-feature-branch
```
7. Create a new pull request to the original project repository. In the pull request, describe the changes you've made and why they're necessary.
The project maintainers will review your changes and provide feedback or merge them into the main branch.

---

## 📄 License

This project is licensed under the `ℹ️  MIT` License. See the [LICENSE](LICENSE) file for additional info.