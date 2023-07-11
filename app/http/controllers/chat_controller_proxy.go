package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/otiai10/openaigo"
	gogpt "github.com/sashabaranov/go-gpt3"
	"openai-api/app/utils/u_http"
)

// QuestionProxy 自定义代理请求的接口地址 - 自定义传openai的官方key  + 自定义代理中转服务器
type QuestionProxy struct {
	QuestionCustom
	ProxyUrl string `json:"proxy_url"`
}

// chatWithGpt35Proxy 获取一个聊天的回答
func (c *ChatController) chatWithGpt35Proxy(ctx *gin.Context, question *QuestionProxy) (string, error) {
	req := openaigo.ChatRequest{
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
		return "", err
	}

	var resp gogpt.ChatCompletionResponse
	err = json.Unmarshal(*body, &resp)

	if err != nil {
		return "", err
	}

	if len(resp.Choices) == 0 {
		return "", err
	}

	resultText := resp.Choices[0].Message.Content
	return resultText, nil
}
