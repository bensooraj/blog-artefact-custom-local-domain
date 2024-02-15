package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
)

func okHandler(w http.ResponseWriter, r *http.Request) {
	slog.Info("Header: ", slog.String("X-API-ServerName", r.Header.Get("X-API-ServerName")))

	fmt.Fprintf(w, "%d OK: %s\n", http.StatusOK, r.Header.Get("X-API-ServerName"))
}

func notOkHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not OK", http.StatusInternalServerError)
}

func main() {
	// Handlers
	http.HandleFunc("/", okHandler)
	http.HandleFunc("/ok", okHandler)
	http.HandleFunc("/notok", notOkHandler)

	port := 443
	certFile := "/app/certificate.pem"
	keyFile := "/app/privatekey.pem"

	slog.Info("server started on localhost.", slog.String("protocol", "https"), slog.Int("port", port))

	// Start the server with HTTPS
	log.Fatal(http.ListenAndServeTLS(fmt.Sprintf(":%d", port), certFile, keyFile, nil))
}
