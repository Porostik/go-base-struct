package handler

import (
	"architecture/internal/core"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserService interface {
	GetUserById(context context.Context, id int) core.User
}

type UserHandler struct {
	service UserService
}

func NewUserHandler(service UserService) *UserHandler {
	return &UserHandler{
		service,
	}
}

func (h *UserHandler) Get(c *gin.Context) {
	paramId := c.Param("id")

	if paramId == "" {
		errorResponse(c, http.StatusBadRequest, "missing id param")
	}

	id, err := strconv.Atoi(paramId)

	if err != nil {
		errorResponse(c, http.StatusBadRequest, "missing id param")
	}

	user := h.service.GetUserById(context.Background(), id)

	c.JSON(http.StatusOK, user)
}

func (u *UserHandler) Init(api *gin.RouterGroup) {
	user := api.Group("/user")
	{
		user.GET("/:id", u.Get)
	}
}
