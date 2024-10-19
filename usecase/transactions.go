package usecase

import (
	"test-mnc/entity"
)

func (u Usecase) Transactions(phoneNumber string) ([]interface{}, error) {

	list := make([]interface{}, 0)
	list = append(list, u.getTopUpTransactions(phoneNumber))
	list = append(list, u.getPaymentTransactions(phoneNumber))
	list = append(list, u.getTransferTransactions(phoneNumber))

	return list, nil
}

func (u Usecase) getTopUpTransactions(phoneNumber string) []entity.Topup {
	list := make([]entity.Topup, 0)

	rows, e := u.db.Model(&entity.Topup{}).Where("phone_number = ?", phoneNumber).Rows()
	if e != nil {
		u.logger.Error(e)
	}
	defer rows.Close()

	for rows.Next() {
		var topup entity.Topup
		u.db.ScanRows(rows, &topup)
		list = append(list, topup)
	}
	return list
}

func (u Usecase) getPaymentTransactions(phoneNumber string) []entity.Payment {
	list := make([]entity.Payment, 0)

	rows, e := u.db.Model(&entity.Payment{}).Where("phone_number = ?", phoneNumber).Rows()
	if e != nil {
		u.logger.Error(e)
	}
	defer rows.Close()

	for rows.Next() {
		var payment entity.Payment
		u.db.ScanRows(rows, &payment)
		list = append(list, payment)
	}
	return list
}

func (u Usecase) getTransferTransactions(phoneNumber string) []entity.Transfer {
	list := make([]entity.Transfer, 0)

	rows, e := u.db.Model(&entity.Transfer{}).Where("phone_number = ?", phoneNumber).Rows()
	if e != nil {
		u.logger.Error(e)
	}
	defer rows.Close()

	for rows.Next() {
		var transfer entity.Transfer
		u.db.ScanRows(rows, &transfer)
		list = append(list, transfer)
	}
	return list
}
