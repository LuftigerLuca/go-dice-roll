# dice-roller

[![Go Version](https://img.shields.io/github/go-mod/go-version/LuftigerLuca/go-dice-roll)](https://github.com/LuftigerLuca/go-dice-roll)
[![Repo Size](https://img.shields.io/github/repo-size/LuftigerLuca/go-dice-roll)](https://github.com/LuftigerLuca/go-dice-roll)

Simple dice rolling HTTP API built with Go.

## Usage

```bash
go run main.go
```

Server starts on `localhost:8080`.

### Endpoint

**`GET /roll/{sides}`**

Roll a die with the given number of sides.

**Example:**

```bash
curl "localhost:8080/roll/6"
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

| Condition | Status | Message |
|---|---|---|
| `sides` is not a number | `400` | `"strconv.Atoi error message"` |
| `sides` is zero or negative | `400` | `"sides must be a number!"` |

## Build

```bash
go build -o dice-roller
```
