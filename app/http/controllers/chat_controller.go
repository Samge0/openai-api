package controllers

import (
	"github.com/gin-gonic/gin"
	gogpt "github.com/sashabaranov/go-gpt3"
	"net/http"
	"openai-api/app/config"
	"openai-api/app/utils"
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

type QuestionCustom struct {
	gogpt.ChatCompletionRequest
	ApiKey string `json:"api_key"`
}

// HandlerChatCustom 回复 - 自定义传参数
func (c *ChatController) HandlerChatCustom(ctx *gin.Context) {
	defer utils.TimeCost(time.Now())
	question := &QuestionCustom{}
	err := ctx.BindJSON(question)
	if err != nil {
		c.ResponseJson(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	if question.ApiKey == "" {
		c.ResponseJson(ctx, http.StatusBadRequest, "request api_key is empty", nil)
		return
	}
	if len(question.Messages) == 0 || question.Messages[0].Content == "" {
		c.ResponseJson(ctx, http.StatusBadRequest, "request Messages.Content is empty", nil)
		return
	}

	logger.Info("request prompt is %s", question.Messages[0].Content)

	var resultText string
	if question.Model == gogpt.GPT3Dot5Turbo || question.Model == gogpt.GPT3Dot5Turbo0301 {
		resultText, err = c.chatWithGpt35Custom(ctx, question)
	} else {
		resultText, err = c.chatWithGpt30Custom(ctx, question)
	}
	if err != nil {
		logger.Danger("request err is %s", err)
		return
	}
	logger.Info("Response resultText is %s", resultText)
	c.ResponseJson(ctx, http.StatusOK, "", resultText)
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
	prompt := question.Prompt
	if !strings.HasSuffix(prompt, "。") && !strings.HasSuffix(prompt, "?") && !strings.HasSuffix(prompt, "？") {
		prompt = prompt + "。\n"
	}
	if cnf.BotDesc != "" {
		prompt = cnf.BotDesc + "\n" + prompt
	}
	logger.Info("request prompt is %s", prompt)

	var resultText string
	if cnf.Model == gogpt.GPT3Dot5Turbo || cnf.Model == gogpt.GPT3Dot5Turbo0301 {
		resultText, err = c.chatWithGpt35(ctx, cnf, prompt)
	} else {
		resultText, err = c.chatWithGpt30(ctx, cnf, prompt)
	}
	if err != nil {
		logger.Danger("request err is %s", err)
		return
	}
	logger.Info("Response resultText is %s", resultText)
	c.ResponseJson(ctx, http.StatusOK, "", resultText)
}

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

// / chatWithGpt35Custom chatGpt3.5模型 - 自定义请求参数
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
