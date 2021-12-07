package main

import (
	"os"
	"reflect"

	"github.com/fatih/color"
	"github.com/xitongsys/parquet-go-source/local"
	"github.com/xitongsys/parquet-go/reader"
)

// tableHeaderStyle is the color stylers for the table header
var (
	tableHeaderStyle   = []color.Attribute{color.FgMagenta, color.Italic}
	tableFirstRowStyle = []color.Attribute{color.FgYellow}
)

type StructField struct {
	Name string       // Name of the struct field
	Type reflect.Type // Field's type (as a reflect.Type)
}

func (sf StructField) isPtr() bool {
	return sf.Type.Kind() == reflect.Ptr
}

func (sf StructField) getDerefType() reflect.Type {
	if sf.isPtr() {
		return sf.Type.Elem()
	}
	return sf.Type
}

// getStructFields returns a list of the names of
// the fields in a struct.
func getStructFields(t reflect.Type) []StructField {
	var fields []StructField
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fields = append(fields, StructField{
			Name: f.Name,
			Type: f.Type,
		})
	}
	return fields
}

// doesFileExist checks if a file exists and is not a directory before
// try using it to prevent further errors.
func doesFileExist(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// readParquetFile opens a local parquet file and returns a
// reader (if possible), as well as a function to close the
// local file reader.
func readParquetFile(path string) (*reader.ParquetReader, func() error, error) {
	fr, err := local.NewLocalFileReader(TestFilePath)
	if err != nil {
		return nil, nil, err
	}

	pr, err := reader.NewParquetReader(fr, nil, 1)
	if err != nil {
		return nil, nil, err
	}
	return pr, fr.Close, nil
}

// getNumRows returns the number of rows in
// a parquet file as an integer.
func getNumRows(pr *reader.ParquetReader) int {
	return int(pr.GetNumRows())
}

// prependIndex adds an index number to the beginning of a
// slice of interface{} data.
//
// This is a utility function used to add row index numbers
// when generating tables.
func prependIndex(idx int, data []interface{}) []interface{} {
	full := make([]interface{}, len(data)+1)
	full[0] = idx
	for i, d := range data {
		full[i+1] = d
	}
	return full
}

// formatRow
func formatRow(cols []StructField, val reflect.Value) []interface{} {
	var data []interface{}
	for _, c := range cols {
		v := val.FieldByName(c.Name)
		switch {
		case c.isPtr() && v.IsNil():
			data = append(data, "")
		case c.isPtr():
			data = append(data, v.Elem().Interface())
		default:
			data = append(data, v.Interface())
		}
	}
	return data
}
