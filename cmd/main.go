package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Li-ReDBox/dateinterval"
)

// usage prints help a message
func usage() {
	fmt.Println("The application accepts exactly two arguments in the format of d[d]/m[m]/yyyy seperated by a space")
}

func main() {
	flag.Usage = usage
	flag.Parse()

	inputs := flag.Args()
	if len(inputs) != 2 {
		usage()
		os.Exit(1)
	}

	var parsed []dateinterval.Date

	for i, v := range inputs {
		d, err := dateinterval.CreateDate(v)

		if err != nil {
			fmt.Printf("Input %d '%s' is not valid.\n%s\n", i+1, v, err)
		} else {
			parsed = append(parsed, d)

		}
	}

	if len(parsed) != 2 {
		os.Exit(1)
	}

	fmt.Println("Days between", parsed[0], "and", parsed[1], "is", parsed[0].Interval(parsed[1]))
}
