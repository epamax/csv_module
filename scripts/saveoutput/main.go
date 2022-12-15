package main

import (
	"fmt"
	"example.com/read_csv"
)

type inputTest struct {
	CalendarYearID int // how would calendar year be input?
	ModelYearID    int // top row, model year <= calendar year
	AgeFractionID  float64
	//	SizeClassID int
	//	TierID      string
	// VPOP float64
}

func main() {

	var dir string = "C:\\Users\\Public\\dozermodel2023\\calculator\\ageDistributionCalculator" // will be from call
	
	data_array := read_csv.ImportData(&inputTest{}, dir)
	fmt.Println(data_array)
	convData_Array := read_csv.ArrToValues(data_array)
	fmt.Println(convData_Array)

	fields := read_csv.GetFields(data_array[0])
		// or -> fields := getFields(data_array[0])
	fmt.Println("Struct fields:", fields)
	//fmt.Println(data_array)
	// fmt.Println(reflect.ValueOf(data_array[0]).Elem().Type())
	//fmt.Println("Value of first point:", getValue(data_array[0]))

	fmt.Println(read_csv.GetValue(data_array[0]).CalendarYearID)

}
