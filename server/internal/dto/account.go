package dto

type AccountData struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type AccountPost struct {
	BaseData
	AccountData
}

type AccountPut struct {
	BaseData
	Login    *string `json:"login,omitempty"`
	Password *string `json:"password,omitempty"`
}

type AccountsGet struct {
	BaseData
	AccountData
	Metadata Metadata `json:"metadata"`
}

type AccountDelete struct {
	ID *int `json:"id,omitempty"`
}
