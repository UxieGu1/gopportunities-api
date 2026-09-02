package user

import (
	"net/http"
	"strconv"

	"github.com/UxieGu1/gopportunities-api/internal/schemas"
	"github.com/gin-gonic/gin"
)

func UpdateUserHandler(context *gin.Context) {
	var request UpdateUserRequest

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

	idStr := context.Param("id")
	if idStr == "" {
		sendError(context, http.StatusBadRequest, errParamIsRequired("id", "pathParameter").Error())
		return
	}

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		sendError(context, http.StatusBadRequest, "invalid id format")
		return
	}

	user, err := userService.GetByID(uint(id))
	if err != nil {
		sendError(context, http.StatusNotFound, "user not found")
		return
	}

	if request.Email != "" {
		user.Email = request.Email
	}
	if request.Role != "" {
		user.Role = request.Role
	}

	if err := userService.Update(user, request.Password); err != nil {
		logger.Errorf("error updating user: %v", err.Error())
		sendError(context, http.StatusInternalServerError, "error updating user")
		return
	}

	response := schemas.UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	sendSuccess(context, "update-user", response)
}
