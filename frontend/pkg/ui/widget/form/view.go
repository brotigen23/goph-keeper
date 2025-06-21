package form

func (m Form[T]) View() string {

	var frame string
	for i := range m.inputs {
		styleInput := m.style.Input
		stylePrompt := m.style.Prompt
		if i == m.focus {
			styleInput = m.style.InputFocus
			stylePrompt = m.style.PromptFocus
		}

		input := m.inputs[i]

		frame += stylePrompt.Render(m.fieldNames[i])
		frame += "\n"

		frame += styleInput.Render(input.View())
		frame += "\n\n"

	}
	frame += "\n\n"
	frame += "Press ESC to quit form"
	return frame
}
