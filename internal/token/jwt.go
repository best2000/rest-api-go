package tkn

import (
	"rest-api/internal/config"

	"github.com/go-chi/jwtauth/v5"
)

var TokenAuth *jwtauth.JWTAuth

type TokenPayload struct {
	Userlogin string
}

func New(conf config.Config) {
	TokenAuth = jwtauth.New("HS256", []byte(conf.App.TokenSecret), nil)
}

func Decode(tkn string) (TokenPayload, error) {
	payload, err := TokenAuth.Decode(tkn)
	if err != nil {
		return TokenPayload{}, err
	}

	userLogin, _ := payload.Get("user_login")

	return TokenPayload{
		Userlogin: userLogin.(string),
	}, err
}
