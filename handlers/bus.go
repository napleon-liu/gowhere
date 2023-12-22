package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"travel_booking/services"
	"travel_booking/typ"
)

func CreateBus(ctx *gin.Context) {
	var createBus typ.CreateBusReq
	_ = ctx.ShouldBindJSON(&createBus)
	err := services.CreateBus(createBus)
	ctx.JSON(http.StatusOK, &typ.T{
		MetaData: &typ.MetaData{
			Code: http.StatusOK,
			Msg:  err.Error(),
		},
	})
}

func DeleteBus(ctx *gin.Context) {
	var deleteBus typ.DeleteBusReq
	_ = ctx.ShouldBindJSON(&deleteBus)
	err := services.DeleteBus(deleteBus)
	ctx.JSON(http.StatusOK, &typ.T{
		MetaData: &typ.MetaData{
			Code: http.StatusOK,
			Msg:  err.Error(),
		},
	})
}

func UpdateBus(ctx *gin.Context) {
	var updateBus typ.UpdateBusReq
	_ = ctx.ShouldBindJSON(&updateBus)
	err := services.UpdateBus(updateBus)
	ctx.JSON(http.StatusOK, &typ.T{
		MetaData: &typ.MetaData{
			Code: http.StatusOK,
			Msg:  err.Error(),
		},
	})
}

func GetBus(ctx *gin.Context) {
	bus, err := services.GetBus()
	ctx.JSON(http.StatusOK, &typ.T{
		MetaData: &typ.MetaData{
			Code: http.StatusOK,
			Msg:  err.Error(),
		},
		Result: bus,
	})
}
