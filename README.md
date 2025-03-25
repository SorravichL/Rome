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

## 📦 Deployment (Render.com)

Set environment variables in **Render Dashboard**:

- `DATABASE_URL`
- `GO_BACKEND_URL` (in TS service)
- `TS_BACKEND_URL` (in Go service)
- `PORT`

Make sure to add:
- **Build Command:** `npm install && npm run build`
- **Start Command:** `npm run dev`

---

## 📄 License

MIT — Built with ❤️ for Intania Hackathon.