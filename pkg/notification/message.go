package notification

import (
	v1 "github.com/mohammadVatandoost/ingbusiness/api/services/notification/v1"
)

func BuildNotification(notifications ...*v1.Notification) []*v1.Notification {
	return notifications
}
