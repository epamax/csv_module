package main

import (
	"fmt"
	"math/rand"
)

/*

func WriteDataToCSV(csvFile string) {
	Generate rows for each combination of:
	SourceTypeIDs 11, 21, 31, 32, 41, 42, 43, 51, 52, 53, 54, 61, 62
	RegClassIDs 10, 20, 30, 41, 42, 46, 47, 48, 49
	FuelTypeIDs 1, 2, 3, 5, 9
	ModelYearIDs 1960-2060, inclusive
	OpModeIDs 0,1,11,12,13,14,15,16,21,22,23,24,25,26,27,28,29,30,33,35,36
	EmissionRate and EmissionRateIM should be random values
}
*/

func main() {

	//sourceTypeIDs := []int{11, 21, 31, 32, 41, 42, 43, 51, 52, 53, 54, 61, 62}
	//regClassIDs := []int{10, 20, 30, 41, 42, 46, 47, 48, 49}
	//fuelTypeIDs := []int{1, 2, 3, 5, 9}
	//modelYearIDs := intSliceRange(1960, 2060, true)
	//opModeIDs := []int{0, 1, 11, 12, 13, 14, 15, 16, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 35, 36}

	singleFl := randFloat(1)
	flList := randFloat(10)

	fmt.Println(singleFl)
	fmt.Println(flList)

	// empData := [][]string{}
	// 	{"Name", "City", "Skills"},
	// 	{"Smith", "Newyork", "Java"},
	// 	{"William", "Paris", "Golang"},
	// 	{"Rose", "London", "PHP"},

	// csvFile, err := os.Create("employee.csv")

	// if err != nil {
	// 	log.Fatalf("failed creating file: %s", err)
	// }
	// csvFile.Close()

	// csvwriter := csv.NewWriter(csvFile)

	// for _, empRow := range empData {
	// 	_ = csvwriter.Write(empRow)
	// }

	// csvwriter.Flush()
	// csvFile.Close()
}

func buildTable() {
	// iterate through keyValue struct for all combinations
}

func intSliceRange(start int, end int, inclusive bool) []int {
	var intSlice = []int{}
	// create range of values because go cant do that ????????

	if inclusive {

		for i := start; i <= end; i++ {
			intSlice = append(intSlice, i)
		}

	} else {
		for i := start; i < end; i++ {
			intSlice = append(intSlice, i)
		}

	}
	return intSlice
}

func randFloat(length int) (randFl any) {
	if length == 1 {
		randFl := rand.Float64()
		return randFl
	} else {
		randFl := make([]float64, length)
		for i := 0; i < length; i++ {
			randFl[i] = rand.Float64()
		}
		return randFl
	}
	return randFl
}
