package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/Webblurt/garantex-usdt-rates/internal/config"
	grpcsvc "github.com/Webblurt/garantex-usdt-rates/internal/grpc"
	"github.com/Webblurt/garantex-usdt-rates/internal/storage"
	ratespb "github.com/Webblurt/garantex-usdt-rates/proto/ratespb"
	"google.golang.org/grpc"
)

func main() {
	// Load config
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// Connect to PostgreSQL
	store, err := storage.New(cfg.Postgres.DSN)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	log.Println("Connected to PostgreSQL")

	// Create gRPC server
	grpcServer := grpc.NewServer()
	rateServer := grpcsvc.NewRateServer(store, cfg.Grinex.URL)
	ratespb.RegisterRateServiceServer(grpcServer, rateServer)

	// Listen on port
	listenAddr := ":" + cfg.Server.Port
	lis, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatalf("failed to listen on %s: %v", listenAddr, err)
	}
	log.Printf("gRPC server is running on %s", listenAddr)

	// Graceful shutdown setup
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve gRPC: %v", err)
		}
	}()

	// Wait for shutdown signal
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop
	log.Println("Shutting down gRPC server...")
	grpcServer.GracefulStop()
	log.Println("Server stopped gracefully")
}
