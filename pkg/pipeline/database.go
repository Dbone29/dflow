package pipeline

var InitDatabase = &Pipeline{}

type InitDatabasePayload struct {
	Entities []interface{}
}
