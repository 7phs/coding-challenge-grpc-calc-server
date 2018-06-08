# Arithmetic calculation service

A service is computing a several arithmetic operations such as add, subtract, divide, multiply
and fibonacci number.

A service is implementing gRPC API.

The current package contains as a service and a client to call it. 

# Requirements

## Environment

A service and a client was testing on the following environment:

* Go 1.10+
* protobuf compiler 3.5.1 
* **godep** v80 devel +c1d9d1f305 or **dep** devel +1174ad3a8f 

## Project path

An expected path to build a service and a client is ```$GOPATH/calcService```

## Preparing

To install all dependencies run one of the command:

```
godep restore
```

or

```
dep ensure
```

## gRPC

Run a following command if changing a gRPC service definition in a file ```api/calc.proto```:

```
go generate ./...
```

# Server

A service supports configuration of a listening port using an environment variable ```PORT```

Default port is 9090.

Another supported parameter is a log level. Using an  environment variable ```LOG_LEVEL```
and values ```debug```, ```info``` (default), ```warning``` and ```error```.

## Build

```
go build -o=bin/server ./server/
```

## Running

Default parameters:
```
server
```

Changinh parameters:
```
PORT=12321 LOG_LEVEL=debug server
```

# Client

A client supports configuration of a local service port using an environment variable ```PORT```

Default service port is 9090.

Another parameter is a timeout of a request in milliseconds.
Changing an environment variable ```TIMEOUT```

## Build
```
go build -o=bin/client ./client/
```

## Running

```
client CMD ARG [ARG ARG ... ARG]
```

* CMD - a name of arithmetical operation: ```add```, ```sub```, ```mul```, ```div```, ```fib```.
* ARG - an integer numeric applied for an arithmetical operation

Default parameters:

```
client add 4 5
``` 

Specific parameters:

```
PORT=12321 TIMEOUT=300 client fib 56
``` 


# Trade-off

1. Supports only integer arguments.
2. Clients supports only local service and configuration only port number.