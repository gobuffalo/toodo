package oncer

import (
	"fmt"
	"io"
	"os"
	"runtime"
)

const deprecated = "DEPRECATED"

var deprecationWriter io.Writer = os.Stdout

func Deprecate(depth int, name string, msg string) {
	Do(deprecated+name, func() {
		if depth <= 0 {
			depth = 5
		}
		i := depth
		for i > 0 {
			_, _, line, _ := runtime.Caller(i)
			if line > 0 {
				depth = i
				break
			}
			i--
		}
		_, file, line, _ := runtime.Caller(depth)
		fmt.Fprintf(deprecationWriter, "[%s] %s has been deprecated. (%s:%d)\n", deprecated, name, file, line)
		if len(msg) > 0 {
			fmt.Fprintf(deprecationWriter, "\t%s\n", msg)
		}
	})
}
