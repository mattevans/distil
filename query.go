package distil

import "fmt"

// Query ...
type Query struct {
	Filters []Filter
}

// Filter represents fields needed to distil a dataset.
type Filter struct {
	Field    string      `json:"field"`
	Operator Operator    `json:"operator"`
	Value    interface{} `json:"value"`
}

// Operator represents a filter operator that will be used to distil a dataset.
type Operator struct {
	Code string `json:"code"`
	Type string `json:"type"`
}

// distiller returns a distiller instance based on the filter operator.
func (f *Filter) distiller() (distiller, error) {
	switch f.Operator.Code {
	case "eq":
		return &eq{}, nil
	case "not_eq":
		return &notEq{}, nil
	case "matches":
		return &matches{}, nil
	case "does_not_match":
		return &notMatches{}, nil
	case "is_null":
		return &isNull{}, nil
	case "not_null":
		return &notNull{}, nil
	case "lt":
		return &lt{}, nil
	case "gt":
		return &gt{}, nil
	case "lteq":
		return &lteq{}, nil
	case "gteq":
		return &gteq{}, nil
	case "contains":
		return &contains{}, nil
	case "does_not_contain":
		return &notContains{}, nil
	default:
		return nil, fmt.Errorf("Unknown operator ID: %v", f.Operator)
	}
}
