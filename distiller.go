package distil

import (
	"errors"
	"fmt"
	"reflect"
	"time"

	"github.com/shopspring/decimal"
)

// The different operator types available.
const (
	operatorTypeString  = "string"
	operatorTypeNumber  = "number"
	operatorTypeDateime = "datetime"
	operatorTypeArray   = "array"
)

// Define methods distiller should implement.
type distiller interface {
	Filter(map[string]interface{}, *Filter) error
	Result() []map[string]interface{}
}

// validateDataType ensures filter value passed has correct datatype for the
// filters operator.
func validateDataType(filter *Filter) error {
	// Retrieve value.
	datum := filter.Value

	// Handle all our operator types.
	switch filter.Operator.Type {
	case operatorTypeString:
		_, err := castToString(datum)
		if err != nil {
			return fmt.Errorf("Invalid datatype. Expected string datatype, got %T for `%v`", datum, filter.Field)
		}
	case operatorTypeDateime:
		_, err := castToDatetime(datum)
		if err != nil {
			return fmt.Errorf("Invalid datatype. Expected datetime datatype, got %T for `%v`", datum, filter.Field)
		}
	case operatorTypeNumber:
		_, err := castToNumber(datum)
		if err != nil {
			return fmt.Errorf("Invalid datatype. Expected number datatype, got %T for `%v`", datum, filter.Field)
		}
	case operatorTypeArray:
		_, err := castToSlice(datum)
		if err != nil {

			return fmt.Errorf("Invalid datatype. Expected array datatype, got %T for `%v` %v", datum, filter.Field, err)
		}
	default:
		return fmt.Errorf("Unknown field type: %s", filter.Operator.Type)
	}
	return nil
}

// castToDatetime attempts to cast given interface{} to a string.
func castToString(datum interface{}) (*string, error) {
	str, ok := datum.(string)
	if !ok {
		return nil, fmt.Errorf("Invalid datatype. Expected string datatype, got %T", datum)
	}
	return &str, nil
}

// castToNumber attempts to cast given interface{} to decimal.Decimal.
func castToNumber(datum interface{}) (*decimal.Decimal, error) {
	var err error
	var d decimal.Decimal

	// Handle different float/integer datatypes, casting to float64.
	switch datumTyped := datum.(type) {
	case int:
		d = decimal.NewFromFloat(float64(datumTyped))
	case int32:
		d = decimal.NewFromFloat(float64(datumTyped))
	case int64:
		d = decimal.NewFromFloat(float64(datumTyped))
	case float32:
		d = decimal.NewFromFloat(float64(datumTyped))
	case float64:
		d = decimal.NewFromFloat(datumTyped)
	default:
		return nil, fmt.Errorf("Invalid datatype. Expected number datatype, got %T", datum)
	}
	return &d, err
}

// castToDatetime attempts to cast given interface{} to a datetime.
func castToDatetime(datum interface{}) (*time.Time, error) {
	var t *time.Time

	// Handle the different time.Time possibities, casting to *time.Time.
	switch datumTyped := datum.(type) {
	case time.Time:
		t = &datumTyped
	case *time.Time:
		if datumTyped == nil {
			return nil, errors.New("Invalid datatype. Received nil *time.Time for datetime field")
		}
		t = datumTyped
	case string:
		parsed, err := time.Parse(time.RFC3339, datumTyped)
		if err != nil {
			return nil, errors.New("Invalid datatype. Invalid date string passed for datetime field. RFC3339 datetime string required.")
		}
		t = &parsed
	default:
		return nil, fmt.Errorf("Invalid datatype. Expected string or time.Time, got %T", datum)
	}
	return t, nil
}

// castToSlice attempts to cast given interface{} to []string{}.
func castToSlice(datum interface{}) ([]*string, error) {
	var err error
	var results []*string
	// Handle different float/integer datatypes, casting to float64.
	switch reflect.TypeOf(datum).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(datum)
		for i := 0; i < s.Len(); i++ {
			str, e := castToString(s.Index(i).String())
			if e != nil {
				return nil, e
			}
			results = append(results, str)
		}
	default:
		return nil, fmt.Errorf("Invalid datatype. Expected array datatype, got %T", datum)
	}
	return results, err
}
