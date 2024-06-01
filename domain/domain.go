package domain

type User struct {
	Id UserId `json:"id"`
	Name UserName `json:"name"`
	Email UserEmail `json:"email"`
	PhoneNumber UserPhoneNumber `json:"phone_number"`
}

type Todo struct {
	Id TodoId `json:"id"`
	Title TodoTitle `json:"title"`
	Person TodoPerson `json:"person"`
	Done TodoDone `json:"done"`
}

type UserId struct {
	Value uint `json:"value"`
}

type UserName struct {
	Value string `json:"value"`
}

type UserEmail struct {
	Value string `json:"value"`
}

type UserPhoneNumber struct{
	Value string `json:"value"`
}

type TodoId struct {
	Value uint `json:"value"`
}

type TodoTitle struct {
	Value string `json:"value"`
}

type TodoPerson struct{
	Value string `json:"value"`
}

type TodoDone struct {
	Value bool `json:"value"`
}

type CreateUser struct {
	Name UserName
	Email UserEmail
	PhoneNumber UserPhoneNumber
}

type CreateTodo struct {
	Title TodoTitle
	Person TodoPerson
}

type UpdateTodo struct {
	Title TodoTitle `json:"title"`
	Person TodoPerson `json:"person"`
	Done TodoDone `json:"done"`
}

func NewUser(name string, email string,phoneNumber string) User {
	return User{
		Name: UserName{Value: name},
		Email: UserEmail{Value: email},
		PhoneNumber: UserPhoneNumber{Value: phoneNumber},
	}
}

func NewTodo(title string, done bool) Todo {
	return Todo{
		Title: TodoTitle{Value: title},
		Done: TodoDone{Value: done},
	}
}