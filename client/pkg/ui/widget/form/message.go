package form

type SubmitFormMsg[T any] struct {
	Data T
}

type SubmitFormErrorMsg struct {
	Field string
	Error error
}

type CloseMsg struct{}
