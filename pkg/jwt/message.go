package jwt

type Message struct {
	UserID    int32  `json:"user_id"`
	UserName  string `json:"user_name"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Timestamp int64  `json:"timestamp"`
}
