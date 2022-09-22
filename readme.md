# Altera Golang Mini Course (AGMC)

This project is a RESTful API with an agnostic approach, so developers only need to think about the business process.

## Features

- Support [RESTful API](https://en.wikipedia.org/wiki/Representational_state_transfer). e.g. User & Book.
- Support Object Relational Mapping ([ORM](https://en.wikipedia.org/wiki/Object%E2%80%93relational_mapping)) concept.
- Implement clean architecture ([Hexagonal Architecture](https://en.wikipedia.org/wiki/Hexagonal_architecture_(software))) on main project.
- Implement common design pattern.

## Tech

This project uses a number of open source project to work properly:

- [Go](https://go.dev/) - Programming language
- [Docker](https://www.docker.com/) - Containerization
- [MySQL](https://www.mysql.com/) - Relational database
- [Heroku](https://www.heroku.com) - Deployment image to server

## Installation

This project requires [go](https://go.dev/) version go1.19 to run.

Initiate new table and start the server.

```sh
cd alterra-agmc
go run main.go -migrate=migrate
```