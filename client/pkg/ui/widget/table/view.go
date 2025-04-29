package table

const help = "Use arrows or j k to navigate\nEnter to choose\nc to create\nd to delete"

// Return representation of table
func (t Model[T]) View() string {
	var frame string

	frame += t.table.View()

	frame += "\n\n"

	frame += help

	return frame
}
