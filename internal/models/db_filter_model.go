package models

type SearchModel struct {
	Column    string `json:"column,omitempty"`
	Condition string `json:"condition,omitempty"`
	Value     string `json:"value,omitempty"`
}

type SortByModel struct {
	Column string `json:"column,omitempty"`
	Order  string `json:"order,omitempty" default:"asc"`
}

type DBFilterModel struct {
	Limit  int32       `json:"limit,omitempty"`
	Offset int32       `json:"offset,omitempty"`
	SortBy SortByModel `json:"sort_by,omitempty"`
	Search SearchModel `json:"search,omitempty"`
}
