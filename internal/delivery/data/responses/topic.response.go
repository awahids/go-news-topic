package response

type TopicResponse struct {
	Id    uint   `json:"id"`
	UUID  string `json:"uuid"`
	Title string `json:"title"`
	Value string `json:"value"`
}
