package test_tahap_2

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"net/http"
	"test-mnc/handler"
	"test-mnc/request"
	"test-mnc/response"
	"test-mnc/usecase"
	"test-mnc/util"
)

type Payment struct {
	db *gorm.DB
}

func NewPaymentHandler(db *gorm.DB) *Payment {
	return &Payment{db: db}
}

func (p Payment) Endpoint() (method, path string) {
	method = echo.POST
	path = "/pay"
	return
}

func (p Payment) Handler(ctx echo.Context) error {
	logger := ctx.Get("logger").(*log.Entry)

	handler.AuthenticationHandler(ctx)

	tokenString := ctx.Request().Header.Get("Authorization")
	tokenString = tokenString[len("Bearer "):]
	phoneNumber, err := util.ExtractPhoneNumber(tokenString)
	if err != nil {
		return err
	}

	var request *request.Payment

	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"message": "Invalid input"})
	}

	defer func() {
		if panicMessage := recover(); panicMessage != nil {
			logger.Error("Panic captured: ", panicMessage)
			ctx.JSON(http.StatusInternalServerError, "")
		}
	}()

	request.PhoneNumber = phoneNumber

	resp, err := usecase.NewUsecase(p.db, logger).Payment(request)

	if err != nil {
		return ctx.JSON(http.StatusOK, response.Response{Message: err.Error()})
	}

	responseRegister := response.Response{
		Status: "SUCCESS",
		Result: resp,
	}

	return ctx.JSON(http.StatusOK, responseRegister)
}
