package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"gopkg.in/telebot.v3"
)

func main() {
	token := os.Getenv("BOT_TOKEN")

	pref := telebot.Settings{
		Token:  token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := telebot.NewBot(pref)
	if err != nil {
		fmt.Println("Ошибка при создании бота:", err)
		return
	}

	b.Handle("/luck", func(c telebot.Context) error {
		rand.Seed(time.Now().UnixNano())
		luckScore := rand.Intn(101) 

		var result string
		switch {
		case luckScore == 100:
			result = "Ты выбьешь всё лучшее. Удача: 100%"
		case luckScore > 50:
			result = fmt.Sprintf("Повезло. Удача: %d%%", luckScore)
		case luckScore > 10:
			result = fmt.Sprintf("Неплохо. Удача: %d%%", luckScore)
		default:
			result = fmt.Sprintf("Сегодня не твой день. Удача: %d%%", luckScore)
		}

		return c.Send(result)
	})

	b.Start()
}
