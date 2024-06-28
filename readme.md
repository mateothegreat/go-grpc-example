# Golang gRPC example

This project is divided into two parts:

* [pkg/server](pkg/server)
* [pkg/client](pkg/client)

These are separated into two different packages to make it easier to publish modules in the future for re-usability.

## Testing

Fist, start the server:

```bash
cd pkg/server
go test -v
```

Second, start the client:

```bash
cd pkg/client
go test -v
```