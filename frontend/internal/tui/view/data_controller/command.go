package datacontroller

import tea "github.com/charmbracelet/bubbletea"

func (d DataController[T]) CRUDCmd(action int, item T) tea.Cmd {
	return func() tea.Msg {
		return CRUDMsg[T]{
			Action: action,
			Item:   item,
		}
	}
}

func SendMsg(msg tea.Msg) tea.Cmd {
	return func() tea.Msg {
		return msg
	}
}
