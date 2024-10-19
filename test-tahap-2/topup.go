package test_tahap_2

import (
	"encoding/json"
	"fmt"
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

type Topup struct {
	db *gorm.DB
}

func NewTopupHandler(db *gorm.DB) *Topup {
	return &Topup{db: db}
}

func (r Topup) Endpoint() (method, path string) {
	method = echo.POST
	path = "/topup"
	return
}

func (r Topup) Handler(ctx echo.Context) error {
	logger := ctx.Get("logger").(*log.Entry)

	handler.AuthenticationHandler(ctx)

	tokenString := ctx.Request().Header.Get("Authorization")
	tokenString = tokenString[len("Bearer "):]
	phoneNumber, err := util.ExtractPhoneNumber(tokenString)
	if err != nil {
		return err
	}

	var request *request.Topup

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

	resp, err := usecase.NewUsecase(r.db, logger).Topup(request)

	if err != nil {
		return ctx.JSON(http.StatusOK, response.Response{Message: err.Error()})
	}

	responseRegister := response.Response{
		Status: "SUCCESS",
		Result: resp,
	}

	return ctx.JSON(http.StatusOK, responseRegister)
}

func InitTopupRequest(ctx echo.Context) (request *request.Topup) {

	requestBody := make(map[string]interface{})
	err := json.NewDecoder(ctx.Request().Body).Decode(&requestBody)
	if err != nil {
		log.Error(err.Error())
	}
	req, e := json.Marshal(requestBody)
	fmt.Println("errorrr", e)

	json.Unmarshal(req, &request)

	return
}
