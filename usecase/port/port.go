package port

import "practice/domain"

type TodoPort interface {
	GetAllUser() ([]domain.User, error)
	GetAll() ([]domain.Todo, error)
	GetById(id domain.TodoId) (domain.Todo, error)
	RegistUser(user domain.CreateUser)(domain.User,error)
	Create(todo domain.CreateTodo) (domain.Todo, error)
	Update(id domain.TodoId, todo domain.UpdateTodo) (error)
	Delete(id domain.TodoId) (error)
}