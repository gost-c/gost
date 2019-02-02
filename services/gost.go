package services

import (
	"github.com/zcong1993/gost/db"
	"github.com/zcong1993/gost/models"
	"github.com/zcong1993/libgo/gin/ginhelper"
)

func GetGosts(limit, offset int) ([]*models.Gost, int, error) {
	var gosts []*models.Gost
	count, err := ginhelper.PaginationQuery(db.ORM.Model(new(models.Gost)).Where("private = ?", 0).Preload("User").Preload("Files").Order("created_at desc"), &gosts, limit, offset)
	return gosts, count, err
}

func GetUserPublicGosts(userID interface{}, limit, offset int) ([]*models.Gost, int, error) {
	var gosts []*models.Gost
	count, err := ginhelper.PaginationQuery(db.ORM.Model(new(models.Gost)).Where("user_id = ?", userID).Where("private = ?", 0).Preload("User").Preload("Files").Order("created_at desc"), &gosts, limit, offset)
	return gosts, count, err
}

func DeleteUserGost(user *models.User, id string) error {
	var gost models.Gost
	err := db.ORM.Where("id = ?", id).Where("user_id = ?", user.ID).First(&gost).Error
	if err != nil {
		return err
	}
	return db.ORM.Delete(&gost).Error
}
