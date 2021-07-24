package printer

type RegisterRequest struct {
	Uuid string `json:"uuid"`
}

type RegisterResponse struct {
	Uuid string `json:"uuid"`
}
