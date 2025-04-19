package domain

type Metadata struct {
	BaseData

	Metadata string `json:"data" table:"true" `
}
