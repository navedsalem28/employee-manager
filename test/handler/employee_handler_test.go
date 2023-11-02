package handler

import (
	"employee-manager/internal/handler"
	"employee-manager/internal/model"
	"employee-manager/internal/repository"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func TestCRUDEmployeeWithMockRepository(t *testing.T) {
	// Create a new Echo instance
	AlreadyCreatedEmployee := &model.Employee{}
	e := echo.New()

	var er repository.EmployeeRepository

	er = repository.NewEmployeeMockRepository()

	// Create the Employee repository and handler
	eh := handler.NewEmployeeHandler(er)

	// Test creating an employee
	t.Run("CreateEmployee", func(t *testing.T) {
		createEmployeeRequest := `{"Name": "John Doe", "Position": "Engineer", "Department": "IT", "Email": "john@example.com"}`
		req := httptest.NewRequest(http.MethodPost, "/employees", strings.NewReader(createEmployeeRequest))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		err := eh.CreateEmployee(c)
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), AlreadyCreatedEmployee))
		retrievedEmployee, err := er.GetEmployeeByID(strconv.Itoa(int(AlreadyCreatedEmployee.ID)))
		assert.NoError(t, err)
		assert.Equal(t, AlreadyCreatedEmployee.Name, retrievedEmployee.Name)
	})

	// Test reading an employee
	t.Run("GetEmployee", func(t *testing.T) {

		// Send a GET request to retrieve the employee
		getRequest := httptest.NewRequest(http.MethodGet, "/employees/"+strconv.Itoa(int(AlreadyCreatedEmployee.ID)), nil)
		getResponse := httptest.NewRecorder()
		getContext := e.NewContext(getRequest, getResponse)
		// Set a parameter in the Echo context to simulate the employee ID in the URL
		getContext.SetParamNames("id")
		getContext.SetParamValues(strconv.Itoa(int(AlreadyCreatedEmployee.ID)))

		// Call the GetEmployee handler with the Echo context
		_ = eh.GetEmployee(getContext)
		assert.Equal(t, http.StatusOK, getResponse.Code)
		retrievedEmployee := &model.Employee{}
		assert.NoError(t, json.Unmarshal(getResponse.Body.Bytes(), retrievedEmployee))
		assert.Equal(t, AlreadyCreatedEmployee.Name, retrievedEmployee.Name)
	})

	// Test updating an employee
	t.Run("UpdateEmployee", func(t *testing.T) {

		// Send a PUT request to update the employee
		updateEmployeeRequest := `{"Name": "New Name", "Position": "New Position", "Department": "New Department", "Email": "new@example.com"}`
		updateRequest := httptest.NewRequest(http.MethodPut, "/employees/"+strconv.Itoa(int(AlreadyCreatedEmployee.ID)), strings.NewReader(updateEmployeeRequest))
		updateRequest.Header.Set("Content-Type", "application/json")
		updateResponse := httptest.NewRecorder()
		updateContext := e.NewContext(updateRequest, updateResponse)
		// Set a parameter in the Echo context to simulate the employee ID in the URL
		updateContext.SetParamNames("id")
		updateContext.SetParamValues(strconv.Itoa(int(AlreadyCreatedEmployee.ID)))
		_ = eh.UpdateEmployee(updateContext)
		assert.Equal(t, http.StatusOK, updateResponse.Code)

		// Verify that the employee has been updated
		getRequest := httptest.NewRequest(http.MethodGet, "/employees/"+strconv.Itoa(int(AlreadyCreatedEmployee.ID)), nil)
		getResponse := httptest.NewRecorder()
		getContext := e.NewContext(getRequest, getResponse)
		// Set a parameter in the Echo context to simulate the employee ID in the URL
		getContext.SetParamNames("id")
		getContext.SetParamValues(strconv.Itoa(int(AlreadyCreatedEmployee.ID)))
		_ = eh.GetEmployee(getContext)
		assert.Equal(t, http.StatusOK, getResponse.Code)
		retrievedEmployee := &model.Employee{}
		assert.NoError(t, json.Unmarshal(getResponse.Body.Bytes(), retrievedEmployee))
		assert.Equal(t, "New Name", retrievedEmployee.Name)
		assert.Equal(t, "New Position", retrievedEmployee.Position)
	})

	// Test deleting an employee
	t.Run("DeleteEmployee", func(t *testing.T) {

		// Send a DELETE request to delete the employee
		deleteRequest := httptest.NewRequest(http.MethodDelete, "/employees/"+strconv.Itoa(int(AlreadyCreatedEmployee.ID)), nil)
		deleteResponse := httptest.NewRecorder()
		deleteContext := e.NewContext(deleteRequest, deleteResponse)
		// Set a parameter in the Echo context to simulate the employee ID in the URL
		deleteContext.SetParamNames("id")
		deleteContext.SetParamValues(strconv.Itoa(int(AlreadyCreatedEmployee.ID)))
		_ = eh.DeleteEmployee(deleteContext)
		assert.Equal(t, http.StatusNoContent, deleteResponse.Code)

		// Verify that the employee has been deleted
		getRequest := httptest.NewRequest(http.MethodGet, "/employees/"+strconv.Itoa(int(AlreadyCreatedEmployee.ID)), nil)
		getResponse := httptest.NewRecorder()
		getContext := e.NewContext(getRequest, getResponse)
		_ = eh.GetEmployee(getContext)
		assert.Equal(t, http.StatusNotFound, getResponse.Code)
	})
}
