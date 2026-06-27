package handler

type createArticleRequest struct {
	Author string `json:"author"`
	Topic  string `json:"topic"`
	Text   string `json:"text"`
}

type updateArticleArticleRequest struct {
	Author string `json:"author"`
	Topic  string `json:"topic"`
	Text   string `json:"text"`
}
