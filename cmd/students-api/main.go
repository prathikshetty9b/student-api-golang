package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/prathikshetty9b/students-api/pkg/config"
	"github.com/prathikshetty9b/students-api/pkg/http/handlers/student"
	"github.com/prathikshetty9b/students-api/pkg/storage/sqllite"
)

func main() {
	// Load configuration
	cfg := config.MustLoad()

	// database connection
	db, err := sqllite.New(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %s", err.Error())
	}

	slog.Info("database initialized", slog.String("env", cfg.Env), slog.String("storage_path", cfg.StoragePath))

	// Setup router
	router := http.NewServeMux()
	router.HandleFunc("POST /api/students", student.New(db))

	// Setup server
	server := http.Server{
		Addr:    cfg.HTTPServer.Addr,
		Handler: router,
	}

	fmt.Printf("Server started at %s\n", cfg.HTTPServer.Addr)

	// Channel to listen for OS signals
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Start server in a goroutine
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %s", err.Error())
		}
	}()

	// Block until an OS signal is received
	<-done

	// Graceful shutdown
	slog.Info("Shutting down server")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Failed to shutdown server", slog.String("error", err.Error()))
	} else {
		slog.Info("Server shutdown successfully")
	}
}
