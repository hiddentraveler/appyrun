package server

import (
	"appyrun/internal/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	h := *handlers.NewHandlers(s)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"https://*", "http://*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	//clerkKey := os.Getenv("CLERK_SECRET_KEY")
	//config := &clerk.ClientConfig{}
	//config.Key = &clerkKey
	//if clerkKey == "" {
	//	log.Fatal("CLERK_SECRET_KEY not set")
	//}
	//
	//client := user.NewClient(config)
	e.GET("/", h.HomeHandler)

	e.GET("/health", h.HealthHandler)
	//fileServer := http.FileServer(http.Dir("static"))
	//e.GET("/static/", http.StripPrefix("/static/", fileServer))
	e.Static("/static", "static")

	e.GET("/login", h.LoginHandler)

	return e
}
