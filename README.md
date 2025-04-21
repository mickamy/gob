# 🐘 gob

> A lightweight database management CLI/Library.
>

`gob` helps you manage your database with simple commands like `create`, `migrate`, `drop`, and more — inspired by tools
like Rails’ `db:*` commands.

---

## ✨ Features

- ✅ Interactive `gob init` to scaffold `.gob.yaml`
- ✅ Easily create and drop databases
- ✅ Run migrations using `migrate` or `generate migrate`
- ✅ Supports MySQL and PostgreSQL
- ✅ CLI and YAML-based — no magic involved

---

## 📦 Installation

```bash
# Install gob into your project
go get -tool github.com/mickamy/gob/cmd/gob@latest

# or install it globally
go install github.com/mickamy/gob/cmd/gob@latest
```

---

## ⚙️ Set-up

Initialize a project with:

```bash
gob init
```

This will walk you through DB config like:

```yaml
# .gob.yaml
database:
  driver: postgres
  host: localhost
  port: 5432
  user: gob
  password: password
  name: gob

migrations:
  dir: migrations
  ext: sql
  seq: false
```

---

## 🚀 Usage

### Create a database

```bash
gob create
```

### Drop a database

```bash
gob drop
```

### Run migrations

```bash
gob migrate
```

This uses the `migrate` binary under the hood (with fallback to go tool migrate if available)

---

## 🛠 Generate migration file

```bash
gob g migration create_users
```

This forwards to `migrate create` to create a new migration file.

```
migrations/000001_create_users.up.sql
migrations/000001_create_users.down.sql
```

---

## 📚 Using as a Library

You can also use `gob` as a Go package to manage your database programmatically:

```bash
go get github.com/mickamy/gob@latest
```

### 1. Load the config

```go
import (
  "log"

  "github.com/mickamy/gob/config"
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

  "github.com/mickamy/gob"
)

err := gob.Create(cfg)
if errors.Is(err, gob.ErrCreateDatabaseExists) {
  log.Println("database already exists, skipping.")
} else if err != nil {
  log.Fatal("failed to create database:", err)
}
```

### 3. Drop a database

```go
import (
  "log"

  "github.com/mickamy/gob"
)

if err := gob.Drop(cfg); err != nil {
  log.Fatal("failed to drop database:", err)
}

```

### 4. Run migrations

```go
import (
  "errors"
  "log"

  "github.com/mickamy/gob"
)

err := gob.Migrate(cfg)
if errors.Is(err, gob.ErrMigrateNoChange) {
  log.Println"no new migrations to apply.")
} else if err != nil {
  log.Fatal("migration failed:", err)
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
