package lockcheck_test

import (
	"testing"

	"github.com/uplo-tech/analyze/lockcheck"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), lockcheck.Analyzer, "a")
}
