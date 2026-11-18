package parser

import (
	"encoding/xml"
	"os"
)

type ValCurs struct {
	XMLName xml.Name `xml:"ValCurs"`
	Date    string   `xml:"Date,attr"`
	Valutes []Valute `xml:"Valute"`
}

type Valute struct {
	CharCode string `xml:"CharCode"`
	Nominal  int    `xml:"Nominal"`
	Value    string `xml:"Value"`
}

func ParseFile(path string) (*ValCurs, error) {
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
