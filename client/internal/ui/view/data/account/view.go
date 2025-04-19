package account

func (m model) View() string {
	if m.isLoading {
		return "Loading..."
	}
	if m.data == nil {
		return "No content" + "\n" + `Press "C" to create new row`
	}
	var frame string
	frame += m.table.View()
	return frame
}
