package login

type Request struct {
	Login    string `json:"login" example:"user"`
	Password string `json:"password" example:"pass"`
} //@name LoginRequest
