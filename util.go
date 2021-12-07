package main

import "os"

func doesFileExist(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func readParquetFile() {
	// ...
}

func getSchema() {
	// ...
}

type displayTable struct {
}
