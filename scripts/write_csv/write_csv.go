package write_csv

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"reflect"
	"strings"
	"time"
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

/*
var sourceTypeIDs = []int{11, 21, 31, 32, 41, 42, 43, 51, 52, 53, 54, 61, 62}
var regClassIDs = []int{10, 20, 30, 41, 42, 46, 47, 48, 49}
var fuelTypeIDs = []int{1, 2, 3, 5, 9}
var modelYearIDs = intSliceRange(1960, 2060, true)
var opModeIDs = []int{0, 1, 11, 12, 13, 14, 15, 16, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 33, 35, 36}

func main() {

	data_table := [][]string{}
	output_table := [][]string{}
	header := getHeader(&KeyValue{})

	// want to be table to build the data from the data_table, idk if its useful at all
	addToTable(&data_table, sourceTypeIDs) // this is literally just append function
	addToTable(&data_table, regClassIDs)
	addToTable(&data_table, fuelTypeIDs)
	addToTable(&data_table, modelYearIDs)
	addToTable(&data_table, opModeIDs)

	// data_table more like input table, then create table of all
	// combinations, and pass to write csv to loop through

	for _, sourceType := range sourceTypeIDs {
		for _, regClass := range regClassIDs {
			for _, fuelType := range fuelTypeIDs {
				for _, modelYear := range modelYearIDs {
					for _, opMode := range opModeIDs {
						emissionRate, emissionRateIM := randFloat(1), randFloat(1)
						row := fmt.Sprintf("%v, %v, %v, %v, %v, %v, %v", sourceType, regClass, fuelType, modelYear, opMode, emissionRate, emissionRateIM)
						split_row := strings.Split(row, ",")
						addToTable(&output_table, split_row)
					}
				}
			}
		}
	}

	//fmt.Println(len(output_table))

	WriteDataToCSV(output_table, header) // (header, data, fn)

}
*/

func addToTable[t any](data_table *[][]string, dataToAdd []t) {
	joinedData := strings.Trim(strings.Replace(fmt.Sprint(dataToAdd), " ", ",", -1), "[]")
	sepData := strings.Fields(joinedData)
	*data_table = append(*data_table, sepData)

}

func WriteDataToCSV(data [][]string, header []string) {
	start := time.Now()
	// iterate through keyValue struct for all combinations

	if fileExists("benchmarkSample.csv") {
		panic("File already exists.")
	}

	csvFile, err := os.Create("benchmarkSample.csv")

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	csvwriter := csv.NewWriter(csvFile)

	_ = csvwriter.Write(header)
	// TODO: call from outside, maybe split functions, WriteRow or something
	for _, row := range data {
		_ = csvwriter.Write(row)
	}

	csvwriter.Flush()
	csvFile.Close()
	duration := time.Since(start)
	fmt.Println("Total time spent: ", duration)
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

func getHeader(emptyStruct any) []string {
	fields := []string{}
	structElements := reflect.ValueOf(emptyStruct).Elem()

	for i := 0; i < structElements.NumField(); i++ {
		varName := structElements.Type().Field(i).Name
		fields = append(fields, varName)
	}
	return fields
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
