package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"io"
	"net/http"
	"puzzle-hackathon-backend/models"
	"puzzle-hackathon-backend/services/employee"
)

type EmployeeController interface {
	CreateEmployee(c *gin.Context)
	GetEmployee(c *gin.Context)
	GetEmployees(c *gin.Context)
	UpdateEmployee(c *gin.Context)
	DeleteEmployee(c *gin.Context)
}

type employeeController struct {
	employeeService *employee.EmployeeService
}

func NewEmployeesController(employeeService *employee.EmployeeService) EmployeeController {
	return &employeeController{
		employeeService: employeeService,
	}
}

func (u *employeeController) CreateEmployee(c *gin.Context) {
	payload := models.Employee{}
	requestBody, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := json.Unmarshal(requestBody, &payload); err != nil {
		panic(err)
	}

	employee, err := u.employeeService.CreateEmployee(c, payload)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
	}
	c.JSON(http.StatusOK, employee)
}

func (u *employeeController) GetEmployee(c *gin.Context) {
	employeeId := c.Param("id")
	if employeeId == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "employee id must be provided"})
		return
	}

	payload := models.Employee{}
	employeeUid, err := uuid.FromString(employeeId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "employee uid must be provided"})
		return
	}
	payload.ID = employeeUid

	patients, err := u.employeeService.GetEmployee(c, payload)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
	}
	c.JSON(http.StatusOK, patients)
}

func (u *employeeController) GetEmployees(c *gin.Context) {
	employees, err := u.employeeService.GetEmployees(c)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
	}
	c.JSON(http.StatusOK, employees)
}

func (u *employeeController) UpdateEmployee(c *gin.Context) {
	employeeId := c.Param("id")
	if employeeId == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "employee id must be provided"})
		return
	}

	payload := models.Employee{}
	employeeUid, err := uuid.FromString(employeeId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "employee uid must be provided"})
		return
	}
	payload.ID = employeeUid
	requestBody, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := json.Unmarshal(requestBody, &payload); err != nil {
		panic(err)
	}

	patients, err := u.employeeService.UpdateEmployee(c, payload)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
	}
	c.JSON(http.StatusOK, patients)
}

func (u *employeeController) DeleteEmployee(c *gin.Context) {
	employeeId := c.Param("id")
	if employeeId == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "employee id must be provided"})
		return
	}

	payload := models.Employee{}
	employeeUid, err := uuid.FromString(employeeId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "employee id must be provided"})
		return
	}
	payload.ID = employeeUid

	patients, err := u.employeeService.DeleteEmployee(c, payload)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
	}
	c.JSON(http.StatusOK, patients)
}
