# 🚀 Pure Go HTTP API

A comprehensive learning project demonstrating core HTTP server development with Go's built-in `net/http` package. Master the fundamentals of building production-ready APIs without external frameworks.

---

## 📋 Table of Contents

- [Overview](#overview)
- [Project Structure](#project-structure)
- [Prerequisites](#prerequisites)
- [Quick Start](#quick-start)
- [Modules](#modules)
- [Learning Objectives](#learning-objectives)
- [API Examples](#api-examples)
- [Development](#development)
- [Best Practices](#best-practices)
- [Resources](#resources)

---

## 🎯 Overview

This project is a **structured, progressive learning path** for Go developers aiming to understand HTTP server fundamentals. Each module builds upon previous concepts, from basic server setup to handling external API calls with proper error management.

**Key Highlights:**
- ✅ Zero external web framework dependencies
- ✅ Pure `net/http` standard library implementation
- ✅ Progressive complexity (8 well-organized modules)
- ✅ Production-ready patterns and error handling
- ✅ JSON encoding/decoding best practices
- ✅ Real-world external API integration

---

## 📁 Project Structure

```
pure-go-http-api/
├── go.mod
├── README.md
├── 01_basic_http_server/          # Foundation: hello world HTTP server
├── 02_http_multiple_routes/       # Route handling and query parameters
├── 03_json_encoder_encode_details/# JSON encoding with http.ResponseWriter
├── 04_json_decoder_decode_details/# JSON decoding from request bodies
├── 05_http_get/                   # Making HTTP GET requests
├── 06_reading_response_body/      # Parsing external response data
├── 07_json_unmarshal_into_struct/ # Unmarshaling JSON into Go structs
└── 08_external_api/               # Complete workflow: fetch & serve external data
```

---

## 🔧 Prerequisites

- **Go 1.21+** (this project uses Go 1.26.1)
- Basic understanding of Go syntax
- A terminal/command line interface
- Optional: `curl` or Postman for testing API endpoints

**Verify Go installation:**
```bash
go version
```

---

## 🚀 Quick Start

### 1. Clone or Navigate to Project

```bash
cd pure-go-http-api
```

### 2. Build & Run Any Module

Each module runs independently. To run a specific example:

```bash
# Run the basic HTTP server (Module 01)
cd 01_basic_http_server
go run main.go

# In a new terminal, test the endpoint
curl http://localhost:8080/hello
```

---

## 📚 Modules

### **01: Basic HTTP Server**
**Goal:** Establish a simple HTTP server with a single route.

- **Port:** `:8080`
- **Route:** `GET /hello`
- **Response:** Plain text message
- **Concepts:** Handler functions, HTTP methods, `http.ListenAndServe`

```bash
cd 01_basic_http_server && go run main.go
curl http://localhost:8080/hello
```

---

### **02: HTTP Multiple Routes**
**Goal:** Handle multiple routes and query parameters.

- **Port:** `:5000`
- **Routes:**
  - `GET /` – Welcome message
  - `GET /hello?name=YourName` – Personalized greeting
- **Concepts:** Route registration, query parameter extraction, URL parsing

```bash
cd 02_http_multiple_routes && go run main.go
curl "http://localhost:5000/hello?name=Sasmith"
```

---

### **03: JSON Encoder - Encode Details**
**Goal:** Learn to encode Go structs/maps to JSON in HTTP responses.

- **Port:** `:5000`
- **Route:** `GET /ok`
- **Response:** JSON object with status, message, and timestamp
- **Concepts:** `json.NewEncoder`, `http.ResponseWriter`, Content-Type headers

```bash
cd 03_json_encoder_encode_details && go run main.go
curl http://localhost:5000/ok | jq .
```

---

### **04: JSON Decoder - Decode Details**
**Goal:** Parse JSON from incoming request bodies into Go structs.

- **Port:** `:5000`
- **Route:** `POST /decode`
- **Request:** JSON body with user data
- **Response:** Decoded object confirmation
- **Concepts:** `json.NewDecoder`, request body parsing, struct tags

```bash
cd 04_json_decoder_decode_details && go run main.go
curl -X POST http://localhost:5000/decode \
  -H "Content-Type: application/json" \
  -d '{"name":"Sasmith","age":25}'
```

---

### **05: HTTP GET**
**Goal:** Make outbound HTTP GET requests to external services.

- **Port:** `:5000`
- **Concepts:** `http.Get`, response handling, error management

```bash
cd 05_http_get && go run main.go
curl http://localhost:5000/fetch
```

---

### **06: Reading Response Body**
**Goal:** Properly read and handle response bodies from external APIs.

- **Port:** `:5000`
- **Concepts:** `io.ReadAll`, defer patterns, resource cleanup

```bash
cd 06_reading_response_body && go run main.go
curl http://localhost:5000/data
```

---

### **07: JSON Unmarshal Into Struct**
**Goal:** Convert external JSON API responses into typed Go structs.

- **Port:** `:5000`
- **Concepts:** `json.Unmarshal`, struct field mapping, type safety

```bash
cd 07_json_unmarshal_into_struct && go run main.go
curl http://localhost:5000/unmarshal
```

---

### **08: External API (Complete Example)**
**Goal:** Integrate all concepts: fetch cat facts from an external API and serve them.

- **Port:** `:5000`
- **Route:** `GET /external`
- **External Source:** [catfact.ninja](https://catfact.ninja)
- **Response:** JSON with cat fact data wrapped in metadata
- **Concepts:** End-to-end API integration, error handling, proper status codes

```bash
cd 08_external_api && go run main.go
curl http://localhost:5000/external | jq .
```

**Sample Response:**
```json
{
  "ok": true,
  "time": "2026-04-22T12:30:45.123Z",
  "external": {
    "source": "catfact.ninja",
    "fact": "Cats can rotate their ears independently...",
    "length": 42
  }
}
```

---

## 🎓 Learning Objectives

By completing this project, you will understand:

- ✅ How HTTP servers work under the hood
- ✅ Request routing and handler functions
- ✅ HTTP methods, status codes, and headers
- ✅ JSON encoding and decoding patterns
- ✅ Query parameters and form data
- ✅ Making external API calls
- ✅ Error handling and logging
- ✅ Resource cleanup with defer
- ✅ Response content negotiation
- ✅ Best practices for production APIs

---

## 📡 API Examples

### Testing with cURL

```bash
# Simple GET request
curl http://localhost:5000/hello

# GET with query parameters
curl "http://localhost:5000/hello?name=Alice&age=30"

# POST with JSON body
curl -X POST http://localhost:5000/endpoint \
  -H "Content-Type: application/json" \
  -d '{"key": "value"}'

# Pretty-print JSON response
curl http://localhost:5000/endpoint | jq .
```

### Testing with Postman

1. Open **Postman**
2. Create a new request
3. Set **Method** to `GET` or `POST`
4. Enter **URL:** `http://localhost:5000/endpoint`
5. Add **Headers** if needed (e.g., `Content-Type: application/json`)
6. Add **Body** for POST requests
7. Click **Send**

---

## 🛠️ Development

### Run All Modules
```bash
# Terminal 1
cd 01_basic_http_server && go run main.go

# Terminal 2 (new terminal)
cd 02_http_multiple_routes && go run main.go

# And so on for each module...
```

### Verify Code Compiles
```bash
go test ./...
```

### Build Executable
```bash
cd <module_folder>
go build -o server main.go
./server
```

### Format and Lint
```bash
go fmt ./...
go vet ./...
```

---

## ✨ Best Practices Demonstrated

| Practice | Module | Example |
|----------|--------|---------|
| **Error Handling** | 08 | Checking `err` before processing |
| **Method Validation** | 01 | `if r.Method != http.MethodGet` |
| **Resource Cleanup** | 08 | `defer res.Body.Close()` |
| **JSON Encoding** | 03 | `json.NewEncoder(w).Encode(data)` |
| **Header Management** | 03 | `w.Header().Set("Content-Type", ...)` |
| **Query Parameters** | 02 | `r.URL.Query().Get("param")` |
| **Status Codes** | 08 | `http.StatusOK`, `http.StatusBadGateway` |
| **Type Safety** | 07 | Struct tags for JSON unmarshaling |

---

## 📖 Resources

### Official Go Documentation
- [net/http Package](https://pkg.go.dev/net/http)
- [encoding/json Package](https://pkg.go.dev/encoding/json)
- [Building Web Applications in Go](https://golang.org/doc/articles/wiki/)

### HTTP & REST Fundamentals
- [HTTP Status Codes Reference](https://httpwg.org/specs/rfc7231.html#status.codes)
- [REST API Best Practices](https://restfulapi.net/)
- [JSON Schema](https://json-schema.org/)

### Go Learning Paths
- [Golang Official Tour](https://tour.golang.org)
- [Effective Go](https://golang.org/doc/effective_go)
- [The Go Programming Language (Book)](https://www.gopl.io/)

---

## 🤝 Contributing

This is a personal learning project. Feel free to:
- Experiment with each module
- Extend examples with new features
- Modify responses and test different scenarios
- Add comments to deepen understanding

---

## 📝 License

This project is open source and available for educational purposes.

---

## 🎉 Next Steps

After completing this project:

1. **Explore Web Frameworks** → Learn frameworks like [Gin](https://gin-gonic.com/), [Echo](https://echo.labstack.com/), or [Fiber](https://gofiber.io/)
2. **Add Database Integration** → Connect to PostgreSQL, MongoDB, etc.
3. **Implement Authentication** → Add JWT or OAuth2
4. **Build Microservices** → Deploy APIs in containers (Docker, Kubernetes)
5. **Performance Testing** → Use tools like `wrk` or `ab` to benchmark

---

**Happy Learning! 🚀**

For questions or feedback, review the code comments and experiment with modifications to deepen your understanding of Go's HTTP capabilities.

