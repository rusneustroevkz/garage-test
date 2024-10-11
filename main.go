package main

import (
	"awesomeProject/telegram"
	"context"
	"fmt"
	"os"
	"os/signal"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err := telegram.NewTelegram(ctx)
	if err != nil {
		fmt.Printf("cannot start telegram bot: %v", err)
		os.Exit(1)
	}

	<-ctx.Done()
}
