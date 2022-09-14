package main

import (
	"fmt"
	"gopkg.in/telegram-bot-api.v4"
	"log"
	"net/http"

	"github.com/LittleMikle/Simple_Golang_bot/Internal"
	"github.com/LittleMikle/Simple_Golang_bot/entity"
	cfg "github.com/LittleMikle/Simple_Golang_bot/pkg/config"
)

func main() {
	config := cfg.New()
	if err := config.Load("./configs", "config", "yml"); err != nil {
		log.Fatal(err)
	}
	bot, err := tgbotapi.NewBotAPI(config.TelegramBotToken)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Authorized on account %s\n", bot.Self.UserName)

	_, err = bot.SetWebhook(tgbotapi.NewWebhook(config.WebhookURL))
	if err != nil {
		panic(err)
	}

	updates := bot.ListenForWebhook("/")

	go func() {
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			fmt.Println("Я обосрался")
		}
	}()
	fmt.Println("start listen :8080")

	for update := range updates {
		if url, ok := entity.Rss[update.Message.Text]; ok {
			rss, err := Internal.GetNews(url)
			if err != nil {
				bot.Send(tgbotapi.NewMessage(
					update.Message.Chat.ID,
					"sorry, error happend",
				))
			}
			for _, item := range rss.Items {
				bot.Send(tgbotapi.NewMessage(
					update.Message.Chat.ID,
					item.URL+"\n"+item.Title,
				))
			}
		} else {
			bot.Send(tgbotapi.NewMessage(
				update.Message.Chat.ID,
				`Поддерживается только Habr`,
			))
		}
	}
}
