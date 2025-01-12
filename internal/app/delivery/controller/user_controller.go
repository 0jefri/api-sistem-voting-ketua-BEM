package controller

import (
	"errors"
	"net/http"
	"time"

	"github.com/api-voting/internal/app/service"
	"github.com/api-voting/internal/model"
	"github.com/api-voting/internal/model/dto"
	"github.com/api-voting/utils/common"
	"github.com/api-voting/utils/exception"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	service service.UserService
	// auth        service.AuthService
}

func NewUserController(service service.UserService) *UserController {
	return &UserController{
		service: service,
		// auth:        auth,
	}
}

func (ctr *UserController) Registration(c *gin.Context) {
	payload := model.User{}

	payload.ID = common.GenerateUUID()
	payload.RegistrationDate = time.Now()
	payload.LastLogin = time.Now()

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"status":  exception.StatusBadRequest,
			"message": exception.FieldErrors(err),
		})
		return
	}

	_, err := ctr.service.RegisterNewUser(&payload)

	if err != nil {
		if errors.Is(err, exception.ErrFailedCreate) {
			c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Status:  exception.StatusInternalServer,
				Message: exception.ErrFailedCreate.Error(),
			})
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Status:  exception.StatusInternalServer,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, dto.Response{
		Code:    http.StatusCreated,
		Status:  exception.StatusSuccess,
		Message: "Register Successful",
		Data:    payload,
	})
}
