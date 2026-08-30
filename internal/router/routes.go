package router

import (
	docs "github.com/UxieGu1/gopportunities-api/docs"
	"github.com/UxieGu1/gopportunities-api/internal/handler/company"
	"github.com/UxieGu1/gopportunities-api/internal/handler/opening"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	opening.InitializeHandler()
	company.InitializeHandler()

	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath

	v1 := router.Group(basePath)
	{

		//User

		//Opening
		v1.GET("/openings", opening.ListOpeningsHandler)
		v1.GET("/openings/:id", opening.ShowOpeningHandler)
		v1.POST("/openings", opening.CreateOpeningHandler)
		v1.PUT("/openings/:id", opening.UpdateOpeningHandler)
		v1.DELETE("/openings/:id", opening.DeleteOpeningHandler)
        
		//Company
		v1.GET("/companies", company.ListCompaniesHandler)
		v1.GET("/companies/:id", company.ShowCompanyHandler)
		v1.POST("/companies", company.CreateCompanyHandler)
		v1.PUT("/companies/:id", company.UpdateCompanyHandler)
		v1.DELETE("/companies/:id", company.DeleteCompanyHandler)

		//Candidate

		//Application
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}