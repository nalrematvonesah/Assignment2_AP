Here is a **clean, professional `README.md`** that you can **paste directly** into your project.
It is written in **academic + industry style**, easy to defend, and matches your assignment **exactly**.

---

```md
# Assignment 2 – Concurrent Web Server in Go

**Course:** Advanced Programming 1  
**Assignment:** 2  
**Student:** Taubakabyl Nurlybek  
**Group:** SE2499  

---

## 📌 Project Overview

This project implements a **concurrent web server** using the Go programming language and the standard `net/http` package.

The server supports multiple REST endpoints, safely handles concurrent requests using mutexes, runs a background worker using goroutines, and shuts down gracefully using context and OS signals.

The project demonstrates understanding of:
- Go concurrency model
- Goroutines and mutexes
- Context-based cancellation
- RESTful API design
- Clean project structure

---

## 🏗️ Project Structure

```

assignment2/
├── cmd/
│   └── server/
│       └── main.go          # Application entry point
├── internal/
│   ├── server/
│   │   ├── server.go        # Server state and statistics
│   │   ├── handlers.go     # HTTP handlers
│   │   ├── worker.go       # Background worker
│   └── storage/
│       └── memory.go       # In-memory storage with mutex
├── go.mod
└── README.md

````

### Structure explanation
- `cmd/server` – contains the `main` package and application startup logic  
- `internal/server` – HTTP handlers, concurrency logic, background worker  
- `internal/storage` – thread-safe in-memory database  
- `go.mod` – Go module definition  

This structure follows **standard Go project conventions**.

---

## ⚙️ Technologies Used

- Go 1.22+
- net/http
- sync.Mutex
- goroutines
- context
- OS signal handling

No external libraries are used.

---

## 🌐 API Endpoints

### `POST /data`
Stores key-value data in memory.

**Request body (JSON):**
```json
{
  "key": "value"
}
````

**Response:**

```json
{
  "status": "stored"
}
```

---

### `GET /data`

Returns all stored data.

**Response:**

```json
{
  "name": "Alice",
  "city": "Almaty"
}
```

---

### `DELETE /data/{key}`

Deletes a value by key.

**Response:**

```json
{
  "deleted": "name"
}
```

---

### `GET /stats`

Returns server statistics.

**Response:**

```json
{
  "total_requests": 5,
  "database_size": 1,
  "uptime_seconds": 42
}
```

---

## 🔒 Thread Safety

* All shared resources are protected using `sync.Mutex`
* The in-memory database is isolated in a separate storage layer
* Request counters and server stats are accessed safely
* The project passes Go race-condition checks

---

## 🔄 Background Worker

A background goroutine:

* Starts when the server starts
* Logs server statistics every 5 seconds
* Stops automatically when the server shuts down

This is implemented using:

* `time.Ticker`
* `context.Context`

---

## 🛑 Graceful Shutdown

The server supports **graceful shutdown**:

* OS signals (`Ctrl + C`) are captured
* Active requests are allowed to complete
* Background worker stops cleanly
* The server shuts down without data corruption

This is implemented using:

* `signal.NotifyContext`
* `http.Server.Shutdown`

---

## ▶️ How to Run the Project

### 1. Check Go version

```bash
go version
```

### 2. Download dependencies

```bash
go mod tidy
```

### 3. Run the server

```bash
go run ./cmd/server
```

Server will start on:

```
http://localhost:8080
```

---

## 🧪 Example Testing Commands

```bash
curl -X POST http://localhost:8080/data \
-H "Content-Type: application/json" \
-d '{"name":"Alice"}'
```

```bash
curl http://localhost:8080/data
```

```bash
curl http://localhost:8080/stats
```

---

## 🧠 Defense Notes

Key points to explain during defense:

* Why mutexes are required in concurrent servers
* How `net/http` handles requests using goroutines
* Difference between channels and context
* Benefits of graceful shutdown
* Reason for separating storage, server logic, and main entry point

---

## ✅ Conclusion

This project fulfills all assignment requirements:

* Concurrent web server
* Thread-safe shared state
* Background worker
* Graceful shutdown
* Clean and professional architecture

The solution is simple, reliable, and easy to maintain.

---
