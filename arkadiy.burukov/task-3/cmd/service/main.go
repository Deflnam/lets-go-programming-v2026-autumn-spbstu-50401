package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/Deflnam/task-3/internal/config"
	"github.com/Deflnam/task-3/internal/parser"
)

type OutputValute struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

func main() {
	configPath := flag.String("config", "", "path to config file")
	flag.Parse()

	if *configPath == "" {
		panic("config path is required")
	}

	cfg, err := config.Load(*configPath)
	if err != nil {
		panic(err)
	}

	valCurs, err := parser.Parse(cfg.InputFile)
	if err != nil {
		panic(err)
	}

	output := make([]OutputValute, 0, len(valCurs.Valutes))

	for _, valute := range valCurs.Valutes {
		val := strings.ReplaceAll(valute.Value, ",", ".")
		rate, err := strconv.ParseFloat(val, 64)
		if err != nil {
			panic(fmt.Errorf("parse float: %w", err))
		}

		output = append(output, OutputValute{
			NumCode:  valute.NumCode,
			CharCode: valute.CharCode,
			Value:    rate,
		})
	}

	sort.Slice(output, func(i, j int) bool {
		return output[i].Value > output[j].Value
	})

	jsonData, err := json.MarshalIndent(output, "", "    ")
	if err != nil {
		panic(fmt.Errorf("marshal json: %w", err))
	}

	dir := filepath.Dir(cfg.OutputFile)
	if err := os.MkdirAll(dir, 0750); err != nil {
		panic(fmt.Errorf("create dir: %w", err))
	}
	if err := os.WriteFile(cfg.OutputFile, jsonData, 0600); err != nil {
		panic(fmt.Errorf("write file: %w", err))
	}
}
