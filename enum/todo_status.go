package enum

//ProductType product type
type TodoStatus string

//const available value for enum
const (
	NotStarted      TodoStatus = "not_started"
	InProgress 		TodoStatus = "in_progress"
	Completed      	TodoStatus = "completed"
)
