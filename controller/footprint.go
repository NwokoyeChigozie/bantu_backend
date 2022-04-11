package controller

// import (
// 	"net/http"
// 	"time"

// 	"github.com/UfiairENE/bantu_solution/pkg/router"
// 	"github.com/gin-gonic/gin"
// 	"github.com/go-playground/validator/v10"
// )

// type Controller struct {
// 	Validate *validator.Validate
// }

// type form struct {
// 	Current_time time.Time
// 	Ip_address   string
// 	Device_info  string
// 	Browser_type string
// 	Longitude    uint
// 	Latitude     uint
// 	City         string
// 	Country      string
// }

// func AddFootprint(c *gin.Context) {
// 	db := router.ConnectToDB()
// 	req := form{}
// 	err := c.ShouldBind(&req)
// 	if err != nil {
// 		c.JSON(http.StatusBadGateway, struct{ message string }{message: err.Error()})
// 		return
// 	}
// 	currenttime := req.Current_time
// 	ipaddress := req.Ip_address
// 	deviceinfo := req.Device_info
// 	browsertype := req.Browser_type
// 	longitude := req.Longitude
// 	latitude := req.Latitude
// 	city := req.City
// 	country := req.Country
// 	tx := db.Exec("INSERT INTO Footprint (currenttime, ipaddress, deviceinfo, latitude, longitude, browsertype, city, country) VALUES(?, ?, ?, ?, ?, ?, ?, ?)", currenttime, ipaddress, deviceinfo, latitude, longitude, city, country, browsertype)
// 	if tx.Error != nil {
// 		panic(tx.Error)
// 	}
// 	c.JSON(http.StatusOK, struct{ message string }{message: "successfuly added to database"})

// }
