package controller

import (
	"net/http"
	"time"

	"github.com/UfiairENE/bantu_solution/pkg/router"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Controller struct {
	Validate *validator.Validate
}

type form struct {
	current_time time.Time
	ip_address   string
	device_info  string
	browser_type string
	longitude    uint
	latitude     uint
	city         string
	country      string
}

func AddFootprint(c *gin.Context) {
	db := router.ConnectToDB()
	req := form{}
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadGateway, struct{ message string }{message: err.Error()})
		return
	}
	currenttime := req.current_time
	ipaddress := req.ip_address
	deviceinfo := req.device_info
	browsertype := req.browser_type
	longitude := req.longitude
	latitude := req.latitude
	city := req.city
	country := req.country
	tx := db.Exec("INSERT INTO Footprint (currenttime, ipaddress, deviceinfo, latitude, longitude, browsertype, city, country) VALUES(?, ?, ?, ?, ?, ?, ?, ?)", currenttime, ipaddress, deviceinfo, latitude, longitude, city, country, browsertype)
	if tx.Error != nil {
		panic(tx.Error)
	}
	c.JSON(http.StatusOK, struct{ message string }{message: "successfuly added to database"})

}
