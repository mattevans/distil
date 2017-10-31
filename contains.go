package distil

import (
	"fmt"
	"reflect"
	"strings"
)

type contains struct {
	rows []map[string]interface{}
}

func (a *contains) Filter(row map[string]interface{}, filter *Filter) error {
	// Check filter.Value datatypes can be used with our filter.Operator.
	err := validateDataType(filter)
	if err != nil {
		return err
	}

	// Ensure our value has datatype of array.
	if reflect.ValueOf(filter.Value).Kind() != reflect.Slice {
		return fmt.Errorf("Invalid datatype. Expected array datatype for `%v`, got %T", filter.Field, row[filter.Field])
	}

	// Skip if the field is nil.
	if row[filter.Field] == nil {
		return nil
	}

	// Build a map of our 'contains' values.
	valueMap := map[interface{}]bool{}
	for _, v := range filter.Value.([]interface{}) {
		switch value := v.(type) {
		case string:
			// Ensure values are mapped as case-insenstive.
			valueMap[strings.ToUpper(value)] = true
			break
		case float32, float64, int, int8, int16, int32, int64:
			valueMap[value.(float64)] = true
			break
		default:
			return fmt.Errorf("Invalid datatype. Expected string datatype for array values of `%v`, got %T", filter.Field, row[filter.Field])
		}
	}

	// Do we have a match?
	switch value := row[filter.Field].(type) {
	case string:
		// Ensure values are checked as case-insenstive.
		if valueMap[strings.ToUpper(value)] {
			a.rows = append(a.rows, row)
		}
		break
	case float32, float64, int, int8, int16, int32, int64:
		if valueMap[value] {
			a.rows = append(a.rows, row)
		}
	default:
		return fmt.Errorf("Invalid datatype. Expected string datatype for array values of `%v`, got %T", filter.Field, row[filter.Field])
	}

	return nil
}

// Result implements distiller().
func (a *contains) Result() []map[string]interface{} {
	return a.rows
}
