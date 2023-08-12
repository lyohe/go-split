package main

import (
	"flag"
	"fmt"
	"os"
)

func usage() {
	fmt.Fprintln(os.Stderr, "usage: split [-l line_count] [-a suffix_length] [file [prefix]]\n       split -b byte_count[K|k|M|m|G|g] [-a suffix_length] [file [prefix]]\n       split -n chunk_count [-a suffix_length] [file [prefix]]\n       split -p pattern [-a suffix_length] [file [prefix]]")
	flag.PrintDefaults()
}
