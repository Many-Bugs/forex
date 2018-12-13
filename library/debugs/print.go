package debugs

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strings"
)

func DebuggingPrint(header string, values ...interface{}) {
	if !strings.HasSuffix(header, "\n") {
		header += "\n"
	}
	fmt.Fprintf(os.Stderr, "[Debug] "+header, values...)
}

func PrintStructureWithField(obj interface{}) {
	fmt.Printf("%#v\n", obj)
}

func PrettyPrintValue(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		fmt.Println(string(b))
	}
	return
}

// obj = object pointer
func InspectStruct(obj interface{}) {

	val := reflect.ValueOf(obj)

	if val.Kind() == reflect.Interface && !val.IsNil() {
		elm := val.Elem()
		if !elm.IsNil() && elm.Kind() == reflect.Ptr && elm.Elem().Kind() == reflect.Ptr {
			val = elm
		}
	}
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)
		address := "not-addressable"

		if valueField.Kind() == reflect.Interface && !valueField.IsNil() {
			if elm := valueField.Elem(); elm.Kind() == reflect.Ptr {
				if !elm.IsNil() && elm.Elem().Kind() == reflect.Ptr {
					valueField = elm
				}
			}
		}
		if valueField.Kind() == reflect.Ptr {
			valueField = valueField.Elem()
		}
		if valueField.CanAddr() {
			address = fmt.Sprint(valueField.Addr().Pointer())
		}

		fmt.Printf("Field Name: %s,\n"+
			" Field Value: %v,\n"+
			" Address: %v,\n"+
			" Field type: %v,\n"+
			" Field kind: %v\n\n",
			typeField.Name,
			valueField.Interface(),
			address,
			typeField.Type,
			valueField.Kind())

		if valueField.Kind() == reflect.Struct {
			InspectStruct(valueField.Interface())
		}
	}
}
