package distil

type queryProcessor struct {
	dataset *Dataset
	query   *Query
	err     error
	results []map[string]interface{}
}

// Run executes the query against the dataset.
func (p *queryProcessor) Run() ([]map[string]interface{}, error) {
	p.distil()
	return p.results, p.err
}

// distil is responsible for filtering the dataset's rows.
func (p *queryProcessor) distil() {
	if p.err != nil {
		return
	}

	// Create distillers for each of the filters.
	filters := make([]distiller, len(p.query.Filters))
	for i, filter := range p.query.Filters {
		filters[i], p.err = filter.distiller()
		if p.err != nil {
			return
		}
	}

	// Apply each filter to the dataset, reducing each interation.
	rangeData := p.dataset.Rows
	for i := range p.query.Filters {
		// Get filter.
		filter := &p.query.Filters[i]

		// Range each row, applying filter.
		for _, row := range rangeData {
			p.err = filters[i].Filter(row, filter)
			if p.err != nil {
				return
			}
		}

		// Update range data to be result of completed filter. This applies filters
		// to reducing dataset.
		rangeData = filters[i].Result()
	}

	// Append filtered data to result set.
	p.results = append(p.results, rangeData...)
}
