package decoder

import (
	"encoding/json"
	"reflect"
)

// type Key struct {
// 	SourceTypeID int
// 	RegClassID   int
// 	FuelTypeID   int
// 	ModelYearID  int
// 	OpModeID     int
// }

// type Value struct {
// 	EmissionRate   float64
// 	EmissionRateIM float64
// }

// func main() {

// 	types := typeAssertion(&Value{})

// 	typeTest := types["EmissionRate"]

// 	// now want to be able to

// 	varTest := reflect.New(typeTest).Interface()
// 	// fmt.Println(varTest.Interface())
// 	_ = Set(&varTest, "1.23423")

// 	val := InterToVal(varTest)
// 	fmt.Println(val)
// }

func Set(v interface{}, s string) error {
	return json.Unmarshal([]byte(s), v)
}

func InterToVal(val any) reflect.Value {
	return reflect.ValueOf(val).Elem()
}

func TypeAssertion(emptyStruct interface{}) map[string]reflect.Type {

	types := make(map[string]reflect.Type)

	structElements := reflect.ValueOf(emptyStruct).Elem()
	structTypes := structElements.Type()

	for j := 0; j < structElements.NumField(); j++ {
		varName := structTypes.Field(j).Name
		varType := structTypes.Field(j).Type

		types[varName] = varType

		// name_type := []string{varName, fmt.Sprint(varType)}
		// fields = append(fields, name_type)
	}
	return types
}
