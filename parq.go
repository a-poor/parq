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

func formatRow(cols []string, val reflect.Value, index int) []interface{} {
	var data []interface{}
	data = append(data, index)
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
	sfs := make([]string, len(fs))
	for i, f := range fs {
		sfs[i] = f.Name
	}

	ifs := make([]interface{}, len(fs)+1)
	ifs[0] = ""
	for i, r := range sfs {
		ifs[i+1] = r
	}
	headerFmt := color.New(tableHeaderStyle...).SprintfFunc()
	firstRowFmt := color.New(tableFirstRowStyle...).SprintfFunc()
	tbl := table.New(ifs...)
	tbl.WithHeaderFormatter(headerFmt)
	tbl.WithFirstColumnFormatter(firstRowFmt)

	for i := 0; i < 5; i++ {
		r0 := res[i]
		e := reflect.ValueOf(r0)
		row := formatRow(sfs, e, i)
		tbl.AddRow(row...)
	}

	tbl.Print()

	pr.ReadStop()
	fr.Close()

	fmt.Println()
	fmt.Println()

	t2 := table.New("Column Name", "Data Type")
	t2.WithHeaderFormatter(color.New(color.FgMagenta, color.Bold).SprintfFunc())
	for _, f := range fs {
		t2.AddRow(f.Name, f.Type.String())
	}
	t2.Print()
}

const appDescription = `parq is a tool for exploring parquet files.
		
It helps with viewing data in a parquet file, viewing a
file's schema, and converting data to/from parquet files.

Read more here: https://github.com/a-poor/parq
File issues here: https://github.com/a-poor/parq/issues
`

const cmdSchemaDesc = `Prints a table showing a parquet file's column names and data types.

Expects FILENAME to be a valid path to a parquet file with at least 
one row.

Example:

	$ parq schema path/to/iris.parquet

	Column Name   Data Type  
	Sepal_length  float64    
	Sepal_width   float64    
	Petal_length  float64    
	Petal_width   float64    
	Species       string

`

func main() {
	app := &cli.App{
		Name:      "parq",
		Usage:     "A tool for exploring parquet files",
		Version:   Version,
		Copyright: "Copyright (c) 2021 Austin Poor",
		Authors: []*cli.Author{
			{
				Name:  "Austin Poor",
				Email: "code@austinpoor.com",
			},
		},
		Description: appDescription,
		Commands: []*cli.Command{
			{
				Name:        "schema",
				Aliases:     []string{"s"},
				Usage:       "Shows a parquet file's column names and data types.",
				ArgsUsage:   "FILENAME",
				Flags:       []cli.Flag{},
				Description: cmdSchemaDesc,
				Action: func(c *cli.Context) error {
					// Check the number of arguments
					if c.NArg() < 1 {
						return cli.Exit("No file specified", 1)
					}
					if c.NArg() > 1 {
						return cli.Exit("Too many arguments. Expected 1.", 1)
					}

					// Get the file name argument
					fileName := c.Args().Get(0)

					// Read in the parquet file
					err := printParquetSchema(fileName)

					// Check return errors
					if err == ErrFileNotFound {
						return cli.Exit("Error: Can't find the specified file.", 1)
					}
					if err == ErrCantReadFile {
						return cli.Exit("Error: Can't read the specified file as parquet.", 1)
					}
					if err == ErrNoRowsInFile {
						return cli.Exit("Error: The specified file doesn't have any rows to read.", 1)
					}
					if err != nil {
						cli.Exit(err, 1)
					}

					return nil
				},
			},
			{
				Name:      "show-all",
				Aliases:   []string{"a"},
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
				Name:      "head",
				Aliases:   []string{"h"},
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
					fmt.Println("You're running 'head'")
					return nil
				},
			},
			{
				Name:      "tail",
				Aliases:   []string{"t"},
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
					fmt.Println("You're running 'tail'")
					return nil
				},
			},
			{
				Name:      "random",
				Aliases:   []string{"r"},
				Usage:     "Shows the n random rows of a parquet file.",
				ArgsUsage: "FILENAME",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:    "n-rows",
						Value:   10,
						Aliases: []string{"n"},
						Usage:   "Number of rows to show",
					},
					// Add --seed flag
				},
				Action: func(c *cli.Context) error {
					fmt.Println("You're running 'random'")
					return nil
				},
			},
			{
				Name:    "convert",
				Aliases: []string{"c"},
				Usage:   "Convert a parquet file to/from another format.",
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
	// Run the app and check for an error
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

	// log.Println(app.Name)
	// runSandbox()
}
