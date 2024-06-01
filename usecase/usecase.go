package usecase

import (
	"practice/domain"
	"practice/usecase/port"
)

type TodoUsecase struct {
	todoPort port.TodoPort
}

func (u TodoUsecase) GetAllUser() ([]domain.User, error) {
	users, err :=  u.todoPort.GetAllUser()

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u TodoUsecase) GetAll() ([]domain.Todo, error) {
	todos, err :=  u.todoPort.GetAll()

	if err != nil {
		return nil, err
	}

	return todos, nil
}

func (u TodoUsecase) GetById(id domain.TodoId) (domain.Todo, error) {
	todo, err := u.todoPort.GetById(id)

	if err != nil {
		return domain.Todo{}, err
	}

	return todo, nil
}

func (u TodoUsecase) RegistUser(user domain.CreateUser) (domain.User, error) {
	newUser, err := u.todoPort.RegistUser(user)

	if err != nil {
		return domain.User{}, err
	}

	return newUser, nil
}

func (u TodoUsecase) Create(todo domain.CreateTodo) (domain.Todo, error) {
	newTodo, err := u.todoPort.Create(todo)

	if err != nil {
		return domain.Todo{}, err
	}

	return newTodo, nil
}

func (u TodoUsecase) Update(id domain.TodoId, todo domain.UpdateTodo) (error) {
	err := u.todoPort.Update(id, todo)

	if err != nil {
		return err
	}

	return nil
}

func (u TodoUsecase) Delete(id domain.TodoId) (error) {
	err := u.todoPort.Delete(id)

	if err != nil {
		return err
	}

	return nil
}

func ProvideTodoUsecase(todoPort port.TodoPort) TodoUsecase {
	return TodoUsecase{todoPort}
}