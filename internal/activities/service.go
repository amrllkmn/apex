package activities

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
	svc.repository.CreateOne()
	return "Created activity", nil
}

// Take data from handler and call read an Activity
func (svc *Service) GetActivity(id int) (string, error) {
	svc.repository.FindById()
	return "Got activity", nil
}

func (svc *Service) GetActivities() []Activity {
	svc.repository.FindAll()
	return []Activity{}
}

// Take data from handler and call update an Activity
func (svc *Service) UpdateActivity(id int, activity *Activity) (string, error) {
	svc.repository.UpdateById()
	return "Updated activity", nil
}

// Take data from handler and call delete an Activity
func (svc *Service) DeleteActivity(id int) (string, error) {
	svc.repository.DeleteById()
	return "Deleted activity", nil
}

// the Service struct
type Service struct {
	repository ActivityRepository
}

func NewService(activityRepo ActivityRepository) ActivityService {
	return &Service{
		repository: activityRepo,
	}
}
