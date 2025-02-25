package main

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/grokkos/go-isa-retail-service/internal/api/handler"
	"github.com/grokkos/go-isa-retail-service/internal/repository"
	"github.com/grokkos/go-isa-retail-service/internal/service"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Initialize repositories
	customerRepo := repository.NewInMemoryCustomerRepository()
	fundRepo := repository.NewInMemoryFundRepository()
	investmentRepo := repository.NewInMemoryInvestmentRepository()

	// Initialize services
	fundService := service.NewFundService(fundRepo)
	investmentService := service.NewInvestmentService(investmentRepo, customerRepo, fundRepo)

	// Initialize handlers
	fundHandler := handler.NewFundHandler(fundService)
	investmentHandler := handler.NewInvestmentHandler(investmentService, fundService)

	// Set up router
	r := mux.NewRouter()

	// API version prefix
	api := r.PathPrefix("/api/v1").Subrouter()

	// Fund routes
	api.HandleFunc("/funds", fundHandler.ListFunds).Methods("GET")
	api.HandleFunc("/funds/{id}", fundHandler.GetFund).Methods("GET")

	// Investment routes
	api.HandleFunc("/investments", investmentHandler.CreateInvestment).Methods("POST")
	api.HandleFunc("/investments/{id}", investmentHandler.GetInvestment).Methods("GET")
	api.HandleFunc("/customers/{id}/investments", investmentHandler.GetCustomerInvestments).Methods("GET")

	// Configure server
	srv := &http.Server{
		Handler:      r,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// Start server in a goroutine
	go func() {
		log.Println("Retail ISA API starting on port 8080")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error starting server: %s", err)
		}
	}()

	// Set up graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// Create context with timeout for shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %s", err)
	}

	log.Println("Server gracefully stopped")
}
