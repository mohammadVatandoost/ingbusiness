// Code generated by sqlc. DO NOT EDIT.

package savedmessages

import (
	"time"
)

type SavedMessage struct {
	ID             int32
	Message        string
	OrganizationID int32
	WriterID       int32
	CreateTime     time.Time
	UpdateTime     time.Time
}
