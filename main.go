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
	)
	flag.BoolVar(&millis, "millis", false, "Parse epoch time in milliseconds.")
	flag.BoolVar(&micros, "micros", false, "Parse epoch time in microseconds.")
	flag.Parse()

	r := bufio.NewReader(os.Stdin)
ReadLoop:
	for {
		line, err := r.ReadString(0x0A)
		if err == io.EOF {
			break ReadLoop
		}
		if err != nil {
			panic(err)
		}
		line = strings.TrimSpace(line)

		number, err := strconv.ParseInt(line, 10, 64)
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
		fmt.Println(t.String())
	}
}
