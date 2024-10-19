package usecase

import (
	"errors"
	"github.com/jinzhu/gorm"
	"test-mnc/entity"
	"test-mnc/request"
	"test-mnc/response"
	"test-mnc/util"
)

func (u Usecase) Login(request *request.RegisterRequest) (token response.Login, e error) {
	var user entity.User
	result := u.db.Where("phone_number = ? AND pin = ?", request.PhoneNumber, request.Pin).First(&user)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return token, errors.New("Phone Number and PIN doesn't match.")
		} else {
			u.logger.Error("Error querying user:", result.Error)
		}
	}
	return util.CreateToken(request.PhoneNumber)
}
