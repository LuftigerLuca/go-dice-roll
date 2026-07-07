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
    "result": 4
}
```

## Build

```bash
go build -o dice-roller
```
