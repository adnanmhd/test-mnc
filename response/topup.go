package response

type Topup struct {
	TopUpId     string `json:"top_up_id"`
	AmountTopUp int    `json:"amount_top_up"`
	BaseTrxResponse
}
