# Luna CLI

**Luna** is a lightweight database management CLI tool for:

- Database migrations
- Seeding fake data
- Running raw SQL commands

It is built with Go using `cobra` and `gofakeit` under the hood.

---

## Installation

```bash
go build -o luna
```

Creates a binary called `luna` in your working directory.

---

## Commands

### 1. `luna migrate`

Manage database migrations.

**Subcommands:**

| Command | Description |
|:--------|:------------|
| `luna migrate up` | Apply all up migrations |
| `luna migrate down` | Roll back migrations (optional) |
| `luna migrate force` | Force a migration to a specific version (for dirty databases) |
| `luna migrate create <name>` | Create a new empty migration file |

**Example usage:**

```bash
luna migrate up
luna migrate create add_users_table
```

---

### 2. `luna seed`

Populate the database with fake data using seed files.

**Usage:**

```bash
luna seed run <seedfile> <count>
```

| Parameter | Meaning |
|:----------|:--------|
| `seedfile` | Name of your seed YAML or config file (without extension) |
| `count` | How many fake rows to generate and insert |

**Example:**

```bash
luna seed run users 10
```

✅  
Generates 10 fake users based on your `users.yaml` seed file  
and inserts them into the `users` table.

**Example `users.yaml` file:**

```yaml
email: email
first_name: firstname
last_name: lastname
password: password|length=12
gender: gender
```

✅  
Supports faker parameters (like `length=12`, `words=5`, etc.)

---

### 3. `luna db`

Run raw SQL commands directly against your database.

**Usage:**

```bash
luna db raw "<sql_query>"
```

**Example:**

```bash
luna db raw "SELECT * FROM users;"
```

✅  
Executes raw SQL and prints the output.

---

## Configuration

Luna expects a simple `config.yaml` file with database connection settings:

```yaml
db_url: "postgres://admin:password@localhost:5432/yourdb?sslmode=disable"
```

✅  
Load your database URL securely from configuration.

---

## Example Workflows

**Migrate database:**

```bash
luna migrate up
```

**Seed 50 fake users:**

```bash
luna seed run users 50
```

**Run manual SQL to see tables:**

```bash
luna db raw "SELECT tablename FROM pg_tables WHERE schemaname='public';"
```

✅

---

## Roadmap

- [ ] Add migration rollback (`down`) support
- [ ] Add database backup/restore CLI
- [ ] Add custom faker templates
- [ ] Add automatic seeding on environment setup

---

# Summary

✅  
Luna is designed to be:

- Lightweight
- Practical
- Extendable
- Fake-data and migration friendly for local development

✅  
Built for fast dev workflows.
