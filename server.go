package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"
)

func main() {
	port := flag.Int("port", 8080, "Port to listen on")
	filePath := flag.String("file", "", "File to serve")
	uuid := flag.String("uuid", "", "UUID for the URL path")
	downloadName := flag.String("name", "", "Download filename")
	expiration := flag.Int("expire", 600, "Expiration in seconds")
	flag.Parse()

	if *filePath == "" || *uuid == "" {
		fmt.Fprintln(os.Stderr, "Usage: nget-server -port PORT -file FILE -uuid UUID -name NAME -expire SECONDS")
		os.Exit(1)
	}

	if *downloadName == "" {
		*downloadName = filepath.Base(*filePath)
	}

	// Setup HTTP server
	mux := http.NewServeMux()

	// Serve the file only on the UUID path
	mux.HandleFunc("/"+*uuid, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, *downloadName))
		w.Header().Set("Content-Type", "application/octet-stream")
		http.ServeFile(w, r, *filePath)
	})

	// 404 for everything else
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/"+*uuid {
			http.NotFound(w, r)
			return
		}
	})

	addr := fmt.Sprintf(":%d", *port)
	server := &http.Server{Addr: addr, Handler: mux}

	// Handle shutdown signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChan
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		server.Shutdown(ctx)
	}()

	// Auto-shutdown timer
	go func() {
		time.Sleep(time.Duration(*expiration) * time.Second)
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		server.Shutdown(ctx)
	}()

	fmt.Fprintf(os.Stderr, "Serving %s on port %d (expires in %ds)\n", *downloadName, *port, *expiration)

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
