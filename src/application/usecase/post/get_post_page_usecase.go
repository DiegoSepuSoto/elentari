package post

import "elentari/src/domain/models"

func (u *UseCase) GetPostPage(postID string) (*models.PostPage, error)  {
	return u.postRepository.GetDetailedPost(postID)
}
