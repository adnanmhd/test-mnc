package usecase

import (
	"errors"
	"github.com/google/uuid"
	"test-mnc/entity"
	"test-mnc/request"
	"test-mnc/response"
	"time"
)

func (u Usecase) Topup(request *request.Topup) (resp response.Topup, e error) {
	topup := setTopup(request)
	result := u.db.Create(topup)

	if result.Error != nil {
		u.logger.Error(result.Error)
		return resp, errors.New("Top Up Failed")

	}
	return setResponse(topup), nil
}

func setResponse(topup *entity.Topup) (resp response.Topup) {
	resp.TopUpId = topup.TopupId
	resp.AmountTopUp = topup.AmountTopup
	resp.BalanceBefore = topup.BalanceBefore
	resp.BalanceAfter = topup.BalanceAfter
	resp.CreatedDate = topup.CreatedDate
	return
}

func setTopup(request *request.Topup) *entity.Topup {
	topup := new(entity.Topup)
	topup.TopupId = uuid.New().String()
	topup.AmountTopup = request.Amount
	topup.BalanceAfter = request.Amount
	topup.BalanceBefore = 0
	topup.CreatedDate = time.Now().String()
	topup.PhoneNumber = request.PhoneNumber
	return topup
}
