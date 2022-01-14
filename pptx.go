package main

import (
	"archive/zip"
	"encoding/xml"
	"io"
)

type pptx struct {
	path   string
	slides []*zip.File
	zip    *zip.ReadCloser
}

func (p pptx) extractTexts() (texts [][]string) {
	for _, s := range p.slides {
		var slideTexts []string
		r, err := s.Open()
		if err != nil {
			panic(err)
		}
		d := xml.NewDecoder(r)
		for {
			tok, err := d.Token()
			if err == io.EOF {
				break
			}
			if err != nil {
				panic(err)
			}
			t, ok := tok.(xml.StartElement)
			if !ok {
				continue
			}
			if t.Name.Local == "t" {
				charDataTok, err := d.Token()
				if err != nil {
					panic(err)
				}
				charData, ok := charDataTok.(xml.CharData)
				if !ok {
					continue
				}
				slideTexts = append(slideTexts, string(charData))
			}
		}

		if len(slideTexts) != 0 {
			texts = append(texts, slideTexts)
		}
	}

	return
}
