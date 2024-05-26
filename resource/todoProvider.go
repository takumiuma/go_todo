package resource

import (
	"gorm.io/gorm"
)

type TodoDriver interface {
	GetAllUser() ([]User, error)
	GetAll() ([]Todo, error)
	GetById(id uint) (Todo, error)
	Create(todo CreateTodo) (error)
	Update(id uint, todo UpdateTodo) (error)
	Delete(id uint) (error)
}

type TodoDriverImpl struct {
	conn *gorm.DB
}

func (t TodoDriverImpl) GetAllUser() ([]User, error) {
	users := []User{}
	t.conn.Find(&users)

	return users, nil
}

func (t TodoDriverImpl) GetAll() ([]Todo, error) {
	todos := []Todo{}
	t.conn.Find(&todos)

	return todos, nil
}

func (t TodoDriverImpl) GetById(id uint) (Todo, error) {
	todo := Todo{}
	t.conn.First(&todo, id)
	return todo, nil
}

func (t TodoDriverImpl) Create(todo CreateTodo) (error) {
	err := t.conn.Create(&todo)

	if err != nil {
		return err.Error
	}

	return nil
}

func (t TodoDriverImpl) Update(id uint, todo UpdateTodo) (error) {
	err := t.conn.Model(&todo).Where("id = ?", id).Select("updated_at", "title", "person", "done").Updates(todo)

	if err != nil {
		return err.Error
	}
	
	return nil
}

func (t TodoDriverImpl) Delete(id uint) (error) {
	todo := Todo{}
	err := t.conn.Delete(&todo, id)

	if err != nil {
		return err.Error
	}

	return nil
}

type User struct {
    gorm.Model
	Name 		string `gorm:"size:255" json:"name"`
	Email 			string  `gorm:"size:100" json:"email"`
	PhoneNumber		string `gorm:"size:100" json:"phone_number"`
}

type Todo struct {
    gorm.Model
	Title 		string `gorm:"size:255" json:"title"`
	Person		string `gorm:"size:100" json:"person"`
	Done 			bool  `gorm:"default:false" json:"done"`
}

type CreateTodo struct {
    gorm.Model
	Title string `json:"title"`
	Person	string `json:"person"`
	Done  bool   `json:"done"`
}

func (CreateTodo) TableName() string {
	return "todos"
}

type UpdateTodo struct {
    gorm.Model
	Title string `json:"title"`
	Person	string `json:"person"`
	Done  bool   `json:"done"`
}

func (UpdateTodo) TableName() string {
	return "todos"
}

func NewTodo(id int, title string, done bool) Todo {
	return Todo{
		Title: title,
		Done: done,
	}
}

func ProvideTodoDriver(conn *gorm.DB) TodoDriver {
	return TodoDriverImpl{conn: conn}
}

