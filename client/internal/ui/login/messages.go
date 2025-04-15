package login

// Message to change page from root model
type LoginSuccessMgs struct{}

type ServerResponse struct {
	msg        string
	statusCode int
}

// Message to check server connection
type PingServerErr error
