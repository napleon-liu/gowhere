package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"travel_booking/services"
	"travel_booking/typ"
)

func CreateCustomer(ctx *gin.Context) {
	var createCustomer typ.CreateCustomerReq
	_ = ctx.ShouldBindJSON(&createCustomer)
	err := services.CreateCustomer(createCustomer)
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
			Msg:  err.Error(),
		},
	})
}

func DeleteCustomer(ctx *gin.Context) {
	var deleteCustomer typ.DeleteCustomerReq
	_ = ctx.ShouldBindJSON(&deleteCustomer)
	err := services.DeleteCustomer(deleteCustomer)
	ctx.JSON(http.StatusOK, &typ.T{
		MetaData: &typ.MetaData{
			Code: http.StatusOK,
			Msg:  err.Error(),
		},
	})
}

func UpdateCustomer(ctx *gin.Context) {
	var updateCustomer typ.UpdateCustomerReq
	_ = ctx.ShouldBindJSON(&updateCustomer)
	err := services.UpdateCustomer(updateCustomer)
	ctx.JSON(http.StatusOK, &typ.T{
		MetaData: &typ.MetaData{
			Code: http.StatusOK,
			Msg:  err.Error(),
		},
	})
}

func GetCustomer(ctx *gin.Context) {
	res, err := services.GetCustomer()
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

func GetRoute(ctx *gin.Context) {

}
