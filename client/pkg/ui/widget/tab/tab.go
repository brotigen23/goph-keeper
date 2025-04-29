package tab

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Tab struct {
	tabsNames  []string
	currentTab int
	tabs       []tea.Model
}

func New(tabsName []string, tabs []tea.Model) *Tab {
	if len(tabsName) != len(tabs) {
		return nil
	}
	return &Tab{
		tabsNames:  tabsName,
		tabs:       tabs,
		currentTab: 0,
	}
}

func (t Tab) Init() tea.Cmd {
	cmds := []tea.Cmd{}
	for i := range t.tabs {
		cmds = append(cmds, t.tabs[i].Init())
	}
	return tea.Batch(cmds...)
}

func (t Tab) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	switch typedMsg := msg.(type) {
	case tea.KeyMsg:
		s := typedMsg.String()
		if strings.HasPrefix(s, "alt+") {
			d := s[len(s)-1]
			if d >= '1' && d <= '9' {
				t.currentTab = int(d - '1')
			}
		}
	}

	model, cmd := t.tabs[0].Update(msg)
	t.tabs[t.currentTab] = model
	cmds = append(cmds, cmd)
	return t, tea.Batch(cmds...)
}

func (t Tab) View() string {
	return t.renderTabsHeader() + "\n\n" + t.tabs[t.currentTab].View()
}

func (t Tab) renderTabsHeader() string {
	var tabsHeader string
	for i, tab := range t.tabsNames {
		if i == int(t.currentTab) {
			tabsHeader += lipgloss.NewStyle().Bold(true).Render("["+tab+"]") + " "
		} else {
			tabsHeader += tab + " "
		}
	}
	return tabsHeader
}

func (t Tab) GetCurrentTab() string {
	return t.tabsNames[t.currentTab]
}
