package postgres

// Users table
// ----------------------------------------------------------------
type userColumns struct {
	ID, CreatedAt, UpdatedAt string
	Login, Password          string
}

var UsersTable = struct {
	Name    string
	Columns userColumns
}{
	Name: "users",
	Columns: userColumns{
		ID: "id", CreatedAt: "created_at", UpdatedAt: "updated_at",
		Login: "login", Password: "password",
	},
}

// Accounts table
// ----------------------------------------------------------------
type accountColumns struct {
	ID, CreatedAt, UpdatedAt string
	UserID, Metadata         string
	Login, Password          string
}

var AccountsTable = struct {
	Name    string
	Columns accountColumns
}{
	Name: "accounts_data",
	Columns: accountColumns{
		ID: "id", CreatedAt: "created_at", UpdatedAt: "updated_at",
		UserID: "user_id", Metadata: "metadata",
		Login: "login", Password: "password",
	},
}

// Text table
// ----------------------------------------------------------------
type textColumns struct {
	ID, CreatedAt, UpdatedAt string
	UserID, Metadata         string
	Data                     string
}

var TextTable = struct {
	Name    string
	Columns textColumns
}{
	Name: "text_data",
	Columns: textColumns{
		ID: "id", CreatedAt: "created_at", UpdatedAt: "updated_at",
		UserID: "user_id", Metadata: "metadata",
		Data: "data",
	},
}

// Binary table
// ----------------------------------------------------------------
type binaryColumns struct {
	ID, CreatedAt, UpdatedAt string
	UserID, Metadata         string
	Data                     string
}

var BinaryTable = struct {
	Name    string
	Columns binaryColumns
}{
	Name: "binary_data",
	Columns: binaryColumns{
		ID: "id", CreatedAt: "created_at", UpdatedAt: "updated_at",
		UserID: "user_id", Metadata: "metadata",
		Data: "data",
	},
}

// Cards table
// ----------------------------------------------------------------
type cardsColumns struct {
	ID, CreatedAt, UpdatedAt string
	UserID, Metadata         string
	Number, CardholderName   string
	ExpiresAt, CVV           string
}

var CardsTable = struct {
	Name    string
	Columns cardsColumns
}{
	Name: "cards_data",
	Columns: cardsColumns{
		ID: "id", CreatedAt: "created_at", UpdatedAt: "updated_at",
		UserID: "user_id", Metadata: "metadata",
		Number: "number", CardholderName: "cardholder_name",
		ExpiresAt: "expire", CVV: "cvv",
	},
}
