package main

import (
	"fmt"
	"log"
	"os"
	"reflect"

	"github.com/fatih/color"
	"github.com/rodaine/table"
	"github.com/urfave/cli/v2"
	"github.com/xitongsys/parquet-go-source/local"
	"github.com/xitongsys/parquet-go/reader"
)

// Version is the current version of the CLI
const Version = "0.0.1"

const TestFilePath = "examples/data/iris.parquet"

// getStructFields returns a list of the names of
// the fields in a struct.
func getStructFields(t reflect.Type) []string {
	var fields []string
	for i := 0; i < t.NumField(); i++ {
		fields = append(fields, t.Field(i).Name)
	}
	return fields
}

func formatRow(cols []string, val reflect.Value) []interface{} {
	var data []interface{}
	for _, c := range cols {
		v := val.FieldByName(c)
		if v.IsNil() {
			data = append(data, "")
		} else {
			data = append(data, v.Elem().Interface())
		}
	}
	return data
}

func runSandbox() {
	log.Println("Starting...")

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
	fmt.Println()

	res, err := pr.ReadByNumber(num)
	if err != nil {
		log.Println("Can't read", err)
		return
	}

	t := pr.ObjType
	fs := getStructFields(t)

	ifs := make([]interface{}, len(fs))
	for i, r := range fs {
		ifs[i] = r
	}
	headerFmt := color.New(color.FgMagenta, color.Italic).SprintfFunc()
	tbl := table.New(ifs...)
	tbl.WithHeaderFormatter(headerFmt)

	for i := 0; i < 5; i++ {
		r0 := res[i]
		e := reflect.ValueOf(r0)
		row := formatRow(fs, e)
		tbl.AddRow(row...)
	}

	tbl.Print()

	pr.ReadStop()
	fr.Close()
}

func main() {
	app := &cli.App{
		Name:  "parq",
		Usage: "A tool for exploring parquet files",
		Commands: []*cli.Command{
			{
				Name:      "head",
				Usage:     "Shows the first n rows of a parquet file.",
				ArgsUsage: "FILENAME",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:    "n-rows",
						Value:   10,
						Aliases: []string{"n"},
						Usage:   "Number of rows to show",
					},
				},
				Action: func(c *cli.Context) error {
					return nil
				},
			},
			{
				Name:      "tail",
				Usage:     "Shows the last n rows of a parquet file",
				ArgsUsage: "FILENAME",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:    "n-rows",
						Value:   10,
						Aliases: []string{"n"},
						Usage:   "Number of rows to show",
					},
				},
				Action: func(c *cli.Context) error {
					return nil
				},
			},
			{
				Name:      "random",
				Usage:     "Shows the n random rows of a parquet file.",
				ArgsUsage: "FILENAME",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:    "n-rows",
						Value:   10,
						Aliases: []string{"n"},
						Usage:   "Number of rows to show",
					},
				},
				Action: func(c *cli.Context) error {
					return nil
				},
			},
			{
				Name:      "show",
				Aliases:   []string{"all"},
				Usage:     "Shows all rows of a parquet file.",
				ArgsUsage: "FILENAME",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:    "n-rows",
						Value:   10,
						Aliases: []string{"n"},
						Usage:   "Number of rows to show",
					},
				},
				Action: func(c *cli.Context) error {
					return nil
				},
			},
			{
				Name:      "schema",
				Usage:     "Shows a parquet file's column names and data types.",
				ArgsUsage: "FILENAME",
				Flags:     []cli.Flag{},
				Action: func(c *cli.Context) error {
					return nil
				},
			},
			{
				Name:  "convert",
				Usage: "Convert a parquet file to/from another format.",
				Subcommands: []*cli.Command{
					{
						Name:      "to",
						Usage:     "Convert a parquet file to another format.",
						ArgsUsage: "FILENAME",
						Flags:     []cli.Flag{},
						Action: func(c *cli.Context) error {
							return nil
						},
					},
					{
						Name:      "from",
						Usage:     "Convert a parquet file from another format.",
						UsageText: "parq convert from [OPTIONS] <FILENAME>",
						ArgsUsage: "FILENAME",
						Flags:     []cli.Flag{},
						Action: func(c *cli.Context) error {
							return nil
						},
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
