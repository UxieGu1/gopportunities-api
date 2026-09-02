package user

import (
	"net/http"

	"github.com/UxieGu1/gopportunities-api/internal/schemas"
	"github.com/gin-gonic/gin"
)

// @Summary Registro de Novo Usuário
// @Description Cria um novo usuário (registro de conta)
// @Tags Users
// @Accept json
// @Produce json
// @Param request body RegisterUserRequest true "Dados de registro"
// @Success 201 {object} schemas.UserResponse
// @Failure 400 {object} ErrorResponse
// @Failure 409 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /register [post]
func RegisterUserHandler(context *gin.Context) {
	var request RegisterUserRequest

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
		Email: request.Email,
		Role:  request.Role,
	}

	if user.Role == "" {
		user.Role = "CANDIDATE"
	}

	if err := userService.Create(&user, request.Password); err != nil {
		logger.Errorf("error registering user: %v", err.Error())
		sendError(context, http.StatusInternalServerError, "erro ao registrar usuário no banco de dados")
		return
	}

	response := schemas.UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	sendSuccess(context, "register-user", response)
}
