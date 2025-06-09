package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/riadalimemmedov/api-server-project/api-server-project/internal/config"
	"github.com/riadalimemmedov/api-server-project/api-server-project/internal/router"
	logger "github.com/riadalimemmedov/api-server-project/api-server-project/pkg/log"
	"go.uber.org/zap"
)

func main() {
	// Load configuration using your existing Load() method
	cfg := config.Load()

	// Setup router
	r := router.SetupRouter()

	// Create server address from config
	addr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)

	// Create HTTP server
	srv := &http.Server{
		Addr:    addr,
		Handler: r,
		// Production-ready timeouts
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Start server in a goroutine
	go func() {
		logger.GetLogger().Info("Server starting on", zap.String("address", addr))

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.GetLogger().Error("Failed to start server", zap.Error(err))
		}
	}()

	// Setup graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive our signal
	<-quit
	logger.GetLogger().Info("Shutting down server...")

	// Create a deadline for shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.GetLogger().Error("Server forced to shutdown", zap.Error(err))
	}

	logger.GetLogger().Error("Server exited gracefully")
}
