package controller_todos_test

import (
	"net/http"
	"testing"

	controller_todos "gin-todo-list/src/controller/todos"
	"gin-todo-list/src/mock"
	"gin-todo-list/src/model"
	"gin-todo-list/src/util"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestDeleteTodoSuccess(t *testing.T) {
	resForAdd := mock.GetResponse()
	cForAdd := mock.GetGinContext(resForAdd)
	userId := util.GetNewUserId()
	todoId := util.GetNewTodoId()
	gormDB := mock.GetMockGorm(t)

	cForAdd.Set("userId", userId)

	fake := model.Add{
		Title:       "123",
		Description: "456",
	}

	cForAdd.Request = &http.Request{
		Header: make(http.Header),
		Body:   mock.GetRequsetBody(fake),
	}

	controller_todos.Add(controller_todos.AddProps{
		Db:           gormDB,
		GetNewTodoId: mock.UtilGetNewTodoId(todoId),
	})(cForAdd)

	assert.Equal(t, http.StatusOK, resForAdd.Code)

	resForDelete := mock.GetResponse()
	cForDelete := mock.GetGinContext(resForDelete)
	cForDelete.Set("userId", userId)

	cForDelete.Params = []gin.Param{
		{Key: "todoId", Value: todoId.String()},
	}

	cForDelete.Request = &http.Request{
		Header: make(http.Header),
	}

	controller_todos.DeleteById(controller_todos.DeleteProps{
		Db: gormDB,
	})(cForDelete)

	var resBody SuccessTodoAPIResponse
	mock.GetResponseBody(resForDelete.Body.Bytes(), &resBody)

	assert.Equal(t, http.StatusOK, resForDelete.Code)
	assert.Equal(t, todoId, resBody.Data.ID)
	assert.Equal(t, fake.Title, resBody.Data.Title)
	assert.Equal(t, fake.Description, resBody.Data.Description)
}

func TestDeleteTodoFailByNotExist(t *testing.T) {
	res := mock.GetResponse()
	c := mock.GetGinContext(res)
	userId := util.GetNewUserId()
	todoId := util.GetNewTodoId()
	c.Set("userId", userId)

	c.Params = []gin.Param{
		{Key: "todoId", Value: todoId.String()},
	}

	c.Request = &http.Request{
		Header: make(http.Header),
	}

	gormDB := mock.GetMockGorm(t)

	defer func() {
		r := recover()
		if r == nil {
			t.Errorf("The code did not panic")
		}
		err := r.(*model.ApiError)
		assert.Equal(t, http.StatusServiceUnavailable, err.StatusCode)
		assert.Equal(t, model.ERROR_DELETE_TODO_NOT_EXIST, err.ErrorType)
	}()

	controller_todos.DeleteById(controller_todos.DeleteProps{
		Db: gormDB,
	})(c)
}
