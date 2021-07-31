package models

type HomePage struct {
	ServicesWithPosts []*ServiceWithPosts `json:"services_with_posts"`
}

type ServiceWithPosts struct {
	ID              string            `json:"id"`
	Abbreviation    string            `json:"abbreviation"`
	PostsForPreview []*PostForPreview `json:"posts_for_preview"`
}

type PostForPreview struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Summary  string `json:"summary"`
	ImageURL string `json:"image_url"`
}
