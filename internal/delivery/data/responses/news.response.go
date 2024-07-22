package response

type NewsResponse struct {
	Id      uint            `json:"id"`
	UUID    string          `json:"uuid"`
	Title   string          `json:"title"`
	Content string          `json:"content"`
	Status  string          `json:"status"`
	Topics  []TopicResponse `json:"topics"`
}
