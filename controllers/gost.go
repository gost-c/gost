package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/zcong1993/gost/common"
	"github.com/zcong1993/gost/db"
	"github.com/zcong1993/gost/logger"
	"github.com/zcong1993/gost/middlewares"
	"github.com/zcong1993/gost/models"
	"github.com/zcong1993/gost/services"
	"github.com/zcong1993/gost/utils"
	"github.com/zcong1993/libgo/gin/ginerr"
	"github.com/zcong1993/libgo/gin/ginhelper"
	"github.com/zcong1993/libgo/validator"
	"net/http"
)

var log = logger.Logger

func CreateGost() ginerr.ApiController {
	return func(ctx *gin.Context) ginerr.ApiError {
		var f models.Gost
		err := ctx.ShouldBindJSON(&f)
		if err != nil {
			return common.CreateInvalidErr(validator.NormalizeErr(err))
		}

		user := middlewares.MustGetUser(ctx)
		f.User = user

		err = db.ORM.Create(&f).Error
		if err != nil {
			log.Error("create gost error: ", err)
			return common.INTERNAL_ERROR
		}
		var gost common.GostResp
		utils.MustCopy(&gost, f)

		ctx.JSON(http.StatusOK, gost)
		return nil
	}
}

func GetAllGosts() ginerr.ApiController {
	return func(ctx *gin.Context) ginerr.ApiError {
		limit, offset := common.Paginator.ParsePagination(ctx)
		gosts, count, err := services.GetGosts(limit, offset)
		if err != nil {
			log.Error("get all gosts error: ", err)
			return common.INTERNAL_ERROR
		}
		gs := make([]common.GostResp, len(gosts))
		utils.MustCopy(&gs, gosts)
		ginhelper.ResponsePagination(ctx, count, gs)
		return nil
	}
}

func DeleteGost() ginerr.ApiController {
	return func(ctx *gin.Context) ginerr.ApiError {
		id := ctx.Param("id")
		user := middlewares.MustGetUser(ctx)
		err := services.DeleteUserGost(user, id)

		if err != nil {
			if gorm.IsRecordNotFoundError(err) {
				return common.NOT_FOUND_ERROR
			}
			log.Error("delete gost error: ", err)
			return common.INTERNAL_ERROR
		}
		ctx.Status(http.StatusNoContent)
		return nil
	}
}

func RetrieveGost() ginerr.ApiController {
	return func(ctx *gin.Context) ginerr.ApiError {
		id := ctx.Param("id")
		var gost models.Gost
		err := db.ORM.Where("id = ?", id).Preload("User").Preload("Files").First(&gost).Error
		if err != nil {
			if gorm.IsRecordNotFoundError(err) {
				return common.NOT_FOUND_ERROR
			}
			log.Error("retrieve gost error: ", err)
			return common.INTERNAL_ERROR
		}
		var g common.GostResp
		utils.MustCopy(&g, &gost)
		ctx.JSON(http.StatusOK, g)
		return nil
	}
}
