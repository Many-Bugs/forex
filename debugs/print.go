package debugs

import (
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
