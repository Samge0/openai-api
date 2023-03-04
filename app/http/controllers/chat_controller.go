package controllers

import (
	"github.com/gin-gonic/gin"
	gogpt "github.com/sashabaranov/go-gpt3"
	"net/http"
	"openai-api/app/utils"
	"openai-api/config"
	"openai-api/pkg/logger"
	"strings"
	"time"
)

// ChatController 首页控制器
type ChatController struct {
	BaseController
}

// NewChatController 创建控制器
func NewChatController() *ChatController {
	return &ChatController{}
}

type Question struct {
	Prompt string `json:"prompt"`
}

// HandlerChat 回复
func (c *ChatController) HandlerChat(ctx *gin.Context) {
	defer utils.TimeCost(time.Now())
	question := &Question{}
	err := ctx.BindJSON(question)
	if err != nil {
		c.ResponseJson(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	if question.Prompt == "" {
		c.ResponseJson(ctx, http.StatusBadRequest, "request text is empty", nil)
		return
	}

	cnf := config.LoadConfig()
	client := gogpt.NewClient(cnf.ApiKey)
	prompt := question.Prompt
	if !strings.HasSuffix(prompt, "。") && !strings.HasSuffix(prompt, "?") && !strings.HasSuffix(prompt, "？") {
		prompt = prompt + "。\n"
	}
	prompt = cnf.BotDesc + "\n" + prompt
	logger.Info("request prompt is %s", prompt)
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

	resp, err := client.CreateChatCompletion(ctx, req)
	if err != nil {
		c.ResponseJson(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	if len(resp.Choices) == 0 {
		c.ResponseJson(ctx, http.StatusInternalServerError, "无结果", nil)
		return
	}

	resultText := resp.Choices[0].Message
	logger.Info("Response resultText is %s", resultText)
	c.ResponseJson(ctx, http.StatusOK, "", resultText)
}
