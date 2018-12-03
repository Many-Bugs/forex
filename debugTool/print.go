package debugTool

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	line = 0
)

func DebuggingPrint(header string, values ...interface{}) {
	if !strings.HasSuffix(header, "\n") {
		header += "\n"
	}
	fmt.Fprintf(os.Stderr, strconv.Itoa(line+1)+": "+"[FOREX-debug] "+header, values...)
}
