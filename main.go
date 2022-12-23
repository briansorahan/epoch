package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	var (
		millis, micros bool
		field          int
	)
	flag.IntVar(&field, "f", 0, "Field number (0-based index).")
	flag.BoolVar(&millis, "millis", false, "Parse epoch time in milliseconds.")
	flag.BoolVar(&micros, "micros", false, "Parse epoch time in microseconds.")
	flag.Parse()

	var (
		lineNo = 1
		r      = bufio.NewReader(os.Stdin)
	)
ReadLoop:
	for {
		line, err := r.ReadString(0x0A)
		if err == io.EOF {
			break ReadLoop
		}
		if err != nil {
			panic(err)
		}
		fields := strings.Fields(line)
		if len(fields) <= field {
			panic(fmt.Sprintf("field %d is greater than number of fields (%d) on line %d", field, len(fields), lineNo))
		}
		epochString := strings.TrimSpace(fields[field])

		number, err := strconv.ParseInt(epochString, 10, 64)
		if err != nil {
			panic(err)
		}
		var t time.Time

		if millis {
			t = time.UnixMilli(number)
		} else if micros {
			t = time.UnixMicro(number)
		} else {
			t = time.Unix(number, 0)
		}
		fields[field] = t.String()

		fmt.Println(strings.Join(fields, " "))
	}
}
