package activities

import "fmt"

type ActivityRepository interface {
	FindById()
	FindAll()
	UpdateById()
	DeleteById()
	CreateOne()
}

func (repo *Repository) FindById() {
	fmt.Println("I'm finding by id")
}
func (repo *Repository) FindAll() {
	fmt.Println("I'm finding all")
}
func (repo *Repository) UpdateById() {
	fmt.Println("I'm updating by id")
}
func (repo *Repository) DeleteById() {
	fmt.Println("I'm deleting by id")
}

func (repo *Repository) CreateOne() {
	fmt.Println("I'm creating one")
}

type Repository struct {
	DB int
}

func NewRepo(db int) ActivityRepository {
	return &Repository{
		DB: db,
	}
}
