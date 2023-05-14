package main

import (
	"net/http"

	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance for serving content
	echoMain := echo.New()
	echoMain.HideBanner = true

	// Echo instance for serving prometheus
	echoProm := echo.New()
	echoProm.HideBanner = true

	// Middleware
	promMid := prometheus.NewPrometheus("echomain", nil)

	// Scrape metrics from Main Server
	echoMain.Use(promMid.HandlerFunc)
	// w3svc structured logs
	echoMain.Use(middleware.Logger())
	// Server stays up!
	echoMain.Use(middleware.Recover())

	// Setup metrics endpoint at another port
	promMid.SetMetricsPath(echoProm)

	// Routes
	echoMain.GET("/", hello)
	echoMain.GET("/livez", livez)
	echoMain.GET("/readyz", readyz)
	echoMain.GET("/panic", serverError)

	// Spawn a separate goroutine for the metrics endpoint
	go func() { echoProm.Logger.Fatal(echoProm.Start(":9360")) }()
	// Start our main server
	echoMain.Logger.Fatal(echoMain.Start(":8080"))
}

// Handler for content
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func serverError(c echo.Context) error {
	panic("This is a simulated server error")
}

// TODO
// Consider if/when these routes should be logged and metric'ed

// Liveness probe
func livez(c echo.Context) error {
	return c.String(http.StatusOK, "Alive!")
}

// Readiness probe
func readyz(c echo.Context) error {
	return c.String(http.StatusOK, "Ready!")
}
