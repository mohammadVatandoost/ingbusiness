package authentication

type Config struct {
	JwtSecret         string
	GoogleKey         string
	GoogleSecret      string
	GoogleCallbackUrl string
	EnableSSL         bool
}
