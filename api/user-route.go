package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

var baseURL = "user"

func route(router *gin.Engine) *gin.RouterGroup {
	return router.Group(baseURL)
}

func CreateUserRoutes(router *gin.Engine) {

	handler, err := NewHandler()
	if err != nil {
		fmt.Printf("an error occurs when creating user route")
	}
	r := route(router)

	r.GET("/", handler.Get)
	r.POST("/", handler.Create)

}
