# dice-roller

Simple dice rolling HTTP API built with Go and Gin.

## Usage

```bash
go run main.go
```

Server starts on `localhost:8080`.

### Endpoint

**`GET /roll?sides=<number>`**

Roll a die with the given number of sides.

**Example:**

```bash
curl "localhost:8080/roll?sides=6"
```

**Response:**

```json
{
    "result": 4,
    "sides": 6,
    "average": 3.5,
    "probability": 0.17
}
```

**Error responses:**

| Query | Status | Message |
|---|---|---|
| `sides` is missing or not a number | `400` | `"Sides must be a number!"` |
| `sides` is zero or negative | `400` | `"Sides must not be zero or lower"` |

## Build

```bash
go build -o dice-roller
```
