package distil

import (
	"fmt"
	"strings"
)

type startsWith struct {
	rows []map[string]interface{}
}

func (a *startsWith) Filter(row map[string]interface{}, filter *Filter) error {
	// Check filter.Value datatypes can be used with our filter.Operator.
	err := validateDataType(filter)
	if err != nil {
		return err
	}

	switch filter.Operator.Type {
	case operatorTypeString:
		// Ensure values are compared case-insensitive.
		s, substr := strings.ToUpper(row[filter.Field].(string)), strings.ToUpper(filter.Value.(string))

		// Match.
		if strings.HasPrefix(s, substr) {
			a.rows = append(a.rows, row)
		}
	case operatorTypeNumber:
		// Cast to number.
		datum, err := castToNumber(filter.Value)
		if err != nil {
			return err
		}
		baseline, err := castToNumber(row[filter.Field])
		if err != nil {
			return err
		}
		if strings.HasPrefix(baseline.String(), datum.String()) {
			a.rows = append(a.rows, row)
		}
	default:
		return fmt.Errorf("Invalid datatype. Expected string/number when applying %v operator, got %T for `%v`", filter.Operator.Type, filter.Value, filter.Field)
	}
	return nil
}

// Result implements distiller().
func (a *startsWith) Result() []map[string]interface{} {
	return a.rows
}
