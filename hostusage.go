package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	//	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func processFile(in io.Reader, out io.Writer) {
	r := csv.NewReader(in)
	r.LazyQuotes = true
	r.FieldsPerRecord = -1

	w := csv.NewWriter(out)
	defer w.Flush()

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		if err := w.Write(record); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
		//		fmt.Println(len(record))
	}
}

func main() {
	flag.Parse()

	for _, a := range flag.Args() {
		i, err := os.Open(a)
		if err != nil {
			log.Fatal(err)
		}
		defer i.Close()

		o, err := os.Create(strings.TrimSuffix(a, filepath.Ext(a)) + " fixed" + filepath.Ext(a))
		if err != nil {
			log.Fatal(err)
		}
		defer o.Close()

		r := bufio.NewReader(i)
		f, err := r.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		log.Print(f)
		processFile(r, os.Stdout)
	}
}
