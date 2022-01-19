package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"go.uber.org/zap"
)

type response struct {
	Message string `json:"message"`
}

var (
	port   string
	logger *zap.Logger
)

// logger, _ = zap.NewProduction()

type apiHandler struct{}

func (apiHandler) ServeHTTP(http.ResponseWriter, *http.Request) {}

func init() {
	port = os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName("github.com/mozart409/golang-microservices-demo"),
		newrelic.ConfigLicense("eu01xxddd2334e240a7e656a54a0163ab085NRAL"),
		newrelic.ConfigDistributedTracerEnabled(true),
	)
}

func handler() {

	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	// Handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		logger.Info("Handle '/' endpoint successfully")

		w.Header().Set("Content-Type", "application/json; charset=utf-8") // normal header
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response{
			Message: `OK Success`,
		})
	})

	http.HandleFunc(newrelic.WrapHandleFunc(app, "/users", usersHandler))

	// Run Server
	// http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", port), nil)
}

const message = "Hello World from mux!"

func main() {

	/* tlsConfig := &tls.Config{
		// Causes servers to use Go's default ciphersuite preferences,
		// which are tuned to avoid attacks. Does nothing on clients.
		PreferServerCipherSuites: true,
		// Only use curves which have assembly implementations
		CurvePreferences: []tls.CurveID{
			tls.CurveP256,
			tls.X25519, // Go 1.8 only
		},

		MinVersion: tls.VersionTLS12,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305, // Go 1.8 only
			tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,   // Go 1.8 only
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,

			// Best disabled, as they don't provide Forward Secrecy,
			// but might be necessary for some clients
			// tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			// tls.TLS_RSA_WITH_AES_128_GCM_SHA256,
		},
	} */

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(message))
	})

	srv := &http.Server{
		Addr:         "0.0.0.0:8080",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		//TLSConfig:    tlsConfig,
		Handler: mux,
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
