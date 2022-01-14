package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	inputPathPtr := flag.String("i", "", "path to pptx file")
	outputPathPtr := flag.String("o", "", "path to output file")
	flag.Parse()
	p := parse(*inputPathPtr)
	fmt.Printf("%v Slides in %s\n", len(p.slides), p.path)
	texts := p.extractTexts()
	f, err := os.Create(*outputPathPtr)
	if err != nil {
		panic(err)
	}
	w := bufio.NewWriter(f)
	for _, t := range texts {
		for _, s := range t {
			_, err := w.WriteString(s + "\n")
			if err != nil {
				panic(err)
			}
		}
	}
	w.Flush()
	fmt.Printf("Texts saved to %s\n", *outputPathPtr)
}
