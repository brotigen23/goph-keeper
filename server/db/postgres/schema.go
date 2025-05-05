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

// Metadata table
// ----------------------------------------------------------------
type metadataColumns struct {
	ID, CreatedAt, UpdatedAt string
	Data                     string
}

var MetadataTable = struct {
	Name    string
	Columns metadataColumns
}{
	Name: "metadata",
	Columns: metadataColumns{
		ID: "id", CreatedAt: "created_at", UpdatedAt: "updated_at",
		Data: "data",
	},
}

// TODO: other tables
