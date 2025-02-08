package response

type ProgramResponse struct {
    Code string
    Message string
    Programs []Program
}

type Program struct {
    ID          uint 
    UUID        string 
    CreatedAt   *time.Time
	UpdatedAt   *time.Time
	DeletedAt   *time.Time
}