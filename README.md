# parq

[![Test](https://github.com/a-poor/parq/actions/workflows/test.yml/badge.svg)](https://github.com/a-poor/parq/actions/workflows/test.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/a-poor/parq.svg)](https://pkg.go.dev/github.com/a-poor/parq)
[![parq](https://snapcraft.io/parq/badge.svg)](https://snapcraft.io/parq)
[![parq](https://snapcraft.io/parq/trending.svg?name=0)](https://snapcraft.io/parq)

_created by Austin Poor_

A CLI for examining parquet files.

## About

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
   v0.1.0

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

## Installation

### Homebrew

```bash
brew tap a-poor/parq
brew install parq
```

### Snap

[![Get it from the Snap Store](https://snapcraft.io/static/images/badges/en/snap-store-white.svg)](https://snapcraft.io/parq)

```bash
snap install parq
```

_NOTE: There seem to be some issues currently with the snap install (see [here](https://github.com/a-poor/parq/issues/8)). Try installing with [Homebrew](https://github.com/a-poor/parq#homebrew), [go install](https://github.com/a-poor/parq#go-install), or from the [releases page](https://github.com/a-poor/parq/releases)._

## Go Install

```bash
go install github.com/a-poor/parq@latest
```

## Precompiled Binaries

Check out the [repo's releases page](https://github.com/a-poor/parq/releases).

## License

[MIT License](./LICENSE)

## Contributing

Contributions are welcome!

* Suggest new features
* Report bugs
* Add docs
* Add tests

Or, just say hi and let me know if this app has been helpful!


