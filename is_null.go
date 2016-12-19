package distil

type isNull struct {
	rows []map[string]interface{}
}

// Filter implements distiller().
func (a *isNull) Filter(row map[string]interface{}, filter *Filter) error {
	if row[filter.Field] == nil {
		a.rows = append(a.rows, row)
	}
	return nil
}

// Result implements distiller().
func (a *isNull) Result() []map[string]interface{} {
	return a.rows
}
