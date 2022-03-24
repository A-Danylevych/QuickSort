package output

import (
	"github.com/olekukonko/tablewriter"
	"os"
	"strconv"
)

type Table struct {
	table tablewriter.Table
	row   []string
}

func NewTable(initGoroutines, endGoroutines int) *Table {
	table := tablewriter.NewWriter(os.Stdout)

	count := endGoroutines - initGoroutines
	header := make([]string, 2+count, 2+count)
	header[0] = "Size"
	header[1] = "Linear"
	index := 2

	for i := initGoroutines; i < endGoroutines; i++ {
		header[index] = "Goroutines - " + strconv.Itoa(i)
	}
	table.SetHeader(header)

	return &Table{
		table: *table,
		row:   make([]string, 0, index),
	}
}

func (t *Table) AddIntElement(element int) {
	t.row = append(t.row, strconv.Itoa(element))
}

func (t *Table) AddDoubleElement(element float64) {
	t.AddIntElement(int(element))
}

func (t *Table) AddRow() {
	t.table.Append(t.row)
	t.row = make([]string, 0, cap(t.row))
}

func (t *Table) Render() {
	t.table.Render()
}
