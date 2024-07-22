package dtos

type CreateNewsRequest struct {
	Title   string      `json:"title" validate:"required"`
	Content string      `json:"content" validate:"required"`
	Status  string      `json:"status" validate:"required,oneof=published draft"`
	Topics  []TopicUuid `json:"topics"`
}

type UpdateNewsRequest struct {
	Title   string      `json:"title"`
	Content string      `json:"content"`
	Status  string      `json:"status"`
	Topics  []TopicUuid `json:"topics"`
}

type UpdateNewsStatus struct {
	Status string `json:"status" validate:"required,oneof=published draft"`
}

type TopicUuid struct {
	Uuid string `json:"uuid" validate:"required"`
}

type FilterNewsRequest struct {
	Title  *string `json:"title"`
	Topic  *string `json:"topic"`
	Status *string `json:"status"`
}
