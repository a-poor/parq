package main

const appDescription = `parq is a tool for exploring parquet files.
		
parq helps with viewing data in a parquet file, viewing a
file's schema, and converting data to/from parquet files.

Read more here: https://github.com/a-poor/parq
Submit issues here: https://github.com/a-poor/parq/issues
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


Related Commands: show, head, tail, random
`

const cmdShowDesc = `Prints the full data contained in the specified parquet file as a formatted table.

Expects FILENAME to be a valid path to a parquet file with at least 
one row.

Example:

	$ parq show path/to/iris.parquet

	     Sepal_length  Sepal_width  Petal_length  Petal_width  Species  
	0    5.1           3.5          1.4           0.2          setosa   
	1    4.9           3            1.4           0.2          setosa   
	2    4.7           3.2          1.3           0.2          setosa   
	3    4.6           3.1          1.5           0.2          setosa   
	4    5             3.6          1.4           0.2          setosa   
	5    5.4           3.9          1.7           0.4          setosa   
	...
	145  6.7           3            5.2           2.3          virginica  
	146  6.3           2.5          5             1.9          virginica  
	147  6.5           3            5.2           2            virginica  
	148  6.2           3.4          5.4           2.3          virginica  
	149  5.9           3            5.1           1.8          virginica


Related Commands: schema, head, tail, random
`

const cmdHeadDesc = `Prints the first "n-rows" rows of data contained in the specified parquet file as a formatted table.

Expects FILENAME to be a valid path to a parquet file with at least 
one row.

Example:

	$ parq head path/to/iris.parquet

	   Sepal_length  Sepal_width  Petal_length  Petal_width  Species  
	0  5.1           3.5          1.4           0.2          setosa   
	1  4.9           3            1.4           0.2          setosa   
	2  4.7           3.2          1.3           0.2          setosa   
	3  4.6           3.1          1.5           0.2          setosa   
	4  5             3.6          1.4           0.2          setosa   
	5  5.4           3.9          1.7           0.4          setosa   
	6  4.6           3.4          1.4           0.3          setosa   
	7  5             3.4          1.5           0.2          setosa   
	8  4.4           2.9          1.4           0.2          setosa   
	9  4.9           3.1          1.5           0.1          setosa


Related Commands: schema, show, tail, random
`

const cmdTailDesc = `Prints the last "n-rows" rows of data contained in the specified parquet file as a formatted table.

Expects FILENAME to be a valid path to a parquet file with at least 
one row.

Example:

	$ parq head path/to/iris.parquet

	     Sepal_length  Sepal_width  Petal_length  Petal_width  Species    
	140  6.7           3.1          5.6           2.4          virginica  
	141  6.9           3.1          5.1           2.3          virginica  
	142  5.8           2.7          5.1           1.9          virginica  
	143  6.8           3.2          5.9           2.3          virginica  
	144  6.7           3.3          5.7           2.5          virginica  
	145  6.7           3            5.2           2.3          virginica  
	146  6.3           2.5          5             1.9          virginica  
	147  6.5           3            5.2           2            virginica  
	148  6.2           3.4          5.4           2.3          virginica  
	149  5.9           3            5.1           1.8          virginica


Related Commands: schema, show, head, random
`

const cmdRandomDesc = `Prints "n-rows" rows randomly selected from the specified parquet file as a formatted table.

Note: Currently, the rows will be randomly selected WITH replacement.

Expects FILENAME to be a valid path to a parquet file with at least 
one row.

The RNG can be seeded using the "--seed" flag. If set to the default value of 0, the current system time will be used.

Example:

	$ parq head path/to/iris.parquet

	     Sepal_length  Sepal_width  Petal_length  Petal_width  Species    
	140  6.7           3.1          5.6           2.4          virginica  
	141  6.9           3.1          5.1           2.3          virginica  
	142  5.8           2.7          5.1           1.9          virginica  
	143  6.8           3.2          5.9           2.3          virginica  
	144  6.7           3.3          5.7           2.5          virginica  
	145  6.7           3            5.2           2.3          virginica  
	146  6.3           2.5          5             1.9          virginica  
	147  6.5           3            5.2           2            virginica  
	148  6.2           3.4          5.4           2.3          virginica  
	149  5.9           3            5.1           1.8          virginica


Related Commands: schema, show, head, tail
`
