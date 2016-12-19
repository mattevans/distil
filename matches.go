package distil

import (
	"fmt"
	"reflect"
	"strings"
)

type matches struct {
	rows []map[string]interface{}
}

func (a *matches) Filter(row map[string]interface{}, filter *Filter) error {
	// Check filter.Value datatypes can be used with our filter.Operator.
	err := validateDataType(filter)
	if err != nil {
		return err
	}

	// Ensure our value has datatype of string.
	if reflect.TypeOf(row[filter.Field]).String() != operatorTypeString {
		return fmt.Errorf("Invalid datatype. Expected string datatype for `%v`, got %T", filter.Field, row[filter.Field])
	}

	// Ensure values are matched case-insenstive.
	s, substr := strings.ToUpper(row[filter.Field].(string)), strings.ToUpper(filter.Value.(string))

	// Match.
	if strings.Contains(s, substr) {
		a.rows = append(a.rows, row)
	}
	return nil
}

// Result implements distiller().
func (a *matches) Result() []map[string]interface{} {
	return a.rows
}
