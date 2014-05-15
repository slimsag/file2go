// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// file2go is an simple tool to convert an binary file to an Go source file.
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var (
	input, output, pkg, v string
	force                 bool
)

func init() {
	flag.StringVar(&input, "i", "", "Input file path")
	flag.StringVar(&output, "o", "", "Output file path")
	flag.StringVar(&pkg, "package", "", "Package name for output file")
	flag.StringVar(&v, "var", "", "Variable name for []byte in output file")
	flag.BoolVar(&force, "f", false, "Force writing over existing output files")
}

func main() {
	log.SetFlags(0)
	flag.Parse()

	if len(output) == 0 || len(input) == 0 || len(pkg) == 0 || len(v) == 0 {
		log.Println("Must specify input and output files.\n")
		flag.PrintDefaults()
		os.Exit(1)
	}

	_, err := os.Stat(output)
	if err == nil && !force {
		log.Println(err)
		log.Println("")
		log.Fatal("Output file already exists. (Use -f to overwrite).")
		os.Exit(1)
	}

	inputFile, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}

	outputFile, err := os.Create(output)
	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(outputFile, "package %s\n\n", pkg)
	fmt.Fprintf(outputFile, "var %s []byte = []byte{\n", v)
	fmt.Fprintf(outputFile, "\t")

	col := 4
	for _, b := range data {
		x := fmt.Sprintf("%d, ", b)
		col += len(x)
		if col >= 80 {
			fmt.Fprintf(outputFile, "\n\t")
			col = 4
		}

		fmt.Fprintf(outputFile, x)
	}

	fmt.Fprintf(outputFile, "\n}")
}
