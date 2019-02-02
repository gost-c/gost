package services

import (
	"github.com/zcong1993/gost/db"
	"github.com/zcong1993/gost/models"
	"github.com/zcong1993/libgo/utils"
	"time"
)

const TOKEN_EXPIRE = time.Duration(time.Hour * 24 * 180)

func GetUserByToken(token string) (*models.User, error) {
	var tk models.Token
	err := db.ORM.Preload("User").Where("token = ?", token).Where("created_at > ?", time.Now().Add(-TOKEN_EXPIRE)).First(&tk).Error
	if err != nil {
		return nil, err
	}
	return tk.User, nil
}

func GetUserByName(name string) (*models.User, error) {
	var user models.User
	err := db.ORM.Where("username = ?", name).First(&user).Error
	return &user, err
}

func CreateUser(user *models.User) error {
	user.Password = utils.HashPassword(user.Password)
	return db.ORM.Create(user).Error
}

func CreateToken(user models.User) (*models.Token, error) {
	tk, err := utils.GenerateRandomStringURLSafe(32)
	if err != nil {
		return nil, err
	}
	token := &models.Token{
		Token: tk,
		User:  &user,
	}
	err = db.ORM.Create(token).Error
	return token, err
}

func GostByUser(user models.User) ([]*models.Gost, error) {
	var gosts []*models.Gost
	err := db.ORM.Model(&user).Preload("User").Preload("Files").Related(&gosts).Order("created_at desc").Error
	return gosts, err
}

func DeleteToken(token string) error {
	var tk models.Token
	err := db.ORM.Where("token = ?", token).First(&tk).Error
	if err != nil {
		return err
	}
	return db.ORM.Delete(&tk).Error
}
