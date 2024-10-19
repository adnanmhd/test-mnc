package usecase

import (
	"errors"
	"github.com/google/uuid"
	"test-mnc/entity"
	"test-mnc/request"
	"test-mnc/response"
	"time"
)

func (u Usecase) Transfer(request *request.Transfer) (resp response.Transfer, e error) {
	payment := setTransfer(request)
	result := u.db.Create(payment)

	if result.Error != nil {
		u.logger.Error(result.Error)
		return resp, errors.New("payment Failed")

	}
	return setResponseTransfer(payment), nil
}

func setResponseTransfer(payment *entity.Transfer) (resp response.Transfer) {
	resp.TransferId = payment.TransferId
	resp.Amount = payment.Amount
	resp.BalanceBefore = payment.BalanceBefore
	resp.BalanceAfter = payment.BalanceAfter
	resp.CreatedDate = payment.CreatedDate
	resp.Remarks = payment.Remarks
	return
}

func setTransfer(request *request.Transfer) *entity.Transfer {
	transfer := new(entity.Transfer)
	transfer.TransferId = uuid.New().String()
	transfer.Amount = request.Amount
	transfer.BalanceAfter = request.Amount
	transfer.BalanceBefore = 0
	transfer.CreatedDate = time.Now().String()
	transfer.Remarks = request.Remarks
	transfer.TargetUser = request.TargetUser
	transfer.PhoneNumber = request.PhoneNumber
	return transfer
}
