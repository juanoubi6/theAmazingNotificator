package main

import (
	"theAmazingNotificator/app"
	"theAmazingNotificator/app/common"
	"theAmazingNotificator/app/router"
)

func main() {
	common.ConnectToDatabase()
	common.ConnectToRabbitMQ()
	go app.ConsumeCommentCommentNotificationQueue()
	go app.ConsumePostCommentNotificationQueue()
	go app.ConsumePostVoteNotificationQueue()
	router.CreateRouter()
	router.RunRouter()
}
