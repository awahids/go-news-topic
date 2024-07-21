package dtos

type CreateTopicRequest struct {
	Title string `json:"title" validate:"required,min=3,max=255"`
	Value string `json:"value"`
}

type UpdateTopicRequest struct {
	Title string `json:"title"`
}
