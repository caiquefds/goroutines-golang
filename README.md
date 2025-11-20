# Goroutines Golang Example: Event Processing API

This is a simple Go application demonstrating the use of **goroutines**, **channels**, **SQLite**, and the **Gin web framework** to handle event processing asynchronously. Each incoming event is processed by worker goroutines and saved to a SQLite database.

---

## Features

- Asynchronous event processing using goroutines and channels.
- Persistent storage with SQLite using GORM ORM.
- REST API for sending events (`POST /event`).
- Multiple workers processing events concurrently.
- Logging for successful and failed event insertions.

---
## How It Works

1. **Database Setup**  
   The application connects to a SQLite database (`events.db`) and auto-migrates the `Event` model.

2. **Event Workers**  
   - Five worker goroutines are started at runtime.  
   - Each worker listens on the `eventChan` channel.  
   - When a new event arrives, a worker saves it to the database and logs the result.

3. **HTTP API**  
   - A POST endpoint `/event` receives events in JSON format.
   - Example request body:
     ```json
     {
       "id": "24fba774-63e6-4da9-9842-7d6df4e60922",
       "value": 5000
     }
     ```
   - The event is sent to the `eventChan` channel for workers to process.

4. **Concurrency**  
   - Multiple workers can process events simultaneously.
   - Channels ensure safe communication between the main thread and worker goroutines.

---

## Run application

Make sure you have **Go 1.20+** installed.

1. Clean and add dependencies :
   ```bash
   go mod tidy
2. Run application:
   ```bash
   go run main.go