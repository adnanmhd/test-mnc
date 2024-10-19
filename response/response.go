package response

type Response struct {
	Status  string      `json:"status,omitempty"`
	Message string      `json:"message,omitempty"`
	Result  interface{} `json:"result,omitempty"`
}
