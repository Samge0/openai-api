package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"openai-api/app/config"
	"openai-api/app/utils"
	"openai-api/pkg/logger"
	"strings"
	"time"
)

// HandlerChatProxy 回复 - 使用自定义的代理服务器
func (c *ChatController) HandlerChatProxy(ctx *gin.Context) {
	defer utils.TimeCost(time.Now())
	question := &QuestionProxy{}
	err := ctx.BindJSON(question)
	if err != nil {
		c.ResponseJson(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	if question.ApiKey == "" {
		c.ResponseJson(ctx, http.StatusBadRequest, "request api_key is empty", nil)
		return
	}
	if question.ProxyUrl == "" {
		c.ResponseJson(ctx, http.StatusBadRequest, "request proxy_url is empty", nil)
		return
	}
	if len(question.Messages) == 0 || question.Messages[0].Content == "" {
		c.ResponseJson(ctx, http.StatusBadRequest, "request Messages.Content is empty", nil)
		return
	}

	logger.Info("request prompt is ", question.Messages[0].Content)

	var resultText string
	resultText, err = c.chatWithGpt35Proxy(ctx, question)
	if err != nil {
		logger.Danger("HandlerChatProxy request err is ", err)
		c.ResponseJson(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}
	logger.Info("Response resultText is ", resultText)
	c.ResponseJson(ctx, http.StatusOK, "success", resultText)
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

	logger.Info("request prompt is ", question.Messages[0].Content)

	var resultText string
	if strings.HasPrefix(question.Model, "gpt-") {
		resultText, err = c.chatWithGpt35Custom(ctx, question)
	} else {
		resultText, err = c.chatWithGpt30Custom(ctx, question)
	}
	if err != nil {
		logger.Danger("HandlerChatCustom request err is ", err)
		c.ResponseJson(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	logger.Info("Response resultText is ", resultText)
	c.ResponseJson(ctx, http.StatusOK, "success", resultText)
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
	logger.Info("request prompt is ", prompt)

	var resultText string
	if strings.HasPrefix(cnf.Model, "gpt-") {
		resultText, err = c.chatWithGpt35(ctx, cnf, prompt)
	} else {
		resultText, err = c.chatWithGpt30(ctx, cnf, prompt)
	}
	if err != nil {
		logger.Danger("HandlerChat request err is ", err)
		c.ResponseJson(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	logger.Info("Response resultText is ", resultText)
	c.ResponseJson(ctx, http.StatusOK, "success", resultText)
}
