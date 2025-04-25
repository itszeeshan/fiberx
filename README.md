# ⚡ FiberX CLI - ⚠️ WARNING: NOT READY FOR PRODUCTION APPLICATIONS YET! PLEASE DO CHECK AGAIN LATER! STAR IT SO YOU DONT MISS OUT :) YOU CAN USE CORE COMMANDS JUST TO TEST, THOSE ARE READY BUT STILL THERE MIGHT BE EDGE CASES AS I'VEN'T WROTE ANY TEST YET(PS: I KNOW THIS IS BAD 😄). WOULD LOVE A FEEDBACK ❤️

**Speed up GoFiber development without the magic.**  
FiberX is a CLI tool that generates clean, idiomatic Go code for your [Fiber](https://github.com/gofiber/fiber) projects — so you can focus on business logic, not boilerplate.

---

## 🚀 Quick Start

```bash
# Create a new Fiber project with common packages
fiberx new my-api

# Generate a handler with CRUD methods
fiberx add handler user --methods=GET,POST,PUT,DELETE

# Add JWT auth middleware (customize as needed)
fiberx add middleware auth-jwt
```

---

## 🧠 Why FiberX?

Go and Fiber thrive on simplicity — but even purists get tired of:

- Repeating the same handler/service patterns
- Rewriting middleware like auth, logging, rate limiting
- Copy/pasting Dockerfiles, CI configs, or Viper loaders

**FiberX helps by generating:**

✅ Real Go code — no hidden frameworks  
✅ Your structure — flat or layered  
✅ Boilerplate fast — and editable anytime

---

## ✨ Features

### 🛠️ Core Commands

| Command                        | Description                               |
| ------------------------------ | ----------------------------------------- |
| `fiberx new <name>`            | Scaffold a new Fiber app                  |
| `fiberx add handler <name>`    | Add handler with route stubs              |
| `fiberx add middleware <name>` | Generate middleware boilerplate           |
| `fiberx add service <name>`    | Add a service/business logic file         |
| `fiberx dev`                   | Start server with hot-reload (Air/Reflex) |

---

### 📦 Optional Templates

Generate starter code for:

- **Auth**: JWT, API keys, OAuth2
- **Databases**: GORM, SQLx, Ent, MongoDB
- **Infra**: Docker, Kubernetes, GitHub Actions
- **Observability**: Prometheus, OpenTelemetry
- **Cache**: Redis, Memcached
- **Queues**: Kafka, NATS
- **Background Jobs**: Asynq, Custom Workers

```bash
# Example: Add Redis caching boilerplate
fiberx add cache redis
```

---

## 🧰 Installation

```bash
go install github.com/itszeeshan/fiberx@latest
```

---

## ⚙️ Usage

### 1. Start a New Project

```bash
fiberx new my-api --with=postgres,viper
```

Creates:

```
my-api/
├── handlers/      # HTTP handlers
├── middlewares/   # Custom middleware
├── services/      # Business logic
├── config/        # Config loader (Viper)
├── cmd/
│   └── main.go    # Entrypoint
└── .gitignore
```

### 2. Add Features

```bash
# Add CRUD user handler
fiberx add handler user --methods=GET,POST,PUT,DELETE

# Add rate limiter middleware
fiberx add middleware rate-limit
```

### 3. Run & Build

```bash
# Hot-reload in development
fiberx dev

# Build production binary
fiberx build
```

---

## 🧭 Philosophy

### ✅ What FiberX Is

- A productivity booster for Fiber projects
- A code generator — you own the code
- A consistency layer for teams

### ❌ What FiberX Isn't

- Not a framework (no runtime deps)
- Not an architecture enforcer (flat/layered = your choice)
- Not magic — you still write Go code

---

## ⚙️ Optional Config: `.fiberxrc`

```json
{
  "port": 3000,
  "orm": "ent",
  "lint": "golangci-lint",
  "structure": "flat" // or "layered"
}
```

---

## 🤝 Contributing

Found a bug? Open an issue.  
Want a new template? Fork and PR.

We ❤️ code generators, not frameworks.

---

## 🧪 License

MIT License  
Go Report Card | Contributions welcome!

**FiberX gives you wings — not handcuffs. 🚀**
