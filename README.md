# parq

_created by Austin Poor_

A CLI for examining parquet files.

`parq` has the following capabilities:
* _schema_: Shows a parquet file's column names and data types.
* _show_: Shows all rows of a parquet file.
* _head_: Shows the first n rows of a parquet file.
* _tail_: Shows the last n rows of a parquet file
* _random_: Shows the n random rows of a parquet file.
* _convert_: (TODO) Convert a parquet file to/from another format

```terminal
$ parq 
NAME:
   parq - A tool for exploring parquet files

USAGE:
   parq [global options] command [command options] [arguments...]

VERSION:
   v0.0.6

DESCRIPTION:
   parq is a tool for exploring parquet files.
       
   parq helps with viewing data in a parquet file, viewing a
   file's schema, and converting data to/from parquet files.
   
   Read more here: https://github.com/a-poor/parq
   Submit issues here: https://github.com/a-poor/parq/issues

AUTHOR:
   Austin Poor <code@austinpoor.com>

COMMANDS:
   schema, s     Shows a parquet file's column names and data types.
   show, all, a  Shows all rows of a parquet file.
   head, h       Shows the first n rows of a parquet file.
   tail, t       Shows the last n rows of a parquet file
   random, r     Shows the n random rows of a parquet file.
   convert, c    Convert a parquet file to/from another format.
   help, h       Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help (default: false)
   --version, -v  print the version (default: false)

COPYRIGHT:
   Copyright (c) 2021 Austin Poor
```
