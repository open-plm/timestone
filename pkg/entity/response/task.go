package response

type TaskResponse struct {
    Code string
    Message string
    Tasks []Task
}

type Task struct {
    ID          uint 
    UUID        string 
    CreatedAt   *time.Time
	UpdatedAt   *time.Time
	DeletedAt   *time.Time
}