package controller_auth

import (
	"net/http"

	"gin-todo-list/src/model"
	"gin-todo-list/src/util"

	"github.com/gin-gonic/gin"
)

func SignOut(c *gin.Context) {
	util.DeleteAuth(c)
	util.ApiOnSuccess(c, &model.ApiSuccess{
		StatusCode: http.StatusOK,
		Data:       "Sign out successfully",
	})
}
