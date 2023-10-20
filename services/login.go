package services

import (
	"errors"

	"github.com/datrine/repositories"
	"github.com/datrine/utils"
	"golang.org/x/crypto/bcrypt"
)

func BasicLoginService(data *utils.BasicLoginDto) (*Tokens, error) {
	user, err := repositories.GetUserByEmail(data.Email)
	if err != nil {
		return nil, errors.New("User does not exist.")
	}
	err =
		bcrypt.CompareHashAndPassword([]byte(user.PassHash), []byte(data.Password))
	if err != nil {
		println(err.Error())
		return nil, err
	}
	token, err := utils.GenerateAccessToken(&utils.TokenPayload{UserName: user.UserName, ID: user.ID, Email: user.Email})
	if err != nil {
		println(err.Error())
		return nil, err
	}
	return &Tokens{
		AccessToken: token,
	}, err
}

type Tokens struct {
	AccessToken string
}
