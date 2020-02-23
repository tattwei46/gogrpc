package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-grpc/api"
	"go-grpc/constant"
	"go-grpc/user"
	"log"
)

func main() {

	go func() {
		s, _ := user.NewService()
		log.Fatal(user.ListenGRPC(s, constant.UserPort))
	}()

	router := setupRouter()
	if err := router.Run(fmt.Sprintf("%s:%d", constant.HostURL, constant.HostPort)); err != nil {
		panic(err.Error())
	}
}

func setupRouter() *gin.Engine {
	gin.DisableConsoleColor()
	r := gin.Default()
	api.CreateUserRoutes(r)
	return r
}
