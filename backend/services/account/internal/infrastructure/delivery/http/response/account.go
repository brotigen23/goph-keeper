package response

type Account struct {
	ID       int
	Login    string
	Password string
}

type AccountPost struct {
	Account
}
