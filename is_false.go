package distil

type isFalse struct {
	rows []map[string]interface{}
}

// Filter implements distiller().
func (a *isFalse) Filter(row map[string]interface{}, filter *Filter) error {
	if row[filter.Field] == false {
		a.rows = append(a.rows, row)
	}
	return nil
}

// Result implements distiller().
func (a *isFalse) Result() []map[string]interface{} {
	return a.rows
}
