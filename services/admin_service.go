package services

import "news-go/repositories"

func BanUser(userID uint) error {
	return repositories.AdminDeleteNews(userID)
}

func DeleteNews(postID uint) error {
	return repositories.AdminDeleteNews(postID)
}
