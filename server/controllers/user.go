package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"io"
	"net/http"
	"puzzle-hackathon-backend/models"
	"puzzle-hackathon-backend/services/user"
)

type UserController interface {
	CreateUser(c *gin.Context)
	GetUser(c *gin.Context)
	GetUsers(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}

type userController struct {
	userService *user.UserService
}

func NewUsersController(userService *user.UserService) UserController {
	return &userController{
		userService: userService,
	}
}

func (u *userController) CreateUser(c *gin.Context) {
	payload := models.User{}
	requestBody, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := json.Unmarshal(requestBody, &payload); err != nil {
		panic(err)
	}

	user, err := u.userService.CreateUser(c, payload)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
	}
	c.JSON(http.StatusOK, user)
}

func (u *userController) GetUser(c *gin.Context) {
	userId := c.Param("id")
	if userId == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "user id must be provided"})
		return
	}

	payload := models.User{}
	userUid, err := uuid.FromString(userId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "user uid must be provided"})
		return
	}
	payload.ID = userUid

	patients, err := u.userService.GetUser(c, payload)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
	}
	c.JSON(http.StatusOK, patients)
}

func (u *userController) GetUsers(c *gin.Context) {
	users, err := u.userService.GetUsers(c)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
	}
	c.JSON(http.StatusOK, users)
}

func (u *userController) UpdateUser(c *gin.Context) {
	userId := c.Param("id")
	if userId == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "user id must be provided"})
		return
	}

	payload := models.User{}
	userUid, err := uuid.FromString(userId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "user uid must be provided"})
		return
	}
	payload.ID = userUid
	requestBody, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := json.Unmarshal(requestBody, &payload); err != nil {
		panic(err)
	}

	patients, err := u.userService.UpdateUser(c, payload)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
	}
	c.JSON(http.StatusOK, patients)
}

func (u *userController) DeleteUser(c *gin.Context) {
	userId := c.Param("id")
	if userId == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "user id must be provided"})
		return
	}

	payload := models.User{}
	userUid, err := uuid.FromString(userId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "user id must be provided"})
		return
	}
	payload.ID = userUid

	patients, err := u.userService.DeleteUser(c, payload)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
	}
	c.JSON(http.StatusOK, patients)
}
