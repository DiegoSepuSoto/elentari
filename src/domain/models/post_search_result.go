package models

type PostSearchResult struct {
	Posts []*PostForPreview `json:"posts"`
}
