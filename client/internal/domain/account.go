package domain

type AccountData struct {
	BaseData
	Login    string `json:"login" table:"true"`
	Password string `json:"password" table:"true"`
	Metadata `json:"metadata"`
}
