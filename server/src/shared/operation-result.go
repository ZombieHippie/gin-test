package shared

// DatabaseResult is for clarifying for the below constants
type DatabaseResult int

// Result constants
const (
	ResultSuccess       = iota
	ResultErrorNotFound = iota
	ResultErrorFound    = iota
	ResultErrorDatabase = iota
)
