package shared

// Filter model
type Filter struct {
	Page   int
	Limit  int
	Offset int
	Sort   string
	SortBy string
}
