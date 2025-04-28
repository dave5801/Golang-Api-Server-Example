# Golang API Server Example

A simple HTTP web server written in Go (Golang) that accepts JSON objects via POST requests and appends them to an in-memory root node with a string parameter. It also supports returning all stored items via GET requests.

## ✨ Features

- Accepts JSON POST requests and stores objects
- Returns all stored objects on GET requests
- Thread-safe (uses sync.Mutex)
- Clean project structure with multiple files
- Simple Makefile for easy build and run commands

## 📂 Project Structure

```
/server-project
├── go.mod
├── main.go        # Entry point
├── types.go       # Defines Root and Item types
├── handlers.go    # HTTP handler functions
├── Makefile       # Helper commands
└── README.md
```

## 🚀 Getting Started

### Prerequisites

- Go 1.20 or higher
- Git

### Installation

1. Clone the repository:
```bash
git clone https://github.com/your-username/server-project.git
cd server-project
```

2. Initialize the module (if needed):
```bash
go mod init server-project
```

3. Run the server:
```bash
make run
```

The server will start at: `http://localhost:8080/`

## 📬 API Endpoints

### GET /
Returns the entire root node (name + all stored items).

**Example Response:**
```json
{
  "name": "My Root Node",
  "items": [
    {
      "title": "Task 1",
      "done": false
    }
  ]
}
```

### POST /
Accepts a JSON object and appends it to the root.

**Example Request:**
```bash
curl -X POST http://localhost:8080/ \
  -H "Content-Type: application/json" \
  -d '{"title":"Learn Go","done":false}'
```

## 🛠 Makefile Commands

| Command | Description |
|---------|-------------|
| `make run` | Run the server |
| `make build` | Build the server binary |
| `make clean` | Remove the binary |
| `make fmt` | Format all Go source files |

## ⚡ Future Improvements

- Add environment variable configuration (port, environment, etc.)
- Add /items sub-path for better routing
- Add PUT/DELETE methods for updating/removing items
- Add basic persistent storage (e.g., save to a JSON file)

## 🧑‍💻 Author

David Franklin, https://github.com/dave5801

## 📜 License

This project is licensed under the MIT License.