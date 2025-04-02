package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	firebase "firebase.google.com/go/v4"
	"github.com/gin-gonic/gin"
	"github.com/notify-vital/backend/config"
	"github.com/notify-vital/backend/internal/auth"
	"github.com/notify-vital/backend/internal/handlers"
	"github.com/notify-vital/backend/internal/middleware"
	"google.golang.org/api/option"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Set Gin mode based on environment
	if cfg.Server.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Initialize Firebase
	ctx := context.Background()
	opt := option.WithCredentialsFile(cfg.Firebase.CredentialsFile)
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("Error initializing Firebase app: %v", err)
	}

	// Initialize Firestore client
	firestoreClient, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("Error initializing Firestore client: %v", err)
	}
	defer firestoreClient.Close()

	// Initialize Firebase Auth service
	firebaseAuth, err := auth.NewFirebaseAuth(app)
	if err != nil {
		log.Fatalf("Error initializing Firebase Auth: %v", err)
	}

	// Initialize user repository
	userRepository := auth.NewUserRepository(firestoreClient)

	// Initialize auth handler
	authHandler := handlers.NewAuthHandler(firebaseAuth, userRepository)

	// Setup router
	router := gin.Default()

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"time":   time.Now().Format(time.RFC3339),
		})
	})

	// API routes
	apiV1 := router.Group("/api/v1")

	// Auth routes (no auth required)
	authRoutes := apiV1.Group("/auth")
	{
		authRoutes.POST("/signup", authHandler.SignUp)
		authRoutes.POST("/signin", authHandler.SignIn)
	}

	// Protected routes (auth required)
	protectedRoutes := apiV1.Group("/")
	protectedRoutes.Use(middleware.AuthMiddleware(firebaseAuth))
	{
		// User profile
		protectedRoutes.GET("/profile", authHandler.GetProfile)
		protectedRoutes.PUT("/profile", authHandler.UpdateProfile)
		protectedRoutes.DELETE("/account", authHandler.DeleteAccount)

		// Add other protected routes here
	}

	// Start server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.Server.Port),
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// Graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
