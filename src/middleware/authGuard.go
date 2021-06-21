package middleware

import (
	"net/http"

	"gin-todo-list/src/config"
	"gin-todo-list/src/model"
	"gin-todo-list/src/util"

	"github.com/gin-gonic/gin"
)

func authGuard() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := util.GetUserIdByToken(c)
		if err != nil {
			util.DeleteAuth(c)
			util.ApiOnError(&model.ApiError{
				StatusCode: http.StatusBadRequest,
				ErrorType:  model.ErrorType(err.Error()),
				Error:      err,
			})
		}

		user := model.User{
			ID: id,
		}

		if err := config.GetDB().First(&user, "id = ?", id).Error; err != nil {
			util.DeleteAuth(c)
			util.ApiOnError(&model.ApiError{
				StatusCode: http.StatusServiceUnavailable,
				ErrorType:  model.ERROR_SIGN_IN_FAILED,
				Error:      err,
			})
		}

		c.Set("userId", id)

		c.Next()
	}
}

func UseAuthGuard(r *gin.RouterGroup) gin.IRoutes {
	return r.Use(authGuard())
}
