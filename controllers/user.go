package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/zcong1993/gost/common"
	"github.com/zcong1993/gost/middlewares"
	"github.com/zcong1993/gost/models"
	"github.com/zcong1993/gost/services"
	"github.com/zcong1993/gost/utils"
	"github.com/zcong1993/libgo/gin/ginerr"
	"github.com/zcong1993/libgo/gin/ginhelper"
	utils2 "github.com/zcong1993/libgo/utils"
	"github.com/zcong1993/libgo/validator"
	"net/http"
)

func Register() ginerr.ApiController {
	return func(ctx *gin.Context) ginerr.ApiError {
		var f models.User
		err := ctx.ShouldBindJSON(&f)
		if err != nil {
			return common.CreateInvalidErr(validator.NormalizeErr(err))
		}

		err = services.CreateUser(&f)
		if err != nil {
			if utils.IsDuplicateError(err) {
				return common.DUPLICATE_USER
			}
			log.Error("register user error: ", err)
			return common.INTERNAL_ERROR
		}

		tk, err := services.CreateToken(f)

		if err != nil {
			log.Error("create token error: ", err)
			return common.INTERNAL_ERROR
		}

		ctx.JSON(http.StatusCreated, common.TokenResp{Token: tk.Token})
		return nil
	}
}

func Login() ginerr.ApiController {
	return func(ctx *gin.Context) ginerr.ApiError {
		var f models.User
		err := ctx.ShouldBindJSON(&f)
		if err != nil {
			return common.CreateInvalidErr(validator.NormalizeErr(err))
		}

		user, err := services.GetUserByName(f.Username)
		if err != nil {
			if gorm.IsRecordNotFoundError(err) {
				return common.NOT_FOUND_ERROR
			}
			log.Error("get user by name error: ", err)
			return common.INTERNAL_ERROR
		}
		if !utils2.ComparePassword(f.Password, user.Password) {
			return common.INVALID_USERNAME_OR_PASSWORD
		}

		tk, err := services.CreateToken(*user)

		if err != nil {
			log.Error("create user token error: ", err)
			return common.INTERNAL_ERROR
		}

		ctx.JSON(http.StatusCreated, common.TokenResp{Token: tk.Token})
		return nil
	}
}

func RevokeToken() ginerr.ApiController {
	return func(ctx *gin.Context) ginerr.ApiError {
		token := ctx.Param("token")
		err := services.DeleteToken(token)
		if err != nil {
			if gorm.IsRecordNotFoundError(err) {
				return common.NOT_FOUND_ERROR
			}
			log.Error("invoke user token error: ", err)
			return common.INTERNAL_ERROR
		}
		ctx.Status(http.StatusNoContent)
		return nil
	}
}

func Me() ginerr.ApiController {
	return func(ctx *gin.Context) ginerr.ApiError {
		user := middlewares.MustGetUser(ctx)
		ctx.JSON(http.StatusOK, user)
		return nil
	}
}

func MyGosts() ginerr.ApiController {
	return func(ctx *gin.Context) ginerr.ApiError {
		user := middlewares.MustGetUser(ctx)
		gosts, err := services.GostByUser(*user)
		if err != nil {
			log.Error("get gosts by user error: ", err)
			return common.INTERNAL_ERROR
		}
		gostResp := make([]common.GostResp, len(gosts))
		utils.MustCopy(&gostResp, gosts)
		ctx.JSON(http.StatusOK, gostResp)
		return nil
	}
}

func UserGosts() ginerr.ApiController {
	return func(ctx *gin.Context) ginerr.ApiError {
		userID := ctx.Param("id")
		limit, offset := common.Paginator.ParsePagination(ctx)
		gosts, count, err := services.GetUserPublicGosts(userID, limit, offset)
		if err != nil {
			log.Error("get user public gosts error: ", err)
			return common.INTERNAL_ERROR
		}
		gs := make([]common.GostResp, len(gosts))
		utils.MustCopy(&gs, gosts)
		ginhelper.ResponsePagination(ctx, count, gs)
		return nil
	}
}
