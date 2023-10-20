package utils

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type TokenPayload struct {
	Email    string
	UserName string
	ID       uint
}

type CustomClaims struct {
	jwt.RegisteredClaims
	Email    string
	UserName string
	ID       uint
}

func GenerateAccessToken(payload *TokenPayload) (string, error) {
	var (
		t *jwt.Token
	)
	key, err := fetchSecretKey(false)
	if err != nil {
		return "", err
	}
	claims := new(CustomClaims)
	claims.ID = payload.ID
	claims.Email = payload.Email
	claims.UserName = payload.UserName
	t = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	s, err := t.SignedString(key)
	if err != nil {
		return "", err
	}
	println(s)
	return s, nil
}

func VerifyAccessToken(tokenString string) (*TokenPayload, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		key, err := fetchSecretKey(true)
		return key, err
	})
	if err != nil {
		println(err.Error())
		return nil, err
	}
	customClaims, _ := token.Claims.(*CustomClaims)
	payload := &TokenPayload{
		Email:    customClaims.Email,
		ID:       customClaims.ID,
		UserName: customClaims.UserName,
	}
	return payload, nil
}

func fetchSecretKey(throwErrorIfNotFound bool) ([]byte, error) {
	secretKeyAsBase64, found := os.LookupEnv("SECRET_KEY")
	if found && secretKeyAsBase64 != "" {
		rawKey, err := base64.StdEncoding.DecodeString(secretKeyAsBase64)
		println(rawKey)
		if err == nil {
			return rawKey, nil
		}
		if throwErrorIfNotFound {
			erruu := os.Unsetenv("SECRET_KEY")
			errjoij := errors.Join(erruu, err)
			println(errjoij.Error())
			return nil, errjoij
		}
	}
	// Create a byte slice to store the random bytes
	secretKey := make([]byte, 32)

	// Read random bytes from a cryptographically secure source
	_, err := rand.Read(secretKey)
	if err != nil {
		return nil, err
	}

	var step1 []byte
	sha := sha256.New()
	sha.Write(secretKey)
	step1 = sha.Sum(nil)
	hm := hmac.New(sha256.New, step1)
	summed := hm.Sum(step1)
	b64 := base64.StdEncoding.EncodeToString(summed)
	os.Setenv("SECRET_KEY", b64)
	return summed, nil
}

func GetSessionUser(c *gin.Context) (*TokenPayload, error) {

	payload, found := c.Get("user")
	if !found {
		return nil, errors.New("unable to process user")
	}
	authUser, ok := payload.(*TokenPayload)
	if !ok {
		return nil, errors.New("unable to process user")
	}
	return authUser, nil
}
