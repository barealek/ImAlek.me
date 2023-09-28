package main

import (
	"fmt"
	"os"

	"github.com/barealek/hjemmeside/internal/controllers/health"
	"github.com/barealek/hjemmeside/internal/controllers/root"
	"github.com/barealek/hjemmeside/shutdown"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"

	gojson "github.com/goccy/go-json"
)

func main() {
	// setup exit code for graceful shutdown
	var exitCode int
	defer func() {
		os.Exit(exitCode)
	}()

	// run the server
	cleanup, err := run()

	// run the cleanup after the server is terminated
	defer cleanup()

	if err != nil {
		fmt.Printf("error: %v", err)
		exitCode = 1
		return
	}

	// ensure the server is shutdown gracefully & app runs
	shutdown.Gracefully()
}

func run() (func(), error) {
	app, err := buildServer()
	if err != nil {
		return nil, err
	}

	// start the server
	go func() {
		app.Listen("0.0.0.0:3000")
	}()

	// return a function to close the server and database
	return func() {
		app.Shutdown()
	}, nil
}

func buildServer() (*fiber.App, error) {

	// create the fiber app
	app := fiber.New(
		fiber.Config{
			JSONEncoder: gojson.Marshal,
			JSONDecoder: gojson.Unmarshal,
		})

	// add middleware
	app.Use(cors.New())
	app.Use(recover.New())

	// Register controllers
	root.InitController(app)
	health.InitController(app)

	return app, nil
}
