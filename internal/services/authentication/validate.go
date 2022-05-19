package authentication

import "github.com/mohammadVatandoost/ingbusiness/pkg/jwt"

func (s *Service) ValidateJWT(token string) (*jwt.Message, error) {
	return jwt.Parse(token, s.jwtSecret)
}
