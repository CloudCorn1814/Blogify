package handler

type createArticleRequest struct {
	Author string `json:"author" validate:"required"`
	Topic  string `json:"topic" validate:"required"`
	Text   string `json:"text" validate:"required"`
}

type updateArticleArticleRequest struct {
	Author string `json:"author" validate:"required"`
	Topic  string `json:"topic" validate:"required"`
	Text   string `json:"text" validate:"required"`
}
