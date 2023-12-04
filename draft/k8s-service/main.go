package main

import (
	"fmt"
	"net/http"
	"os"
	"sync/atomic"

	"github.com/joho/godotenv"
)

var appIsReady atomic.Uint64

func main() {
	godotenv.Load()
	printServiceDetails()

	deploymentPort := mustParseEnv("PORT", "8080", false)

	mux := http.NewServeMux()

	mux.HandleFunc("/health/ready", healthReadyHandler)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, World!")
	})

	// Print bootup sequence
	fmt.Println("Ready to receive traffic")
	appIsReady.Add(1)
	http.ListenAndServe(":"+deploymentPort, mux)

}

func healthReadyHandler(w http.ResponseWriter, r *http.Request) {
	if appIsReady.Load() == 0 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Not Ready")
	} else {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "OK")
	}
}

// printServiceDetails does log directly to stdout, not to the logging service
func printServiceDetails() {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println("hostname:", hostname)
	}

	usedEnvironmentVariables := []string{"PORT", "RED_ENVIRONMENT", "KUBERNETES_DEPLOYMENT", "KUBERNETES_NAMESPACE"}
	for _, value := range usedEnvironmentVariables {
		fmt.Printf("%s=%s\n", value, os.Getenv(value))
	}

	// Custom buildinfo.json here
	// to print the pipeline build number, repo name and git hash

}
