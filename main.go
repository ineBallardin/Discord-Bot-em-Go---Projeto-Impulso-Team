package main

import (
	"discord_bot/cmd"
	"discord_bot/config"
	"discord_bot/discord"
	"net/http"
)

type Message struct {
	Content string `json:"content"`
}

func main() {
	configInstance := config.LoadConfig()

	registry := cmd.NewCommandRegistry()
	registry.RegistryCommand(cmd.NewPingCommand(configInstance, &http.Client{}))

	discordClient := discord.GetDiscordClient(configInstance, registry)
	discordClient.RegisterSlashCommands()

	if err := discordClient.ConnectToGateway(); err != nil {
		discordClient.HandleError(err)
	}

	select {}
}
