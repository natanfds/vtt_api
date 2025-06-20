package messagehandlers

import (
	"fmt"
	"strings"

	"vtt_api/models"
)

func DetectCommand(message models.Message) (string, error) {
	firstChar := string(message.Content[0])
	if firstChar == "/" {
		command := strings.ToLower(strings.ReplaceAll(strings.Split(message.Content, " ")[0], "/", ""))
		switch true {
		case command == "roll" || command == "r":
			return "roll", nil
		default:
			return "", fmt.Errorf("invalid command")
		}

	}
	return "", nil
}
