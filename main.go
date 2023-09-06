package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/dikyayodihamzah/connex-bot/config"
	"github.com/dikyayodihamzah/connex-bot/controller"
	"github.com/dikyayodihamzah/connex-bot/exception"
	"github.com/dikyayodihamzah/connex-bot/repository"
	"github.com/dikyayodihamzah/connex-bot/service"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func routes(db *mongo.Database) {
	app := fiber.New()
	server := config.NewServerConfig()

	bot := config.NewBot()

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	msgService := service.NewMessageService(bot, userRepo)

	controller := controller.NewMessageController(msgService)
	controller.NewRouter(app)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil && update.Message.IsCommand() {
			switch update.Message.Command() {
			case "start":
				err := userService.AddTelegramChatID(update.Message.Chat.ID, update.Message.From.UserName)
				exception.PanicIfError(err)

				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hello, I'm CONNEX BOT. I will inform you about the latest news from CONNEX.")
				bot.Send(msg)

			default:
				continue
			}

			log.Printf("Message from [%s] %s", update.Message.From.UserName, update.Message.Text)
		}
	}

	err := app.Listen(server.URI)
	log.Fatal(err)
}

func main() {
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	consumer := config.NewKafkaConsumer()
	topics := []string{config.KafkaTopicAuth, config.KafkaTopicProfile, config.KafkaTopicUser}
	if err := consumer.SubscribeTopics(topics, nil); err != nil {
		log.Fatal(err)
	}

	db := config.NewDB()
	userRepo := repository.NewUserRepository(db)
	userConsumer := service.NewUserConsumerService(userRepo)

	// Listen routes
	go routes(db)

	run := true
	for run {
		select {
		case sig := <-sigchan:
			log.Printf("Caught signal %v: terminating\n", sig)
			run = false
		default:
			ev := consumer.Poll(100)
			if ev == nil {
				continue
			}

			switch e := ev.(type) {
			case *kafka.Message:
				log.Printf("Message on %s:\n", e.TopicPartition)
				if e.Headers != nil {
					log.Printf("Headers: %v\n", e.Headers)
				}

				method := fmt.Sprintf("%v", e.Headers)

				switch method {

				// User
				case "[method=\"POST.USER\"]":
					err := userConsumer.Create(e.Value)
					exception.PanicIfError(err)

				case "[method=\"PUT.USER\"]":
					err := userConsumer.Update(e.Value)
					exception.PanicIfError(err)

				case "[method=\"DELETE.USER\"]":
					err := userConsumer.Delete(e.Value)
					exception.PanicIfError(err)
				}

			case kafka.Error:
				fmt.Fprintf(os.Stderr, "Error: %v: %v\n", e.Code(), e)
				if e.Code() == kafka.ErrAllBrokersDown {
					run = false
				}

			default:
				log.Printf("Ignored %v\n", e)
			}
		}
	}

	log.Printf("Closing consumer\n")
}
