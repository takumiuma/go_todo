package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"practice/domain"
	"practice/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TodoHandler struct {
	todoUsecase usecase.TodoUsecase
}

type User struct {
	Name string `json:"name"`
	Email string `json:"email"`
	PhoneNumber string `json:"phone_number"`
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

func (h TodoHandler) RegistUser(c *gin.Context) {
	var param User
	if err := json.NewDecoder(c.Request.Body).Decode(&param); err != nil {
		log.Fatal(err)
	}

	if param.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "name is required",
		})
		return
	}
	if param.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "email is required",
		})
		return
	}
	if param.PhoneNumber == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "phone_number is required",
		})
		return
	}

	user := domain.CreateUser{
		Name: domain.UserName{Value: param.Name},
		Email: domain.UserEmail{Value: param.Email},
		PhoneNumber: domain.UserPhoneNumber{Value: param.PhoneNumber},
	}

	newUser, err := h.todoUsecase.RegistUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	response := UserResponse{
		User: newUser,
	}

	c.JSON(http.StatusOK, response)
}

func (h TodoHandler) Create(c *gin.Context) {
	var param Todo
	if err := json.NewDecoder(c.Request.Body).Decode(&param); err != nil {
		log.Fatal(err)
	}

	// validate := validator.New()  //インスタンス生成
	// errors := validate.Struct(param) //バリデーションを実行し、NGの場合、ここでエラーが返る。
	// if(errors != nil) {
	// 		c.JSON(http.StatusBadRequest, gin.H{
	// 			"message": "goバリデーターだよ",
	// 		})
	// 	return
	// }

	if param.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "title is required",
		})
		return
	}

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

type UserResponse struct {
	User domain.User `json:"user"`
}

type TodoResponse struct {
	Todo domain.Todo `json:"todo"`
}