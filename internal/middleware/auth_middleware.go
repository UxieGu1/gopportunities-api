package middleware

import (
	"net/http"
	"strings"

	"github.com/UxieGu1/gopportunities-api/internal/auth"
	"github.com/UxieGu1/gopportunities-api/internal/config"
	"github.com/gin-gonic/gin"
)

var logger *config.Logger

func init() {
	logger = config.GetLogger("middleware")
}

const (
	UserIDKey   = "userID"
	UserRoleKey = "userRole"
)

func AuthRequired() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			logger.Warnf("token não fornecido - IP: %s", ctx.ClientIP())
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "token não fornecido",
				"code":  "MISSING_TOKEN",
			})
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			logger.Warnf("formato de token inválido - IP: %s", ctx.ClientIP())
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "formato de token inválido (use: Bearer <token>)",
				"code":  "INVALID_TOKEN_FORMAT",
			})
			return
		}

		claims, err := auth.ValidateToken(parts[1])
		if err != nil {
			logger.Warnf("token inválido ou expirado - IP: %s, erro: %v", ctx.ClientIP(), err)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "token inválido ou expirado",
				"code":  "INVALID_TOKEN",
			})
			return
		}

		ctx.Set(UserIDKey, claims.UserID)
		ctx.Set(UserRoleKey, claims.Role)
		logger.Debugf("autenticação bem-sucedida - userID: %d, role: %s", claims.UserID, claims.Role)
		ctx.Next()
	}
}

func RoleRequired(requiredRoles ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		role, exists := ctx.Get(UserRoleKey)
		if !exists {
			logger.Warnf("role não encontrado no contexto - IP: %s", ctx.ClientIP())
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "usuário não autenticado",
				"code":  "NOT_AUTHENTICATED",
			})
			return
		}

		userRole, ok := role.(string)
		if !ok {
			logger.Errorf("role inválido no contexto")
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "erro interno",
				"code":  "INTERNAL_ERROR",
			})
			return
		}

		hasRole := false
		for _, allowed := range requiredRoles {
			if userRole == allowed {
				hasRole = true
				break
			}
		}

		if !hasRole {
			userID, _ := ctx.Get(UserIDKey)
			logger.Warnf("acesso negado - userID: %v, role: %s, roles esperados: %v", userID, userRole, requiredRoles)
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "acesso negado: permissão insuficiente",
				"code":  "FORBIDDEN",
			})
			return
		}

		ctx.Next()
	}
}

func GetUserID(ctx *gin.Context) uint {
	id, exists := ctx.Get(UserIDKey)
	if !exists {
		return 0
	}
	userID, ok := id.(uint)
	if !ok {
		return 0
	}
	return userID
}

func GetUserRole(ctx *gin.Context) string {
	role, exists := ctx.Get(UserRoleKey)
	if !exists {
		return ""
	}
	userRole, ok := role.(string)
	if !ok {
		return ""
	}
	return userRole
}