package main

import (
	"encoding/json"
	"flag"
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

	var output []OutputValute
	for _, v := range valCurs.Valutes {
		val := strings.ReplaceAll(v.Value, ",", ".")
		rate, err := strconv.ParseFloat(val, 64)
		if err != nil {
			panic(err)
		}
		output = append(output, OutputValute{
			NumCode:  v.NumCode,
			CharCode: v.CharCode,
			Value:    rate,
		})
	}

	sort.Slice(output, func(i, j int) bool {
		return output[i].Value > output[j].Value
	})

	jsonData, err := json.MarshalIndent(output, "", "    ")
	if err != nil {
		panic(err)
	}

	dir := filepath.Dir(cfg.OutputFile)
	if err := os.MkdirAll(dir, 0750); err != nil {
		panic(err)
	}

	if err := os.WriteFile(cfg.OutputFile, jsonData, 0644); err != nil {
		panic(err)
	}
}
