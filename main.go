package main

import (
	"encoding/json"
	"log/slog"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type RollResponse struct {
	Result      int     `json:"result"`
	Sides       int     `json:"sides"`
	Average     float64 `json:"average"`
	Probability float64 `json:"probability"`
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /roll/{sides}", getRoll)

	addr := ":8080"
	slog.Info("starting the webserver", "address", addr)
	err := http.ListenAndServe(addr, middleware(mux))
	if err != nil {
		slog.Error("error while starting the webserver!", "error", err)
	}
}

func getRoll(w http.ResponseWriter, r *http.Request) {
	sidesStr := r.PathValue("sides")
	sides, err := strconv.Atoi(sidesStr)
	if err != nil {
		res := ErrorResponse{
			Error: err.Error(),
		}
		writeResponse(http.StatusBadRequest, w, res)
		return
	}

	if sides <= 0 {
		res := ErrorResponse{
			Error: "sides must be a number!",
		}
		writeResponse(http.StatusBadRequest, w, res)
		return
	}

	result := rand.Intn(sides) + 1
	res := RollResponse{
		Result:      result,
		Sides:       sides,
		Average:     float64(1+sides) / 2,
		Probability: math.Round((1.0/float64(sides))*100) / 100,
	}
	writeResponse(http.StatusOK, w, res)
}

func writeResponse(code int, writer http.ResponseWriter, response any) {
	writer.WriteHeader(code)
	if err := json.NewEncoder(writer).Encode(response); err != nil {
		slog.Warn("an error occurred!", "error", err)
	}
}

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		defer func() {
			err := recover()
			if err != nil {
				slog.Warn("web middleware recovery caught something!", "error", err)
				http.Error(w, "There was an internal server error", http.StatusInternalServerError)
			}
		}()

		start := time.Now()
		next.ServeHTTP(w, r)
		difference := time.Since(start)
		slog.Info(r.Method, "path", r.URL, "duration", difference)
	})
}
