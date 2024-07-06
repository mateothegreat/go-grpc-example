# Golang gRPC example

This project is divided into two parts:

* [protos](protos)
* [server](server)
* [client](client)

These are separated into two different packages to make it easier to publish modules in the future for re-usability.

## Generate proto files

```bash
cd protos
make compile
```

## Testing

Fist, start the server:

```bash
cd server
go test -v
```

Second, start the client:

```bash
cd client
go test -v
```