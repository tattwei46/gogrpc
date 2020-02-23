package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-grpc/constant"
	"go-grpc/user"
	"net/http"
)

type Handler struct {
	client *user.Client
}

func NewHandler() (*Handler, error) {
	c, err := user.NewClient(fmt.Sprintf("%s:%d", constant.HostURL, constant.UserPort))
	if err != nil {
		return &Handler{}, err
	}
	return &Handler{client: c}, nil
}

func (h *Handler) Get(c *gin.Context) {
	res, err := h.client.GetUser(c, "1")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"message": res})
}

func (h *Handler) Create(c *gin.Context) {
	var user user.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	err = h.client.CreateUser(c, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
