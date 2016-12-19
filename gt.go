package distil

import "fmt"

type gt struct {
	rows []map[string]interface{}
}

// Filter implements distiller().
func (a *gt) Filter(row map[string]interface{}, filter *Filter) error {
	// Check filter.Value datatypes can be used with our filter.Operator.
	err := validateDataType(filter)
	if err != nil {
		return err
	}

	// Ensure the values we're checking aren't nil.
	if filter.Value == nil || row[filter.Field] == nil {
		return nil
	}

	// Handle equality check for the different operator types this filter can
	// action.
	switch filter.Operator.Type {
	case operatorTypeNumber:
		// Cast to number,
		datum, err := castToNumber(filter.Value)
		if err != nil {
			return err
		}
		baseline, err := castToNumber(row[filter.Field])
		if err != nil {
			return err
		}
		// Compare.
		if (*baseline).Cmp(*datum) > 0 {
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
		// Compare.
		if baseline.After(*datum) {
			a.rows = append(a.rows, row)
		}
		break
	default:
		return fmt.Errorf("Invalid datatype. Expected string/dateime when applying %v operator, got %T for `%v`", filter.Operator.Type, filter.Value, filter.Field)
	}

	return nil
}

// Result implements distiller().
func (a *gt) Result() []map[string]interface{} {
	return a.rows
}
