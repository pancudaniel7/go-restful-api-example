# Go Restful API Example

The project is an **API** built with Go, designed to provide an example of implementing a **RESTful API**
structure. This document provides instructions on how to set up the project, manage dependencies, run the application,
and use the included Postman collection for API testing.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing
purposes.

### Prerequisites

Before going further make sure that you have this tools installed on your machine.

 - Go
 - Docker
 - Docker-compose
 - Dlv

### Installing

First, clone the repository to your local machine:

```bash
git clone https://github.com/pancudaniel7/go-restful-api-example.git
```

Navigate to the project directory:

```bash
cd go-restful-api-example/
```

This project uses go mod for managing dependencies. To install all the necessary dependencies, run the following command
in the project **root** directory:

```bash
go mod tidy
```

This will download and install the required dependencies.

### Run the Database
To run the **database** you can run the following command in the project root directory:

```bash
docker-compose -f deployments/docker-compose.yml up -d database
```

### Running the Application

To start the server, run the following command in the project root directory:

```bash
go run cmd/main.go
```

This will start the application on <default port. :8080>. You can access the API
at `http://localhost:8080`.

### Running the Application using remote debug port

In order to remote debug the application you can run this and attach a debug interface to 2345 port.
```bash
dlv debug cmd/go-restful-api-example/main.go --headless --listen=:2345 --api-version=2 --log
```

## Using the Postman Collection

A Postman collection and environment variables are included in the `docs/postman` folder for testing the API endpoints.

### Importing the Collection and Environment Variables

1. Open **Postman** and click on the `Import` button.
2. Choose `File` and then `Upload Files`.
3. Navigate to your project's `docs/postman` directory.
4. Select both the collection and environment variables JSON files and click `Open` to import them into Postman.

### Setting Up Environment Variables

After importing, ensure the environment is selected in Postman:

1. In the top right corner, find the environment dropdown and select the environment you imported.
2. Review and adjust the variables as necessary for your local setup (e.g., the base URL of the API).

### Using the Collection

With the collection and environment set up, you can now use the requests defined in the Postman collection to test the
API:

1. Expand the imported collection on the left sidebar.
2. Click on an API request to view its details.
3. Hit the `Send` button to execute the request and see the response.

### RESTful API and Hypermedia

The highest maturity level in REST, known as **Level3**: **Hypermedia** Controls" (based on the Richardson Maturity Model),
emphasizes the use of Hypermedia as the Engine of Application State (HATEOAS). 
At this stage, server responses include
hypermedia links that guide clients towards other relevant actions and resources. This design principle abstracts the API into a
self-descriptive, navigable format. 
It enables clients to dynamically discover and interact with the API's capabilities without prior
knowledge of its structure, making applications more adaptive, scalable, and easier to evolve. This level epitomizes the
full realization of REST by enhancing client-server interaction with comprehensive, context-aware hypermedia.

---

This project was created to demonstrate the principles of building a RESTful API with Go, showcasing effective practices
in structuring and testing API endpoints.
