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
	w.Write([]byte(helloWorld()))
}

func helloWorld() string {
	return "Hello, World!"
}
