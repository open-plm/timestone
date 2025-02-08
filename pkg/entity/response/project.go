package response

type ProjectResponse struct {
    Code string
    Message string
    Projects []Project
}

type Project struct {
    ID          uint 
    UUID        string 
    CreatedAt   *time.Time
	UpdatedAt   *time.Time
	DeletedAt   *time.Time
}