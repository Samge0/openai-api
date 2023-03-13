package controllers

import (
	"github.com/gin-gonic/gin"
	gogpt "github.com/sashabaranov/go-gpt3"
	"net/http"
	"openai-api/app/config"
	"openai-api/app/utils"
)

// chatWithGpt35 chatGpt3.5模型
func (c *ChatController) chatWithGpt35(ctx *gin.Context, cnf *config.Configuration, prompt string) (string, error) {
	req := gogpt.ChatCompletionRequest{
		Model:            cnf.Model,
		MaxTokens:        cnf.MaxTokens,
		TopP:             cnf.TopP,
		FrequencyPenalty: cnf.FrequencyPenalty,
		PresencePenalty:  cnf.PresencePenalty,
		Messages: []gogpt.ChatCompletionMessage{
			{
				Role:    "user",
				Content: prompt,
			},
		},
	}

	client := gogpt.NewClient(utils.GetRandomApiKey())
	resp, err := client.CreateChatCompletion(ctx, req)
	if err != nil {
		c.ResponseJson(ctx, http.StatusInternalServerError, err.Error(), nil)
		return "", err
	}

	if len(resp.Choices) == 0 {
		c.ResponseJson(ctx, http.StatusInternalServerError, "无结果", nil)
		return "", err
	}

	resultText := resp.Choices[0].Message.Content
	return resultText, nil
}

// chatWithGpt30 chatGpt3.0模型
func (c *ChatController) chatWithGpt30(ctx *gin.Context, cnf *config.Configuration, prompt string) (string, error) {
	req := gogpt.CompletionRequest{
		Model:            cnf.Model,
		MaxTokens:        cnf.MaxTokens,
		TopP:             cnf.TopP,
		FrequencyPenalty: cnf.FrequencyPenalty,
		PresencePenalty:  cnf.PresencePenalty,
		Prompt:           prompt,
	}

	client := gogpt.NewClient(utils.GetRandomApiKey())
	resp, err := client.CreateCompletion(ctx, req)
	if err != nil {
		c.ResponseJson(ctx, http.StatusInternalServerError, err.Error(), nil)
		return "", err
	}

	if len(resp.Choices) == 0 {
		c.ResponseJson(ctx, http.StatusInternalServerError, "无结果", nil)
		return "", err
	}

	resultText := resp.Choices[0].Text
	return resultText, nil
}
