package domain

type AccountData struct {
	BaseData

	Login    string `json:"login" table:"true" form:"true"`
	Password string `json:"password" table:"true" form:"true"`

	Metadata `json:"metadata"`
}

func (d AccountData) GetID() int {
	return d.ID
}
