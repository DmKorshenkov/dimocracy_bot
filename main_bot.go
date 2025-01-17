package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/DmKorshenkov/helper/bot/fnc"
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
	log.Println(os.Chdir("/data"))
	log.Println(fnc.StartStr())
	b.Start(ctx)

}

func handler(ctx context.Context, b *bot.Bot, update *models.Update) {

	log.Println("end handler")
}

func admin(get *models.Message) bool {
	return get.From.ID == 404531178 && get.Text == "start"
}
