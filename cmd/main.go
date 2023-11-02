package main

import (
	"employee-manager/configs"
	"employee-manager/internal/handler"
	"employee-manager/internal/repository"
	"employee-manager/logger"
	"employee-manager/pkg/database"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func main() {
	// Load configuration
	configs.InConfig()

	// Load Logger
	logger.InitLogger(viper.GetString("logger.level"), viper.GetString("logger.console"))

	// Initialize the Environment with the configuration
	env := viper.GetString("env.state")

	// Initialize the repository based on the environment
	var er repository.EmployeeRepository
	// Initialize the database with the configuration
	if env == "prod" {
		database.InitDBFromViper()
		er = repository.NewEmployeeMySQLRepository(database.DB)
	} else {
		er = repository.NewEmployeeMockRepository()
	}

	// Create a new Echo instance
	e := echo.New()
	// Create the Employee handler
	eh := handler.NewEmployeeHandler(er)

	// Define routes for CRUD operations
	e.POST("/employees", eh.CreateEmployee)
	e.GET("/employees/:id", eh.GetEmployee)
	e.GET("/employees", eh.GetEmployees)
	e.PUT("/employees/:id", eh.UpdateEmployee)
	e.DELETE("/employees/:id", eh.DeleteEmployee)

	// Start the Echo server
	logger.Info("==============================")
	logger.Info("==============================")
	logger.Info("Ready To Serve")
	logger.Info("==============================")
	logger.Info("==============================")
	e.Start(":8080")
}
