package models

type PostCategory struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type PostPage struct {
	ID         string          `json:"id"`
	Title      string          `json:"title"`
	Image      string          `json:"image"`
	Body       string          `json:"body"`
	ServiceID  string          `json:"service_id"`
	Categories []*PostCategory `json:"categories"`
}
