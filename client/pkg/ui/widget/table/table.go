package table

import (
	"fmt"
	"reflect"

	"github.com/charmbracelet/bubbles/table"
	"github.com/mattn/go-runewidth"
)

type Table[T any] struct {
	table          table.Model
	maxColumnWidth int
}

// Creates a new table with columns of T struct, with no rows
func New[T any]() *Table[T] {
	columns, _ := generateColumnsAndRows[T](nil)

	// Create table
	table := table.New(
		table.WithColumns(columns),
	)

	return &Table[T]{
		maxColumnWidth: 50,
		table:          table,
	}
}

// Return representation of table
func (t Table[T]) View() string {
	return t.table.View()
}

// Add rows into table
func (t *Table[T]) Refresh(rows []T) {
	// Create new rows
	_, newRows := generateColumnsAndRows(rows)
	// Add rows to table
	t.table.SetRows(newRows)
	t.autoResizeWidth()
}

func (t *Table[T]) autoResizeWidth() {
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

func (t *Table[T]) CursorUp() {
	cursor := t.table.Cursor()
	t.table.SetCursor(cursor - 1)
}
func (t *Table[T]) CursorDown() {
	cursor := t.table.Cursor()
	t.table.SetCursor(cursor + 1)
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
