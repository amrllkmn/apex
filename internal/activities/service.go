package activities

import "fmt"

// the Activity Model
type Activity struct {
	id    int
	count int
	Title string `json:"title"`
}

// the Activity DB Layer
// Create an Activity
// Read an Activity
// Update an Activity
// Delete an Activity

// the Activity Service interface
type ActivityService interface {
	GetActivities() []Activity
	GetActivity(id int) (string, error)
	CreateActivity(activity *Activity) (string, error)
	UpdateActivity(id int, activity *Activity) (string, error)
	DeleteActivity(id int) (string, error)
}

// Take data from handler and call create an Activity
func (svc *Service) CreateActivity(activity *Activity) (string, error) {
	fmt.Println(*activity)
	return "Created activity", nil
}

// Take data from handler and call read an Activity
func (svc *Service) GetActivity(id int) (string, error) {
	return "Got activity", nil
}

func (svc *Service) GetActivities() []Activity {
	return []Activity{}
}

// Take data from handler and call update an Activity
func (svc *Service) UpdateActivity(id int, activity *Activity) (string, error) {
	fmt.Println(*activity)
	return "Updated activity", nil
}

// Take data from handler and call delete an Activity
func (svc *Service) DeleteActivity(id int) (string, error) {
	return "Deleted activity", nil
}

// the Service struct
type Service struct{}

func NewService() ActivityService {
	return &Service{}
}
