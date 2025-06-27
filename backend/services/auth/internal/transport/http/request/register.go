package request

type Register struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Login struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
