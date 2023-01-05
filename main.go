package main

// TODO: generalize given header, could make map index given field names or struct tags if field different from header etc.
// TODO: could make seperate option to read line by line if want specific column or limit number of lines without reading in all data

import (
	"fmt"
	"math/rand"
	"reflect"
	"strings"
	"time"

	"github.com/epamax/csv_module/read_csv"
	"github.com/epamax/csv_module/write_csv"
)

type KeyValue struct {
	SourceTypeID   int
	RegClassID     int
	FuelTypeID     int
	ModelYearID    int
	OpModeID       int
	EmissionRate   float64
	EmissionRateIM float64
}

type Key struct {
	SourceTypeID int
	RegClassID   int
	FuelTypeID   int
	ModelYearID  int
	OpModeID     int
}

type Value struct {
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
	header := write_csv.GetHeader(&KeyValue{}) // could create struct field from header
	output := "C:\\Users\\mcolli03\\go\\scripts\\data\\"
	fn := "benchmarkTest.csv"

	for _, sourceType := range sourceTypeIDs {
		for _, regClass := range regClassIDs {
			for _, fuelType := range fuelTypeIDs {
				for _, modelYear := range modelYearIDs {
					for _, opMode := range opModeIDs {
						emissionRate, emissionRateIM := randFloat(1), randFloat(1)
						row := fmt.Sprintf("%v, %v, %v, %v, %v, %v, %v", sourceType, regClass, fuelType, modelYear, opMode, emissionRate, emissionRateIM)
						split_row := strings.Split(row, ",")
						write_csv.AddToTable(&output_table, split_row) // why did i recreate append
					}
				}
			}
		}
	}

	// 1,240,785 rows
	write_csv.WriteDataToCSV(output_table, header, output, fn)

	start := time.Now()
	// var dir string = "C:\\Users\\Public\\dozermodel2023\\calculator\\ageDistributionCalculator"

	// d, data_array := read_csv.ImportData[inputTest](&inputTest{}, dir)

	var dir string = output + fn // will be from call

	d, data_array := read_csv.ImportData[KeyValue](&KeyValue{}, dir)

	duration := time.Since(start)
	fmt.Printf("Time to read input file as slice of structs %v: %v \n", dir, duration)
	fmt.Println(data_array[0].SourceTypeID)
	fmt.Println(d.Fields)
	fmt.Println(d.Types)

	start = time.Now()
	emRate, emRateIM := LookupSliceOfStructs(data_array, 62, 49, 9, 2060, 36)
	duration = time.Since(start)

	fmt.Printf("Time taken to locate slice %v, %v: %v \n", emRate, emRateIM, duration)

	valFields := new(read_csv.Data)
	valFields.GetFields(&Value{})

	start = time.Now()
	dataMap := read_csv.BuildMapOfKeysToValues(valFields, dir, &Key{}, &Value{})
	duration = time.Since(start)

	fmt.Println("Time to build map: ", duration)

	start = time.Now()
	values := LookupMapOfKeysToValues(dataMap, 62, 49, 9, 2060, 36)
	duration = time.Since(start)

	fmt.Printf("Time to find values %v from map: %v \n", values, duration)

	start = time.Now()
	nestedMap := read_csv.BuildNestedMap(dir)
	duration = time.Since(start)

	fmt.Printf("Time to build nested map: %v \n", duration)

	start = time.Now()
	nestedVals := LookupNestedMap(nestedMap, "62", "49", "9", "2060", "36")
	duration = time.Since(start)

	fmt.Printf("Time to find values %v from nested map: %v \n", nestedVals, duration)
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

func LookupSliceOfStructs(data []KeyValue, sourceTypeID, regClassID, fuelTypeID, modelYearID, opModeID int) (float64, float64) {
	// return the correct KeyValue in data corresponding to the input values
	for _, val := range data { // this feels wrong help
		// fmt.Println(val)
		if (val.SourceTypeID == sourceTypeID) && (val.RegClassID == regClassID) && (val.FuelTypeID == fuelTypeID) && (val.ModelYearID == modelYearID) && (val.OpModeID == opModeID) {
			return val.EmissionRate, val.EmissionRateIM
		}
	}
	panic("Value not found")
}

func LookupMapOfKeysToValues(data map[string]reflect.Value, sourceTypeID, regClassID, fuelTypeID, modelYearID, opModeID int) reflect.Value {
	// return the correct Value in data corresponding to the input values
	key := fmt.Sprintf("%v, %v, %v, %v, %v", sourceTypeID, regClassID, fuelTypeID, modelYearID, opModeID)
	return data[key]
}

// TODO: fix types data map[int]map[int]map[int]map[int]map[int]Value
func LookupNestedMap(data map[string]map[string]map[string]map[string]map[string][]string, sourceTypeID, regClassID, fuelTypeID, modelYearID, opModeID string) []string {
	// return the correct Value in data corresponding to the input values
	return data[sourceTypeID][regClassID][fuelTypeID][modelYearID][opModeID]
}
