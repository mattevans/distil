package distil

import (
	"encoding/json"
	"testing"

	. "github.com/onsi/gomega"
)

func TestIsTrue(t *testing.T) {
	// Register the test.
	RegisterTestingT(t)

	// Build dummy dataset.
	data := []map[string]interface{}{
		{"location": "Auckland", "active": false},
		{"location": "Wellington", "active": false},
		{"location": "Nelson", "active": true},
		{"location": "Chirstchurch", "active": false},
		{"location": "Queenstown", "active": true},
		{"location": "Invercargill", "active": false},
	}

	// Init distil dataset with our dummy data.
	dataset := NewDataset(data...)

	// Build a distil query.
	query := &Query{}

	// Append the appropriate filters.
	query.Filters = append(query.Filters, Filter{
		Field: "active",
		Value: true,
		Operator: Operator{
			Code: "is_true",
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
		{"location": "Nelson", "active": true},
		{"location": "Queenstown", "active": true},
	}

	// Compare actual/expected result sets.
	rm, _ := json.Marshal(results)
	em, _ := json.Marshal(expected)
	Expect(rm).To(MatchJSON(em))
}
