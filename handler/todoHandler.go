package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"practice/domain"
	"practice/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type TodoHandler struct {
	todoUsecase usecase.TodoUsecase
}
type Todo struct {
	Id int `json:"id"`
	Title string `json:"title" validate:"gt=0,lt=100"`
	Person string `json:"person" validate:"oneof=担当者A 担当者B 担当者C"`
	Done bool `json:"done"`
}

func (h TodoHandler) GetAllUser(c *gin.Context) {
	users, err := h.todoUsecase.GetAllUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	response := UsersResponse{
		Users: users,
	}
	c.JSON(http.StatusOK, response)
}

func (h TodoHandler) GetAll(c *gin.Context) {
	todos, err := h.todoUsecase.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	response := TodosResponse{
		Todos: todos,
	}
	c.JSON(http.StatusOK, response)
}

func (h TodoHandler) GetById(c *gin.Context) {
	paramsId := c.Params.ByName("id")
	uintId, err := strconv.ParseUint(paramsId, 10, 32)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	
	id := domain.TodoId{Value: uint(uintId)}
	todo, err := h.todoUsecase.GetById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	response := TodoResponse{
		Todo: todo,
	}
	c.JSON(http.StatusOK, response)
}

func (h TodoHandler) Create(c *gin.Context) {
	var param Todo
	if err := json.NewDecoder(c.Request.Body).Decode(&param); err != nil {
		log.Fatal(err)
	}

	validate := validator.New()  //インスタンス生成
	errors := validate.Struct(param) //バリデーションを実行し、NGの場合、ここでエラーが返る。
	if(errors != nil) {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "title is required",
			})
		return
	}

	// if param.Title == "" {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"message": "title is required",
	// 	})
	// 	return
	// }

	todo := domain.CreateTodo{
		Title: domain.TodoTitle{Value: param.Title},
		Person: domain.TodoPerson{Value: param.Person},
	}

	newTodo, err := h.todoUsecase.Create(todo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	response := TodoResponse{
		Todo: newTodo,
	}

	c.JSON(http.StatusOK, response)
}

func (h TodoHandler) Update(c *gin.Context) {
	paramsId := c.Params.ByName("id")
	uintId, err := strconv.ParseUint(paramsId, 10, 32)
	id := domain.TodoId{Value: uint(uintId)}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	var param Todo
	if err := json.NewDecoder(c.Request.Body).Decode(&param); err != nil {
		log.Fatal(err)
	}
	if param.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "title is required",
		})
		return
	}

	updateTodo := domain.UpdateTodo{
		Title: domain.TodoTitle{Value: param.Title},
		Person:domain.TodoPerson{Value:param.Person},
		Done: domain.TodoDone{Value: param.Done},
	}

	err = h.todoUsecase.Update(id, updateTodo)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func (h TodoHandler) Delete(c *gin.Context) {
	paramsId := c.Params.ByName("id")
	uintId, err := strconv.ParseUint(paramsId, 10, 32)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H {
			"message": err.Error(),
		})

		return
	}

	id := domain.TodoId{Value: uint(uintId)}
	err = h.todoUsecase.Delete(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H {
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H {
		"message": "success",
	})
}

func ProvideTodoHandler(u usecase.TodoUsecase) *TodoHandler {
	return &TodoHandler{u}
}

type UsersResponse struct {
	Users []domain.User `json:"users"`
}

type TodosResponse struct {
	Todos []domain.Todo `json:"todos"`
}

type TodoResponse struct {
	Todo domain.Todo `json:"todo"`
}