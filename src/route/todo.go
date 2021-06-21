package route

import (
	"gin-todo-list/src/config"
	controller_todos "gin-todo-list/src/controller/todos"
	"gin-todo-list/src/middleware"
	"gin-todo-list/src/util"

	"github.com/gin-gonic/gin"
)

func UseTodos(r *gin.Engine) {
	todos := r.Group("/todos")
	db := config.GetDB()
	middleware.UseAuthGuard(todos)
	{
		todos.GET("", controller_todos.GetList(controller_todos.GetListProps{
			Db: db,
		}))
		todos.POST("", controller_todos.Add(controller_todos.AddProps{
			Db:           db,
			GetNewTodoId: util.GetNewUserId,
		}))
		todos.GET("/:todoId", controller_todos.GetById(controller_todos.GetByIdProps{
			Db: db,
		}))
		todos.PATCH("/:todoId", controller_todos.PatchById(controller_todos.PatchProps{
			Db: db,
		}))
		todos.DELETE("/:todoId", controller_todos.DeleteById(controller_todos.DeleteProps{
			Db: db,
		}))
	}
}
