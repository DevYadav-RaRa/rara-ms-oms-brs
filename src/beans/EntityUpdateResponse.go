package beans
type EntityUpdateResponseStruct struct {
	Entity  string
	Success bool
	Message string
	Order   map[string]interface{}
	Batch   map[string]interface{}
}