package main

import (
	"fmt"
	"math/rand"
	"read_csv"
	"strings"
	"time"
	"write_csv"
)

type inputTest struct {
	CalendarYearID int // how would calendar year be input?
	ModelYearID    int // top row, model year <= calendar year
	AgeFractionID  float64
	//	SizeClassID int
	//	TierID      string
	// VPOP float64
}

type KeyValue struct {
	SourceTypeID   int
	RegClassID     int
	FuelTypeID     int
	ModelYearID    int
	OpModeID       int
	EmissionRate   float64
	EmissionRateIM float64
}

// or read from list of keys? variadic ...

func main() {

	var sourceTypeIDs = []int{11, 21, 31, 32, 41, 42, 43, 51, 52, 53, 54, 61, 62}
	var regClassIDs = []int{10, 20, 30, 41, 42, 46, 47, 48, 49}
	var fuelTypeIDs = []int{1, 2, 3, 5, 9}
	var modelYearIDs = intSliceRange(1960, 2060, true)
	var opModeIDs = []int{0, 1, 11, 12, 13, 14, 15, 16, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 33, 35, 36}

	output_table := [][]string{}
	header := write_csv.GetHeader(&KeyValue{})
	output := "C:\\Users\\mcolli03\\Documents\\testing\\data\\"
	fn := "benchmarkTest.csv"

	for _, sourceType := range sourceTypeIDs {
		for _, regClass := range regClassIDs {
			for _, fuelType := range fuelTypeIDs {
				for _, modelYear := range modelYearIDs {
					for _, opMode := range opModeIDs {
						emissionRate, emissionRateIM := randFloat(1), randFloat(1)
						row := fmt.Sprintf("%v, %v, %v, %v, %v, %v, %v", sourceType, regClass, fuelType, modelYear, opMode, emissionRate, emissionRateIM)
						split_row := strings.Split(row, ",")
						write_csv.AddToTable(&output_table, split_row)
					}
				}
			}
		}
	}

	write_csv.WriteDataToCSV(output_table, header, output, fn)

	start := time.Now()
	// var dir string = "C:\\Users\\Public\\dozermodel2023\\calculator\\ageDistributionCalculator" // will be from call

	// d, data_array := read_csv.ImportData[inputTest](&inputTest{}, dir) // not sure how the internal Data struct works

	var dir string = output + fn // will be from call

	d, data_array := read_csv.ImportData[KeyValue](&KeyValue{}, dir) // not sure how the internal Data struct works

	duration := time.Since(start)
	fmt.Println("Time to read input file %v: %v", dir, duration)
	fmt.Println(data_array[0].SourceTypeID)
	fmt.Println(d.Fields)
	fmt.Println(d.Types)

	// fmt.Println(data_array) // how to deal with this?
	// save original data_array internally to check fields etc.?
	// ! only need to save one interface to check fields
	// I think? brain hurt how do
	// create some kind of interface in package to hold fields and use that to index/get value etc

	// fields := read_csv.GetFields(data_array[0]) // show field types are included, but may be better to make optional so can use in passing to other functions
	// or -> fields := getFields(data_array[0])

	// fieldByIndex := read_csv.FieldByIndex(0, data_array[0]) // returns field name
	// typeByIndex := read_csv.TypeByIndex(0, data_array[0])
	// typeByField := read_csv.TypeByField(fieldByIndex, data_array[0])
	// valByIndex := read_csv.ValByIndex(0, data_array[0])
	// valByField := read_csv.ValByField(fieldByIndex, data_array[0])

	//fmt.Println("Struct fields:", fields)

	// fmt.Println("Field name:", fieldByIndex)
	// fmt.Println("Field type by index:", typeByIndex)
	// fmt.Println("Field type by name:", typeByField)
	// fmt.Println("Field value by index:", valByIndex)
	// fmt.Println("Field value by name:", valByField)

	//fmt.Println(data_array[0].CalendarYearID)

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
}
