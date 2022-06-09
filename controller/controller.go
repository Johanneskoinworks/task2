package controller

import (
	"fmt"
	"net/http"
	"task2/database"
	"task2/models"

	bv "github.com/koinworks/asgard-bivrost/service"
)

func PingHandlerCreateProduct(ctx *bv.Context) bv.Result {

	db := database.GetDB()
	product := models.Product{}
	errCtx := ctx.BodyJSONBind(&product)
	if errCtx != nil {
		return ctx.JSONResponse(http.StatusOK, bv.ResponseBody{
			Message: map[string]string{
				"status": "Error",
				"error":  errCtx.Error(),
			},
			Data: ctx.BodyJSONBind(&product),
		})
	}
	errs := db.Debug().Create(&product).Error
	if errs != nil {
		return ctx.JSONResponse(http.StatusOK, bv.ResponseBody{
			Message: map[string]string{
				"status": "Error",
				"error":  errs.Error(),
			},
			Data: product,
		})
	}
	return ctx.JSONResponse(http.StatusOK, bv.ResponseBody{
		Message: map[string]string{
			"en": "Data added successfully",
			"id": "Data berhasil ditambahkan",
		},
		Data: product,
	})

}
func PingHandlerCreateOrder(ctx *bv.Context) bv.Result {
	db := database.GetDB()
	orders := models.Order{}
	errCtx := ctx.BodyJSONBind(&orders)
	if errCtx != nil {
		return ctx.JSONResponse(http.StatusOK, bv.ResponseBody{
			Message: map[string]string{
				"status": "Error",
				"error":  errCtx.Error(),
			},
			Data: ctx.BodyJSONBind(&orders),
		})
	}
	errs := db.Debug().Create(&orders).Error
	if errs != nil {
		return ctx.JSONResponse(http.StatusOK, bv.ResponseBody{
			Message: map[string]string{
				"status": "Error",
				"error":  errs.Error(),
			},
			Data: orders,
		})
	}
	return ctx.JSONResponse(http.StatusOK, bv.ResponseBody{
		Message: map[string]string{
			"en": "Data added successfully",
			"id": "Data berhasil ditambahkan",
		},
		Data: orders,
	})

}
func PingHandlerList(ctx *bv.Context) bv.Result {
	db := database.GetDB()
	product := []models.Product{}
	result := db.Find(&product).Error
	fmt.Println(result)
	return ctx.JSONResponse(http.StatusOK, bv.ResponseBody{
		Data: product,
	})

}
func PingHandlerOrder(ctx *bv.Context) bv.Result {
	db := database.GetDB()
	orders := []models.Order{}
	result := db.Find(&orders).Error
	fmt.Println(result)
	return ctx.JSONResponse(http.StatusOK, bv.ResponseBody{
		Data: orders,
	})

}
