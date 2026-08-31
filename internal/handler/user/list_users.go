package user

import (
	"net/http"

	"github.com/UxieGu1/gopportunities-api/internal/schemas"
	"github.com/gin-gonic/gin"
)

func ListUsersHandler(context *gin.Context) {
	users, err := userService.GetAll()
	if err != nil {
		sendError(context, http.StatusInternalServerError, "error listing users")
		return
	}

	responses := make([]schemas.UserResponse, 0, len(users))
	for _, u := range users {
		responses = append(responses, schemas.UserResponse{
			ID:        u.ID,
			Email:     u.Email,
			Role:      u.Role,
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
		})
	}

	sendSuccess(context, "list-users", responses)
}