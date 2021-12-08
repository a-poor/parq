package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

// Version is the current version of the CLI
const Version = "v0.1.0"

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
				Name:        "show",
				Aliases:     []string{"all", "a"},
				Usage:       "Shows all rows of a parquet file.",
				Description: cmdShowDesc,
				ArgsUsage:   "FILENAME",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:    "n-rows",
						Value:   10,
						Aliases: []string{"n"},
						Usage:   "Number of rows to show",
					},
				},
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

					// Print the parquet file
					idxr := showAllConfig{}
					err := printParquetFile(fileName, idxr)

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
				Name:        "head",
				Aliases:     []string{"h"},
				Usage:       "Shows the first n rows of a parquet file.",
				Description: cmdHeadDesc,
				ArgsUsage:   "FILENAME",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:    "n-rows",
						Value:   10,
						Aliases: []string{"n"},
						Usage:   "Number of rows to show",
					},
				},
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

					// Get the n-rows argument
					nrows := c.Int("n-rows")
					if nrows <= 0 {
						return cli.Exit("`--n-rows` must be >= 0", 1)
					}

					// Print the parquet file
					idxr := showHeadConfig{
						n: nrows,
					}
					err := printParquetFile(fileName, idxr)

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
				Name:        "tail",
				Aliases:     []string{"t"},
				Usage:       "Shows the last n rows of a parquet file",
				Description: cmdTailDesc,
				ArgsUsage:   "FILENAME",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:    "n-rows",
						Value:   10,
						Aliases: []string{"n"},
						Usage:   "Number of rows to show",
					},
				},
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

					// Get the n-rows argument
					nrows := c.Int("n-rows")
					if nrows <= 0 {
						return cli.Exit("`--n-rows` must be >= 0", 1)
					}

					// Print the parquet file
					idxr := showTailConfig{
						n: nrows,
					}
					err := printParquetFile(fileName, idxr)

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
				Name:        "random",
				Aliases:     []string{"r"},
				Usage:       "Shows the n random rows of a parquet file.",
				Description: cmdRandomDesc,
				ArgsUsage:   "FILENAME",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:    "n-rows",
						Value:   10,
						Aliases: []string{"n"},
						Usage:   "Number of rows to show",
					},
					&cli.IntFlag{
						Name:    "seed",
						Value:   0,
						Aliases: []string{"s"},
						Usage:   "Seed for the random number generator. If 0, the current time will be used.",
					},
				},
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

					// Get the "n-rows" argument
					nrows := c.Int("n-rows")
					if nrows <= 0 {
						return cli.Exit("`--n-rows` must be >= 0", 1)
					}

					// Get the "seed" argument
					seed := c.Int("seed")

					// Print the parquet file
					idxr := showRandomConfig{
						n:    nrows,
						seed: seed,
					}
					err := printParquetFile(fileName, idxr)

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

}
