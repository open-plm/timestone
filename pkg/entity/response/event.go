package response

type EventResponse struct {
    Code string
    Message string
    Events []Event
}

type Event struct {
    ID          uint 
    UUID        string 
    CreatedAt   *time.Time
	UpdatedAt   *time.Time
	DeletedAt   *time.Time
}