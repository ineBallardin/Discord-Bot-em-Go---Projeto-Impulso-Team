package cmd

import (
	"bytes"
	"discord_bot/config"
	"discord_bot/shared"
	"encoding/json"
	"fmt"
	"net/http"
)

type PingCommand struct {
	Config *config.Config
	Client shared.HTTPClient
}

func NewPingCommand(config *config.Config, client shared.HTTPClient) *CommandInfo {
	pingCmd := &PingCommand{
		Config: config,
		Client: client,
	}

	return &CommandInfo{
		Name:        "ping",
		Description: "Responde com Pong!",
		Options:     nil,
		Command:     pingCmd,
	}
}

func (p *PingCommand) Execute(interaction map[string]interface{}) error {
	response := map[string]interface{}{
		"type": 4,
		"data": map[string]interface{}{
			"content": "pong!",
		},
	}
	jsonResponse, _ := json.Marshal(response)

	interactionID := interaction["id"].(string)
	interactionToken := interaction["token"].(string)

	url := fmt.Sprintf("%s/interactions/%s/%s/callback", p.Config.BaseURL, interactionID, interactionToken)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonResponse))

	req.Header.Set("Authorization", "Bot "+p.Config.Token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := p.Client.Do(req)
	if err != nil {
		fmt.Println("error")
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("failed to send message, status code: %d", resp.StatusCode)
	}

	return nil
}
