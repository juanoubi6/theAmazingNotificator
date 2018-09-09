package common

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/streadway/amqp"
	"theAmazingNotificator/app/config"
)

var db *gorm.DB
var rabbitMqConnection *amqp.Connection

func ConnectToDatabase() {
	var err error
	dbname := config.GetConfig().DB_NAME
	dbhost := config.GetConfig().DB_HOST
	dbport := config.GetConfig().DB_PORT
	dbuser := config.GetConfig().DB_USERNAME
	dbpass := config.GetConfig().DB_PASSWORD

	db, err = gorm.Open("mysql", dbuser+":"+dbpass+"@"+"tcp("+dbhost+":"+dbport+")"+"/"+dbname+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

}

func ConnectToRabbitMQ() {
	connection, err := amqp.Dial("amqp://" + config.GetConfig().RABBITMQ_USER + ":" + config.GetConfig().RABBITMQ_PASSWORD + "@" + config.GetConfig().RABBITMQ_HOST + ":" + config.GetConfig().RABBITMQ_PORT + "/")
	if err != nil {
		panic(err)
	}

	rabbitMqConnection = connection
}

func GetDatabase() *gorm.DB {
	return db
}

func GetRabbitMQChannel() *amqp.Channel {
	ch, err := rabbitMqConnection.Channel()
	if err != nil {
		panic(err)
	}

	return ch
}
