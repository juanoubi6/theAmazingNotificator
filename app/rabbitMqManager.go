package app

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"strconv"
	"theAmazingNotificator/app/common"
	"theAmazingNotificator/app/communications/rabbitMQ/tasks"
	"theAmazingNotificator/app/config"
	"theAmazingNotificator/app/models"
)

var workerAmount, _ = strconv.Atoi(config.GetConfig().WORKER_AMOUNT)

func ConsumeCommentCommentNotificationQueue() {

	ch := common.GetRabbitMQChannel()
	defer ch.Close()

	err := ch.ExchangeDeclare(
		config.GetConfig().RABBIT_NOTIFICATION_EXCHANGE,
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.WithFields(log.Fields{
			"place": "declaring exchange",
		}).Info(err.Error())
	}

	//Queue declared but not needed if created previously
	queue, err := ch.QueueDeclare(
		config.GetConfig().RABBIT_COMMENT_COMMENT_QUEUE,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.WithFields(log.Fields{
			"place": "declaring queue",
		}).Info(err.Error())
	}

	err = ch.QueueBind(
		queue.Name,
		"comment-comment",
		config.GetConfig().RABBIT_NOTIFICATION_EXCHANGE,
		false,
		nil,
	)
	if err != nil {
		log.WithFields(log.Fields{
			"place": "binding queue",
		}).Info(err.Error())
	}

	//Allow rabbitMQ to send me as many messages as workers I have
	err = ch.Qos(
		workerAmount,
		0,
		false,
	)
	if err != nil {
		log.WithFields(log.Fields{
			"place": "prefetch count",
		}).Info(err.Error())
	}

	msgs, err := ch.Consume(
		queue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.WithFields(log.Fields{
			"place": "consuming message",
		}).Info(err.Error())
	}

	forever := make(chan bool)

	for i := 0; i < workerAmount; i++ {
		go createCommentCommentNotification(msgs)
	}

	log.Printf("Waiting for comment comment notifications tasks")
	<-forever

}

func ConsumePostCommentNotificationQueue() {

	ch := common.GetRabbitMQChannel()
	defer ch.Close()

	err := ch.ExchangeDeclare(
		config.GetConfig().RABBIT_NOTIFICATION_EXCHANGE,
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.WithFields(log.Fields{
			"place": "declaring exchange",
		}).Info(err.Error())
	}

	//Queue declared but not needed if created previously
	queue, err := ch.QueueDeclare(
		config.GetConfig().RABBIT_POST_COMMENT_QUEUE,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.WithFields(log.Fields{
			"place": "declaring exchange",
		}).Info(err.Error())
	}

	err = ch.QueueBind(
		queue.Name,
		"post-comment",
		config.GetConfig().RABBIT_NOTIFICATION_EXCHANGE,
		false,
		nil,
	)
	if err != nil {
		log.WithFields(log.Fields{
			"place": "binding queue",
		}).Info(err.Error())
	}

	//Allow rabbitMQ to send me as many messages as workers I have
	err = ch.Qos(
		workerAmount,
		0,
		false,
	)
	if err != nil {
		log.WithFields(log.Fields{
			"place": "prefetch count",
		}).Info(err.Error())
	}

	msgs, err := ch.Consume(
		queue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.WithFields(log.Fields{
			"place": "consuming message",
		}).Info(err.Error())
	}

	forever := make(chan bool)

	for i := 0; i < workerAmount; i++ {
		go createPostCommentNotification(msgs)
	}

	log.Printf("Waiting for post comments notification tasks")
	<-forever

}

