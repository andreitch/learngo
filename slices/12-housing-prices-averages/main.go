package main

import (
	"fmt"
	"strconv"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Housing Prices and Averages
//
//  Use the previous exercise to solve this exercise (Housing Prices).
//
//  Your task is to print the averages of the sizes, beds, baths, and prices.
//
//
// EXPECTED OUTPUT
//
//  Location       Size           Beds           Baths          Price
//  ===========================================================================
//  New York       100            2              1              100000
//  New York       150            3              2              200000
//  Paris          200            4              3              400000
//  Istanbul       500            10             5              1000000
//
//  ===========================================================================
//                 237.50         4.75           2.75           425000.00
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
		locations                                                  []string
		sizes, beds, paths, prices                                 []float64
		totalSizes, totalBeds, totalPaths, totalPrices, total, tmp float64
	)
	for _, row := range strings.Split(data, "\n") {
		cols := strings.Split(row, separator)
		locations = append(locations, cols[0])
		tmp, _ = strconv.ParseFloat(cols[1], 64)
		sizes = append(sizes, tmp)
		totalSizes += float64(tmp)
		tmp, _ = strconv.ParseFloat(cols[2], 64)
		beds = append(beds, tmp)
		totalBeds += tmp
		tmp, _ = strconv.ParseFloat(cols[3], 64)
		paths = append(paths, tmp)
		totalPaths += tmp
		tmp, _ = strconv.ParseFloat(cols[4], 64)
		prices = append(prices, tmp)
		totalPrices += tmp
	}

	for _, el := range strings.Split(header, separator) {
		fmt.Printf("%-15s", el)
	}
	fmt.Printf("\n%s\n", strings.Repeat("=", 75))

	for i := range locations {
		fmt.Printf("%-15s", locations[i])
		fmt.Printf("%-15.2f", sizes[i])
		fmt.Printf("%-15.2f", beds[i])
		fmt.Printf("%-15.2f", paths[i])
		fmt.Printf("%-15.2f", prices[i])
		fmt.Println()
	}
	total = float64(len(locations))
	fmt.Printf("\n%s\n", strings.Repeat("=", 75))
	fmt.Printf("%-15s", " ")
	fmt.Printf("%-15.2f", totalSizes/total)
	fmt.Printf("%-15.2f", float64(totalBeds)/total)
	fmt.Printf("%-15.2f", float64(totalPaths)/total)
	fmt.Printf("%-15.2f", float64(totalPrices)/total)
	fmt.Println()
}
