package main

import (
	"archive/zip"
	"regexp"
)

func isSlide(file *zip.File) bool {
	matched, err := regexp.MatchString("slides/slide(\\d+).xml", file.Name)
	if err != nil {
		return false
	}
	return matched
}
