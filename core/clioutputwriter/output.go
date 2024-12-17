package clioutputwriter

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
)

// Format represents the output format.
type Format string

const (
	FormatTable   Format = "table"
	FormatJSON    Format = "json"
	FormatCSV     Format = "csv"
	FormatNuShell Format = "nushell"
)

// TabularData represents the data to be printed in a structured format.
type TabularData struct {
	Headers []string        // Column headers
	Rows    [][]interface{} // Rows of data
}

var (
	ErrUnsupportedFormat = fmt.Errorf("unsupported output format")
)

// FilterColumns will filter out all columns not in the provided list.
func FilterColumns(data TabularData, columns []string) TabularData {
	filteredData := TabularData{
		Headers: columns,
		Rows:    [][]interface{}{},
	}
	for _, row := range data.Rows {
		filteredRow := make([]interface{}, len(columns))
		for i, col := range columns {
			for j, header := range data.Headers {
				if header == col {
					filteredRow[i] = row[j]
					break
				}
			}
		}
		filteredData.Rows = append(filteredData.Rows, filteredRow)
	}
	return filteredData
}

// SupportedOutputFormats returns the list of supported output formats, for use in help text.
func SupportedOutputFormats() []string {
	return []string{
		string(FormatTable),
		string(FormatJSON),
		string(FormatCSV),
		string(FormatNuShell),
	}
}

// DefaultOutputFormat returns the suggested output format
func DefaultOutputFormat() Format {
	// detect nushell
	if _, ok := os.LookupEnv("NU_VERSION"); ok {
		return FormatNuShell
	}

	return FormatTable
}

// PrintData is the main function to print data in the specified format.
func PrintData(w io.Writer, data TabularData, format Format) error {
	switch format {
	case FormatTable:
		return printTable(w, data)
	case FormatJSON:
		return printJSON(w, data)
	case FormatCSV:
		return printCSV(w, data)
	case FormatNuShell:
		return printNuShell(w, data)
	default:
		return errors.Join(ErrUnsupportedFormat, fmt.Errorf(string(format)))
	}
}

// printTable renders the data in a tab separated format.
func printTable(w io.Writer, data TabularData) error {
	tw := tabwriter.NewWriter(w, 1, 1, 1, ' ', 0)
	defer func(tw *tabwriter.Writer) {
		_ = tw.Flush()
	}(tw)

	_, err := fmt.Fprintln(tw, strings.Join(data.Headers, "\t"))
	if err != nil {
		return err
	}
	for _, row := range data.Rows {
		strRow := interfaceToStringRow(row)
		_, err = fmt.Fprintln(tw, strings.Join(strRow, "\t"))
		if err != nil {
			return err
		}
	}

	return nil
}

// printJSON renders the data in JSON format.
func printJSON(w io.Writer, data TabularData) error {
	output := make([]map[string]interface{}, len(data.Rows))
	for i, row := range data.Rows {
		rowMap := make(map[string]interface{})
		for j, header := range data.Headers {
			rowMap[header] = row[j]
		}
		output[i] = rowMap
	}

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")
	return encoder.Encode(output)
}

// printCSV renders the data in CSV format.
func printCSV(w io.Writer, data TabularData) error {
	_, err := fmt.Fprintln(w, strings.Join(data.Headers, ","))
	if err != nil {
		return err
	}
	for _, row := range data.Rows {
		strRow := interfaceToStringRow(row)
		_, err = fmt.Fprintln(w, strings.Join(strRow, ","))
		if err != nil {
			return err
		}
	}
	return nil
}

// printNuShell renders the data in Nushell format. - e.g. [[a b]; [1 2]] where a b is the header and 1 2 is the first row
func printNuShell(w io.Writer, data TabularData) error {
	_, err := fmt.Fprint(w, "[[")
	if err != nil {
		return err
	}

	// Write the header row
	for i, header := range data.Headers {
		_, err = fmt.Fprint(w, strconv.Quote(header))
		if err != nil {
			return err
		}
		if i < len(data.Headers)-1 {
			_, err = fmt.Fprint(w, " ")
			if err != nil {
				return err
			}
		}
	}
	_, err = fmt.Fprint(w, "]; ")
	if err != nil {
		return err
	}

	for i, row := range data.Rows {
		_, err = fmt.Fprint(w, "[")
		if err != nil {
			return err
		}

		strRow := interfaceToStringRow(row)
		for j, cell := range strRow {
			_, err = fmt.Fprint(w, strconv.Quote(cell))
			if err != nil {
				return err
			}
			if j < len(strRow)-1 {
				_, err = fmt.Fprint(w, " ")
				if err != nil {
					return err
				}
			}
		}

		_, err = fmt.Fprint(w, "]")
		if err != nil {
			return err
		}
		if i < len(data.Rows)-1 {
			_, err = fmt.Fprint(w, " ")
			if err != nil {
				return err
			}
		}
	}

	_, err = fmt.Fprint(w, "]")
	return err
}

func interfaceToStringRow(row []interface{}) []string {
	strRow := make([]string, len(row))
	for i, v := range row {
		strRow[i] = fmt.Sprintf("%v", v)
	}
	return strRow
}
