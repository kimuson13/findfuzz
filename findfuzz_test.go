package findfuzz_test

import (
	"testing"

	"findfuzz"

	"github.com/google/go-cmp/cmp"
	"github.com/gostaticanalysis/testutil"
	"golang.org/x/tools/go/analysis/analysistest"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := testutil.WithModules(t, analysistest.TestData(), nil)
	r := analysistest.Run(t, testdata, findfuzz.Analyzer, "a")
	expected := findfuzz.Result{
		Events: []findfuzz.Event{{Name: "F"}},
	}
	for _, v := range r {
		if diff := cmp.Diff(&expected, v.Result); diff != "" {
			t.Errorf("diff :%v", diff)
		}
	}
}
