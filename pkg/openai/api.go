package openai

import (
	"context"
	"fmt"

	ai "github.com/sashabaranov/go-openai"
)

func Ask(token, Question string) (string, error) {
	client := ai.NewClient(token)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		ai.ChatCompletionRequest{
			Model: ai.GPT3Dot5Turbo,
			Messages: []ai.ChatCompletionMessage{
				{
					Role:    ai.ChatMessageRoleUser,
					Content: Question,
				},
			},
		},
	)
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
	}
	return resp.Choices[0].Message.Content, err
}
