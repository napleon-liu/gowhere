package handlers

import (
	"github.com/gin-gonic/gin"
	"travel_booking/services"
	"travel_booking/typ"
)

func CreateFlight(c *gin.Context) {
	var flight typ.CreateFlightReq
	if err := c.ShouldBindJSON(&flight); err != nil {
		c.JSON(200, typ.T{
			MetaData: &typ.MetaData{
				Code: 400,
				Msg:  err.Error(),
			},
		})
		return
	}

	err := services.CreateFlight(&flight)
	c.JSON(200, typ.T{
		MetaData: &typ.MetaData{
			Code: 200,
			Msg:  err.Error(),
		},
	})
}

func UpdateFlight(c *gin.Context) {
	var flight typ.UpdateFlightReq
	if err := c.ShouldBindJSON(&flight); err != nil {
		c.JSON(200, typ.T{
			MetaData: &typ.MetaData{
				Code: 400,
				Msg:  err.Error(),
			},
		})
		return
	}

	err := services.UpdateFlight(&flight)
	c.JSON(200, typ.T{
		MetaData: &typ.MetaData{
			Code: 200,
			Msg:  err.Error(),
		},
	})
}

func DeleteFlight(c *gin.Context) {
	var flight typ.DeleteFlightReq
	if err := c.ShouldBindJSON(&flight); err != nil {
		c.JSON(200, typ.T{
			MetaData: &typ.MetaData{
				Code: 400,
				Msg:  err.Error(),
			},
		})
		return
	}

	err := services.DeleteFlight(&flight)
	c.JSON(200, typ.T{
		MetaData: &typ.MetaData{
			Code: 200,
			Msg:  err.Error(),
		},
	})
}

func GetFlight(c *gin.Context) {
	flights, err := services.GetFlight()
	if err != nil {
		return
	}
	c.JSON(200, typ.T{
		MetaData: &typ.MetaData{
			Code: 200,
			Msg:  "",
		},
		Result: flights,
	})
}
