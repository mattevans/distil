package distil

import (
	"fmt"
	"strings"
)

type notEq struct {
	rows []map[string]interface{}
}

// Filter implements distiller().
func (a *notEq) Filter(row map[string]interface{}, filter *Filter) error {
	// Check filter.Value datatypes can be used with our filter.Operator.
	err := validateDataType(filter)
	if err != nil {
		return err
	}

	// Handle equality check for the different operator types this filter can
	// action.
	switch filter.Operator.Type {
	case operatorTypeString:
		// Ensure values are checked as case-insensitive.
		s, substr := strings.ToUpper(row[filter.Field].(string)), strings.ToUpper(filter.Value.(string))

		// Check equality.
		if s != substr {
			a.rows = append(a.rows, row)
		}
		break
	case operatorTypeDateime:
		// Cast to time.Time.
		datum, err := castToDatetime(filter.Value)
		if err != nil {
			return err
		}
		baseline, err := castToDatetime(row[filter.Field])
		if err != nil {
			return err
		}
		// Check equality.
		if !baseline.Equal(*datum) {
			a.rows = append(a.rows, row)
		}
		break
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
		// Check equality.
		if !(*baseline).Equals(*datum) {
			a.rows = append(a.rows, row)
		}
	default:
		return fmt.Errorf("Invalid datatype. Expected string/number/datetime when applying %v operator, got %T for `%v`", filter.Operator.Type, filter.Value, filter.Field)
	}

	return nil
}

// Result implements distiller().
func (a *notEq) Result() []map[string]interface{} {
	return a.rows
}
