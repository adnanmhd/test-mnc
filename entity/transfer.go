package entity

type Transfer struct {
	TransferId    string `gorm:"column:transfer_id;primaryKey"`
	Amount        int    `gorm:"column:amount"`
	Remarks       string `gorm:"column:remarks"`
	BalanceBefore int    `gorm:"column:balance_before"`
	BalanceAfter  int    `gorm:"column:balance_after"`
	TargetUser    string `gorm:"column:target_user"`
	PhoneNumber   string `gorm:"column:phone_number"`
	CreatedDate   string `gorm:"column:created_date"`
}

func (t Transfer) TableName() string {
	return "transfer"
}
