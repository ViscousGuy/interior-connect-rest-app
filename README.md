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
# Documentation 

You can find the documentation [on this website](https://documenter.getpostman.com/view/34685195/2sA3JKehv9).
# Routes

The application defines the following routes:

**Basic Routes**

- **GET /:** Renders the home page.

**Routes (grouped)**

- **GET /furnitures:** Displays a list of all furnitures.
- **GET /furnitures/:slug** Displays a furniture.
- **GET /materials:** Displays a list of all materials for a furniture.
- **GET /contractors:** Displays a list of all contractors.
- **GET /contractors/:slug:** Displays a contractor.
- **GET /projects:** Displays a list of all projects.
- **GET /projects/:slug:** Displays a project.
