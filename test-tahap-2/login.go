package test_tahap_2

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"net/http"
	"test-mnc/response"
	"test-mnc/usecase"
)

type Login struct {
	db *gorm.DB
}

func NewLoginHandler(db *gorm.DB) *Login {
	return &Login{db: db}
}

func (L Login) Endpoint() (method, path string) {
	method = echo.POST
	path = "/login"
	return
}

func (L Login) Handler(ctx echo.Context) error {
	logger := ctx.Get("logger").(*log.Entry)

	request := InitRequest(ctx)

	defer func() {
		if panicMessage := recover(); panicMessage != nil {
			logger.Error("Panic captured: ", panicMessage)
			ctx.JSON(http.StatusInternalServerError, "")
		}
	}()

	token, err := usecase.NewUsecase(L.db, logger).Login(request)

	if err != nil {
		return ctx.JSON(http.StatusOK, response.Response{Message: err.Error()})
	}

	responseRegister := response.Response{
		Status: "SUCCESS",
		Result: token,
	}

	return ctx.JSON(http.StatusOK, responseRegister)
}
