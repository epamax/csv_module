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

// or read from list of keys? variadic ...

func main() {

	var dir string = "C:\\Users\\Public\\dozermodel2023\\calculator\\ageDistributionCalculator" // will be from call

	d, data_array := read_csv.ImportData[inputTest](&inputTest{}, dir) // not sure how the internal Data struct works
	fmt.Println(data_array)
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

	fmt.Println(data_array[0].CalendarYearID)

}
