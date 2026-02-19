package main

import (
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/patrickjane/playerlist-bot/internal/config"
	"github.com/patrickjane/playerlist-bot/internal/discord"
	"github.com/patrickjane/playerlist-bot/internal/model"
	"github.com/patrickjane/playerlist-bot/internal/rcon"
)

var version = ""

func main() {
	slog.Info(fmt.Sprintf("PlayerlistBot %s", version))
	slog.Info("https://github.com/patrickjane/playerlist-bot")

	cfg := config.ParseConfig()

	slog.Info("Monitoring the following servers via RCON:")

	for _, s := range cfg.Rcon.Servers {
		slog.Info(fmt.Sprintf("   %s at %s", s.Name, s.Address))
	}

	slog.Info(fmt.Sprintf("Query servers every %d seconds", cfg.Rcon.QueryEverySeconds))

	errorChan := make(chan error)
	updateChan := make(chan map[string]*model.ServerInfo, 100)

	slog.Info("Connecting to discord")

	discordBot := discord.NewBot(cfg.Discord)

	go func() {
		err := discordBot.Start(updateChan)

		if err != nil {
			slog.Error(fmt.Sprintf("Failed to start discord bot: %s", err))
			os.Exit(1)
		}
	}()

	slog.Info("Creating RCON reader")

	go func() {
		err := rcon.Run(cfg.Rcon, updateChan, errorChan)

		if err != nil {
			slog.Error(fmt.Sprintf("Failed to start RCON connection(s): %s", err))
			os.Exit(1)
		}
	}()

	slog.Info("Successfully started.")

	sigShutdown := make(chan os.Signal, 1)
	signal.Notify(sigShutdown, syscall.SIGTERM, syscall.SIGINT)

	for {
		select {
		case <-sigShutdown:
			slog.Info("Shutting down.")
			discordBot.Stop()
			return
		case err := <-errorChan:
			slog.Error(fmt.Sprintf("RCON error: %s", err))
		}
	}
}
