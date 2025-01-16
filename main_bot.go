package main

import (
	"context"
	"log"
	"os"
	"os/signal"

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
	err = os.Chdir("./data")
	if err != nil {
		log.Println(err.Error())
	}
	os.Create("weight.json")
	os.Create("rate.json")
	os.Create("ratetmp.json")
	os.Create("mealtake.json")
	b.Start(ctx)

}

func handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	if admin(update.Message) {
		dir, _ := os.Getwd()
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: 404531178,
			Text:   dir,
		})
		dirs, _ := os.ReadDir("./")
		for _, d := range dirs {
			b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID: 404531178,
				Text:   d.Name(),
			})
			if d.Name() == "data" {
				err := os.Chdir(d.Name())
				if err != nil {
					log.Println(err.Error())
				}
				dir, _ = os.Getwd()
				b.SendMessage(ctx, &bot.SendMessageParams{
					ChatID: 404531178,
					Text:   dir,
				})
			}
		}
	}
	log.Println("end handler")
}

func admin(get *models.Message) bool {
	return get.From.ID == 404531178
}
