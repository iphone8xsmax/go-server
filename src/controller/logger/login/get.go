// Copyright 2019 Axetroy. All rights reserved. MIT license.
package login

import (
	"errors"
	"github.com/axetroy/go-server/src/controller"
	"github.com/axetroy/go-server/src/exception"
	"github.com/axetroy/go-server/src/model"
	"github.com/axetroy/go-server/src/schema"
	"github.com/axetroy/go-server/src/service/database"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/mitchellh/mapstructure"
	"net/http"
	"time"
)

func GetLatestLoginLog(context controller.Context) (res schema.Response) {
	var (
		err  error
		data = schema.LogLogin{}
	)

	defer func() {
		if r := recover(); r != nil {
			switch t := r.(type) {
			case string:
				err = errors.New(t)
			case error:
				err = t
			default:
				err = exception.Unknown
			}
		}

		if err != nil {
			res.Data = nil
			res.Message = err.Error()
		} else {
			res.Data = data
			res.Status = schema.StatusSuccess
		}
	}()

	logInfo := model.LoginLog{
		Uid: context.Uid,
	}

	query := schema.Query{}

	query.Normalize()

	if err = database.Db.Where(&logInfo).Preload("User").Order(query.Sort).First(&logInfo).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			err = exception.NoData
		}
		return
	}

	if err = mapstructure.Decode(logInfo, &data.LogLoginPure); err != nil {
		return
	}

	if err = mapstructure.Decode(logInfo.User, &data.User); err != nil {
		return
	}

	data.CreatedAt = logInfo.CreatedAt.Format(time.RFC3339Nano)
	data.UpdatedAt = logInfo.UpdatedAt.Format(time.RFC3339Nano)

	return
}

func GetLoginLog(id string) (res schema.Response) {
	var (
		err  error
		data = schema.LogLogin{}
	)

	defer func() {
		if r := recover(); r != nil {
			switch t := r.(type) {
			case string:
				err = errors.New(t)
			case error:
				err = t
			default:
				err = exception.Unknown
			}
		}

		if err != nil {
			res.Data = nil
			res.Message = err.Error()
		} else {
			res.Data = data
			res.Status = schema.StatusSuccess
		}
	}()

	logInfo := model.LoginLog{
		Id: id,
	}

	query := schema.Query{}

	query.Normalize()

	if err = database.Db.Where(&logInfo).Preload("User").Order(query.Sort).First(&logInfo).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			err = exception.NoData
		}
		return
	}

	if err = mapstructure.Decode(logInfo, &data.LogLoginPure); err != nil {
		return
	}

	if err = mapstructure.Decode(logInfo.User, &data.User); err != nil {
		return
	}

	data.CreatedAt = logInfo.CreatedAt.Format(time.RFC3339Nano)
	data.UpdatedAt = logInfo.UpdatedAt.Format(time.RFC3339Nano)

	return
}

func GetLoginLogRouter(context *gin.Context) {
	var (
		err error
		res = schema.Response{}
	)

	defer func() {
		if err != nil {
			res.Data = nil
			res.Message = err.Error()
		}
		context.JSON(http.StatusOK, res)
	}()

	id := context.Param("log_id")

	res = GetLoginLog(id)
}