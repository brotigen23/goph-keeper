package form

import (
	"fmt"
	"reflect"
	"strconv"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Form[T]) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmds := []tea.Cmd{}
	switch formMsg := msg.(type) {
	case tea.KeyMsg:
		switch formMsg.String() {
		case "tab":
			m.changeInputDowm(1)
		case "shift+tab":
			m.changeInputDowm(-1)
		case "enter":
			if m.focus < len(m.inputs)-1 {
				m.changeInputDowm(1)
			} else {
				cmds = append(cmds, m.submit)
			}
		case "esc":
			cmds = append(cmds, func() tea.Msg {
				return CloseMsg{}
			})
		}
	}
	cmds = append(cmds, m.updateInputs(msg))
	return m, tea.Batch(cmds...)
}

func (m *Form[T]) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.inputs))

	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

func (m *Form[T]) changeInputDowm(d int) {
	switch d > 0 {
	case true:
		m.focus++
		if m.focus > len(m.inputs)-1 {
			m.focus = 0
		}
	case false:
		m.focus--
		if m.focus < 0 {
			m.focus = len(m.inputs) - 1
		}

	default:
		break
	}
	for i := range m.inputs {
		if i == m.focus {
			m.inputs[m.focus].Focus()
			continue
		}
		m.inputs[i].Blur()
	}
}

func (m Form[T]) submit() tea.Msg {
	val := reflect.ValueOf(&m.data).Elem()

	for i, name := range m.fieldNames {
		field := val.FieldByName(name)
		if !field.IsValid() || !field.CanSet() {
			continue
		}

		inputValue := m.inputs[i].Value()

		switch field.Kind() {
		case reflect.String:
			field.SetString(inputValue)

		case reflect.Int, reflect.Int64:
			if parsed, err := strconv.ParseInt(inputValue, 10, 64); err == nil {
				field.SetInt(parsed)
			} else {
				return SubmitFormErrorMsg{Field: name, Error: fmt.Errorf("не удалось распарсить int: %w", err)}
			}

		case reflect.Float64:
			if parsed, err := strconv.ParseFloat(inputValue, 64); err == nil {
				field.SetFloat(parsed)
			} else {
				return SubmitFormErrorMsg{Field: name, Error: fmt.Errorf("не удалось распарсить float: %w", err)}
			}

		case reflect.Bool:
			if parsed, err := strconv.ParseBool(inputValue); err == nil {
				field.SetBool(parsed)
			} else {
				return SubmitFormErrorMsg{Field: name, Error: fmt.Errorf("не удалось распарсить bool: %w", err)}
			}

		case reflect.Struct:
			if field.Type() == reflect.TypeOf(time.Time{}) {
				if t, err := time.Parse(time.RFC3339, inputValue); err == nil {
					field.Set(reflect.ValueOf(t))
				} else {
					return SubmitFormErrorMsg{Field: name, Error: fmt.Errorf("не удалось распарсить дату: %w", err)}
				}
			}
		}
	}

	return SubmitFormMsg[T]{Data: m.data}
}

func getAllFields(t reflect.Type) []reflect.StructField {
	var fields []reflect.StructField

	for i := range t.NumField() {
		f := t.Field(i)

		if f.Anonymous && f.Type.Kind() == reflect.Struct {
			fields = append(fields, getAllFields(f.Type)...)
		} else {
			fields = append(fields, f)
		}
	}

	return fields
}
