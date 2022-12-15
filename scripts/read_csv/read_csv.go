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

// TODO: re-assign structs and values to return? doesn't seem to retain the fields after returning
// TODO: also indexby function to create a map to index the 
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

// type Data struct {
// 	fields interface{}
// 	types ...
// 	indexed map[string]interface{} 
// }

func assignFields(emptyStruct interface{}, row []string) interface{} {

	structElements := reflect.ValueOf(emptyStruct).Elem()
	structTypes := structElements.Type()
	dataCopy := reflect.New(structTypes).Interface()
	// dataCopy := emptyStruct
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

// func (d *Data), what does * do here
// d.fields := getfields, map with field, index?
// d.types = := gettypes, map with type, index?
// d.indexby := map[indexbystring]value

// fieldbyindex := d.fields[index] = field
// doesn't matter data_array index, will all be same
// but matters for value

// valbyindex := data_array[index]
// want to be able to do something like data_array.getvalue(index = 0, field = 0)
// easier to iterate through whole array

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
	
	// d.data_array := data_array[0]

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

// index of?

func TypeByField(field string, dataStruct interface{}) reflect.Type {
	return reflect.ValueOf(dataStruct).Elem().FieldByName(field).Type()
}

func TypeByIndex(index int, dataStruct interface{}) reflect.Type {
	return reflect.ValueOf(dataStruct).Elem().Field(index).Type()
}

func ValByField(field string, dataStruct interface{}) reflect.Value {
	return reflect.ValueOf(dataStruct).Elem().FieldByName(field)
}

func ValByIndex(index int, dataStruct interface{}) reflect.Value {
	return reflect.ValueOf(dataStruct).Elem().Field(index)
}

func FieldByIndex(index int, dataStruct interface{}) string {
	structElements := reflect.ValueOf(dataStruct).Elem()
	structTypes := structElements.Type()

	if index  < structElements.NumField() {
		varName := structTypes.Field(index).Name
		return varName
		// varType := structTypes.Field(j).Type.Kind()
	} else {
		fmt.Println("Index out of bounds for struct of length", structElements.NumField())
		return ""
		}
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

