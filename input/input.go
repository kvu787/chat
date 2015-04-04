package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var usage string = `usage: input [-h] [-f filepath] [-p prefix]

Input reads from standard input and appends text to a specified file.

The -f flag specifies the file to append to. 
If unspecified, filepath is 'cht'.

The -p flag specifies a prefix to be added to each line of text appended to 
the file. If unspecified, prefix is the empty string.`

func main() {
	// read args
	var filepath string
	var prefix string
	var isHelp bool
	flag.StringVar(&filepath, "f", "cht", "")
	flag.StringVar(&prefix, "p", "", "")
	flag.BoolVar(&isHelp, "h", false, "")
	flag.Parse()

	if isHelp {
		fmt.Println(usage)
		os.Exit(0)
	}

	// open chat stream as writer
	f, err := os.OpenFile(filepath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", "err when opening chat stream", err)
		os.Exit(1)
	}
	bf := bufio.NewWriter(f)

	// open stdin as scanner
	scanner := bufio.NewScanner(os.Stdin)

	for {
		// read user input
		if !scanner.Scan() {
			err := scanner.Err()
			if err == nil {
				fmt.Println("EOF encountered, exiting")
				os.Exit(0)
			} else {
				fmt.Fprintf(os.Stderr, "%s: %v\n", "err when reading input", err)
				os.Exit(1)
			}
		}

		// write user input
		_, err := bf.WriteString(fmt.Sprintf("%s: %s\n", prefix, scanner.Text()))
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v\n", "err when writing to chat stream", err)
			os.Exit(1)
		}
		err = bf.Flush()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v\n", "err when flushing write", err)
			os.Exit(1)
		}
	}
}
