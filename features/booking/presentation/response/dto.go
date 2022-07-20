package response

type Booking struct {
	Token       string `json:"token"`
	RedirectURL string `json:"redirect_url"`
}
