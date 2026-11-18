package parser

import (
	"encoding/xml"
	"os"
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
	if err := xml.Unmarshal(data, &vc); err != nil {
		return nil, err
	}
	return &vc, nil
}
