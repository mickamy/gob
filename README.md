# 🐘 gob

> A lightweight database management CLI for Go projects.
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
go get -tool github.com/mickamy/gob@latest
# Install gob globally
go install github.com/mickamy/gob@latest
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
gob g migrate -ext sql -dir migrations create_users
```

This forwards to `golang-migrate create` to create a new migration file.

```
migrations/000001_create_users.up.sql
migrations/000001_create_users.down.sql
```

---
## 🧪 Supported drivers

- ✅ MySQL
- ✅ PostgreSQL

---

## 📄 License

[MIT](./LICENSE)
