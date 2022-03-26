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

const requiredCols = 2

func NewTable(initGoroutines, endGoroutines int) *Table {
	table := tablewriter.NewWriter(os.Stdout)

	header := makeHeader(initGoroutines, endGoroutines)

	table.SetHeader(header)

	return &Table{
		table: *table,
		row:   make([]string, 0, requiredCols+endGoroutines-initGoroutines),
	}
}

func makeHeader(initGoroutines, endGoroutines int) []string {
	count := requiredCols + endGoroutines - initGoroutines
	header := make([]string, count, count)
	header[0] = "Size"
	header[1] = "Sequential\n (time in microseconds)"
	index := requiredCols

	for i := initGoroutines; i < endGoroutines; i++ {
		header[index] = "Goroutines - " + strconv.Itoa(i)
		index++
	}
	return header
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
