package datacontroller

import (
	"reflect"

	ui "github.com/brotigen23/goph-keeper/client/internal/tui"
	"github.com/brotigen23/goph-keeper/client/pkg/ui/widget/form"
	tea "github.com/charmbracelet/bubbletea"
)

func (m DataController[T]) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmds := []tea.Cmd{}
	t := reflect.TypeOf(msg)
	m.logger.Info("message", "msg", t.Name())
	switch m.state {
	case View:
		switch viewMsg := msg.(type) {
		case tea.KeyMsg:
			switch viewMsg.String() {
			case "c":
				m.logger.Info("Create new data")
				m.state = Create
				m.form = form.New[T]()
			case "d":
				currentItem := m.table.GetCurrentItem()
				m.logger.Info("Delete data", "item", currentItem)
				if currentItem == nil {
					break
				}
				cmds = append(cmds, m.CRUDCmd(Delete, *currentItem))
			case "enter":
				currentItem := m.table.GetCurrentItem()
				m.logger.Info("Edit exist data", "item", currentItem)
				if currentItem == nil {
					break
				}
				m.state = Edit
				m.form = form.NewWithData(*currentItem)
			}
		case ui.FetchSuccessMsg[T]:
			m.logger.Info("Fetch data success")
			m.table.Refresh(viewMsg.Data)
		}
		table, cmd := m.table.Update(msg)
		m.table = table
		cmds = append(cmds, cmd)
	case Create, Edit:
		switch editMsg := msg.(type) {
		case form.CloseMsg:
			m.logger.Info("Close form msg")
			m.state = View
			m.form = nil
			return m, tea.Batch(cmds...)
		case form.SubmitFormMsg[T]:
			m.logger.Info("submit form", "data", editMsg.Data)
			cmds = append(cmds, m.CRUDCmd(m.state, editMsg.Data))
		}
		form, cmd := m.form.Update(msg)
		m.form = form
		cmds = append(cmds, cmd)
	}
	return m, tea.Batch(cmds...)
}
