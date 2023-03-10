package main

import (
	"errors"
	"fmt"
	"os/exec"
	"log"
	"net/http"
	"os"
)

func startServer(w http.ResponseWriter, r *http.Request) {
cmd := exec.Command("java", "-jar", "server.jar")

    err := cmd.Run()

    if err != nil {
        log.Fatal(err)
    }
}
func main() {
	http.HandleFunc("/start", startServer)

	err := http.ListenAndServe(":3333", nil)

	if errors.Is(err, http.ErrServerClosed) {
    		fmt.Printf("server closed\n")
    	} else if err != nil {
    		fmt.Printf("error starting server: %s\n", err)
    		os.Exit(1)
    	}
	}