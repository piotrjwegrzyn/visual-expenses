package main

import "log/slog"

func main() {
	slog.Info(helloWorld())
}

func helloWorld() string {
	return "Hello, World!"
}
