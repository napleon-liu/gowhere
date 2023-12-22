package handlers

import (
	"github.com/gin-gonic/gin"
	"travel_booking/services"
	"travel_booking/typ"
)

func CreateBooking(c *gin.Context) {
	var req typ.CreateReservationReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message": "invalid request",
		})
		return
	}
	if err := services.CreateReservation(req); err != nil {
		c.JSON(500, gin.H{
			"message": "internal error",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "success",
	})

}
