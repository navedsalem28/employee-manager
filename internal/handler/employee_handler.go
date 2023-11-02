package handler

import (
	"employee-manager/internal/model"
	"employee-manager/internal/repository"
	"github.com/labstack/echo/v4"
	"net/http"
)

type EmployeeHandler struct {
	er repository.EmployeeRepository
}

func NewEmployeeHandler(er repository.EmployeeRepository) *EmployeeHandler {
	return &EmployeeHandler{er}
}

// Implement HTTP handlers for CRUD operations using the repository.

func (eh *EmployeeHandler) CreateEmployee(c echo.Context) error {
	employee := new(model.Employee)
	if err := c.Bind(employee); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}
	if err := eh.er.CreateEmployee(employee); err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to create employee")
	}
	return c.JSON(http.StatusCreated, employee)
}

func (eh *EmployeeHandler) GetEmployee(c echo.Context) error {
	id := c.Param("id")

	employee, err := eh.er.GetEmployeeByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Employee not found")
	}
	return c.JSON(http.StatusOK, employee)
}

func (eh *EmployeeHandler) GetEmployees(c echo.Context) error {
	employee, err := eh.er.GetEmployees()
	if err != nil || len(employee) == 0 {
		return c.JSON(http.StatusNotFound, "Employee not found")
	}
	return c.JSON(http.StatusOK, employee)
}

func (eh *EmployeeHandler) UpdateEmployee(c echo.Context) error {
	id := c.Param("id")
	employee, err := eh.er.GetEmployeeByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Employee not found")
	}

	if err := c.Bind(employee); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	if err := eh.er.UpdateEmployee(employee); err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to update employee")
	}
	return c.JSON(http.StatusOK, employee)
}

func (eh *EmployeeHandler) DeleteEmployee(c echo.Context) error {
	id := c.Param("id")
	if err := eh.er.DeleteEmployee(id); err != nil {
		return c.JSON(http.StatusNotFound, "Employee not found")
	}
	return c.JSON(http.StatusNoContent, nil)
}
