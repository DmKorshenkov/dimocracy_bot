package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/DmKorshenkov/helper/bot/fnc"
	"github.com/DmKorshenkov/helper/bot/in"
	"github.com/DmKorshenkov/helper/bot/o"
	"github.com/DmKorshenkov/helper/bot/t"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(handler),
	}
	b, err := bot.New(t.Token(), opts...)
	if err != nil {
		// panics for the sake of simplicity.
		// you should handle this error properly in your code.
		print("!")
		panic(err)
	}
	print("now b.Start!!!!\n")
	b.Start(ctx)

}

func handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	var ch = make(chan string, 2)
	go fnc.UpDayRate(ch)
	if update.Message.Chat.ID == 404531178 && update.Message.Text == "start" {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: 404531178,
			Text:   fnc.Start(),
		})
	} else {
		answer := in.In(update.Message.Text)
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   answer,
		})
	}
	if val, ok := <-ch; ok {
		fmt.Println(val)
	}
	if update.Message.Chat.ID == 404531178 && update.Message.Text == "tmp" {
		rate := o.MemRate()
		rate.Str()
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: 404531178,
			Text:   rate.Str(),
		})
	}

}
