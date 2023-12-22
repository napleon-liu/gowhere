package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"travel_booking/databases"
	"travel_booking/handlers"
	"travel_booking/middlewares"
)

func main() {
	if err := databases.Init(); err != nil {
		log.Println("Mysql init error")
	}

	r := gin.Default()

	r.Use(middlewares.Cors())

	router := r.Group("/travel_booking/v1")
	{
		router.GET("/customer", handlers.GetCustomer)
		router.POST("/customer", handlers.CreateCustomer)
		router.PUT("/customer", handlers.UpdateCustomer)
		router.DELETE("/customer", handlers.DeleteCustomer)
		router.GET("/route", handlers.GetRoute)

		router.GET("/bus", handlers.GetBus)
		router.POST("/bus", handlers.CreateBus)
		router.PUT("/bus", handlers.UpdateBus)
		router.DELETE("/bus", handlers.DeleteBus)

		router.GET("/flight", handlers.GetFlight)
		router.POST("/flight", handlers.CreateFlight)
		router.PUT("/flight", handlers.UpdateFlight)
		router.DELETE("/flight", handlers.DeleteFlight)

		router.GET("/hotel", handlers.GetHotel)
		router.POST("/hotel", handlers.CreateHotel)
		router.PUT("/hotel", handlers.UpdateHotel)
		router.DELETE("/hotel", handlers.DeleteHotel)

		router.POST("/booking", handlers.CreateBooking)
	}

	_ = r.Run(":80")

}
