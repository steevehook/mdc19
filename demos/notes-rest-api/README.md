# Dummy Notes REST API

## Overview

This demo provides you a dummy **Notes** example of **REST API** in Go
going all they way up from **developing** the API
and **covering** it with **tests**

This API will also make use of some third party libraries
to give you an insight on how you can make use of them
and tweak them for your needs

We'll be mostly using these

- `HTTP router`
- `Assertion library`
- `Database driver`

The rest will be **pure Go code** without any magic
or things which you won't be able to **understand**


## Steps to run

```bash
# Download all go libraries used
go get ./...

# Compile & run your program
go run *.go

# run unit tests
./coverage.sh

# run unit & integration tests
./coverage.sh --intgr

# run unit tests and show HTML coverage report
./coverage.sh --html
```
 
## More info

- [HTTP Router](https://github.com/julienschmidt/httprouter)
- [Go MySQL Driver](https://github.com/go-sql-driver/mysql)
- [Testify](https://github.com/stretchr/testify)

## Back to MDC19

[Home](https://github.com/steevehook/mdc19)
