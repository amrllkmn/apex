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
}

// the Service struct
type Service struct{}

// Take data from handler and call create an Activity
// Take data from handler and call read an Activity
func (svc *Service) GetActivity(id int) (string, error) {
	return "Got activity", nil
}

func (svc *Service) CreateActivity(activity *Activity) (string, error) {
	fmt.Println(*activity)
	return "Created activity", nil
}

func (svc *Service) GetActivities() []Activity {
	return []Activity{}
}

func NewService() ActivityService {
	return &Service{}
}

// Take data from handler and call update an Activity
// Take data from handler and call delete an Activity
