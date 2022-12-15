package read_csv

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
)

// framework_io.ReadCSV(filepath, &dataTemplate{})
// return slice of templates
// make([]inputTest, 0)

// also errors

// vpop uses matrix with all years etc
// will tags be useful? think about it
// specifiy sperator
// how many rows to read
// header?
// look at other data readers
// NaNs?

// keys? index slice how? for larger files, return map instead of slices



func assignFields(emptyStruct interface{}, row []string) interface{} {

	structElements := reflect.ValueOf(emptyStruct).Elem()
	structTypes := structElements.Type()
	dataCopy := reflect.New(structTypes).Interface()
	// dataCopy := reflect.New(reflect.TypeOf(emptyStruct))

	for j := 0; j < structElements.NumField(); j++ {
		varName := structTypes.Field(j).Name
		varType := structTypes.Field(j).Type.Kind()

		switch varType {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
			reflect.Bool:
			val, _ := strconv.ParseInt(row[j], 10, 64) // can reflect type of struct field, but not sure how to make work for the vpop data
			//fmt.Println("int")
			//dataCopy.FieldByName(varName) = 5
			reflect.ValueOf(dataCopy).Elem().FieldByName(varName).SetInt(val)
		case reflect.Float32, reflect.Float64:
			val, _ := strconv.ParseFloat(row[j], 64)
			reflect.ValueOf(dataCopy).Elem().FieldByName(varName).SetFloat(val)
			//fmt.Println("float")
		case reflect.String:
			reflect.ValueOf(dataCopy).Elem().FieldByName(varName).SetString(row[j])
			//fmt.Println("string")
		default:
			fmt.Println("Data type not supported")
		}
	}
	return dataCopy
}

func readInput(emptyStruct interface{}, data [][]string) []interface{} {

	// use reflection to grab each struct member's name & type

	data_array := make([]interface{}, 0)

	for i := 0; i < len(data); i++ {
		if i == 0 {
			// header := data[i]
			// fmt.Println(header)
			continue
		} else {
			row := data[i]
			// data_copy := dereferenceIfPtr(assignFields(emptyStruct, row))
			data_copy := assignFields(emptyStruct, row)
			data_array = append(data_array, data_copy)
			// fmt.Println(reflect.ValueOf(inputAtom).Elem().FieldByName(varName).Set())
		}
	}
	return data_array
}

func GetValue(value interface{}) interface{} {

	if reflect.TypeOf(value).Kind() == reflect.Ptr {

		return reflect.ValueOf(value).Elem().Interface()

	} else {

		return value

	}

}

func GetFields(value interface{}) [][]string {
	fields := [][]string{}

	structElements := reflect.ValueOf(value).Elem()
	structTypes := structElements.Type()

	for j := 0; j < structElements.NumField(); j++ {
		varName := structTypes.Field(j).Name
		varType := structTypes.Field(j).Type.Kind()
		name_type := []string{varName, fmt.Sprint(varType)}
		fields = append(fields, name_type)
	}
	return fields
}

func ImportData(emptyStruct interface{}, dir string) []interface{} {

	// TODO: check if dir or single file

	// to iterate over multiple files

	//files, err := ioutil.ReadDir(dir)

	// for _, file := range files {
	// fn := strings.Join([]string{dir, file.Name()}, "\\")

	files, err := filepath.Glob(dir + "\\*.csv")

	if err != nil {
		log.Fatal(err)
	}

	fn := files[0] // change

	f, err := os.Open(fn)

	if err != nil {
		log.Fatal(err)
	}

	// remember to close the file at the end of the program
	defer f.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()

	if err != nil {
		log.Fatal(err)
	}
	data_array := readInput(emptyStruct, data)



	return data_array
}

func ArrToValues(data_array []interface{}) []interface{} {

	convData_Array := make([]interface{}, len(data_array))
	for i := 0; i < len(data_array); i++ {
		convData_Array[i] = GetValue(data_array[i])
	}
	return convData_Array
}

