package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// Register routes.
	registerLoginRoutes(mux)
	registerCalendarRoutes(mux)
	registerDashboardRoutes(mux)

	// Setup the static files.
	fs := http.FileServer(http.Dir("web/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	// Start the server.
	fmt.Println("Listening on :8080")
	http.ListenAndServe(":8080", mux)
}
