package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"travel_booking/services"
	"travel_booking/typ"
)

func CreateHotel(ctx *gin.Context) {
	var createHotel typ.CreateHotelReq
	_ = ctx.ShouldBindJSON(&createHotel)
	err := services.CreateHotel(createHotel)
	ctx.JSON(http.StatusOK, &typ.T{
		MetaData: &typ.MetaData{
			Code: http.StatusOK,
			Msg:  err.Error(),
		},
	})
}

func DeleteHotel(ctx *gin.Context) {
	var deleteHotel typ.DeleteHotelReq
	_ = ctx.ShouldBindJSON(&deleteHotel)
	err := services.DeleteHotel(deleteHotel)
	ctx.JSON(http.StatusOK, &typ.T{
		MetaData: &typ.MetaData{
			Code: http.StatusOK,
			Msg:  err.Error(),
		},
	})
}

func UpdateHotel(ctx *gin.Context) {
	var updateHotel typ.UpdateHotelReq
	_ = ctx.ShouldBindJSON(&updateHotel)
	err := services.UpdateHotel(updateHotel)
	ctx.JSON(http.StatusOK, &typ.T{
		MetaData: &typ.MetaData{
			Code: http.StatusOK,
			Msg:  err.Error(),
		},
	})
}

func GetHotel(ctx *gin.Context) {
	res, err := services.GetHotel()
	if err != nil {
		ctx.JSON(http.StatusOK, &typ.T{
			MetaData: &typ.MetaData{
				Code: http.StatusOK,
				Msg:  err.Error(),
			},
		})
		return
	}
	ctx.JSON(http.StatusOK, &typ.T{
		MetaData: &typ.MetaData{
			Code: http.StatusOK,
			Msg:  "",
		},
		Result: res,
	})
}
