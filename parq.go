package main

import (
	"fmt"
	"log"
	"reflect"

	"github.com/xitongsys/parquet-go-source/local"
	"github.com/xitongsys/parquet-go/reader"
)

const TestFilePath = "examples/data/iris.parquet"

func getStructFields(t reflect.Type) []string {
	var fields []string
	for i := 0; i < t.NumField(); i++ {
		fields = append(fields, t.Field(i).Name)
	}
	return fields
}

func main() {
	fmt.Println("Starting...")

	fr, err := local.NewLocalFileReader(TestFilePath)
	if err != nil {
		log.Println("Can't open file")
		return
	}

	pr, err := reader.NewParquetReader(fr, nil, 4)
	if err != nil {
		log.Println("Can't create parquet reader", err)
		return
	}

	num := int(pr.GetNumRows())
	log.Printf("There are %d rows in the file\n", num)

	res, err := pr.ReadByNumber(num)
	if err != nil {
		log.Println("Can't read", err)
		return
	}

	// log.Printf("%T\n", res)
	log.Printf("%+v\n", res[0])

	t := pr.ObjType
	log.Printf("%+v\n", t)
	fmt.Println()
	fs := getStructFields(t)
	fmt.Printf("Struct fields: %v\n", fs)

	r0 := res[0]
	e := reflect.ValueOf(r0)
	for _, f := range fs {
		v := e.FieldByName(f)
		fmt.Printf("%s = %v\n", f, v.Elem().Interface())
	}

	// tbl := table.New("ID", "Name", "Score", "Added")
	// fmt.Println(tbl)

	pr.ReadStop()
	fr.Close()
}
