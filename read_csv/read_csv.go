package read_csv

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/epamax/csv_module/decoder"
)

// TODO: re-assign structs and values to return? doesn't seem to retain the fields after returning
// TODO: also indexby function to create a map to index the
// TODO: could have seperate option depending on limiting amount of data you want to read (i.e. read all vs read row by row)
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

type Data struct {
	Fields map[int]string
	Types  map[int]reflect.Kind
	// indexed map[string]interface{}
	prevValue []string
}

func assignFields(emptyStruct interface{}, row []string) interface{} {

	structElements := reflect.ValueOf(emptyStruct).Elem()
	structTypes := structElements.Type()
	dataCopy := reflect.New(structTypes).Interface()
	// dataCopy := emptyStruct
	// dataCopy := reflect.New(reflect.TypeOf(emptyStruct))

	for j := 0; j < structElements.NumField(); j++ {
		varName := structTypes.Field(j).Name
		varType := structTypes.Field(j).Type.Kind()
		//fmt.Println(j)
		//fmt.Println(row[j])
		//fmt.Println(varName)

		switch varType {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
			reflect.Bool:
			val, _ := strconv.ParseInt(row[j], 0, 64) // can reflect type of struct field, but not sure how to make work for the vpop data
			// fmt.Println(val, err, row[j])
			reflect.ValueOf(dataCopy).Elem().FieldByName(varName).SetInt(val)
		case reflect.Float32, reflect.Float64:
			val, _ := strconv.ParseFloat(row[j], 64)
			// fmt.Println(val, row[j])
			reflect.ValueOf(dataCopy).Elem().FieldByName(varName).SetFloat(val)
			// fmt.Println(val)
		case reflect.String:
			reflect.ValueOf(dataCopy).Elem().FieldByName(varName).SetString(row[j])
			// fmt.Println(row[j])
		default:
			fmt.Println("Data type not supported")
		}

		// fmt.Println(dataCopy)
	}
	return dataCopy
}

// func (d *Data) readInput(emptyStruct interface{}, data [][]string) []interface{} {
func readInput(emptyStruct interface{}, data [][]string) []interface{} {

	// use reflection to grab each struct member's name & type

	data_array := make([]interface{}, 0)

	for i := 0; i < len(data); i++ {
		if i == 0 {
			// header := data[i]
			// fmt.Println(header)
			continue
		}

		row := strings.Fields(strings.Trim(strings.Replace(fmt.Sprint(data[i]), "  ", " ", -1), "[]"))
		// data_copy := dereferenceIfPtr(assignFields(emptyStruct, row))
		data_copy := assignFields(emptyStruct, row)
		// fmt.Println(data_copy)
		data_array = append(data_array, data_copy)
		// fmt.Println(reflect.ValueOf(inputAtom).Elem().FieldByName(varName).Set())
		// d.getFields(data_copy)
	}

	// d.data_array := data_array[0]

	return data_array
}

func BuildMapOfKeysToValues(d *Data, dir string, keyStruct interface{}, valueStruct interface{}) map[string]reflect.Value {
	// parse CSV file into map[Key]Value
	var newVar any
	var elemSlice reflect.Value

	data := readFile(dir)

	data_map := make(map[string]reflect.Value)

	keyElements := reflect.ValueOf(keyStruct).Elem()
	valElements := reflect.ValueOf(valueStruct).Elem()

	types := decoder.TypeAssertion(valueStruct)
	sameType := checkSameType(types)
	varType := types[d.Fields[0]]
	newVar = reflect.New(varType).Interface()

	for i, row := range data {
		var key string

		// initialize new slice of values for each row
		if sameType {
			elemSlice = reflect.MakeSlice(reflect.SliceOf(varType), 0, valElements.NumField())
		} else {
			panic("Cannot make slice of multiple types")
		}

		// header
		if i == 0 {
			continue
		}

		// go through all keys and create and then through values and add
		for j := 0; j < keyElements.NumField(); j++ {
			key += row[j] + ","
		}
		key = strings.Trim(key, ", ")

		// have to declare slice before so doesnt keep getting re initialized or something
		// if len types > 1, check if all same since can't return slice of different types
		// maybe idea for later, if they are different types return as string and parse after?

		for k := keyElements.NumField(); k < (keyElements.NumField() + valElements.NumField()); k++ {
			index := k - keyElements.NumField()

			varType := types[d.Fields[index]]
			newVar = reflect.New(varType).Interface()

			_ = decoder.Set(&newVar, row[k])
			val := decoder.InterToVal(newVar)
			elemSlice = reflect.Append(elemSlice, val)
		}
		// fmt.Println(elemSlice)
		data_map[key] = elemSlice
	}
	return data_map
}

