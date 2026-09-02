package user

import (
	"net/http"

	"github.com/UxieGu1/gopportunities-api/internal/auth"
	"github.com/UxieGu1/gopportunities-api/internal/schemas"
	"github.com/gin-gonic/gin"
)

// @Summary Login do Usuário
// @Description Realiza a autenticação e retorna os dados do usuário
// @Tags Users
// @Accept json
// @Produce json
// @Param request body LoginUserRequest true "Credenciais de login"
// @Success 200 {object} ShowUserResponse
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Router /login [post]
func LoginUserHandler(ctx *gin.Context) {
	var request LoginUserRequest

	if err := ctx.BindJSON(&request); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	user, err := userService.Login(request.Email, request.Password)
	if err != nil {
		logger.Errorf("authentication error: %v", err.Error())
		sendError(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	token, err := auth.GenerateToken(user.ID, user.Role)
	if err != nil {
		logger.Errorf("error generating token: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "erro interno ao gerar credencial")
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "login realizado com sucesso",
		"data": schemas.UserResponse{
			ID:        user.ID,
			Email:     user.Email,
			Role:      user.Role,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
		"token": token, 
	})

	response := schemas.UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	sendSuccess(ctx, "login-user", response)
}