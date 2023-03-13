package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	gogpt "github.com/sashabaranov/go-gpt3"
	"net/http"
	"openai-api/app/utils/u_http"
)

// chatWithGpt35Proxy 获取一个聊天的回答
func (c *ChatController) chatWithGpt35Proxy(ctx *gin.Context, question *QuestionProxy) (string, error) {
	req := gogpt.ChatCompletionRequest{
		Model:            question.Model,
		MaxTokens:        question.MaxTokens,
		TopP:             question.TopP,
		FrequencyPenalty: question.FrequencyPenalty,
		PresencePenalty:  question.PresencePenalty,
		Messages:         question.Messages,
	}
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	body, err := u_http.Post(question.ProxyUrl, headers, req)
	if err != nil {
		c.ResponseJson(ctx, http.StatusInternalServerError, err.Error(), nil)
		return "", err
	}

	var resp gogpt.ChatCompletionResponse
	err = json.Unmarshal(*body, &resp)

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
