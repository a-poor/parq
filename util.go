package main

import (
	"os"
	"reflect"

	"github.com/xitongsys/parquet-go-source/local"
	"github.com/xitongsys/parquet-go/reader"
)

type StructFields struct {
	Name string
	Type reflect.Type
}

// getStructFields returns a list of the names of
// the fields in a struct.
func getStructFields(t reflect.Type) []StructFields {
	var fields []StructFields
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fields = append(fields, StructFields{
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

// readParquetFile
func readParquetFile(path string) (*reader.ParquetReader, error) {
	fr, err := local.NewLocalFileReader(TestFilePath)
	if err != nil {
		return nil, err
	}

	pr, err := reader.NewParquetReader(fr, nil, 4)
	if err != nil {
		return nil, err
	}
	return pr, nil
}

func getSchema() {
	// ...
}

type displayTable struct {
}
