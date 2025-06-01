# Proxy Server

A simple HTTP proxy server written in Go. It accepts JSON-based requests to forward HTTP calls to external services and returns structured JSON responses. It also stores logs of all requests and responses in memory.

## Features

- Accepts method, URL, and headers via JSON
- Sends HTTP request to external API
- Returns status code, headers, and response length
- Stores logs using sync.Map
- Allows retrieving logs by ID

## Example Request

**POST** `/proxy`

```json
{
  "method": "GET",
  "url": "https://jsonplaceholder.typicode.com/posts/1",
  "headers": {
    "Accept": "application/json"
  }
}
```

**Response:**

```json
{
  "id": "abc123",
  "status": 200,
  "headers": { ... },
  "length": 292
}
```

**GET** `/logs/abc123`  
Returns the full request and response log by ID

## Run with Docker

```bash
docker compose up --build
```

## Port

Runs on http://localhost:8080
