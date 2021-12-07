package main

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/fatih/color"
	"github.com/rodaine/table"
)

var (
	ErrFileNotFound = errors.New("file not found")
	ErrCantReadFile = errors.New("can't read file")
	ErrNoRowsInFile = errors.New("no rows in file")
)

func printParquetSchema(path string) error {
	// Verrify the file exists
	if path == "" || !doesFileExist(path) {
		return ErrFileNotFound
	}

	// Read in the parquet file
	pr, err := readParquetFile(path)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrCantReadFile, err)
	}

	// TODO: Check the number of rows so the read
	// doesn't fail
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

func printParquetFile(path string) error {
	// Verrify the file exists
	if !doesFileExist(path) {
		return ErrFileNotFound
	}

	// Read in the parquet file
	pr, err := readParquetFile(path)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrCantReadFile, err)
	}

	fmt.Println(pr)

	return nil
}
