package distil

import (
	"encoding/json"
	"testing"

	. "github.com/onsi/gomega"
)

func TestMatches(t *testing.T) {
	// Register the test.
	RegisterTestingT(t)

	// Build dummy dataset.
	data := []map[string]interface{}{
		{"location": "Auckland", "department": "Engineering", "team": "Security", "salary": 120000, "prev_employers": 7, "start_date": "2016-01-15T12:00:00Z"},
		{"location": "Auckland", "department": "Engineering", "team": "Security", "salary": 140000, "prev_employers": 6, "start_date": "2016-01-07T12:00:00Z"},
		{"location": "Auckland", "department": nil, "team": nil, "salary": 125000, "prev_employers": 7, "start_date": "2016-01-17T12:00:00Z"},
		{"location": "Auckland", "department": "Engineering", "team": "Security", "salary": 80000, "prev_employers": 3, "start_date": "2016-03-23T12:00:00Z"},
		{"location": "Auckland", "department": "Marketing", "team": "Content", "salary": 90000, "prev_employers": 6, "start_date": "2016-01-15T12:00:00Z"},
		{"location": "Auckland", "department": "Marketing", "team": "Content", "salary": 150000, "prev_employers": 2, "start_date": "2016-01-04T12:00:00Z"},
		{"location": "Wellington", "department": "Engineering", "team": "Security", "salary": 120000, "prev_employers": 6, "start_date": "2016-01-23T12:00:00Z"},
		{"location": "Wellington", "department": "Engineering", "team": "Security", "salary": 160000, "prev_employers": 4, "start_date": "2016-03-23T12:00:00Z"},
	}

	// Init distil dataset with our dummy data.
	dataset := NewDataset(data...)

	// Build a distil query.
	query := &Query{}

	// Append the appropriate filters.
	query.Filters = append(query.Filters, Filter{
		Field: "location",
		Value: "weLL",
		Operator: Operator{
			Code: "matches",
			Type: "string",
		},
	})

	// Run the query.
	results, err := dataset.Run(query)
	if err != nil {
		t.Fatalf("Unexpected error running query: %s", err.Error())
	}

	// Handle empty result set.
	if results == nil {
		t.Fatalf("Unexpectedly got an empty resultset running query")
	}

	// Build expected result set.
	expected := []map[string]interface{}{
		{"location": "Wellington", "department": "Engineering", "team": "Security", "salary": 120000, "prev_employers": 6, "start_date": "2016-01-23T12:00:00Z"},
		{"location": "Wellington", "department": "Engineering", "team": "Security", "salary": 160000, "prev_employers": 4, "start_date": "2016-03-23T12:00:00Z"},
	}

	// Compare actual/expected result sets.
	rm, _ := json.Marshal(results)
	em, _ := json.Marshal(expected)
	Expect(rm).To(MatchJSON(em))
}
