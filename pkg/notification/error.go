package notification

import v1 "github.com/mohammadVatandoost/ingbusiness/api/services/notification/v1"

func MakeError(title string, message string) *v1.Notification {
	return &v1.Notification{
		Title:   title,
		Message: message,
		Icon:    "",
		Color:   "red",
	}
}
