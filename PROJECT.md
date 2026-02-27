# PROJECT.md — Piggy Bank Specification

> This is the single source of truth for all decisions, architecture, and scope for the Piggy Bank project. Update this file as decisions change or features are completed. Use it to re-orient any AI assistant or collaborator at the start of a session.

---

## Project Overview

Piggy Bank is a self-hosted personal finance web app for individual use. It is not an accounting system and has no tax/legal compliance features. The goal is to give the user full visibility and control over their money — logging every transaction, enforcing budgets, tracking goals, and initiating real payments through M-Pesa and NCBA without leaving the app.

---

## Stack Decisions

| Concern | Choice | Why |
|---|---|---|
| Backend language | Go | Performance, concurrency for payment callbacks, idiomatic control |
| HTTP framework | Gin | Larger community, better middleware ecosystem than Chi for this use case |
| Database | PostgreSQL | Mature, handles financial decimal types correctly, great tooling |
| DB interaction | GORM (start), migrate to SQLC later | GORM for speed of development early on |
| Migrations | golang-migrate | Versioned SQL migration files |
| Auth | JWT (golang-jwt/jwt) | Stateless, works well for a single-user or small-user app |
| Frontend | React + Vite | Vite over CRA for speed; React for eventual React Native transition |
| Styling | TailwindCSS | Utility-first, fast to build with |
| M-Pesa | Daraja API (direct HTTP) | Official Safaricom API |
| NCBA | NCBA API (applied for payment-tier access) | Balance, transfers, transaction history, webhooks |
| Scheduling | gocron | Recurring transactions, monthly report generation |
| Notifications | Africa's Talking (SMS) + SendGrid (email) | Both relevant for Kenyan market |
| Dev tunneling | ngrok | Expose local server for M-Pesa callbacks during development |
| Containerization | Docker | Local dev and production deployment |

**Not using:** Django/Python (Go is better suited here), Create React App (replaced by Vite), SQLite (replaced by PostgreSQL), Flutter (React Native later after React is comfortable)

---

## Separate Projects — DO NOT CONFLATE

**Piggy Bank** — this repo. Personal finance tracker and payment initiator. Tracks all transactions, budgets, goals, accounts.

**Payments Library** — separate repo, separate module. A Go package for M-Pesa and NCBA payment integration for use by ecommerce sites and other apps. It is stateless, has no database, no user model, no business logic beyond the payment communication layer. Piggy Bank may eventually consume it as a dependency, but they are designed independently.

---

## Features

### Core (Build first)
- [ ] User registration and login with JWT auth
- [ ] Multi-account management: NCBA, M-Pesa, Cash
- [ ] Log income (amount, source, account, category, date, note)
- [ ] Log expense (amount, category, payment method, account, date, note)
- [ ] Transaction history with search and filters (date range, category, account, type)
- [ ] Budget creation per category with monthly limits
- [ ] Budget progress tracking with visual indicators
- [ ] Budget alerts when approaching or exceeding limit
- [ ] Dashboard: net balance across accounts, monthly income vs spending, budget health snapshot

### Mid-tier (Build after core is stable)
- [ ] M-Pesa STK push — initiate payment from app, auto-log on callback
- [ ] NCBA — balance check, transfer initiation, transaction history pull, webhooks for auto-logging
- [ ] Goals — target amount, deadline, progress tracking
- [ ] Recurring transactions — flag expected regular transactions
- [ ] Monthly summary report — income, spending by category, savings rate, month-over-month comparison
- [ ] Spending insights — top categories, day-of-week trends, anomalies
- [ ] CSV export of transactions

### Later
- [ ] Portfolio tracker — stocks, crypto, SACCOs, savings accounts
- [ ] Net worth over time (assets minus liabilities)
- [ ] Bill reminders with notifications
- [ ] React Native mobile app

---

## Database Schema

### users
```sql
id            UUID PRIMARY KEY
email         VARCHAR UNIQUE NOT NULL
password_hash VARCHAR NOT NULL
full_name     VARCHAR NOT NULL
currency      VARCHAR DEFAULT 'KES'
created_at    TIMESTAMP
updated_at    TIMESTAMP
```

### accounts
```sql
id         UUID PRIMARY KEY
user_id    UUID REFERENCES users(id)
name       VARCHAR NOT NULL
type       VARCHAR NOT NULL  -- bank | mpesa | cash
balance    NUMERIC(15,2) DEFAULT 0
currency   VARCHAR DEFAULT 'KES'
created_at TIMESTAMP
```

### categories
```sql
id         UUID PRIMARY KEY
user_id    UUID REFERENCES users(id)  -- NULL for system defaults
name       VARCHAR NOT NULL
type       VARCHAR NOT NULL  -- income | expense
color      VARCHAR
icon       VARCHAR
is_default BOOLEAN DEFAULT false
created_at TIMESTAMP
```

### transactions
```sql
id               UUID PRIMARY KEY
user_id          UUID REFERENCES users(id)
account_id       UUID REFERENCES accounts(id)
category_id      UUID REFERENCES categories(id)
amount           NUMERIC(15,2) NOT NULL
type             VARCHAR NOT NULL  -- income | expense | transfer
description      VARCHAR
payment_method   VARCHAR  -- cash | mpesa | card | bank_transfer
reference_id     VARCHAR  -- external transaction ID from M-Pesa or NCBA
status           VARCHAR DEFAULT 'completed'  -- pending | completed | failed
transaction_date TIMESTAMP
created_at       TIMESTAMP
```

