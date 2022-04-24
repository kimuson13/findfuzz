package main

import (
	"findfuzz"

	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(findfuzz.Analyzer) }
