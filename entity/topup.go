package entity

type Topup struct {
	TopupId       string `gorm:"column:top_up_id;primaryKey"`
	AmountTopup   int    `gorm:"column:amount_top_up"`
	BalanceBefore int    `gorm:"column:balance_before"`
	BalanceAfter  int    `gorm:"column:balance_after"`
	PhoneNumber   string `gorm:"column:phone_number"`
	CreatedDate   string `gorm:"column:created_date"`
}

func (t Topup) TableName() string {
	return "topup"
}
