package repository

type NotificationRepository struct {
	tx *gorm.DB
}

func NewNotificationRepository(tx *gorm.DB) *AuthRepository {
	log.Println("notification repository")
	return &NotificationRepository{
		TX: tx,
	}
}

func (n *NotificationRepository) ReadNotification(email string) ([]*model.Notification, error) {
	var notif map[string]interface{}

	err := n.TX.model(&model.Notification{}).Find(&notif).Where("email = ?", email).Error
	if err != nil {
		return nil, err
	}
	return notif, nil
}

func (n *NotificationRepository) NotifyUser(id uuid.UUID) ([]*model.User, error) {
	var user map[string]interface{}
	err := n.TX.Model(&model.User{}).Find(&user, "id", id).Error
	if err != nil {
		return nil, err
	}

	return user, nil
} 

func (n *NotificationRepository) DeleteNotification(id uuid.UUID) error {
	
	return n.TX.Where("id = ?", id).Delete(&model.Notification{}).Error 
}