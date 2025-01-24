package activities

import (
	"fmt"

	"gorm.io/gorm"
)

// the Activity Model
type Activity struct {
	gorm.Model
	Count int    `json:"count"`
	Title string `json:"title"`
}

type ActivityRepository interface {
	FindById()
	FindAll()
	UpdateById()
	DeleteById()
	CreateOne(activity *Activity)
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

func (repo *Repository) CreateOne(activity *Activity) {
	fmt.Println("I'm creating one")
	result := repo.DB.Create(activity)
	if result.Error != nil {
		fmt.Println("Failed to create")
	}
}

type Repository struct {
	DB *gorm.DB
}

func NewRepo(db *gorm.DB) ActivityRepository {
	return &Repository{
		DB: db,
	}
}
