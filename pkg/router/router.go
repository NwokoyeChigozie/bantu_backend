package router

import (
	"net/http"
	"time"

	"github.com/UfiairENE/bantu_solution/pkg/router/connection"
    "github.com/UfiairENE/bantu_solution/pkg/middleware"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Setup(validate *validator.Validate) *gin.Engine {
	r := gin.Default()

	// Middlewares
	r.Use(gin.Logger())
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.CORS())
	r.Use(gzip.Gzip(gzip.DefaultCompression))

	// ApiVersion := "v1"
	// FootprintUrl(r, validate, ApiVersion)

	type form struct {
		Current_time time.Time
		Ip_address   string
		Device_info  string
		Browser_type string
		Longitude    uint
		Latitude     uint
		City         string
		Country      string
	}

	r.POST("/add", func(c *gin.Context) {
		db := connection.ConnectToDB()
		req := form{}
		err := c.ShouldBind(&req)
		if err != nil {
			c.JSON(http.StatusBadGateway, struct{ message string }{message: err.Error()})
			return
		}
		currenttime := req.Current_time
		ipaddress := req.Ip_address
		deviceinfo := req.Device_info
		browsertype := req.Browser_type
		longitude := req.Longitude
		latitude := req.Latitude
		city := req.City
		country := req.Country
		tx := db.Exec("INSERT INTO Footprint (currenttime, ipaddress, deviceinfo, latitude, longitude, browsertype, city, country) VALUES(?, ?, ?, ?, ?, ?, ?, ?)", currenttime, ipaddress, deviceinfo, latitude, longitude, city, country, browsertype)
		if tx.Error != nil {
			panic(tx.Error)
		}
		c.JSON(http.StatusOK, struct{ message string }{message: "successfuly added to database"})

	})

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"name":    "Not Found",
			"message": "Page not found.",
			"code":    400,
			"status":  http.StatusNotFound,
		})
	})

	return r
}