func ConsumePostVoteNotificationQueue() {

	ch := common.GetRabbitMQChannel()
	defer ch.Close()

	err := ch.ExchangeDeclare(
		config.GetConfig().RABBIT_NOTIFICATION_EXCHANGE,
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.WithFields(log.Fields{
			"place": "declaring exchange",
		}).Info(err.Error())
	}

	//Queue declared but not needed if created previously
	queue, err := ch.QueueDeclare(
		config.GetConfig().RABBIT_POST_VOTE_QUEUE,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.WithFields(log.Fields{
			"place": "declaring exchange",
		}).Info(err.Error())
	}

	err = ch.QueueBind(
		queue.Name,
		"post-vote",
		config.GetConfig().RABBIT_NOTIFICATION_EXCHANGE,
		false,
		nil,
	)
	if err != nil {
		log.WithFields(log.Fields{
			"place": "binding queue",
		}).Info(err.Error())
	}

	//Allow rabbitMQ to send me as many messages as workers I have
	err = ch.Qos(
		workerAmount,
		0,
		false,
	)
	if err != nil {
		log.WithFields(log.Fields{
			"place": "prefetch count",
		}).Info(err.Error())
	}

	msgs, err := ch.Consume(
		queue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.WithFields(log.Fields{
			"place": "consuming message",
		}).Info(err.Error())
	}

	forever := make(chan bool)

	for i := 0; i < workerAmount; i++ {
		go createPostVoteNotification(msgs)
	}

	log.Printf("Waiting for post votes notification tasks")
	<-forever

}

func createCommentCommentNotification(messageChannel <-chan amqp.Delivery) {
	for d := range messageChannel {
		var notificationMessageData tasks.CommentCommentNotificationTask
		err := json.Unmarshal(d.Body, &notificationMessageData)
		if err != nil {
			log.WithFields(log.Fields{
				"place": "decoding message body",
			}).Info(err.Error())
		}

		notificationMessage, recipient, err := models.CreateCommentCommentNotificationMessage(notificationMessageData.FatherCommentID, notificationMessageData.CommentID)

		//Create the notification and save it
		newNotification := models.Notification{
			Message:            notificationMessage,
			UserID:             recipient,
			NotificationTypeID: models.CommentComment,
		}

		if err := newNotification.Save(); err != nil {
			log.WithFields(log.Fields{
				"place": "creating notification",
			}).Info(err.Error())
		}

		//Push notification (in the future)

		d.Ack(false)
	}
}

func createPostCommentNotification(messageChannel <-chan amqp.Delivery) {
	for d := range messageChannel {
		var notificationMessageData tasks.PostCommentNotificationTask
		err := json.Unmarshal(d.Body, &notificationMessageData)
		if err != nil {
			log.WithFields(log.Fields{
				"place": "decoding message body",
			}).Info(err.Error())
		}

		notificationMessage, recipient, err := models.CreatePostCommentNotificationMessage(notificationMessageData.PostID, notificationMessageData.CommentID)

		//Create the notification and save it
		newNotification := models.Notification{
			Message:            notificationMessage,
			UserID:             recipient,
			NotificationTypeID: models.PostComment,
		}

		if err := newNotification.Save(); err != nil {
			log.WithFields(log.Fields{
				"place": "creating notification",
			}).Info(err.Error())
		}

		d.Ack(false)
	}
}

func createPostVoteNotification(messageChannel <-chan amqp.Delivery) {
	for d := range messageChannel {
		var notificationMessageData tasks.PostVoteNotificationTask
		err := json.Unmarshal(d.Body, &notificationMessageData)
		if err != nil {
			log.WithFields(log.Fields{
				"place": "decoding message body",
			}).Info(err.Error())
		}

		notificationMessage, recipient, err := models.CreatePostVoteNotificationMessage(notificationMessageData.PostID, notificationMessageData.VotingUserID)

		//Create the notification and save it
		newNotification := models.Notification{
			Message:            notificationMessage,
			UserID:             recipient,
			NotificationTypeID: models.PostVote,
		}

		if err := newNotification.Save(); err != nil {
			log.WithFields(log.Fields{
				"place": "creating notification",
			}).Info(err.Error())
		}

		d.Ack(false)
	}
}
