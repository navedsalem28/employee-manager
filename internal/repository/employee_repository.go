package repository

import (
	"employee-manager/internal/model"
	"errors"
	"gorm.io/gorm"
	"strconv"
)

type EmployeeMockRepository struct {
	employees map[uint]*model.Employee
	autoID    uint
}

func NewEmployeeMockRepository() *EmployeeMockRepository {
	return &EmployeeMockRepository{
		employees: make(map[uint]*model.Employee),
		autoID:    1,
	}
}

func NewEmployeeRepository(db *gorm.DB) *EmployeeMySQLRepository {
	return &EmployeeMySQLRepository{db}
}

type EmployeeMySQLRepository struct {
	db *gorm.DB
}

func NewEmployeeMySQLRepository(db *gorm.DB) *EmployeeMySQLRepository {
	return &EmployeeMySQLRepository{db}
}

type EmployeeRepository interface {
	CreateEmployee(employee *model.Employee) error
	GetEmployeeByID(id string) (*model.Employee, error)
	UpdateEmployee(employee *model.Employee) error
	DeleteEmployee(id string) error
	GetEmployees() ([]model.Employee, error)
}

// CreateEmployee creates a new employee in the MySQL database.
func (r *EmployeeMySQLRepository) CreateEmployee(employee *model.Employee) error {
	if err := r.db.Create(employee).Error; err != nil {
		return err
	}
	return nil
}

// GetEmployeeByID retrieves an employee by ID from the MySQL database.
func (r *EmployeeMySQLRepository) GetEmployeeByID(id string) (*model.Employee, error) {
	employee := new(model.Employee)
	if err := r.db.First(employee, id).Error; err != nil {
		return nil, err
	}
	return employee, nil
}

// UpdateEmployee updates an employee in the MySQL database.
func (r *EmployeeMySQLRepository) UpdateEmployee(employee *model.Employee) error {

	if err := r.db.Save(employee).Error; err != nil {
		return err
	}
	return nil

}

// DeleteEmployee deletes an employee by ID from the MySQL database.
func (r *EmployeeMySQLRepository) DeleteEmployee(id string) error {
	if err := r.db.Where("id = ?", id).Delete(&model.Employee{}).Error; err != nil {
		return err
	}
	return nil
}

// GetEmployees retrieves all employees from the MySQL database.
func (r *EmployeeMySQLRepository) GetEmployees() ([]model.Employee, error) {
	var employees []model.Employee
	if err := r.db.Find(&employees).Error; err != nil {
		return nil, err
	}
	return employees, nil
}

// CreateEmployee creates a new employee in the mock repository.
func (r *EmployeeMockRepository) CreateEmployee(employee *model.Employee) error {
	if employee == nil {
		return errors.New("employee is nil")
	}

	// Generate a unique ID for the new employee
	employee.ID = r.autoID
	r.employees[employee.ID] = employee

	// Increment the autoID for the next employee
	r.autoID++

	return nil
}

// GetEmployeeByID retrieves an employee by ID from the mock repository.
func (r *EmployeeMockRepository) GetEmployeeByID(id string) (*model.Employee, error) {
	// Convert the ID to a uint, assuming it's a string representation of a uint.
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return nil, errors.New("invalid ID format")
	}

	// Look up the employee by ID in the map.
	employee, found := r.employees[uint(idUint)]
	if !found {
		return nil, errors.New("employee not found")
	}

	return employee, nil
}

// UpdateEmployee updates an employee in the mock repository.
func (r *EmployeeMockRepository) UpdateEmployee(employee *model.Employee) error {
	if employee == nil {
		return errors.New("employee is nil")
	}

	// Check if the employee exists in the mock repository.
	_, found := r.employees[employee.ID]
	if !found {
		return errors.New("employee not found")
	}

	// Update the employee's information in the mock repository.
	r.employees[employee.ID] = employee

	return nil
}

// DeleteEmployee deletes an employee by ID from the mock repository.
func (r *EmployeeMockRepository) DeleteEmployee(id string) error {
	// Convert the ID to a uint, assuming it's a string representation of a uint.
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return errors.New("invalid ID format")
	}

	// Check if the employee exists in the mock repository.
	_, found := r.employees[uint(idUint)]
	if !found {
		return errors.New("employee not found")
	}

	// Delete the employee from the mock repository.
	delete(r.employees, uint(idUint))

	return nil
}

// GetEmployees retrieves all employees from the mock repository.
func (r *EmployeeMockRepository) GetEmployees() ([]model.Employee, error) {
	employees := make([]model.Employee, 0)

	for _, employee := range r.employees {
		employees = append(employees, *employee)
	}

	return employees, nil
}
