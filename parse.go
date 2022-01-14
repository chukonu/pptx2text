package main

import (
	"archive/zip"
	"log"
)

func parse(path string) (p pptx) {
	read, err := zip.OpenReader(path)
	if err != nil {
		msg := "Failed to open: %s"
		log.Fatalf(msg, err)
	}

	p.path = path
	p.zip = read

	for _, file := range p.zip.File {
		if isSlide(file) {
			p.slides = append(p.slides, file)
		}
	}

	return
}
