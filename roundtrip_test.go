// Package tsuga_test contains round-trip deserialization tests for the
// generated SDK models.
//
// These tests are intentionally kept out of the OpenAPI generator's reach via
// .openapi-generator-ignore so they survive regeneration. They act as a
// post-generation smoke test: if the generator emits broken oneOf
// deserialization code (e.g. missing discriminator mapping), the tests here
// will fail before any API call is ever made.
//
// Run alongside the standard post-generation checks in generate.sh:
//
//	go test ./...
package tsuga_test

import (
	"encoding/json"
	"testing"

	tsuga "github.com/tsuga-dev/tsuga-go-sdk"
)

// unmarshal deserialises raw JSON into T, failing the test on error.
func unmarshal[T any](t *testing.T, raw string) T {
	t.Helper()
	var v T
	if err := json.Unmarshal([]byte(raw), &v); err != nil {
		t.Fatalf("Unmarshal into %T failed: %v\nJSON: %s", v, err, raw)
	}
	return v
}

// -----------------------------------------------------------------------------
// GraphVisualization — oneOf with 18 subtypes discriminated by "type"
// -----------------------------------------------------------------------------
//
// Before the discriminator mapping fix, UnmarshalJSON tried every subtype and
// returned "data matches more than one schema in oneOf(GraphVisualization)"
// whenever more than one passed structural validation (e.g. timeseries and bar
// both have source+queries). With the fix, the "type" field routes directly to
// the correct subtype.

func TestGraphVisualizationDiscriminator(t *testing.T) {
	cases := []struct {
		name    string
		json    string
		checkFn func(*testing.T, tsuga.GraphVisualization)
	}{
		{
			name: "timeseries",
			json: `{"type":"timeseries","source":"logs","queries":[{"aggregate":{"type":"count"}}]}`,
			checkFn: func(t *testing.T, v tsuga.GraphVisualization) {
				t.Helper()
				if v.InputGraphVisualizationTimeseries == nil {
					t.Fatal("expected InputGraphVisualizationTimeseries to be set")
				}
				if v.InputGraphVisualizationTimeseries.GetType() != "timeseries" {
					t.Errorf("type = %q, want %q", v.InputGraphVisualizationTimeseries.GetType(), "timeseries")
				}
			},
		},
		{
			name: "bar",
			json: `{"type":"bar","source":"metrics","queries":[{"aggregate":{"type":"count"}}]}`,
			checkFn: func(t *testing.T, v tsuga.GraphVisualization) {
				t.Helper()
				if v.InputGraphVisualizationBar == nil {
					t.Fatal("expected InputGraphVisualizationBar to be set")
				}
			},
		},
		{
			name: "top-list",
			json: `{"type":"top-list","source":"traces","queries":[{"aggregate":{"type":"count"}}]}`,
			checkFn: func(t *testing.T, v tsuga.GraphVisualization) {
				t.Helper()
				if v.InputGraphVisualizationTopList == nil {
					t.Fatal("expected InputGraphVisualizationTopList to be set")
				}
			},
		},
		{
			name: "pie",
			json: `{"type":"pie","source":"logs","queries":[{"aggregate":{"type":"count"}}]}`,
			checkFn: func(t *testing.T, v tsuga.GraphVisualization) {
				t.Helper()
				if v.InputGraphVisualizationPie == nil {
					t.Fatal("expected InputGraphVisualizationPie to be set")
				}
			},
		},
		{
			name: "query-value",
			json: `{"type":"query-value","source":"logs","queries":[{"aggregate":{"type":"count"}}]}`,
			checkFn: func(t *testing.T, v tsuga.GraphVisualization) {
				t.Helper()
				if v.InputGraphVisualizationQueryValue == nil {
					t.Fatal("expected InputGraphVisualizationQueryValue to be set")
				}
			},
		},
		{
			name: "list",
			json: `{"type":"list","source":"logs","query":"level:error"}`,
			checkFn: func(t *testing.T, v tsuga.GraphVisualization) {
				t.Helper()
				if v.InputGraphVisualizationList == nil {
					t.Fatal("expected InputGraphVisualizationList to be set")
				}
			},
		},
		{
			name: "note",
			json: `{"type":"note","note":"hello world"}`,
			checkFn: func(t *testing.T, v tsuga.GraphVisualization) {
				t.Helper()
				if v.InputGraphVisualizationNote == nil {
					t.Fatal("expected InputGraphVisualizationNote to be set")
				}
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			v := unmarshal[tsuga.GraphVisualization](t, tc.json)
			tc.checkFn(t, v)

			// Verify only one subtype field is populated.
			// Each field must be checked directly — storing a typed nil pointer
			// in an `any` slice produces a non-nil interface value, which would
			// make every entry appear set (classic Go nil-interface trap).
			populated := 0
			if v.InputGraphVisualizationBar != nil {
				populated++
			}
			if v.InputGraphVisualizationBarConnection != nil {
				populated++
			}
			if v.InputGraphVisualizationDistribution != nil {
				populated++
			}
			if v.InputGraphVisualizationGauge != nil {
				populated++
			}
			if v.InputGraphVisualizationHeatmap != nil {
				populated++
			}
			if v.InputGraphVisualizationList != nil {
				populated++
			}
			if v.InputGraphVisualizationListConnection != nil {
				populated++
			}
			if v.InputGraphVisualizationListLogPatterns != nil {
				populated++
			}
			if v.InputGraphVisualizationNote != nil {
				populated++
			}
			if v.InputGraphVisualizationPie != nil {
				populated++
			}
			if v.InputGraphVisualizationPieConnection != nil {
				populated++
			}
			if v.InputGraphVisualizationQueryValue != nil {
				populated++
			}
			if v.InputGraphVisualizationQueryValueConnection != nil {
				populated++
			}
			if v.InputGraphVisualizationTable != nil {
				populated++
			}
			if v.InputGraphVisualizationTimeseries != nil {
				populated++
			}
			if v.InputGraphVisualizationTimeseriesConnection != nil {
				populated++
			}
			if v.InputGraphVisualizationTopList != nil {
				populated++
			}
			if v.InputGraphVisualizationTopListConnection != nil {
				populated++
			}
			if populated != 1 {
				t.Errorf("expected exactly 1 subtype field set, got %d", populated)
			}
		})
	}
}