func BuildNestedMap(dir string) map[string]map[string]map[string]map[string]map[string][]string { // just gonna hard code this for now to see the time jk idk how
	// parse CSV file into map[Key]Value
	// map[sourceTypeID]map[regClassID]map[fuelTypeID]map[modelYearID]map[opModeID]
	// can check uninitialized map using ==nil, better if not ordered

	data_map := make(map[string]map[string]map[string]map[string]map[string][]string)
	data := readFile(dir)
	d := new(Data)

	for i, row := range data {

		row = strings.Fields(strings.Join(row, " "))

		// header
		if i == 0 {
			continue
		}

		// intialize map (?) i think, is this right, jk overwrites, jk fixed but p ugly
		// will think about later but just to make work, check each previous value
		if i == 1 {
			d.prevValue = []string{}
			d.prevValue = append(d.prevValue, row[0])
			d.prevValue = append(d.prevValue, row[1])
			d.prevValue = append(d.prevValue, row[2])
			d.prevValue = append(d.prevValue, row[3])
			// fmt.Println(d.prevValue)
			data_map[row[0]] = make(map[string]map[string]map[string]map[string][]string)
			data_map[row[0]][row[1]] = make(map[string]map[string]map[string][]string)
			data_map[row[0]][row[1]][row[2]] = make(map[string]map[string][]string)
			data_map[row[0]][row[1]][row[2]][row[3]] = make(map[string][]string)
			data_map[row[0]][row[1]][row[2]][row[3]][row[4]] = []string{row[5], row[6]} //TODO: type conversions
		}

		if d.prevValue[0] != row[0] {
			data_map[row[0]] = make(map[string]map[string]map[string]map[string][]string)
			d.prevValue[0] = row[0]
		} else if d.prevValue[1] != row[1] {
			data_map[row[0]][row[1]] = make(map[string]map[string]map[string][]string)
			d.prevValue[1] = row[1]
		} else if d.prevValue[2] != row[2] {
			data_map[row[0]][row[1]][row[2]] = make(map[string]map[string][]string)
			d.prevValue[2] = row[2]
		} else if d.prevValue[3] != row[3] {
			data_map[row[0]][row[1]][row[2]][row[3]] = make(map[string][]string)
			d.prevValue[3] = row[3]
		} else {
			data_map[row[0]][row[1]][row[2]][row[3]][row[4]] = []string{row[5], row[6]}
		}
		//fmt.Println(data_map)

		// var sourceTypeIDs = []int{11, 21, 31, 32, 41, 42, 43, 51, 52, 53, 54, 61, 62}
		// var regClassIDs = []int{10, 20, 30, 41, 42, 46, 47, 48, 49}
		// var fuelTypeIDs = []int{1, 2, 3, 5, 9}
		// var modelYearIDs = intSliceRange(1960, 2060, true)
		// var opModeIDs = []int{0, 1, 11, 12, 13, 14, 15, 16, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 33, 35, 36}

		// 11 :
		// 	{10: {
		// 		1: {
		// 			1960: {
		// 				0: [0.1243, 0.1234]
		// 				1:
		// 				11:
		// 			}
		// 			1961:
		// 			1962:
		// 		}
		// 		2:
		// 		3:
		// 	}
		// 	20:
		// 30:}
		// 12:
		// 13:

		//data_map[var0][var1][var2][var3][var4] = []float64{var5, var6}
	}
	return data_map
}

func GetValue(value interface{}) interface{} {

	if reflect.TypeOf(value).Kind() == reflect.Ptr {

		return reflect.ValueOf(value).Elem().Interface()

	} else {

		return value

	}

}

func checkSameType(types map[string]reflect.Type) bool {
	//traverse through the map
	var valueCheck reflect.Type
	cnt := 0
	if len(types) > 1 {
		for _, value := range types {
			if cnt == 0 {
				valueCheck = value
				continue
			} else {
				//check if present value is equals to userValue
				if value != valueCheck {
					return false
				}
			}
			cnt++
		}
	}
	return true
}

// func (d *Data) getFields(value interface{}) {
// d...=fields?

func (d *Data) GetFields(value interface{}) {
	// fields := [][]string{}
	fields := make(map[int]string)
	types := make(map[int]reflect.Kind)

	structElements := reflect.ValueOf(value).Elem()
	structTypes := structElements.Type()

	for j := 0; j < structElements.NumField(); j++ {
		varName := structTypes.Field(j).Name
		varType := structTypes.Field(j).Type.Kind()
		fields[j] = varName
		types[j] = varType

		// name_type := []string{varName, fmt.Sprint(varType)}
		// fields = append(fields, name_type)
	}
	d.Fields = fields
	d.Types = types
}

func readFile(dir string) [][]string {

	// TODO: should loop through if directory than one or only do one at a time?

	f, err := os.Open(dir)

	if err != nil {
		fmt.Println("error: ", err)
		log.Fatal(err)
	}

	// fn := files[0] // change

	// This returns an *os.FileInfo type
	fileInfo, err := f.Stat()

	if err != nil {
		fmt.Println("error: ", err)
		log.Fatal(err)
	}

	// IsDir is short for fileInfo.Mode().IsDir()
	if fileInfo.IsDir() {
		panic("Path is a directory")

		// to iterate over multiple files

		//	files, err := ioutil.ReadDir(dir)

		// for _, file := range files {
		// fn := strings.Join([]string{dir, file.Name()}, "\\")

		// files, err := filepath.Glob(dir + "\\*.csv")
	}
	// file is not a directory

	// remember to close the file at the end of the program
	defer f.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	return data
}

func ImportData[t any](emptyStruct interface{}, dir string) (d Data, conv_array []t) {

	data := readFile(dir)

	data_array := readInput(emptyStruct, data)

	conv_array = AssignStruct[t](data_array)

	d.GetFields(data_array[0])

	return d, conv_array
}

func AssignStruct[t any](data_array []any) (conv_array []t) {

	conv_array = make([]t, len(data_array))

	for i := 0; i < len(data_array); i++ {
		conv_array[i] = ConvStruct[t](data_array[i])
	}
	return conv_array
}

func ConvStruct[t any](datum interface{}) (dataStruct t) {

	dataStruct = (GetValue(datum)).(t)

	return dataStruct
}