### budgets
```sql
id          UUID PRIMARY KEY
user_id     UUID REFERENCES users(id)
category_id UUID REFERENCES categories(id)
amount      NUMERIC(15,2) NOT NULL
spent       NUMERIC(15,2) DEFAULT 0
period      VARCHAR DEFAULT 'monthly'  -- monthly | weekly
start_date  DATE
end_date    DATE
created_at  TIMESTAMP
```

### goals
```sql
id             UUID PRIMARY KEY
user_id        UUID REFERENCES users(id)
name           VARCHAR NOT NULL
target_amount  NUMERIC(15,2) NOT NULL
current_amount NUMERIC(15,2) DEFAULT 0
deadline       DATE
created_at     TIMESTAMP
```

**Key rule:** Never store money as FLOAT. Always use NUMERIC(15,2).

---

## API Endpoints

All endpoints are prefixed with `/api/v1`. All except auth routes and the M-Pesa callback require a valid JWT in the `Authorization: Bearer <token>` header.

### Auth
```
POST   /auth/register
POST   /auth/login
POST   /auth/logout
GET    /auth/me
```

### Accounts
```
GET    /accounts
POST   /accounts
GET    /accounts/:id
PUT    /accounts/:id
DELETE /accounts/:id
```

### Transactions
```
GET    /transactions
POST   /transactions
GET    /transactions/:id
PUT    /transactions/:id
DELETE /transactions/:id
GET    /transactions/export
```

### Categories
```
GET    /categories
POST   /categories
PUT    /categories/:id
DELETE /categories/:id
```

### Budgets
```
GET    /budgets
POST   /budgets
GET    /budgets/:id
PUT    /budgets/:id
DELETE /budgets/:id
```

### Goals
```
GET    /goals
POST   /goals
GET    /goals/:id
PUT    /goals/:id
DELETE /goals/:id
```

### Summary & Insights
```
GET    /summary/monthly
GET    /summary/overview
GET    /insights/spending
```

### M-Pesa
```
POST   /mpesa/stk-push       -- initiate payment (authenticated)
POST   /mpesa/callback        -- Safaricom hits this (public, validate payload)
GET    /mpesa/status/:id      -- check transaction status
```

### NCBA
```
GET    /ncba/balance
GET    /ncba/transactions
POST   /ncba/transfer
```

---

## Backend Folder Structure

```
backend/
├── cmd/server/main.go
├── internal/
│   ├── api/
│   │   ├── handlers/         # One file per resource
│   │   ├── middleware/        # auth.go, cors.go, logger.go
│   │   └── router.go
│   ├── auth/                  # jwt.go, password.go
│   ├── config/                # config.go — all env vars loaded here
│   ├── database/
│   │   ├── db.go
│   │   └── migrations/        # Versioned .sql files
│   ├── models/                # One file per model
│   ├── repository/            # DB queries, one file per model
│   ├── services/              # Business logic, one file per domain
│   └── utils/                 # errors.go, response.go, validator.go
├── pkg/summary/
├── tests/
├── .env.example
├── go.mod
├── go.sum
└── Makefile
```

**Layer pattern:** Handler → Service → Repository. Each layer only knows about the layer directly below it. Handlers validate input and return HTTP responses. Services contain business logic. Repositories talk to the database only.

---

## Frontend Folder Structure

```
frontend/
├── src/
│   ├── api/           # client.js + one file per resource
│   ├── components/
│   │   ├── common/    # Button, Input, Modal, Navbar, Sidebar, Spinner
│   │   ├── transactions/
│   │   ├── budgets/
│   │   ├── goals/
│   │   └── charts/
│   ├── pages/         # Login, Register, Dashboard, Transactions, Budgets, Goals, Accounts, Settings
│   ├── context/       # AuthContext, AppContext
│   ├── hooks/         # Custom React hooks wrapping API calls
│   ├── utils/         # formatCurrency.js, formatDate.js, storage.js
│   ├── App.jsx
│   ├── main.jsx
│   └── index.css
├── .env.example
├── index.html
├── package.json
└── vite.config.js
```

---

## Development Timeline (26 days)

| Phase | Days | Focus |
|---|---|---|
| Week 1 | 1–7 | Backend: models, migrations, auth, core CRUD endpoints |
| Week 2 | 8–14 | React basics (days 8–9), frontend: login, dashboard, transactions |
| Week 3 | 15–21 | Budget UI, M-Pesa integration (STK push + callback) |
| Final push | 22–26 | Polish, NCBA integration if access granted, summary views |

**Note:** Apply for NCBA API access on Day 1 — approval is out of your hands once submitted.

---

## Key Development Notes

- Use `numeric(15,2)` for all money columns, never `float`
- M-Pesa callback URL must be publicly accessible — use ngrok during local development
- Validate M-Pesa callback payloads on receipt, don't trust blindly
- NCBA API requires formal application for payment-tier access (not just read access)
- All secrets via environment variables, never hardcoded
- Format code with `gofmt` before committing