package form

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Form[T any] struct {
	data T

	inputs     []textinput.Model
	fieldNames []string
	focus      int

	style Style
}

func NewWithData[T any](data T) *Form[T] {
	val := reflect.ValueOf(data)
	inputs, fieldNames := generateFields[T](val)

	for i, name := range fieldNames {
		field := val.FieldByName(name)
		if field.IsValid() {
			var strValue string

			switch field.Kind() {
			case reflect.String:
				strValue = field.String()
			case reflect.Int, reflect.Int64, reflect.Int32:
				strValue = fmt.Sprintf("%d", field.Int())
			case reflect.Float64, reflect.Float32:
				strValue = fmt.Sprintf("%f", field.Float())
			case reflect.Bool:
				strValue = fmt.Sprintf("%t", field.Bool())
			case reflect.Struct:
				if t, ok := field.Interface().(time.Time); ok {
					strValue = t.String()
				}
			}

			inputs[i].SetValue(strValue)
		}
	}
	if len(inputs) > 0 {
		inputs[0].Focus()
	}

	return &Form[T]{
		data:       data,
		inputs:     inputs,
		fieldNames: fieldNames,
		focus:      0,
		style:      DefaultStyle(),
	}
}

func New[T any]() *Form[T] {
	var zero T
	val := reflect.ValueOf(zero)
	inputs, fieldNames := generateFields[T](val)
	if len(inputs) > 0 {
		inputs[0].Focus()
	}
	ret := &Form[T]{
		inputs:     inputs,
		fieldNames: fieldNames,
		focus:      0,
		style:      DefaultStyle(),
	}
	ret.Init()
	return ret
}

func generateFields[T any](val reflect.Value) ([]textinput.Model, []string) {
	typ := val.Type()
	if val.Kind() == reflect.Pointer {
		val = val.Elem()
		typ = typ.Elem()
	}
	fields := getAllFields(typ)
	var inputs []textinput.Model
	var fieldNames []string

	for _, field := range fields {
		formTag := field.Tag.Get("form")

		if formTag == "" {
			continue
		}

		parts := strings.Split(formTag, ",")
		if len(parts) == 0 || parts[0] != "true" {
			continue
		}

		width := 50
		maxLen := 0

		if len(parts) > 1 {
			if w, err := strconv.Atoi(parts[1]); err == nil {
				width = w
			}
		}
		if len(parts) > 2 {
			if l, err := strconv.Atoi(parts[2]); err == nil {
				maxLen = l
			}
		}

		input := textinput.New()
		input.Placeholder = field.Name
		input.Prompt = ""
		input.Width = width

		if maxLen > 0 {
			input.CharLimit = maxLen
		}

		inputs = append(inputs, input)
		fieldNames = append(fieldNames, field.Name)
	}

	return inputs, fieldNames
}

func (m Form[T]) Init() tea.Cmd {
	return nil
}
