package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/otiai10/openaigo"
	gogpt "github.com/sashabaranov/go-gpt3"
	"openai-api/app/config"
	"openai-api/app/utils"
)

// ChatController 首页控制器
type ChatController struct {
	BaseController
}

// NewChatController 创建控制器
func NewChatController() *ChatController {
	return &ChatController{}
}

// Question 普通封装的查询体 - 只需要传一个提示字段
type Question struct {
	Prompt string `json:"prompt"`
}

// chatWithGpt35 chatGpt3.5模型
func (c *ChatController) chatWithGpt35(ctx *gin.Context, cnf *config.Configuration, prompt string) (string, error) {

	client := openaigo.NewClient(utils.GetRandomApiKey())
	req := openaigo.ChatRequest{
		Model:            cnf.Model,
		MaxTokens:        cnf.MaxTokens,
		TopP:             cnf.TopP,
		FrequencyPenalty: cnf.FrequencyPenalty,
		PresencePenalty:  cnf.PresencePenalty,
		Messages: []openaigo.Message{
			{
				Role:    "user",
				Content: prompt,
			},
		},
	}
	resp, err := client.Chat(ctx, req)
	if err != nil {
		return "", err
	}

	if len(resp.Choices) == 0 {
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
		return "", err
	}

	if len(resp.Choices) == 0 {
		return "", err
	}

	resultText := resp.Choices[0].Text
	return resultText, nil
}
