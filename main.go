package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"://github.com"
)

func main() {
	token := os.Getenv("BOT_TOKEN")

	b, err := telebot.NewBot(telebot.Settings{
		Token:  token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	b.Handle("/luck", func(m *telebot.Message) {
		rand.Seed(time.Now().UnixNano())
		luckScore := rand.Intn(101)
		
		res := fmt.Sprintf("Твоя удача сегодня: %d%%", luckScore)
		b.Send(m.Chat, res)
	})

	b.Start()
}
