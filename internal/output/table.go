package output

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
	"strconv"
)

type Table struct {
	table tablewriter.Table
	row   []string
}

const requiredCols = 2

func NewTable(initGoroutines, endGoroutines, step int) *Table {
	table := tablewriter.NewWriter(os.Stdout)

	header := makeHeader(initGoroutines, endGoroutines, step)

	table.SetHeader(header)

	return &Table{
		table: *table,
		row:   make([]string, 0, requiredCols+endGoroutines-initGoroutines),
	}
}

func makeHeader(initGoroutines, endGoroutines, step int) []string {
	count := requiredCols + ((endGoroutines-initGoroutines)/step+1)*4
	header := make([]string, 0, count)
	header = append(header, "Size")
	header = append(header, "Sequential\n (time in microseconds)")

	for i := initGoroutines; i <= endGoroutines; i += step {
		header = append(header, "Concurrent - "+strconv.Itoa(i))
		header = append(header, "SpeedUp - "+strconv.Itoa(i))
		header = append(header, "Parallel - "+strconv.Itoa(i))
		header = append(header, "SpeedUp - "+strconv.Itoa(i))
	}
	return header
}

func (t *Table) AddIntElement(element int) {
	t.row = append(t.row, strconv.Itoa(element))
}

func (t *Table) AddDoubleElement(element float64) {
	t.row = append(t.row, fmt.Sprintf("%.2f", element))
}

func (t *Table) AddRow() {
	t.table.Append(t.row)
	t.row = make([]string, 0, cap(t.row))
}

func (t *Table) Render() {
	t.table.Render()
}
