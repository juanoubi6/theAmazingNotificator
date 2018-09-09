package models

import (
	"errors"
)

type NotificationType struct {
	ID   uint `gorm:"primary_key"`
	Type string
}

const (
	PostVote       = 1
	PostComment    = 2
	CommentComment = 3
)

func CreateCommentCommentNotificationMessage(fatherCommentID uint,commentID uint)(message string,recipient uint,err error){

	fatherCommentData, found, err := GetCommentById(fatherCommentID)
	if found == false {
		return "",0,errors.New("The commented comment was not found")
	}
	if err != nil {
		return "",0,err
	}

	commentData, found, err := GetCommentById(commentID)
	if found == false {
		return "",0,errors.New("The comment's comment was not found")
	}
	if err != nil {
		return "",0,err
	}

	postData, found, err := GetPostById(fatherCommentData.PostID)
	if found == false {
		return "",0,errors.New("The comment's post was not found")
	}
	if err != nil {
		return "",0,err
	}

	message = "The user " + commentData.Author.Name + " " + commentData.Author.LastName + " has commented your comment on post '" + postData.Title + "'"

	return message,fatherCommentData.Author.ID,nil

}

func CreatePostCommentNotificationMessage(postID uint, commentID uint)(message string,recipient uint,err error){

	postData, found, err := GetPostById(postID)
	if found == false {
		return "",0,errors.New("The commented post was not found")
	}
	if err != nil {
		return "",0,err
	}

	commentData, found, err := GetCommentById(commentID)
	if found == false {
		return "",0,errors.New("The post comment was not found")
	}
	if err != nil {
		return "",0,err
	}

	message = "The user " + commentData.Author.Name + " " + commentData.Author.LastName + " has commented your post " + "'" + postData.Title + "'"

	return message,postData.Author.ID,nil

}

func CreatePostVoteNotificationMessage(postID uint, votingUserID uint)(message string,recipient uint,err error){

	postData, found, err := GetPostById(postID)
	if found == false {
		return "",0,errors.New("The voted post was not found")
	}
	if err != nil {
		return "",0,err
	}

	userData, found, err := GetUserById(votingUserID)
	if found == false {
		return "",0,errors.New("The voting user was not found")
	}
	if err != nil {
		return "",0,err
	}

	message = "The user " + userData.Name + " " + userData.LastName + " has voted your post " + "'" + postData.Title + "'"

	return message,postData.Author.ID,nil

}