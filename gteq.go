package distil

import "fmt"

type gteq struct {
	rows []map[string]interface{}
}

func (a *gteq) Filter(row map[string]interface{}, filter *Filter) error {
	// Check filter.Value datatypes can be used with our filter.Operator.
	err := validateDataType(filter)
	if err != nil {
		return err
	}

	if filter.Value == nil || row[filter.Field] == nil {
		return nil
	}

	// The `lteq` operator can only be run on number/datetime datatypes.
	switch filter.Operator.Type {
	case operatorTypeNumber:
		// Cast to decimal.Decimals and compare.
		datum, err := castToNumber(filter.Value)
		if err != nil {
			return err
		}
		baseline, err := castToNumber(row[filter.Field])
		if err != nil {
			return err
		}
		if ((*baseline).Cmp(*datum) > 0) || (*baseline).Equals(*datum) {
			a.rows = append(a.rows, row)
		}
		break
	case operatorTypeDateime:
		// Cast to time.Time and compare.
		datum, err := castToDatetime(filter.Value)
		if err != nil {
			return err
		}
		baseline, err := castToDatetime(row[filter.Field])
		if err != nil {
			return err
		}
		if baseline.After(*datum) || baseline.Equal(*datum) {
			a.rows = append(a.rows, row)
		}
		break
	default:
		return fmt.Errorf("Invalid datatype. Expected number/dateime when applying %v operator, got %T for `%v`", filter.Operator.Type, filter.Value, filter.Field)
	}

	return nil
}

func (a *gteq) Result() []map[string]interface{} {
	return a.rows
}
