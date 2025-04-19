package form

type SubmitFormMsg[T any] struct {
	Data *T
}
