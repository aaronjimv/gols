### Gols
---

This project is a clone of the Linux shell command `ls` that lists directory contents of files and directories maded with `go` language program.

It utilizes the `github.com/fatih/color` package for colorful output and `github.com/AJRDRGZ/fileinfo` for extracting user and group information. The `main.go` file contains the main functionality, while `gols.go` defines constants and types used in the program. 

#### flags:

* -p: Filter files by a pattern.
* -a: Include hidden files.
* -n: Number of records to display.
* -t: Sort by modification time (oldest first).
* -s: Sort by file size (smallest first).
* -r: Reverse order while sorting.