package notification

import (
	v1 "github.com/mohammadVatandoost/ingbusiness/api/services/notification/v1"
	"google.golang.org/protobuf/encoding/protojson"
)

func BuildNotification(notifications ...*v1.Notification) ([]byte, error) {
	notifs := v1.Notifications{
		Notifications: notifications,
	}
	return protojson.Marshal(&notifs)
}
