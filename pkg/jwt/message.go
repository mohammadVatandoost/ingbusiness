package jwt


"id":        strconv.Itoa(msg.UserID),
"username":  msg.UserName,
"phone":     msg.Phone,
"email":     msg.Email,
"timestamp": time.Now().Unix(),

const name = 

type Message struct {
	UserID int
	UserName string
	Phone string
	Email string
}
