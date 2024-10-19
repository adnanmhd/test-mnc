package request

type BaseTrxRequest struct {
	Amount      int    `json:"amount"`
	Remarks     string `json:"remarks,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
}
