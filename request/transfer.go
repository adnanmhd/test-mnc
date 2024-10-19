package request

type Transfer struct {
	TargetUser string `json:"target_user"`
	BaseTrxRequest
}
