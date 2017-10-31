package distil

import (
	"fmt"
	"strings"
)

type eq struct {
	rows []map[string]interface{}
}

// Filter implements distiller().
func (a *eq) Filter(row map[string]interface{}, filter *Filter) error {
	// Check filter.Value datatypes can be used with our filter.Operator.
	err := validateDataType(filter)
	if err != nil {
		return err
	}

	// Ensure the values we're checking aren't nil.
	if filter.Value == nil || row[filter.Field] == nil {
		return nil
	}

	// Handle equality check for the different operator types.
	switch filter.Operator.Type {
	case operatorTypeString:
		// Ensure values are checked as case-insensitive.
		s, substr := strings.ToUpper(row[filter.Field].(string)), strings.ToUpper(filter.Value.(string))

		// Check equality.
		if s == substr {
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
		if baseline.Equal(*datum) {
			a.rows = append(a.rows, row)
		}
		break
	case operatorTypeNumber:
		// Cast to decimal.
		datum, err := castToNumber(filter.Value)
		if err != nil {
			return err
		}
		baseline, err := castToNumber(row[filter.Field])
		if err != nil {
			return err
		}
		// Check equality.
		if (*baseline).Equals(*datum) {
			a.rows = append(a.rows, row)
		}
	case operatorTypeBoolean:
		// Cast to boolean.
		datum, err := castToBoolean(filter.Value)
		if err != nil {
			return err
		}
		baseline, err := castToBoolean(row[filter.Field])
		if err != nil {
			return err
		}
		// Check equality.
		if datum == baseline {
			a.rows = append(a.rows, row)
		}
	default:
		return fmt.Errorf("Invalid datatype. Expected string/number/datetime when applying %v operator, got %T for `%v`", filter.Operator.Type, filter.Value, filter.Field)
	}

	return nil
}

// Result implements distiller().
func (a *eq) Result() []map[string]interface{} {
	return a.rows
}
