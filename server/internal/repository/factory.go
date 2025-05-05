package repository

type Factory interface {
	NewUserRepository() User
	NewAccountRepository() AccountsData
}
