package usecase

import (
	"errors"
	"github.com/google/uuid"
	"test-mnc/entity"
	"test-mnc/request"
	"test-mnc/response"
	"time"
)

func (u Usecase) Payment(request *request.Payment) (resp response.Payment, e error) {
	payment := setPayment(request)
	result := u.db.Create(payment)

	if result.Error != nil {
		u.logger.Error(result.Error)
		return resp, errors.New("payment Failed")

	}
	return setResponsePayment(payment), nil
}

func setResponsePayment(payment *entity.Payment) (resp response.Payment) {
	resp.PaymentId = payment.PaymentId
	resp.Amount = payment.Amount
	resp.BalanceBefore = payment.BalanceBefore
	resp.BalanceAfter = payment.BalanceAfter
	resp.CreatedDate = payment.CreatedDate
	resp.Remarks = payment.Remarks
	return
}

func setPayment(request *request.Payment) *entity.Payment {
	payment := new(entity.Payment)
	payment.PaymentId = uuid.New().String()
	payment.Amount = request.Amount
	payment.BalanceAfter = request.Amount
	payment.BalanceBefore = 0
	payment.CreatedDate = time.Now().String()
	payment.PhoneNumber = request.PhoneNumber
	payment.Remarks = request.Remarks
	return payment
}
