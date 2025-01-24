package activities

// the Activity DB Layer
// Create an Activity
// Read an Activity
// Update an Activity
// Delete an Activity

// the Activity Service interface
type ActivityService interface {
	GetActivities() ([]Activity, error)
	GetActivity(id int) (Activity, error)
	CreateActivity(activity *Activity) (string, error)
	UpdateActivity(id int, activity *Activity) (string, error)
	DeleteActivity(id int) (string, error)
}

// Take data from handler and call create an Activity
func (svc *Service) CreateActivity(activity *Activity) (string, error) {
	err := svc.repository.CreateOne(activity)
	if err != nil {
		return "Failed to create", err
	}
	return "Activity created", nil
}

// Take data from handler and call read an Activity
func (svc *Service) GetActivity(id int) (Activity, error) {
	activity, err := svc.repository.FindById(id)
	if err != nil {
		return activity, err
	}
	return activity, nil
}

func (svc *Service) GetActivities() ([]Activity, error) {
	var activities []Activity
	err := svc.repository.FindAll(&activities)
	if err != nil {
		return []Activity{}, err
	}
	return activities, nil
}

// Take data from handler and call update an Activity
func (svc *Service) UpdateActivity(id int, activity *Activity) (string, error) {
	updatedActivity, err := svc.repository.UpdateById(id, activity)
	if err != nil {
		return updatedActivity, err
	}
	return updatedActivity, nil
}

// Take data from handler and call delete an Activity
func (svc *Service) DeleteActivity(id int) (string, error) {
	deleteErr := svc.repository.DeleteById(id)
	if deleteErr != nil {
		return "Failed delete", deleteErr
	}
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
