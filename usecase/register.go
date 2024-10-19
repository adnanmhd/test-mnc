package usecase

import (
	"errors"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"test-mnc/entity"
	"test-mnc/request"
	"time"
)

type Usecase struct {
	db     *gorm.DB
	logger *log.Entry
}

func NewUsecase(db *gorm.DB, logger *log.Entry) *Usecase {
	return &Usecase{db: db, logger: logger}
}

func (u Usecase) AddUser(request *request.RegisterRequest) error {

	addUser := u.db.Create(setNewUser(request))
	if addUser.Error != nil {
		u.logger.Error("Failed to save user:", addUser.Error)
		if addUser.Error.Error() == "pq: duplicate key value violates unique constraint \"user_un\"" {
			return errors.New("Phone Number already registered")
		}
	}
	return nil
}
func setNewUser(registerRequest *request.RegisterRequest) *entity.User {
	user := new(entity.User)
	user.UserId = uuid.New().String()
	user.FirstName = registerRequest.FirstName
	user.LastName = registerRequest.LastName
	user.PhoneNumber = registerRequest.PhoneNumber
	user.Address = registerRequest.Address
	user.Pin = registerRequest.Pin
	user.CreatedDate = time.Now().String()
	registerRequest.CreatedDate = user.CreatedDate
	registerRequest.UserId = user.UserId
	return user
}
