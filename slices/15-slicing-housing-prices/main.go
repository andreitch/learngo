package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Slicing the Housing Prices
//
//  We have received housing prices. Your task is to print only the requested
//  columns of data (see the expected output).
//
//  1. Separate the data by the newline ("\n") -> rows
//
//  2. Separate each row of the data by the separator (",") -> columns
//
//  3. Find the from and to positions inside the columns depending
//     on the command-line arguments.
//
//  4. Print only the requested column headers and their data
//
//
// RESTRICTIONS
//
//  + You should use slicing when printing the columns.
//
//
// EXPECTED OUTPUT
//
//  go run main.go
//    Location       Size           Beds           Baths          Price
//
//    New York       100            2              1              100000
//    New York       150            3              2              200000
//    Paris          200            4              3              400000
//    Istanbul       500            10             5              1000000
//
//
//  go run main.go Location
//    Location       Size           Beds           Baths          Price
//
//    New York       100            2              1              100000
//    New York       150            3              2              200000
//    Paris          200            4              3              400000
//    Istanbul       500            10             5              1000000
//
//
//  NOTE : Finds the position of the Size column and slices the columns
//         appropriately.
//
//  go run main.go Size
//    Size           Beds           Baths          Price
//
//    100            2              1              100000
//    150            3              2              200000
//    200            4              3              400000
//    500            10             5              1000000
//
//
//  NOTE : Finds the positions of the Size and Baths columns and
//         slices the columns appropriately.
//
//  go run main.go Size Baths
//    Size           Beds           Baths
//
//    100            2              1
//    150            3              2
//    200            4              3
//    500            10             5
//
//
//  go run main.go Beds Price
//    Beds           Baths          Price
//
//    2              1              100000
//    3              2              200000
//    4              3              400000
//    10             5              1000000
//
//
//  Note : It works even if the given column name does not exist.
//
//  go run main.go Beds NotExist
//    Beds           Baths          Price
//
//    2              1              100000
//    3              2              200000
//    4              3              400000
//    10             5              1000000
//
//
//  go run main.go NotExist NotExist
//    Location       Size           Beds           Baths          Price
//
//    New York       100            2              1              100000
//    New York       150            3              2              200000
//    Paris          200            4              3              400000
//    Istanbul       500            10             5              1000000
//
//
// Note : It works even if the Price's index > Size's index
//
//        In that case, it resets the invalid starting position to 0.
//
//        But it still uses the Size column.
//
//  go run main.go Price Size
//    Location       Size
//
//    New York       100
//    New York       150
//    Paris          200
//    Istanbul       500
//
//
// HINTS
//
//  + strings.Split function can separate a string into a []string
//
// ---------------------------------------------------------

func main() {
	const (
		header = "Location,Size,Beds,Baths,Price"
		data   = `New York,100,2,1,100000
New York,150,3,2,200000
Paris,200,4,3,400000
Istanbul,500,10,5,1000000`

		separator = ","
	)

	var (
		tmp                        int
		locations                  []string
		sizes, beds, paths, prices []int
	)

	for _, row := range strings.Split(data, "\n") {
		cols := strings.Split(row, separator)
		locations = append(locations, cols[0])
		tmp, _ = strconv.Atoi(cols[1])
		sizes = append(sizes, tmp)
		tmp, _ = strconv.Atoi(cols[2])
		beds = append(beds, tmp)
		tmp, _ = strconv.Atoi(cols[3])
		paths = append(paths, tmp)
		tmp, _ = strconv.Atoi(cols[4])
		prices = append(prices, tmp)
	}

	start, end := 0, 5
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "Size":
			start = 1
		case "Beds":
			start = 2
		case "Baths":
			start = 3
		case "Price":
			start = 4
		default:
			start = 0
		}
	}
	if len(os.Args) > 2 {
		switch os.Args[2] {
		case "Size":
			end = 1
		case "Beds":
			end = 2
		case "Baths":
			end = 3
		case "Price":
			end = 4
		default:
			end = 5
		}
	}
	if end < start {
		start = 0
	}

	var br bool
	for i, el := range strings.Split(header, separator) {
		if i < start || i > end {
			continue
		}
		br = true
		fmt.Printf("%-15s", el)
	}
	if br {
		fmt.Println()
		fmt.Println()
	}

	for i := range locations {
		if start < 1 && end >= 0 {
			fmt.Printf("%-15s", locations[i])
		}
		if start < 2 && end >= 1 {
			fmt.Printf("%-15d", sizes[i])
		}
		if start < 3 && end >= 2 {
			fmt.Printf("%-15d", beds[i])
		}
		if start < 4 && end >= 3 {
			fmt.Printf("%-15d", paths[i])
		}
		if start < 5 && end >= 4 {
			fmt.Printf("%-15d", prices[i])
		}
		if br {
			fmt.Println()
		}
	}

}
