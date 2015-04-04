package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"time"
)

var usage string = `usage: display [-h] [-f filepath]

Display reads the specified file and any changes to it and outputs to
standard output.

The -f flag specifies the file to read. 
If unspecified, filepath is 'cht'.`

func main() {
	// read args
	var filepath string
	var isUsage bool
	flag.StringVar(&filepath, "f", "cht", "")
	flag.BoolVar(&isUsage, "h", false, "")
	flag.Parse()

	if isUsage {
		fmt.Println(usage)
		os.Exit(0)
	}

	// open chat stream as scanner
	f, err := os.OpenFile(filepath, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", "err when opening chat stream", err)
		os.Exit(1)
	}
	reader := bufio.NewReader(f)

	// open stdout as writer
	stdout := bufio.NewWriter(os.Stdout)

	for {
		// read from chat stream
		bs, isPrefix, err := reader.ReadLine()

		// write to stdout if text available
		if err == nil {
			text := string(bs)
			if !isPrefix {
				text += "\n"
			}
			_, err := stdout.WriteString(text)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s: %v\n", "err when writing to stdout", err)
				os.Exit(1)
			}
			err = stdout.Flush()
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s: %v\n", "err when flushing write", err)
				os.Exit(1)
			}
		}

		// limit readline calls
		time.Sleep(time.Duration(100 * int64(time.Millisecond)))
	}
}
