package service

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/nikitalier/tenderMonitoring/pkg/models"
)

var jwtKey = []byte("secret")

//Auth ...
func (s *Service) Auth(login string, password string) bool {
	u := s.repository.GetUserByLogin(login)
	if u.Password != password {
		return false
	}
	return true
}

//GenerateJWT ...
func (s *Service) GenerateJWT(user models.User) (token models.JwtToken, err error) {
	var cred models.Credential
	cred.ID = user.ID
	cred.Login = user.Login
	cred.FullName = user.FullName

	userRoles := s.GetUserRoles(cred.ID)
	cred.Authorities = userRoles

	expirationTime := time.Now().Add(3600 * time.Second)

	claims := &models.Claims{
		Credential: cred,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	tokenUnSign := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	token.AccessToken, err = tokenUnSign.SignedString(jwtKey)

	return token, err
}
