package main

import (
	"time"

	"github.com/DavidDamron/SolanaTokenLsitBot/bot"
)

func main() {
	for true {
		bot.Run()
		time.Sleep(100 * time.Second)
	}
}
