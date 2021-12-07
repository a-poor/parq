package main

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/fatih/color"
	"github.com/rodaine/table"
)

var (
	ErrFileNotFound       = errors.New("file not found")
	ErrCantReadFile       = errors.New("can't read file")
	ErrNoRowsInFile       = errors.New("no rows in file")
	ErrReadingParquetFile = errors.New("unable to read parquet file")
)

func printParquetSchema(path string) error {
	// Verrify the file exists
	if path == "" || !doesFileExist(path) {
		return ErrFileNotFound
	}

	// Read in the parquet file
	pr, pfClose, err := readParquetFile(path)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrCantReadFile, err)
	}
	defer pfClose()
	defer pr.ReadStop()

	// Check the number of rows so the read doesn't fail
	nrows := getNumRows(pr)
	if nrows < 1 {
		return ErrNoRowsInFile
	}

	// Read a single row so `pr.ObjType`
	// gets populated
	pr.ReadByNumber(1)

	// Get table's data type
	dt := pr.ObjType

	// Get the schema
	sf := getStructFields(dt)

	// Create a table
	tbl := table.New("Column Name", "Data Type")
	tbl.WithHeaderFormatter(color.New(
		color.FgMagenta,
		color.Bold,
	).SprintfFunc())

	// Populate the table
	for _, f := range sf {
		ft := f.Type
		if ft.Kind() == reflect.Ptr {
			ft = ft.Elem()
		}
		tbl.AddRow(f.Name, ft.String())
	}

	// Print the table
	tbl.Print()
	fmt.Println()

	return nil
}

func printParquetFile(path string, idxr getIndexer) error {
	// Verrify the file exists
	if !doesFileExist(path) {
		return ErrFileNotFound
	}

	// Read in the parquet file
	pr, pfClose, err := readParquetFile(path)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrCantReadFile, err)
	}
	defer pfClose()
	defer pr.ReadStop()

	// Check the number of rows so the read doesn't fail
	nrows := getNumRows(pr)
	if nrows < 1 {
		return ErrNoRowsInFile
	}

	// Read data from the file
	res, err := pr.ReadByNumber(nrows)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrReadingParquetFile, err)
	}

	// Get table's data type
	dt := pr.ObjType

	// Extract fields & col-names
	fs := getStructFields(dt)
	colNames := make([]interface{}, len(fs)+1)
	colNames[0] = "" // Empty first column to store row index
	for i, f := range fs {
		colNames[i+1] = f.Name
	}

	// Create the table
	tbl := table.New(colNames...)

	// Style the table
	headerFmt := color.New(tableHeaderStyle...).SprintfFunc()
	firstRowFmt := color.New(tableFirstRowStyle...).SprintfFunc()
	tbl.WithHeaderFormatter(headerFmt)
	tbl.WithFirstColumnFormatter(firstRowFmt)

	// Indexes to include
	idxs := idxr.getIndexes(nrows)

	// Add the data
	for _, i := range idxs {
		r := res[i]
		e := reflect.ValueOf(r)
		row := formatRow(fs, e)
		irow := prependIndex(i, row)
		tbl.AddRow(irow...)
	}

	// Print the table
	tbl.Print()
	fmt.Println()

	return nil
}
