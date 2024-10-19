package test_tahap_2

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"net/http"
	"test-mnc/handler"
	"test-mnc/response"
	"test-mnc/usecase"
	"test-mnc/util"
)

type Transactions struct {
	db *gorm.DB
}

func NewTransactionsHandler(db *gorm.DB) *Transactions {
	return &Transactions{db: db}
}

func (t Transactions) Endpoint() (method, path string) {
	method = echo.POST
	path = "/transactions"
	return
}

func (t Transactions) Handler(ctx echo.Context) error {
	logger := ctx.Get("logger").(*log.Entry)

	handler.AuthenticationHandler(ctx)

	tokenString := ctx.Request().Header.Get("Authorization")
	tokenString = tokenString[len("Bearer "):]
	phoneNumber, err := util.ExtractPhoneNumber(tokenString)

	defer func() {
		if panicMessage := recover(); panicMessage != nil {
			logger.Error("Panic captured: ", panicMessage)
			ctx.JSON(http.StatusInternalServerError, "")
		}
	}()

	resp, err := usecase.NewUsecase(t.db, logger).Transactions(phoneNumber)

	if err != nil {
		return ctx.JSON(http.StatusOK, response.Response{Message: err.Error()})
	}

	responseRegister := response.Response{
		Status: "SUCCESS",
		Result: resp,
	}

	return ctx.JSON(http.StatusOK, responseRegister)
}
