package datacontroller

func (m DataController[T]) View() string {
	var frame string
	switch m.state {
	case View:
		frame += m.table.View()
	case Edit, Create:
		frame += m.form.View()
	}

	return frame
}
