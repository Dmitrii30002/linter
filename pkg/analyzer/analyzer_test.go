package analyzer_test

import (
	"testing"

	"github.com/dmitrii30002/linter/pkg/analyzer"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAll(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), analyzer.NewAnalyzer(), "tests")
}
