package api

type CashFlowSearchCriteria struct {
	Page  int
	Limit int
	Owner int
}

type Page struct {
	Total int
	Page  int
	Limit int
	Data  []any
}
