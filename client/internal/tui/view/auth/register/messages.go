package register

// Message to change page from root model
type SignUpSuccessMsg struct {
	Username string
}

// Message to check server connection
type PingServerErr error
