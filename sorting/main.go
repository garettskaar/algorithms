package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// User chosen sorting algorithm
	var sortAlgo string
	flag.StringVar(&sortAlgo, "sort", "quick", "The sorting algorithm to use.")
	// User chosen input file
	input_file := flag.String("input", "unsorted_array.txt", "The input array to be sorted.\nThis can be a text file with comma/space seperated values.")
	// Parse command line args
	flag.Parse()
	// Read input file
	data, err := os.ReadFile(*input_file)
	if err != nil {
		log.Fatal(err)
	}
	// Initialize []int from file contents
	sortMe := initArrayFromFile(string(data))
	// User input to lower case
	strings.ToLower(sortAlgo)
	switch sortAlgo {
	case "quick":
		fmt.Println("Using quick sort...")
		//Todo call quick sort here
		printIntArray(sortMe)
	case "selection":
		fmt.Println("Using selection sort...")
		//Todo call selection sort here
	case "merge":
		fmt.Println("Using merge sort...")
		//Todo call merge sort here
	case "insertion":
		fmt.Println("Using insertion sort...")
		//Todo call insertion sort here
	case "bubble":
		fmt.Println("Using bubble sort...")
		//Todo call insertion sort here
	default:
		fmt.Println("Using default quick sort...")
		//Todo call quick sort here
		printIntArray(sortMe)
	}

}
func initArrayFromFile(fileContent string) []int {
	// Slice of number stings
	var nums []string
	if strings.Contains(fileContent, ",") {
		// Comma seperated numbers
		nums = strings.Split(fileContent, " ")
	} else if strings.Contains(fileContent, " ") {
		// Space seperated numbers
		nums = strings.Split(fileContent, " ")
	}
	// int unsorted array
	unsortedArray := make([]int, len(nums))
	// Fill the int array
	for i, n := range nums {
		num, err := strconv.Atoi(n)
		if err != nil {
			log.Fatal(err)
		}
		unsortedArray[i] = num
	}

	return unsortedArray
}

func printIntArray(arr []int) {
	for _, n := range arr {
		fmt.Printf("%d ", n)
	}
}
