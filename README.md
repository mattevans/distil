# distil 💧

[![GoDoc](https://godoc.org/github.com/mattevans/distil?status.svg)](https://godoc.org/github.com/mattevans/distil)
[![Build Status](https://travis-ci.org/mattevans/distil.svg?branch=master)](https://travis-ci.org/mattevans/distil)
[![Go Report Card](https://goreportcard.com/badge/github.com/mattevans/distil)](https://goreportcard.com/report/github.com/mattevans/distil)

In memory dataset filtering.

Installation
-----------------

`go get -u github.com/mattevans/distil`

Example
-------------

Given a dataset...

```go
data := []map[string]interface{}{
	{"location": "New York", "department": "Engineering", "salary": 120000, "start_date": "2016-01-23T12:00:00Z"},
	{"location": "New York", "department": "Engineering", "salary": 80000, "start_date": "2016-03-23T12:00:00Z"},
	{"location": "New York", "department": "Marketing", "salary": 90000, "start_date": "2016-01-23T12:00:00Z"},
	{"location": "New York", "department": "Marketing", "salary": 150000, "start_date": "2016-01-23T12:00:00Z"},
	{"location": "Chicago", "department": "Engineering", "salary": 120000, "start_date": "2016-01-23T12:00:00Z"},
	{"location": "Chicago", "department": "Engineering", "salary": 160000, "start_date": "2016-03-23T12:00:00Z"},
}
```

Initialize distil, build filters and run...

```go
// Init distil dataset.
dataset := NewDataset(data...)

// Build a distil query and apply filters.
query := &Query{}
query.Filters = append(query.Filters, Filter{
	Field: "location",
	Value: "Chicago",
	Operator: Operator{
		Code: "eq",
		Type: "string",
	},
})

// Run it.
results, err := dataset.Run(query)
if err != nil {
  errors.New("Unexpected error running query: %s", err.Error())
}
```

Returns...

```go
[]map[string]interface{}{
	{"location": "Chicago", "department": "Engineering", "salary": 120000, "start_date": "2016-01-23T12:00:00Z"},
	{"location": "Chicago", "department": "Engineering", "salary": 160000, "start_date": "2016-03-23T12:00:00Z"},
}

```

Find a list of available [operators here](https://github.com/mattevans/distil/blob/master/example/operators.json).

Thanks &amp; Acknowledgements :ok_hand:
----------------

The packages's architecture is adapted from
[aggro](https://github.com/snikch/aggro), created by [Mal
Curtis](https://github.com/snikch). :beers:

Contributing
-----------------
If you've found a bug or would like to contribute, please create an issue here on GitHub, or better yet fork the project and submit a pull request!
