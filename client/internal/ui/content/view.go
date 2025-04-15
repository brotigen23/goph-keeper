package content

func (m model) View() string {
	if m.isLoading {
		return "Loading..."
	}
	return StringAccounts(m.table)
}
