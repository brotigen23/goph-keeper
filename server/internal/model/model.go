package model

type Model interface {
	GetID() int
}

func (d AccountData) GetID() int { return d.ID }
func (d TextData) GetID() int    { return d.ID }
func (d BinaryData) GetID() int  { return d.ID }
func (d CardData) GetID() int    { return d.ID }
func (d Metadata) GetID() int    { return d.ID }

type WithMetadata[T Model] struct {
	Data     T
	Metadata Metadata
}
