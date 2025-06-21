package table

import (
	"fmt"
	"reflect"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mattn/go-runewidth"
)

type Model[T any] struct {
	items []T

	table          table.Model
	maxColumnWidth int

	style table.Styles
}

// Creates a new table with columns of T struct, with no rows
func New[T any]() *Model[T] {
	columns, _ := generateColumnsAndRows[T](nil)

	// Create table
	table := table.New(
		table.WithColumns(columns),
		table.WithFocused(true),
	)
	t := table

	t.SetStyles(FromExample())
	return &Model[T]{
		maxColumnWidth: 50,
		table:          t,
	}
}

// Return representation of table
func (t Model[T]) Init() tea.Cmd {
	return nil
}

// Add rows into table
func (t *Model[T]) Refresh(rows []T) {
	t.items = rows
	_, newRows := generateColumnsAndRows(rows)
	t.table.SetRows(newRows)
	t.autoResizeWidth()
}

func (t *Model[T]) autoResizeWidth() {
	columns := t.table.Columns()
	rows := t.table.Rows()

	numCols := len(columns)
	colWidths := make([]int, numCols)

	// Headers
	for i, h := range columns {
		colWidths[i] = runewidth.StringWidth(h.Title)
	}

	// Rows
	for _, row := range rows {
		for i, cell := range row {
			w := runewidth.StringWidth(cell)
			if w > colWidths[i] {
				if w > t.maxColumnWidth {
					colWidths[i] = t.maxColumnWidth
				}
				colWidths[i] = w
			}
		}
	}
	cols := make([]table.Column, numCols)
	for i := range numCols {
		cols[i] = table.Column{
			Title: columns[i].Title,
			Width: colWidths[i] + 2,
		}
	}
	t.table.SetColumns(cols)
}

func (t Model[T]) GetCurrentItem() *T {
	if t.items == nil {
		return nil
	}
	return &t.items[t.table.Cursor()]
}

func generateColumnsAndRows[T any](items []T) ([]table.Column, []table.Row) {
	var columns []table.Column
	var rows []table.Row

	var zero T
	t := reflect.TypeOf(zero)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	fieldNames := extractTableFields(t)

	for _, field := range fieldNames {
		columns = append(columns, table.Column{
			Title: field,
			Width: 20,
		})
	}

	for _, item := range items {
		val := reflect.ValueOf(item)
		if val.Kind() == reflect.Ptr {
			val = val.Elem()
		}

		var row table.Row
		for _, name := range fieldNames {
			field := val.FieldByName(name)
			row = append(row, fmt.Sprintf("%v", field.Interface()))
		}
		rows = append(rows, row)
	}

	return columns, rows
}

func extractTableFields(t reflect.Type) []string {
	var fields []string

	for i := range t.NumField() {
		f := t.Field(i)

		if f.Anonymous && f.Type.Kind() == reflect.Struct {
			fields = append(fields, extractTableFields(f.Type)...)
			continue
		}

		if tag := f.Tag.Get("table"); tag == "true" {
			fields = append(fields, f.Name)
		}
	}

	return fields
}
