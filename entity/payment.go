package entity

type Payment struct {
	PaymentId     string `gorm:"column:payment_id;primaryKey"`
	Amount        int    `gorm:"column:amount"`
	Remarks       string `gorm:"column:remarks"`
	BalanceBefore int    `gorm:"column:balance_before"`
	BalanceAfter  int    `gorm:"column:balance_after"`
	PhoneNumber   string `gorm:"column:phone_number"`
	CreatedDate   string `gorm:"column:created_date"`
}

func (t Payment) TableName() string {
	return "payment"
}
