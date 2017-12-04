# Authenticator

JWT based authentication service with a JSON over HTTP interface.
The main purpose of this project is for me to get familiar with golang. More specifically, to get to know how to organize a simple HTTP server with it.
Feedback is welcome.

# TODO

* CRUD credentials
* Tests
* Configuration management
* Define and implement deployment strategy
* Database schema migrations
* Refactor, the code became a mess :(
* Log

# Build

```
make build
```

# Database

In order to setup the database you need to execute the statements in `database_schema.sql`

# Usage

```bash
go run cmd/authenticator/main.go # OR ./bin/authenticator

curl -X POST -u user:123456 localhost:8080/authenticate
```

## References

* https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1
* https://medium.com/wtf-dial
