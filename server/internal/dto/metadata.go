package dto

type Metadata struct {
	BaseData
	Data string
}

type MetadataPut struct {
	ID   int     `json:"id"`
	Data *string `json:"data,omitempty"`
}

type MetadataDelete struct {
	ID *int `json:"id,omitempty"`
}
