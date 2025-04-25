# 🐘 godb

> A lightweight database management CLI/Library.
>

`godb` helps you manage your database with simple commands like `create`, `migrate`, `drop`, and more — inspired by tools
like Rails’ `db:*` commands.

---

## ✨ Features

- ✅ Interactive `godb init` to scaffold `.godb.yaml`
- ✅ Easily create and drop databases
- ✅ Run migrations using `migrate` or `generate migrate`
- ✅ Supports MySQL and PostgreSQL
- ✅ CLI and YAML-based — no magic involved

---

## 📦 Installation

```bash
# Install godb into your project
go get -tool github.com/mickamy/godb/cmd/godb@latest

# or install it globally
go install github.com/mickamy/godb/cmd/godb@latest
```

---

## ⚙️ Set-up

Initialize a project with:

```bash
godb init
```

This will walk you through DB config like:

```yaml
# .godb.yaml
database:
  driver: postgres
  host: localhost
  port: 5432
  user: godb
  password: password
  name: godb

migrations:
  dir: migrations
  ext: sql
  seq: false
```

---

## 🚀 Usage

### Create a database

```bash
godb create
```

### Drop a database

```bash
godb drop
```

### Run migrations

```bash
godb migrate
```

This uses the `migrate` binary under the hood (with fallback to go tool migrate if available).


### Rollback migrations

```bash
godb rollback --step=1
````

This also uses the `migrate` binary under the hood.

---

## 🛠 Generate migration file

```bash
godb g migration create_users
```

This forwards to `migrate create` to create a new migration file.

```
migrations/000001_create_users.up.sql
migrations/000001_create_users.down.sql
```

---

## 📚 Using as a Library

You can also use `godb` as a Go package to manage your database programmatically:

```bash
go get github.com/mickamy/godb@latest
```

### 1. Load the config

```go
import (
  "log"

  "github.com/mickamy/godb/config"
)

cfg, err := config.Load()
if err != nil {
  log.Fatal("failed to load config:", err)
}
```

### 2. Create a database

```go
import (
  "errors"
  "log"

  "github.com/mickamy/godb"
)

err := godb.Create(cfg)
if errors.Is(err, godb.ErrCreateDatabaseExists) {
  log.Println("database already exists, skipping.")
} else if err != nil {
  log.Fatal("failed to create database:", err)
}
```

### 3. Drop a database

```go
import (
  "log"

  "github.com/mickamy/godb"
)

if err := godb.Drop(cfg, false); err != nil {
  log.Fatal("failed to drop database:", err)
}
```

The second argument is a flag to terminate all the connections to the database before dropping it.

### 4. Run migrations

```go
import (
  "errors"
  "log"

  "github.com/mickamy/godb"
)

err := godb.Migrate(cfg)
if errors.Is(err, godb.ErrMigrateNoChange) {
  log.Println"no new migrations to apply.")
} else if err != nil {
  log.Fatal("migration failed:", err)
}

err = godb.Rollback(cfg, 1)
if errors.Is(err, godb.ErrMigrateNoChange) {
  log.Println("no migrations to rollback.")
} else if err != nil {
  log.Fatal("rollback failed:", err)
}
```

> You can use this to integrate database set-up into your own tooling, tests, or set-up scripts.
>

---

## 🧪 Supported drivers

- ✅ MySQL
- ✅ PostgreSQL

---

## 📄 License

[MIT](./LICENSE)
