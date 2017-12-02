# Authenticator

JWT based authentication service with a JSON over HTTP interface.
The main purpose of this project is for me to get familiar with golang. More specifically, to get to know how to organize a simple HTTP server with it.
Feedback is welcome.

# TODO

* HTTP interface
* CRUD credentials
* Tests
* Configuration management
* Deployment strategy
* Refactor, the code became a mess :(

# Usage

```
curl -X POST  -u user:123456 localhost:8080/authenticate
```

## References

* https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1
* https://medium.com/wtf-dial
