package distil

import (
	"encoding/json"
	"testing"

	. "github.com/onsi/gomega"
)

func TestLtEq(t *testing.T) {
	RegisterTestingT(t)

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

	// Init distil dataset.
	dataset := NewDataset(data...)

	// Build a distil query.
	query := &Query{}
	query.Filters = append(query.Filters, Filter{
		Field: "salary",
		Value: 90000,
		Operator: Operator{
			Code: "lteq",
			Type: "number",
		},
	})
	query.Filters = append(query.Filters, Filter{
		Field: "start_date",
		Value: "2016-03-23T12:00:00Z",
		Operator: Operator{
			Code: "lteq",
			Type: "datetime",
		},
	})

	results, err := dataset.Run(query)
	if err != nil {
		t.Fatalf("Unexpected error running query: %s", err.Error())
	}
	if results == nil {
		t.Fatalf("Unexpectedly got an empty resultset running query")
	}

	expected := []map[string]interface{}{
		{"location": "Auckland", "department": "Engineering", "team": "Security", "salary": 80000, "prev_employers": 3, "start_date": "2016-03-23T12:00:00Z"},
		{"location": "Auckland", "department": "Marketing", "team": "Content", "salary": 90000, "prev_employers": 6, "start_date": "2016-01-15T12:00:00Z"},
	}

	rm, _ := json.Marshal(results)
	em, _ := json.Marshal(expected)
	Expect(rm).To(MatchJSON(em))
}
