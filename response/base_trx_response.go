package response

type BaseTrxResponse struct {
	Amount        int    `json:"amount,omitempty"`
	Remarks       string `json:"remarks,omitempty"`
	BalanceBefore int    `json:"balance_before"`
	BalanceAfter  int    `json:"balance_after"`
	CreatedDate   string `json:"created_date,omitempty"`
}
