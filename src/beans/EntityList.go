package beans

type EntityListStruct struct {
	Entity string
	Limit  int
	Offset int
	Total  int
	Items  []map[string]interface{}
}
