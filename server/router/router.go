package router

import (
	"github.com/gin-gonic/gin"
	"puzzle-hackathon-backend/repositories"
	"puzzle-hackathon-backend/server/controllers"
	"puzzle-hackathon-backend/server/middlewares"
	"puzzle-hackathon-backend/services/employee"
	"puzzle-hackathon-backend/services/user"
	"time"
)

func NewRouter() *gin.Engine {
	router := gin.New()

	// Initialize the rate limiter
	rateLimiter := middlewares.NewRateLimiter(1, time.Second*2, 1)

	// Apply middleware for CORS, request ID generation, and rate limiting
	router.Use(
		middlewares.CORSMiddleware(),
		middlewares.RequestIDMiddleware(),
		middlewares.RateLimiterMiddleware(rateLimiter),
	)

	userRepository := repositories.NewUserRepository()
	userService := user.NewUserService(userRepository)
	userController := controllers.NewUsersController(userService)

	employeeRepository := repositories.NewEmployeeRepository()
	employeeService := employee.NewEmployeeService(employeeRepository)
	employeeController := controllers.NewEmployeesController(employeeService)

	api := router.Group("api")
	{
		v1 := api.Group("v1")
		{
			userRouter := v1.Group("/user")
			{
				userRouter.POST("", userController.CreateUser)
				userRouter.GET("", userController.GetUsers)
				userRouter.GET("/:id", userController.GetUser)
				userRouter.PUT("/:id", userController.UpdateUser)
				userRouter.DELETE("/:id", userController.DeleteUser)
			}
			employeeRouter := v1.Group("/employee")
			{
				employeeRouter.POST("", employeeController.CreateEmployee)
				employeeRouter.GET("", employeeController.GetEmployees)
				employeeRouter.GET("/:id", employeeController.GetEmployee)
				employeeRouter.PUT("/:id", employeeController.UpdateEmployee)
				employeeRouter.DELETE("/:id", employeeController.DeleteEmployee)
			}
		}
	}

	return router
}
