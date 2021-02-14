package main

import (
	"github.com/uplo-tech/analyze/jsontag"
	"github.com/uplo-tech/analyze/lockcheck"
	"github.com/uplo-tech/analyze/responsewritercheck"
	"golang.org/x/tools/go/analysis/multichecker"
)

func main() {
	multichecker.Main(
		lockcheck.Analyzer,
		responsewritercheck.Analyzer,
		jsontag.Analyzer,
	)
}
