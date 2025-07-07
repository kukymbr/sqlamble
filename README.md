<img align="right" width="125" src="sqlamble.png" alt="image with a gopher holding a fried egg">

# sqlamble

> For those who hate SQL inside the Go code.

[![License](https://img.shields.io/github/license/kukymbr/sqlamble.svg)](https://github.com/kukymbr/sqlamble/blob/main/LICENSE)
[![Release](https://img.shields.io/github/release/kukymbr/sqlamble.svg)](https://github.com/kukymbr/sqlamble/releases/latest)
[![GoDoc](https://godoc.org/github.com/kukymbr/sqlamble?status.svg)](https://godoc.org/github.com/kukymbr/sqlamble)
[![GoReport](https://goreportcard.com/badge/github.com/kukymbr/sqlamble)](https://goreportcard.com/report/github.com/kukymbr/sqlamble)
[![Pipeline](https://gitlab.com/kukymbr/sqlamble/badges/main/pipeline.svg)](https://gitlab.com/kukymbr/sqlamble/-/pipelines/)
[![Coverage](https://gitlab.com/kukymbr/sqlamble/badges/main/coverage.svg)](https://gitlab.com/kukymbr/sqlamble/)

The `sqlamble` tool allows you to embed your SQL queries into the Golang code in structural read-only way.

## Why?

When you are looking for a `golang sql embed` in Google, the most popular answer is that kind of examples:

```go
package mypkg

import _ "embed"

//go:embed sql/my_query.sql
var myQuery string
```

And there is nothing wrong with this way of embedding SQL queries, until you have tons of them:

```text
sql
├── users
│   ├── get_list.sql
│   ├── get_user_data.sql
│   └── ... other 100500 queries
├── orders
│   └── ... other 100500 queries
└── ... other 100500 directories
```

The `sqlamble` takes all these files on the `go generate` stage and converts them into the structured set 
of the read-only strings. For example:

```go
query := queries.Users().GetListQuery()
```

## Installation

The go 1.24 is a minimal requirement for the `sqlamble`, so the `go tool` is a preferred way to install:

```shell
go get -tool github.com/kukymbr/sqlamble/cmd/sqlamble@latest
```

## Usage

The `sqlamble --help` output:

```text
Generates structured SQL getters in go code.
See https://github.com/kukymbr/sqlamble for info.

Usage:
  sqlamble [flags]

Flags:
      --ext strings           If set, source files will be filtered by these suffixes in names (default [.sql])
      --fmt string            Formatter used to format generated go files (gofmt|noop) (default "gofmt")
  -h, --help                  help for sqlamble
      --package string        Target package name of the generated code (default "queries")
      --query-suffix string   Suffix for query getter functions (default "Query")
  -s, --silent                Silent mode
      --source string         Directory containing SQL files (default ".")
      --target string         Directory for the generated Go files (default "internal/queries")
  -v, --version               version for sqlamble
```

1. Create sql files directory and put some SQL inside it (any level of subdirectories is supported), for example `sql/`.
2. Add the go file with a `//go:generate` directive, for example `sql/generate.go`:
   ```go
    package sql  

   //go:generate go tool sqlamble --package=queries --target=../internal/queries
   ```
3. Run the `go generate` command:
   ```shell
   go generate ./sql
   ```
4. Use the types, generated into the `queries` package (see the generated `internal/queries/` directory):
   ```go
   package users
   
   func GetUsers() []User {
       query := queries.Users().GetListQuery()   
	   // ... go fetch some users
   }
   
   func GetUser() User {
    query := queries.Users().SingleUser().GetUserDataQuery()
	   // ... go fetch some user data
   }
   ```

See the [example](example) directory for a full example.

### Hidden features

In fact, it's okay to embed any type of string content into the Go code
using the sqlamble, because there is no parsing of the SQL syntax itself.

For example, you could embed set of YAMLs:

```shell
go tool sqlamble --package=configs --target=internal/configs --query-suffix=YAML --ext=.yaml,.yml
```

See the generator's [testdata](internal/generator/testdata/source/yaml) 
and [test code](internal/generator/generator_test.go) for an example.

## Contributing

Please, refer the [CONTRIBUTING.md](CONTRIBUTING.md) doc.

## License

[MIT](LICENSE).