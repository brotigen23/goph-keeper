package login

// Message to change page from root model
type LoginSuccessMsg struct {
	Username string
}

// Message to check server connection
type PingServerErr error
