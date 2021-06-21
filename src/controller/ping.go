package controller

import (
	"net/http"

	"gin-todo-list/src/model"
	"gin-todo-list/src/util"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	util.ApiOnSuccess(c, &model.ApiSuccess{
		StatusCode: http.StatusOK,
		Data:       "Server is working",
	})
}
