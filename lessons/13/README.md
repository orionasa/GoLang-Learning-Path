# Module 13: Database Mastery with sqlc

In this module, we'll introduce `sqlc`, a tool that generates type-safe Go code from your SQL queries.

## Why `sqlc`? (Type-safe, compile-time checked SQL)

While the `database/sql` package provides a generic SQL interface, it can be cumbersome to work with. You have to manually scan rows into structs, and there's no compile-time checking of your SQL queries.

`sqlc` solves these problems by generating Go code directly from your SQL schema and queries. This gives you:

*   **Type safety**: No more `interface{}`. Your database queries become type-safe Go functions.
*   **Compile-time checking**: `sqlc` validates your SQL queries against your database schema.
*   **Readability**: Your SQL queries live in `.sql` files, separate from your Go code.

## Setting up `sqlc` with your database

First, you'll need to install `sqlc`:

```bash
go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
```

Next, create a `sqlc.yaml` configuration file in your project's root directory:

```yaml
version: "2"
sql:
  - engine: "postgresql"
    queries: "db/query/"
    schema: "db/migration/"
    gen:
      go:
        package: "db"
        out: "db/sqlc"
```

This configuration tells `sqlc` where to find your database schema, queries, and where to output the generated Go code.

## Writing SQL queries to generate type-safe Go code

Now, let's create our schema and a query.

`db/migration/001_init_schema.sql`:
```sql
CREATE TABLE authors (
  id   BIGSERIAL PRIMARY KEY,
  name TEXT      NOT NULL,
  bio  TEXT
);
```

`db/query/authors.sql`:
```sql
-- name: GetAuthor :one
SELECT * FROM authors
WHERE id = $1 LIMIT 1;

-- name: ListAuthors :many
SELECT * FROM authors
ORDER BY name;

-- name: CreateAuthor :one
INSERT INTO authors (
  name, bio
) VALUES (
  $1, $2
)
RETURNING *;
```

Now, run `sqlc generate`:

```bash
sqlc generate
```

This will generate a `db/sqlc/authors.sql.go` file with the following functions:

*   `GetAuthor(ctx context.Context, id int64) (Author, error)`
*   `ListAuthors(ctx context.Context) ([]Author, error)`
*   `CreateAuthor(ctx context.Context, arg CreateAuthorParams) (Author, error)`

## Integrating `sqlc` for CRUD operations and transactions

Now you can use the generated code in your application.

```go
package main

import (
	"context"
	"database/sql"
	"fmt"

	"example.com/myproject/db/sqlc"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	ctx := context.Background()
	conn, err := sql.Open("pgx", "user=user password=password dbname=mydb sslmode=disable")
	if err != nil {
		// handle error
	}

	queries := db.New(conn)

	// Create an author
	createdAuthor, err := queries.CreateAuthor(ctx, db.CreateAuthorParams{
		Name: "Brian Kernighan",
		Bio:  sql.NullString{String: "Co-author of The C Programming Language and The Go Programming Language", Valid: true},
	})
	if err != nil {
		// handle error
	}
	fmt.Println(createdAuthor)

	// List all authors
	authors, err := queries.ListAuthors(ctx)
	if err != nil {
		// handle error
	}
	fmt.Println(authors)
}
```

`sqlc` is a powerful tool that can greatly improve your productivity when working with databases in Go.

**Previous:** [Module 12: Middleware and Databases](../12/README.md)
**Next:** [Module 14: Authentication and JWT](../14/README.md)
