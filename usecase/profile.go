package usecase

import (
	"errors"
	"github.com/jinzhu/gorm"
	"test-mnc/entity"
	"test-mnc/request"
	"test-mnc/response"
)

func (u Usecase) UpdateProfile(request *request.Profile) (response.Profile, error) {
	var profile entity.User
	result := u.db.Where("phone_number = ? ", request.PhoneNumber).First(&profile)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return response.Profile{}, errors.New("User not found")
		} else {
			u.logger.Error("Error querying user:", result.Error)
			return response.Profile{}, result.Error
		}
	}
	profile = updateUserData(request, profile)
	if err := u.db.Model(&entity.User{}).Where("user_id = ?", profile.UserId).Updates(profile).Error; err != nil {
		u.logger.Error("Failed to save user:", err.Error())
		return response.Profile{}, errors.New("Failed update profile")
	}

	return setProfileResponse(profile), nil
}
func updateUserData(request *request.Profile, user entity.User) entity.User {
	user.FirstName = request.FirstName
	user.LastName = request.LastName
	user.Address = request.Address
	return user
}

func setProfileResponse(user entity.User) response.Profile {
	var resp response.Profile
	resp.UserId = user.UserId
	resp.FirstName = user.FirstName
	resp.LastName = user.LastName
	resp.Address = user.Address
	return resp
}
