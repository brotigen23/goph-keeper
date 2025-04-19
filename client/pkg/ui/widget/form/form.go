package form

import (
	"reflect"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Form[T any] struct {
	data *T

	inputs     []textinput.Model
	fieldNames []string
	focus      int
}

func NewWithData[T any](data *T) *Form[T] {
	typ := reflect.TypeOf(data)
	val := reflect.ValueOf(data)

	inputs := make([]textinput.Model, typ.NumField())
	fieldNames := make([]string, typ.NumField())

	for i := range typ.NumField() {
		field := typ.Field(i)
		fieldNames[i] = field.Name

		input := textinput.New()
		input.Placeholder = field.Name
		input.Prompt = field.Name + ": "

		if val.Field(i).Kind() == reflect.String {
			input.SetValue(val.Field(i).String())
		}

		if i == 0 {
			input.Focus()
		}

		inputs[i] = input
	}

	return &Form[T]{
		data:       data,
		inputs:     inputs,
		fieldNames: fieldNames,
		focus:      0,
	}
}

func (m Form[T]) submit() tea.Cmd {
	val := reflect.ValueOf(&m.data).Elem()

	for i, name := range m.fieldNames {
		field := val.FieldByName(name)
		if field.IsValid() && field.CanSet() && field.Kind() == reflect.String {
			field.SetString(m.inputs[i].Value())
		}
	}

	return func() tea.Msg {
		return SubmitFormMsg[T]{Data: m.data}
	}
}

func New[T any](inputsText []string) *Form[T] {
	inputs := make([]textinput.Model, len(inputsText))
	for i := range inputs {
		inputs[i] = textinput.New()

		inputs[i].Prompt = inputsText[i]
		inputs[i].Width = 50
	}
	inputs[0].Focus()

	return &Form[T]{
		inputs: inputs,
		focus:  0,
	}
}

func (m Form[T]) Init() tea.Cmd {
	return nil
}

func (m Form[T]) View() string {

	var frame string
	for i := range m.inputs {
		frame += m.inputs[i].Prompt
		frame += m.inputs[i].View()
	}
	return frame
}
func (m *Form[T]) changeInput() {
	m.focus++
	if m.focus > 1 {
		m.focus = 0
	}
	for i := range m.inputs {
		if i == m.focus {
			m.inputs[m.focus].Focus()
			continue
		}
		m.inputs[i].Blur()
	}
}
func (m *Form[T]) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.inputs))

	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

func (m Form[T]) EditConfirm() tea.Msg {
	return nil
}

func (m Form[T]) GetValues() *T {
	return nil
}
