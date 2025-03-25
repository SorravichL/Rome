# 🏛️ Rome Monorepo Project

This is a fullstack monorepo boilerplate built for the Intania Hackathon. It contains two services (`go-backend` and `ts-backend`) communicating with each other using a shared OpenAPI spec for full **type safety**, and supports **CI/CD**, **linting**, **formatting**, and **PostgreSQL** database logging.

---

## 🧱 Project Structure

```
rome-project/
├── go-backend/          # Go service with database logging
├── ts-backend/          # TypeScript service using Express and Prisma
├── shared/              # Shared OpenAPI definition
├── scripts/             # Dev + Generate scripts
├── .github/             # GitHub Actions CI
└── README.md
```

---

## ⚡ Features

- ✅ Two services (Go + TypeScript)
- ✅ Communication with full **Type Safety** via OpenAPI
- ✅ PostgreSQL logging (via Prisma and `database/sql`)
- ✅ Shared OpenAPI spec
- ✅ Linter and Formatter (Go + TypeScript)
- ✅ Unit tests
- ✅ CI (GitHub Actions)
- ✅ CD-ready (Deployable to Render.com)

---

## 🚀 Getting Started

### 1. Clone the repo

```bash
git clone https://github.com/SorravichL/Rome.git
cd Rome
```

### 2. Set up environment

Create two `.env` files:

#### `go-backend/.env`
```env
PORT=5001
TS_BACKEND_URL=http://localhost:5002
DATABASE_URL=your_postgres_connection_string
```

#### `ts-backend/.env`
```env
PORT=5002
GO_BACKEND_URL=http://localhost:5001
DATABASE_URL=your_postgres_connection_string
```

---

### 3. Install dependencies

```bash
cd ts-backend
npm install

cd ../go-backend
go mod tidy
```

---

### 4. Generate OpenAPI types

```bash
chmod +x scripts/generate.sh
./scripts/generate.sh
```

---

### 5. Migrate and generate Prisma client

```bash
cd ts-backend
npx prisma migrate dev --name init
npx prisma generate
```

---

### 6. Start both services

```bash
chmod +x scripts/dev.sh
./scripts/dev.sh
```

---

## ✅ Available Commands

### Lint & Format

```bash
# TypeScript Lint
cd ts-backend && npm run lint

# Go Lint
cd go-backend && golangci-lint run

# TypeScript Format
cd ts-backend && npm run format

# Go Format
cd go-backend && go fmt ./...
```

Or use the one-click all-in-one:

```bash
chmod +x scripts/clean.sh
./scripts/clean.sh
```

---

### Run Unit Tests

```bash
cd ts-backend
npm run test
```

---

## 🛠 Technologies Used

- 🟨 TypeScript + Express + Prisma
- 🟦 Go + `database/sql`
- 📦 PostgreSQL
- 📜 OpenAPI v3
- 🧪 Jest
- 🧹 ESLint + Prettier
- 🧽 golangci-lint
- 🐙 GitHub Actions CI/CD

---

## 🌍 Live Demo (Hosted on Render)

| Service         | URL                                  |
|------------------|--------------------------------------|
| TypeScript (TS)  | [https://rome-ts.onrender.com](https://rome-ts.onrender.com) |
| Go (Golang)      | [https://rome-go.onrender.com](https://rome-go.onrender.com) |

---

## 🗃️ Public PostgreSQL Database

You can test using our shared database. No setup required!

- 🛠️ Prisma Ready
- 📦 Logs messages with `from`, `to`, `message`, and `timestamp`
### 🌐 External (Local Development):

Feel free to use my database for test
DATABASE_URL=postgresql://rome_db_user:FSDH7AOOUOXuUeIGOLPalKLsO0YwOHfS@dpg-cvh85ldrie7s73eld5s0-a.singapore-postgres.render.com/rome_db

### 🔒 Internal (Render.com):
DATABASE_URL=postgresql://rome_db_user:FSDH7AOOUOXuUeIGOLPalKLsO0YwOHfS@dpg-cvh85ldrie7s73eld5s0-a/rome_db


---

## 🔁 API Endpoints & Examples

### 1. `POST /send` (Go)

Forwards a message from Go to TypeScript backend and logs it to the DB.

**URL**: `http://localhost:5001/send`

```json
{
  "from": "go-service",
  "to": "ts-service",
  "message": "Hello from Go!",
  "date": "2025-03-25T12:00:00Z"
}
```

---
### 2. POST /log (Go)
Logs a message received from TypeScript (no forwarding).

URL: http://localhost:5001/log

```json
{
  "from": "ts-service",
  "to": "go-service",
  "message": "Hello from TypeScript!",
  "date": "2025-03-25T12:00:00Z"
}
```
### 3. GET /logs (Go)
Fetches the latest 10 logs from the database.

URL: http://localhost:5001/logs

### 4. POST /send (TypeScript)
Sends a message from TypeScript to Go and logs it in the database.

URL: http://localhost:5002/send

```json
{
  "from": "ts-service",
  "to": "go-service",
  "message": "Hello from TypeScript!",
  "date": "2025-03-25T12:00:00Z"
}
```
### 5. POST /log (TypeScript)
Logs a message from Go into the database.

URL: http://localhost:5002/log

```json
{
  "from": "go-service",
  "to": "ts-service",
  "message": "Hello from Go!",
  "date": "2025-03-25T12:00:00Z"
}
```
### 6. GET /logs (TypeScript)
Fetches the latest 10 logs from the database.

URL: http://localhost:5002/logs


## 📄 License

Built with ❤️ for Intania Hackathon.
