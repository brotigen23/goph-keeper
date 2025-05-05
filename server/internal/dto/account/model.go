package account

type Model struct {
	Login    string `json:"login" example:"user"`
	Password string `json:"password" example:"user"`

	Metadata string `json:"metadata" example:"metadata"`
}
