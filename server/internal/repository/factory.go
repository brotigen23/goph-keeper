package repository

type Factory interface {
	NewUserRepository() User
	NewAccountRepository() Account
	NewTextRepository() Text
	NewBinaryRepository() Binary
	NewCardsRepository() Cards
}
