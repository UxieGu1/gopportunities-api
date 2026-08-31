package user

import (
	"net/http"

	"github.com/UxieGu1/gopportunities-api/internal/schemas"
	"github.com/gin-gonic/gin"
)

func CreateUserHandler(context *gin.Context) {
	var request CreateUserRequest

	if err := context.BindJSON(&request); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(context, http.StatusBadRequest, err.Error())
		return
	}

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(context, http.StatusBadRequest, err.Error())
		return
	}

	
	user := schemas.User{
		Email:        request.Email,
		PasswordHash: request.Password, 
		Role:         request.Role,
	}

	if user.Role == "" {
		user.Role = "CANDIDATE" 
	}

	if err := userService.Create(&user); err != nil {
		logger.Errorf("error creating user: %v", err.Error())
		sendError(context, http.StatusInternalServerError, "error creating user on database")
		return
	}

	response := schemas.UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	sendSuccess(context, "create-user", response)
}