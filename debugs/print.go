package debugs

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func DebuggingPrint(header string, values ...interface{}) {
	if !strings.HasSuffix(header, "\n") {
		header += "\n"
	}
	fmt.Fprintf(os.Stderr, "[FOREX-debug] "+header, values...)
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
