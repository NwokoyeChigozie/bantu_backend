package router

import (
	// "net/http"

	"net/http"

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

	ApiVersion := "v1"
	FootprintUrl(r, validate, ApiVersion)

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

func FootprintUrl(r *gin.Engine, validate *validator.Validate, ApiVersion string) {
	panic("unimplemented")
}
