package main

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"os"
	"sync"
	"testing"

	"employee-manager/internal/handler"    // Import your actual package
	"employee-manager/internal/repository" // Import your actual package
)

var testServer *echo.Echo
var wg sync.WaitGroup

func setupTestServer() {
	testServer = echo.New()

	// Create a mock repository for testing
	mockRepo := repository.NewEmployeeMockRepository()
	handler.NewEmployeeHandler(mockRepo)

	// Define routes for CRUD operations
	testServer.POST("/employees", CreateEmployee)

	// Start the test server on a separate Goroutine
	go func() {
		err := testServer.Start(":8081")
		if err != nil {
			_ = fmt.Errorf("failed to start the test server: %v", err)
		}
	}()
}

func shutdownTestServer() {
	if testServer != nil {
		_ = testServer.Close()
	}
}

func TestMain(m *testing.M) {
	// Set up the test server
	setupTestServer()

	// Run the tests
	code := m.Run()

	// Shutdown the test server
	shutdownTestServer()

	// Exit with the test code
	os.Exit(code)
}

func CreateEmployee(c echo.Context) error {
	// Your implementation for creating an employee
	return c.JSON(http.StatusCreated, "Employee created")
}
