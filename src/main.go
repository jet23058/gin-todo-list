package main

import (
	"fmt"
	"log"

	"gin-todo-list/src/config"
	"gin-todo-list/src/controller"
	"gin-todo-list/src/middleware"
	"gin-todo-list/src/route"

	"github.com/gin-gonic/gin"
)

func useGlobalMiddlewares(r *gin.Engine) {
	middleware.UseCustomLogger(r)
	middleware.UseLogger(r)
	middleware.UseErrorHandler(r)
}

func main() {
	config.InitOs()

	r := gin.New()

	useGlobalMiddlewares(r)

	r.GET("/", controller.Ping)

	route.UseAuth(r)
	route.UseUser(r)
	route.UseTodos(r)

	if err := r.Run(fmt.Sprintf(":%s", config.GetEnv().SERVER_PORT)); err != nil {
		log.Fatal("Unable to start:", err)
	}
}
