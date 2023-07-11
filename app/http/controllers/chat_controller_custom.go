package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/otiai10/openaigo"
	gogpt "github.com/sashabaranov/go-gpt3"
)

// QuestionCustom 自定义传参的请求体 - 自定义传openai的官方key
type QuestionCustom struct {
	openaigo.ChatCompletionRequestBody
	ApiKey string `json:"api_key"`
}

// chatWithGpt35Custom chatGpt3.5模型 - 自定义请求参数
func (c *ChatController) chatWithGpt35Custom(ctx *gin.Context, question *QuestionCustom) (string, error) {

	client := openaigo.NewClient(question.ApiKey)
	req := openaigo.ChatRequest{
		Model:            question.Model,
		MaxTokens:        question.MaxTokens,
		TopP:             question.TopP,
		FrequencyPenalty: question.FrequencyPenalty,
		PresencePenalty:  question.PresencePenalty,
		Messages:         question.Messages,
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
func (c *ChatController) chatWithGpt30Custom(ctx *gin.Context, question *QuestionCustom) (string, error) {
	req := gogpt.CompletionRequest{
		Model:            question.Model,
		MaxTokens:        question.MaxTokens,
		TopP:             question.TopP,
		FrequencyPenalty: question.FrequencyPenalty,
		PresencePenalty:  question.PresencePenalty,
		Prompt:           question.Messages[0].Content,
	}

	client := gogpt.NewClient(question.ApiKey)
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
