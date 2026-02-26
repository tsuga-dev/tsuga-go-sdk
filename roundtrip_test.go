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

// marshalRoundTrip serialises v to JSON then deserialises into a fresh value
// of the same type. Used to verify that Marshal→Unmarshal is lossless.
func marshalRoundTrip[T any](t *testing.T, v T) T {
	t.Helper()
	b, err := json.Marshal(v)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}
	var out T
	if err := json.Unmarshal(b, &out); err != nil {
		t.Fatalf("Unmarshal failed: %v\nJSON: %s", err, b)
	}
	return out
}

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
// GraphVisualization — oneOf with 7 subtypes discriminated by "type"
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
				if v.GraphVisualizationTimeseries == nil {
					t.Fatal("expected GraphVisualizationTimeseries to be set")
				}
				if v.GraphVisualizationTimeseries.GetType() != "timeseries" {
					t.Errorf("type = %q, want %q", v.GraphVisualizationTimeseries.GetType(), "timeseries")
				}
			},
		},
		{
			name: "bar",
			json: `{"type":"bar","source":"metrics","queries":[{"aggregate":{"type":"count"}}]}`,
			checkFn: func(t *testing.T, v tsuga.GraphVisualization) {
				t.Helper()
				if v.GraphVisualizationBar == nil {
					t.Fatal("expected GraphVisualizationBar to be set")
				}
			},
		},
		{
			name: "top-list",
			json: `{"type":"top-list","source":"traces","queries":[{"aggregate":{"type":"count"}}]}`,
			checkFn: func(t *testing.T, v tsuga.GraphVisualization) {
				t.Helper()
				if v.GraphVisualizationTopList == nil {
					t.Fatal("expected GraphVisualizationTopList to be set")
				}
			},
		},
		{
			name: "pie",
			json: `{"type":"pie","source":"logs","queries":[{"aggregate":{"type":"count"}}]}`,
			checkFn: func(t *testing.T, v tsuga.GraphVisualization) {
				t.Helper()
				if v.GraphVisualizationPie == nil {
					t.Fatal("expected GraphVisualizationPie to be set")
				}
			},
		},
		{
			name: "query-value",
			json: `{"type":"query-value","source":"logs","queries":[{"aggregate":{"type":"count"}}]}`,
			checkFn: func(t *testing.T, v tsuga.GraphVisualization) {
				t.Helper()
				if v.GraphVisualizationQueryValue == nil {
					t.Fatal("expected GraphVisualizationQueryValue to be set")
				}
			},
		},
		{
			name: "list",
			json: `{"type":"list","source":"logs","query":"level:error"}`,
			checkFn: func(t *testing.T, v tsuga.GraphVisualization) {
				t.Helper()
				if v.GraphVisualizationList == nil {
					t.Fatal("expected GraphVisualizationList to be set")
				}
			},
		},
		{
			name: "note",
			json: `{"type":"note","note":"hello world"}`,
			checkFn: func(t *testing.T, v tsuga.GraphVisualization) {
				t.Helper()
				if v.GraphVisualizationNote == nil {
					t.Fatal("expected GraphVisualizationNote to be set")
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
			if v.GraphVisualizationTimeseries != nil {
				populated++
			}
			if v.GraphVisualizationBar != nil {
				populated++
			}
			if v.GraphVisualizationTopList != nil {
				populated++
			}
			if v.GraphVisualizationPie != nil {
				populated++
			}
			if v.GraphVisualizationQueryValue != nil {
				populated++
			}
			if v.GraphVisualizationList != nil {
				populated++
			}
			if v.GraphVisualizationNote != nil {
				populated++
			}
			if populated != 1 {
				t.Errorf("expected exactly 1 subtype field set, got %d", populated)
			}
		})
	}
}

// -----------------------------------------------------------------------------
// Aggregate — oneOf with 7 subtypes discriminated by "type"
// -----------------------------------------------------------------------------

func TestAggregateDiscriminator(t *testing.T) {
	cases := []struct {
		name    string
		json    string
		checkFn func(*testing.T, tsuga.Aggregate)
	}{
		{
			name: "count",
			json: `{"type":"count"}`,
			checkFn: func(t *testing.T, a tsuga.Aggregate) {
				t.Helper()
				if a.AggregateCount == nil {
					t.Fatal("expected AggregateCount to be set")
				}
			},
		},
		{
			name: "unique-count",
			json: `{"type":"unique-count","field":"user_id"}`,
			checkFn: func(t *testing.T, a tsuga.Aggregate) {
				t.Helper()
				if a.AggregateUniqueCount == nil {
					t.Fatal("expected AggregateUniqueCount to be set")
				}
			},
		},
		{
			name: "average",
			json: `{"type":"average","field":"duration"}`,
			checkFn: func(t *testing.T, a tsuga.Aggregate) {
				t.Helper()
				if a.AggregateAverage == nil {
					t.Fatal("expected AggregateAverage to be set")
				}
			},
		},
		{
			name: "percentile",
			json: `{"type":"percentile","field":"duration","percentile":99}`,
			checkFn: func(t *testing.T, a tsuga.Aggregate) {
				t.Helper()
				if a.AggregatePercentile == nil {
					t.Fatal("expected AggregatePercentile to be set")
				}
				if a.AggregatePercentile.GetPercentile() != 99 {
					t.Errorf("percentile = %v, want 99", a.AggregatePercentile.GetPercentile())
				}
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			a := unmarshal[tsuga.Aggregate](t, tc.json)
			tc.checkFn(t, a)
		})
	}
}

// -----------------------------------------------------------------------------
// AggregateScalarRequest — round-trip marshal/unmarshal
// -----------------------------------------------------------------------------

func TestAggregateScalarRequestRoundTrip(t *testing.T) {
	tr := tsuga.NewTimeRange(1704067200, 1704153600)
	q := tsuga.NewAggregationQuery(
		tsuga.AggregateCountAsAggregate(tsuga.NewAggregateCount("count")),
	)
	req := tsuga.NewAggregateScalarRequest(*tr, []tsuga.AggregationQuery{*q}, "logs")

	got := marshalRoundTrip(t, req)
	if got.GetDataSource() != "logs" {
		t.Errorf("DataSource = %q, want %q", got.GetDataSource(), "logs")
	}
	if len(got.GetQueries()) != 1 {
		t.Errorf("Queries len = %d, want 1", len(got.GetQueries()))
	}
}
