package response

type Payment struct {
	PaymentId string `json:"payment_id"`
	BaseTrxResponse
}
