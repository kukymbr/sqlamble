# 🍳 sqlamble

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
query := queries.Queries().Users().GetList()
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
      --ext strings      If set, source files will be filtered by these suffixes in names (default [.sql])
  -h, --help             help for sqlamble
      --package string   Target package name of the generated code (default "queries")
      --source string    Directory containing SQL files (default ".")
      --target string    Directory for the generated Go files (default "internal/queries")
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
4. Use the types, generated into the `queries` package inside the `internal/queries` (names are from the example):
   ```go
   package users
   
   func GetUsers() []User {
       query := queries.Queries().Users().GetListQuery()   
	   // ... go fetch some users
   }
   
   func GetUser() User {
    query := queries.Queries().Users().SingleUser().GetUserDataQuery()
	   // ... go fetch some user data
   }
   ```

See the [example](example) directory for a full example.

## Contributing

Please, refer the [CONTRIBUTING.md](CONTRIBUTING.md) doc.

## License

[MIT](LICENSE).