package entity

type User struct {
	UserId      string `gorm:"column:user_id;primaryKey"`
	FirstName   string `gorm:"column:first_name"`
	LastName    string `gorm:"column:last_name"`
	PhoneNumber string `gorm:"column:phone_number;unique"`
	Address     string `gorm:"column:address"`
	Pin         string `gorm:"column:pin"`
	CreatedDate string `gorm:"column:created_date"`
}

func (u User) TableName() string {
	return "user"
}
