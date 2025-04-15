package dto

type Metadata struct {
	BaseData
	Data string
}

type MetadataPut struct {
	ID       int     `json:"id"`
	Metadata *string `json:"data,omitempty"`
}

type MetadataDelete struct {
	ID *int `json:"id,omitempty"`
}
