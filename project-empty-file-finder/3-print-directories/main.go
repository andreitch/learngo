package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// ---------------------------------------------------------
// EXERCISE: Find and write the names of subdirectories to a file
//
//  Create a program that can get multiple directory paths from
//  the command-line, and prints only their subdirectories into a
//  file named: dirs.txt
//
//
//  1. Get the directory paths from command-line
//
//  2. Append the names of subdirectories inside each directory
//     to a byte slice
//
//  3. Write that byte slice to dirs.txt file
//
//
// EXPECTED OUTPUT
//
//   go run main.go
//     Please provide directory paths
//
//   go run main.go dir/ dir2/
//
//   cat dirs.txt
//
//     dir/
//             subdir1/
//             subdir2/
//
//     dir2/
//             subdir1/
//             subdir2/
//             subdir3/
//
//
// HINTS
//
//   ONLY READ THIS IF YOU GET STUCK
//
//   + Get all the files in a directory using ioutil.ReadDir
//     (A directory is also a file)
//
//   + You can use IsDir method of a FileInfo value to detect
//     whether a file is a directory or not.
//
//     Check out its documentation:
//
//     go doc os.FileInfo.IsDir
//
//     # or using godocc
//     godocc os.FileInfo.IsDir
//
//   + You can use '\t' escape sequence for indenting the subdirs.
//
//   + You can find a sample directory structure under:
//     solution/ directory
//
// ---------------------------------------------------------

func main() {
	dirs := os.Args[1:]
	if len(dirs) == 0 {
		fmt.Println("Please provide directory paths")
		return
	}

	var text []byte
	for _, dir := range dirs {
		tmp := listDir(dir, dir+"/", "")
		text = append(text, tmp...)
	}

	ioutil.WriteFile("dirs.txt", text, 0644)
}

func listDir(s string, curr string, space string) []byte {
	text := []byte(space + curr + "\n")
	files, err := ioutil.ReadDir(s)
	if err == nil {
		for i := 0; i < len(files); i++ {
			if files[i].IsDir() {
				tmp := listDir(s+"/"+files[i].Name(), files[i].Name()+"/", space+"   ")
				text = append(text, tmp...)
			} else {
				tmp := "   " + space + files[i].Name() + "\n"
				text = append(text, tmp...)
			}
		}
	}
	return text
}
