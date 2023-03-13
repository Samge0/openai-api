package controllers

import gogpt "github.com/sashabaranov/go-gpt3"

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

// QuestionCustom 自定义传参的请求体 - 自定义传openai的官方key
type QuestionCustom struct {
	gogpt.ChatCompletionRequest
	ApiKey string `json:"api_key"`
}

// QuestionProxy 自定义代理请求的接口地址 - 自定义传openai的官方key  + 自定义代理中转服务器
type QuestionProxy struct {
	QuestionCustom
	ProxyUrl string `json:"proxy_url"`
}
