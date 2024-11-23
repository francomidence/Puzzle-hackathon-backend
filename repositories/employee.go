package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"puzzle-hackathon-backend/models"
	"time"
)

type EmployeeRepository interface {
	CreateEmployee(ctx context.Context, employee *models.Employee) (*models.Employee, error)
	GetEmployees(ctx context.Context) (*[]models.Employee, error)
	GetEmployee(ctx context.Context, employee *models.Employee) (*models.Employee, error)
	UpdateEmployee(ctx context.Context, employee *models.Employee) (*models.Employee, error)
	DeleteEmployee(ctx context.Context, employee *models.Employee) (*models.Employee, error)
}

type employeeRepository struct {
}

func NewEmployeeRepository() EmployeeRepository {
	return &employeeRepository{}
}

func (r *employeeRepository) CreateEmployee(ctx context.Context, employee *models.Employee) (*models.Employee, error) {
	supabase := InitializeSupabaseClient()

	employee.ID = uuid.NewV4()

	employee.CreatedAt = time.Now()
	employee.Active = true

	var results []map[string]interface{}
	err := supabase.DB.From("Employees").Insert(employee).Execute(&results)
	if err != nil {
		fmt.Println("The error is: ", err.Error())
		panic(err) // Consider handling this error more gracefully
	}

	if len(results) > 0 {
		// Convert the first element of 'results' to JSON
		jsonData, err := json.Marshal(results[0])
		if err != nil {
			panic(err) // Consider handling this error more gracefully
		}

		var newEmployee models.Employee // Use a single instance of 'models.Employee'
		err = json.Unmarshal(jsonData, &newEmployee)
		if err != nil {
			panic(err)
		}

		return &newEmployee, nil
	}

	return nil, fmt.Errorf("no employee data returned after insert operation")
}

func (r *employeeRepository) GetEmployees(ctx context.Context) (*[]models.Employee, error) {
	supabase := InitializeSupabaseClient()

	var results []map[string]interface{} // Change to a slice of maps
	err := supabase.DB.From("Employees").Select("*").Eq("active", "true").Execute(&results)
	if err != nil {
		fmt.Println("The error is ", err)
		fmt.Println(err.Error())
		panic(err)
	}

	// Convert to JSON and then to Employee struct
	jsonData, err := json.Marshal(results)
	if err != nil {
		panic(err)
	}

	var employees *[]models.Employee
	err = json.Unmarshal(jsonData, &employees)
	if err != nil {
		panic(err)
	}

	return employees, nil
}

func (r *employeeRepository) GetEmployee(ctx context.Context, employee *models.Employee) (*models.Employee, error) {
	supabase := InitializeSupabaseClient()

	var results []map[string]interface{} // Change to a slice of maps
	err := supabase.DB.From("Employees").Select("*").Eq("id", employee.ID.String()).Execute(&results)
	if err != nil {
		panic(err)
	}

	if len(results) > 0 {
		jsonData, err := json.Marshal(results[0])
		if err != nil {
			panic(err) // Consider handling this error more gracefully
		}

		var updatedEmployee models.Employee
		err = json.Unmarshal(jsonData, &updatedEmployee)
		if err != nil {
			panic(err) // Consider handling this error more gracefully
		}

		return &updatedEmployee, nil
	}

	return nil, fmt.Errorf("no employee")
}

func (r *employeeRepository) UpdateEmployee(ctx context.Context, employee *models.Employee) (*models.Employee, error) {
	supabase := InitializeSupabaseClient()

	// Assign current time to UpdatedAt
	employee.UpdatedAt = time.Now()

	// Create a map of fields to update, excluding 'CreatedAt' and any other fields you don't want to update
	updateFields := map[string]interface{}{
		"updated_at": employee.UpdatedAt,

		"first_name":     employee.FirstName,
		"last_name":      employee.LastName,
		"code":           employee.Code,
		"address":        employee.Address,
		"salary_type":    employee.SalaryType,
		"hourly_rate":    employee.HourlyRate,
		"monthly_salary": employee.MonthlySalary,
		"payment_method": employee.PaymentMethod,

		"active": employee.Active,
	}

	var results []map[string]interface{}
	err := supabase.DB.From("Employees").Update(updateFields).Eq("id", employee.ID.String()).Execute(&results)
	if err != nil {
		panic(err) // Consider handling this error more gracefully
	}

	if len(results) > 0 {
		jsonData, err := json.Marshal(results[0])
		if err != nil {
			panic(err) // Consider handling this error more gracefully
		}

		var updatedEmployee models.Employee
		err = json.Unmarshal(jsonData, &updatedEmployee)
		if err != nil {
			panic(err) // Consider handling this error more gracefully
		}

		return &updatedEmployee, nil
	}

	return nil, fmt.Errorf("no employee data returned after update operation")
}

func (r *employeeRepository) DeleteEmployee(ctx context.Context, employee *models.Employee) (*models.Employee, error) {
	supabase := InitializeSupabaseClient()

	// Assign current time to DeletedAt
	employee.DeletedAt = time.Now()

	updateFields := map[string]interface{}{
		"deleted_at": employee.DeletedAt,
		"active":     false,
	}

	var results []map[string]interface{}
	err := supabase.DB.From("Employees").Update(updateFields).Eq("id", employee.ID.String()).Execute(&results)
	if err != nil {
		panic(err) // Consider handling this error more gracefully
	}

	if len(results) > 0 {
		jsonData, err := json.Marshal(results[0])
		if err != nil {
			panic(err) // Consider handling this error more gracefully
		}

		var updatedEmployee models.Employee
		err = json.Unmarshal(jsonData, &updatedEmployee)
		if err != nil {
			panic(err) // Consider handling this error more gracefully
		}

		return &updatedEmployee, nil
	}

	return nil, fmt.Errorf("no employee data returned after update operation")
}
