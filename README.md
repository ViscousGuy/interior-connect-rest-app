# interior-connect-rest-app

A simple RESTful API built using Go with Beego framework

# Prerequisites

- Ensure you have Go installed on your system. You can verify by running `go version` in your terminal.
- **Dependencies:** This project uses the following external libraries:
  - github.com/beego/beego/v2 (v2.2.1) - A web framework for Go

**Installing Dependencies**

Once you have Go set up, use the following command in your project's root directory to fetch the required dependencies:

```bash
go get -u
```

**Copy the example environment file and configure:**

```bash
cp .env.example .env
```

# Run the Project

```bash
go run main.go
```

# Routes

The application defines the following routes:

**Basic Routes**

- **GET /:** Renders the home page.

**Furniture Routes (grouped)**

- **GET /furniture:** Displays a list of all furnitures.
