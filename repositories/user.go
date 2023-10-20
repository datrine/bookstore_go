package repositories

import (
	"time"

	"github.com/datrine/conn"
	"github.com/datrine/models"
	"github.com/datrine/utils"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	Email       string `validate:"email,required"`
	PassHash    string
	LastName    string
	FirstName   string
	DOB         *time.Time
	AccountID   uint
	PhoneNumber string
}
type Account struct {
	ID   uint
	Role string
}

var validate *validator.Validate

func init() {
	validate = validator.New(validator.WithRequiredStructEnabled())
}

func CreateUser(data *utils.RegisterDto) (interface{}, error) {
	accountData := &Account{Role: "USER"}
	userData := &User{
		Email:       data.Email,
		LastName:    data.LastName,
		FirstName:   data.FirstName,
		PhoneNumber: data.PhoneNumber,
	}
	err := validate.Struct(accountData)
	if err != nil {
		return nil, err
	}
	err = validate.Struct(userData)
	if err != nil {
		return nil, err
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(data.Password), 10)
	if err != nil {
		return nil, err
	}
	userData.PassHash = string(hash)
	db := conn.GetDB()
	if db == nil {
		println("Nil connection established....")
	}
	err = db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&accountData).Error; err != nil {
			return err
		}
		userData.AccountID = accountData.ID
		if err = tx.Create(&userData).Error; err != nil {
			return err
		}
		return nil
	})
	return &userData, err
}

func GetUserByID(userID uint) (interface{}, error) {
	user := &models.User{ID: userID}
	db := conn.GetDB()
	result := db.First(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func GetUserByEmail(email string) (*models.User, error) {
	user := &models.User{Email: email}
	db := conn.GetDB()
	result := db.First(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}
