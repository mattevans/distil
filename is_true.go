package distil

type isTrue struct {
	rows []map[string]interface{}
}

// Filter implements distiller().
func (a *isTrue) Filter(row map[string]interface{}, filter *Filter) error {
	if row[filter.Field] == true {
		a.rows = append(a.rows, row)
	}
	return nil
}

// Result implements distiller().
func (a *isTrue) Result() []map[string]interface{} {
	return a.rows
}
