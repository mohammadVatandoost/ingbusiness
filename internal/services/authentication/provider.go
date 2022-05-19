package authentication

import (
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
	"github.com/mohammadVatandoost/ingbusiness/internal/users"
	"github.com/sirupsen/logrus"
)

type Service struct {
	logger         *logrus.Logger
	usersDirectory users.Querier
	conf           Config
	jwtSecret      []byte
}

func New(logger *logrus.Logger, usersDirectory users.Querier, conf Config) *Service {
	key := conf.GoogleSecret // Replace with your SESSION_SECRET or similar
	maxAge := 86400 * 30     // 30 days
	isProd := conf.EnableSSL // Set to true when serving over https

	store := sessions.NewCookieStore([]byte(key))
	store.MaxAge(maxAge)
	store.Options.Path = "/"
	store.Options.HttpOnly = true // HttpOnly should always be enabled
	store.Options.Secure = isProd

	gothic.Store = store

	goth.UseProviders(
		google.New(conf.GoogleKey, conf.GoogleSecret, conf.GoogleCallbackUrl, "email", "profile"),
	)

	return &Service{
		usersDirectory: usersDirectory,
		conf:           conf,
		jwtSecret:      []byte(conf.JwtSecret),
	}
}
