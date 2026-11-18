package parser

import (
	"bytes"
	"encoding/xml"
	"os"

	"golang.org/x/net/html/charset"
)

type ValCurs struct {
	XMLName xml.Name `xml:"ValCurs"`
	Valutes []Valute `xml:"Valute"`
}

type Valute struct {
	NumCode  int    `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

func Parse(path string) (*ValCurs, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var vc ValCurs
	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = charset.NewReaderLabel

	if err := decoder.Decode(&vc); err != nil {
		return nil, err
	}
	return &vc, nil
}
