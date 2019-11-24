## QBranch compiler

QBranch compiler, short for qbc, provides a command line tool to convert
given Bond files to Go source code files.

### Command line format

Qbc's command line looks like below:

```
qbc -include .;../folder1;../folder2
    -structprefix C
    -fieldprefix F
    <file1.bond> <file2.bond> <file3.bond>

```
