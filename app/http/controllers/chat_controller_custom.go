package controllers

import (
	"github.com/gin-gonic/gin"
	gogpt "github.com/sashabaranov/go-gpt3"
	"net/http"
)

// chatWithGpt35Custom chatGpt3.5模型 - 自定义请求参数
func (c *ChatController) chatWithGpt35Custom(ctx *gin.Context, question *QuestionCustom) (string, error) {
	req := gogpt.ChatCompletionRequest{
		Model:            question.Model,
		MaxTokens:        question.MaxTokens,
		TopP:             question.TopP,
		FrequencyPenalty: question.FrequencyPenalty,
		PresencePenalty:  question.PresencePenalty,
		Messages:         question.Messages,
	}

	client := gogpt.NewClient(question.ApiKey)
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
