package models

type CategoryPostsPage struct {
	ID              string            `json:"id"`
	Name            string            `json:"name"`
	PostsForPreview []*PostForPreview `json:"posts_for_preview"`
}
