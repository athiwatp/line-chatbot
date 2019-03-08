package shared

// Result common result
type Result struct {
	Data  interface{}
	Total int
	Error error
}
