package test_tahap_2

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"net/http"
	"test-mnc/request"
	"test-mnc/response"
	"test-mnc/usecase"
)

type Register struct {
	db *gorm.DB
}

func NewRegisterHandler(db *gorm.DB) *Register {
	return &Register{db: db}
}

func (R Register) Endpoint() (method, path string) {
	method = echo.POST
	path = "/register"
	return
}

func (R Register) Handler(ctx echo.Context) error {
	logger := ctx.Get("logger").(*log.Entry)

	request := InitRequest(ctx)

	defer func() {
		if panicMessage := recover(); panicMessage != nil {
			logger.Error("Panic captured: ", panicMessage)
			ctx.JSON(http.StatusInternalServerError, "")
		}
	}()

	err := usecase.NewUsecase(R.db, logger).AddUser(request)

	if err != nil {
		return ctx.JSON(http.StatusOK, response.Response{Message: err.Error()})
	}

	responseRegister := response.Response{
		Status: "SUCCESS",
		Result: request,
	}

	return ctx.JSON(http.StatusOK, responseRegister)
}

func InitRequest(ctx echo.Context) (request *request.RegisterRequest) {

	requestBody := make(map[string]interface{})
	err := json.NewDecoder(ctx.Request().Body).Decode(&requestBody)
	if err != nil {
		log.Error(err.Error())
	}
	req, _ := json.Marshal(requestBody)

	json.Unmarshal(req, &request)

	return
}
