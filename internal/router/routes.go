package router

import (
	docs "github.com/UxieGu1/gopportunities-api/docs"
	"github.com/UxieGu1/gopportunities-api/internal/handler/application"
	"github.com/UxieGu1/gopportunities-api/internal/handler/candidate"
	"github.com/UxieGu1/gopportunities-api/internal/handler/company"
	"github.com/UxieGu1/gopportunities-api/internal/handler/opening"
	"github.com/UxieGu1/gopportunities-api/internal/handler/user"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {

	opening.InitializeHandler()
	company.InitializeHandler()
	user.InitializeHandler()
	candidate.InitializeHandler()
	application.InitializeHandler()

	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath

	v1 := router.Group(basePath)
	{
		// User
		v1.GET("/users", user.ListUsersHandler)
		v1.GET("/users/:id", user.ShowUserHandler)
		v1.POST("/users", user.CreateUserHandler)
		v1.PUT("/users/:id", user.UpdateUserHandler)
		v1.DELETE("/users/:id", user.DeleteUserHandler)

		// Opening
		v1.GET("/openings", opening.ListOpeningsHandler)
		v1.GET("/openings/:id", opening.ShowOpeningHandler)
		v1.POST("/openings", opening.CreateOpeningHandler)
		v1.PUT("/openings/:id", opening.UpdateOpeningHandler)
		v1.DELETE("/openings/:id", opening.DeleteOpeningHandler)
        
		// Company
		v1.GET("/companies", company.ListCompaniesHandler)
		v1.GET("/companies/:id", company.ShowCompanyHandler)
		v1.POST("/companies", company.CreateCompanyHandler)
		v1.PUT("/companies/:id", company.UpdateCompanyHandler)
		v1.DELETE("/companies/:id", company.DeleteCompanyHandler)

		// Candidate
		v1.GET("/candidates", candidate.ListCandidatesHandler)
		v1.GET("/candidates/:id", candidate.ShowCandidateHandler)
		v1.POST("/candidates", candidate.CreateCandidateHandler)
		v1.PUT("/candidates/:id", candidate.UpdateCandidateHandler)
		v1.DELETE("/candidates/:id", candidate.DeleteCandidateHandler)

		// Application
		v1.GET("/applications", application.ListApplicationsHandler)
		v1.GET("/applications/:id", application.ShowApplicationHandler)
		v1.POST("/applications", application.CreateApplicationHandler)
		v1.PUT("/applications/:id", application.UpdateApplicationHandler)
		v1.DELETE("/applications/:id", application.DeleteApplicationHandler)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}