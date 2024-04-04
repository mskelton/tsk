package printer

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/mattn/go-runewidth"
)

type Row struct {
	Cells     []string
	Highlight bool
}

type Table struct {
	Columns []string
	Rows    []Row
}

// Special implementation of string padding to account for unicode string width
func pad(str string, w int) string {
	return str + strings.Repeat(" ", w-runewidth.StringWidth(str))
}

func (table *Table) Print() {
	widths := make([]int, len(table.Columns))
	boldUnderline := color.New().Add(color.Bold, color.Underline).SprintFunc()

	// Find the maximum width of each column
	for _, row := range table.Rows {
		for i, cell := range row.Cells {
			length := runewidth.StringWidth(cell)
			widths[i] = max(widths[i], length)
		}
	}

	// Calculate the width of each column header, ignoring empty columns
	for i, col := range table.Columns {
		if widths[i] > 0 {
			// Column headers never have Unicode, so `len()` is safe to use
			widths[i] = max(widths[i], len(col))
		}
	}

	// Create the header row, skipping empty columns
	var header []string
	for i, col := range table.Columns {
		if widths[i] > 0 {
			header = append(header, boldUnderline(pad(col, widths[i])))
		}
	}

	fmt.Println(strings.Join(header, " "))

	// Print an ASCII underline if colorization is disabled
	if color.NoColor {
		var underline []string

		for _, width := range widths {
			if width > 0 {
				underline = append(underline, strings.Repeat("-", width))
			}
		}

		fmt.Println(strings.Join(underline, " "))
	}

	for i, row := range table.Rows {
		var cells []string

		for i, cell := range row.Cells {
			if widths[i] > 0 {
				cells = append(cells, pad(cell, widths[i]))
			}
		}

		line := strings.Join(cells, " ")
		if row.Highlight {
			// color.BGRGB(129, 97, 170)
			color.New(color.BgMagenta).Println(line)
		} else if i%2 == 0 {
			fmt.Println(line)
		} else {
			color.New(color.BgBlack).Println(line)
		}
	}
}
