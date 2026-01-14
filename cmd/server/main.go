package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"AdvancedProgramming/internal/server"
)

func main() {
	srv := server.NewServer()

	mux := http.NewServeMux()
	mux.HandleFunc("POST /data", srv.PostData)
	mux.HandleFunc("GET /data", srv.GetData)
	mux.HandleFunc("DELETE /data/{key}", srv.DeleteData)
	mux.HandleFunc("GET /stats", srv.StatsHandler)

	mux.Handle("/", http.FileServer(http.Dir("./web")))

	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	go srv.StartWorker(ctx)

	go func() {
		log.Println("Server started on :8080")
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	<-ctx.Done()
	log.Println("Shutdown initiated")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	httpServer.Shutdown(shutdownCtx)
	log.Println("Server stopped gracefully")
}
