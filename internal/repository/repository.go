package repository

type Authorization interface {
}

type TodoLists interface {
}

type TodoItems interface {
}

type Repository struct {
	Authorization
	TodoLists
	TodoItems
}

func NewRepository() *Repository {
	return &Repository{}
}
