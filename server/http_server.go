package server

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/piyushverma013/token-athena/config"
	"github.com/piyushverma013/token-athena/middleware"
)

type HTTPServer struct {
	Router *gin.Engine
}

func Start(ctx context.Context, appConfig *config.AppConfig) error {
	server := &HTTPServer{}
	gin.SetMode(appConfig.GinMode)
	router := gin.Default()

	router.Use(middleware.HealthCheck())
	server.setUpRoutes(router, appConfig)
	server.Router = router

	return server.Serve(ctx, appConfig)

}

// Start runs the HTTP server on a specific address
func (server *HTTPServer) Serve(ctx context.Context, appConfig *config.AppConfig) error {
	srv := &http.Server{
		Addr:         appConfig.HTTPServerAddress,
		Handler:      server.Router,
		ReadTimeout:  time.Duration(appConfig.AppReadTimeOut) * time.Second,
		WriteTimeout: time.Duration(appConfig.AppWriteTimeOut) * time.Second,
		IdleTimeout:  time.Duration(appConfig.AppIdleTimeOut) * time.Second,
	}

	go func() {
		//service connections
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	<-ctx.Done()
	log.Println(
		"Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)

	defer cancel()

	err := srv.Shutdown(ctx)
	if err != nil {
		log.Printf("Server Shutdown: %v", err)
	}

	<-ctx.Done()
	log.Println("timeout of 1 second.")

	log.Println("server exiting")
	return nil
}
