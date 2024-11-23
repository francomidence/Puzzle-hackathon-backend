package employee

import (
	"context"
	"puzzle-hackathon-backend/models"
	"puzzle-hackathon-backend/repositories"
)

type EmployeeService struct {
	EmployeeRepository repositories.EmployeeRepository
}

func NewEmployeeService(employeeRepository repositories.EmployeeRepository) *EmployeeService {
	return &EmployeeService{
		EmployeeRepository: employeeRepository,
	}
}

func (s *EmployeeService) CreateEmployee(ctx context.Context, payload models.Employee) (*models.Employee, error) {
	employee, err := s.EmployeeRepository.CreateEmployee(ctx, &payload)
	if err != nil {
		panic(err)
	}
	return employee, nil
}

func (s *EmployeeService) GetEmployees(ctx context.Context) (*[]models.Employee, error) {
	employees, err := s.EmployeeRepository.GetEmployees(ctx)
	if err != nil {
		return nil, err
	}
	return employees, nil
}

func (s *EmployeeService) GetEmployee(ctx context.Context, payload models.Employee) (*models.Employee, error) {
	employee, err := s.EmployeeRepository.GetEmployee(ctx, &payload)
	if err != nil {
		return nil, err
	}
	return employee, nil
}

func (s *EmployeeService) UpdateEmployee(ctx context.Context, payload models.Employee) (*models.Employee, error) {
	patient, err := s.EmployeeRepository.UpdateEmployee(ctx, &payload)
	if err != nil {
		panic(err)
	}
	return patient, nil
}

func (s *EmployeeService) DeleteEmployee(ctx context.Context, payload models.Employee) (*models.Employee, error) {
	patient, err := s.EmployeeRepository.DeleteEmployee(ctx, &payload)
	if err != nil {
		panic(err)
	}
	return patient, nil
}
