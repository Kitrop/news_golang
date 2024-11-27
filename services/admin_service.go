package services

import "news-go/repositories"

// BanUser bans a user
func BanUser(userID uint) error {
	return repositories.AdminDeleteNews(userID)
}

// DeleteNews deletes a news item
func DeleteNews(postID uint) error {
	return repositories.AdminDeleteNews(postID)
}
