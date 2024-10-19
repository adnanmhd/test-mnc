package request

type RegisterRequest struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
	Pin         string `json:"pin"`
	UserId      string `json:"user_id"`
	CreatedDate string `json:"created_date"`
}
