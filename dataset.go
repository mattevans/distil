package distil

// Dataset holds our raw data that filters will be applied against.
type Dataset struct {
	Rows []map[string]interface{}
}

// NewDataset creates a new Dataset adding the appropriate rows.
func NewDataset(data ...map[string]interface{}) *Dataset {
	return &Dataset{
		Rows: data,
	}
}

// Run executes the query against the dataset.
func (set *Dataset) Run(query *Query) ([]map[string]interface{}, error) {
	return (&queryProcessor{
		dataset: set,
		query:   query,
	}).Run()
}
