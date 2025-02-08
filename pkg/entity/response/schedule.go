package response

type ScheduleResponse struct {
    Code string
    Message string
    Schedules []Schedule
}

type Schedule struct {
    ID          uint 
    UUID        string 
    CreatedAt   *time.Time
	UpdatedAt   *time.Time
	DeletedAt   *time.Time
}