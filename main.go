package main

import (
	"log/slog"
	"net/http"
)

func main() {
	http.HandleFunc("/", getRoot)

	if err := http.ListenAndServe(":80", nil); err != nil {
		slog.Error("error starting server", slog.Attr{Key: "error", Value: slog.StringValue(err.Error())})
	}
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte(helloWorld())); err != nil {
		slog.Error("error during writing response")
	}
}

func helloWorld() string {
	return "Hello, World!"
}
