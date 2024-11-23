package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type Employee struct {
	ID uuid.UUID `json:"id"`

	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	DeletedAt time.Time `json:"deleted_at,omitempty"`

	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Code          string `json:"code"`
	Address       string `json:"address"`
	SalaryType    string `json:"salary_type"`
	HourlyRate    string `json:"hourly_rate"`
	MonthlySalary string `json:"monthly_salary"`
	PaymentMethod string `json:"payment_method"`

	Active bool `json:"active"`
}
