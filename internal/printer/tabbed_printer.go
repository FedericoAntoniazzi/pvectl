package printer

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
)

// TabbedPrinter abstracts a writer to print out tabstopped text into aligned text
type TabbedPrinter struct {
	writer *tabwriter.Writer
}

// NewTabbedPrinter creates a default TabbedPrinter
func NewTabbedPrinter() *TabbedPrinter {
	return &TabbedPrinter{
		writer: tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', tabwriter.AlignRight),
	}
}

// SetHeaders set titles for each column of the output
func (tp *TabbedPrinter) SetHeaders(headers ...string) {
	fmt.Fprintln(tp.writer, strings.Join(headers, "\t"))
}

// AddRow add rows to the output
func (tp *TabbedPrinter) AddRow(columns ...any) {
	fmt.Fprintln(tp.writer, strings.Join(normalizeStrings(columns), "\t"))
}

// Print flushes the writer buffer to default output
func (tp *TabbedPrinter) Print() {
	tp.writer.Flush()
}

func normalizeStrings(items []any) []string {
	normalized := make([]string, len(items))

	for pos, item := range items {
		normalized[pos] = fmt.Sprint(item)
	}

	return normalized
}
