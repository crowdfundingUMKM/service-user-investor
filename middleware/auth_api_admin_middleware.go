package middleware

import (
	"net/http"
	api_admin "service-user-investor/api/admin"
	"service-user-investor/auth"
	"service-user-investor/core"
	"service-user-investor/helper"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthApiAdminMiddleware(authService auth.Service, userService core.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized0", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		// token, err := authService.ValidateToken(tokenString)
		// if err != nil {
		// 	response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
		// 	c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		// 	return
		// }

		// claim, ok := token.Claims.(jwt.MapClaims)

		// if !ok || !token.Valid {
		// 	response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
		// 	c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		// 	return
		// }

		// adminID := claim["admin_id"].(string)

		if tokenString == "" {
			response := helper.APIResponse("Unauthorized1", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		adminID, err := api_admin.VerifyTokenAdmin(tokenString)

		if err != nil { //wrong token
			response := helper.APIResponse("Unauthorized2", http.StatusUnauthorized, "error", err.Error())
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUserAdmin", api_admin.AdminId{UnixAdmin: adminID})
		c.Next()
	}
}
