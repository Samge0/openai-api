# openai-api
本项目可以一键部署属于自己定制化&简易接口权限验证的 openai api 中间请求接口


> 项目当前默认为示例中AI聊天机器人参数，可以根据自己需求定制化。
> 
> **注意，每个参数都可能影响你得到不一样的聊天效果,改变一个参数你就可能得到另一种回答，所以请自己尝试去调试。文档中有二十多中参数示例，等等等...**
> 
> 详情参考官方详细[参数示例](https://beta.openai.com/examples)

# 项目初衷
> 多个项目需要使用openai的api，将功能独立为中间处理程序，方便多个程序访问。


# 使用前提
> 有openai账号，并且创建好api_key，注册事项可以参考[此文章](https://juejin.cn/post/7173447848292253704) 。

# 快速开始

# 基于源码运行(适合了解go语言编程的同学)

````
# 获取项目
$ git clone https://github.com/samge0/openai-api.git

# 进入项目目录
$ cd openai-api

# 复制配置文件
$ copy config.dev.json config.json

# 启动项目
$ go run main.go
````

# 使用docker运行
你可以使用docker快速运行本项目。
`第一种：基于环境变量运行`

```sh
# 运行项目，环境变量参考下方配置说明
$ docker run -itd --name openai-api --restart=always \
 -e APIKEY=换成你的key \
 -e ACCESS_TOKEN=换成你接口请求的token（请求头中的xxx值，Authorization: Bearer xxx） \
 -e ALLOW_ORIGIN=允许访问的域名+多个用英文逗号分隔 \
 -e BOT_DESC="以下是与AI助手的对话。助手乐于助人，富有创造力，聪明且非常友好。" \
 -e MODEL=text-davinci-003 \
 -e MAX_TOKENS=512 \
 -e TEMPREATURE=0.9 \
 -e TOP_P=1 \
 -e FREQ=0.0 \
 -e PRES=0.6 \
 -p 8080:8080 \
 samge/openai-api:v1
```

运行命令中映射的配置文件参考下边的配置文件说明。

`第二种：基于配置文件挂载运行`

```sh
# 复制配置文件，根据自己实际情况，调整配置里的内容
$ cp config.dev.json `pwd`/docker_data/config.json  # 其中 config.dev.json 从项目的根目录获取

# 运行项目
$ docker run -itd --name openai-api --restart=always -v `pwd`/docker_data/config.json:/app/config.json -p 8080:8080 samge/openai-api:latest
```

其中配置文件参考下边的配置文件说明。



# 配置文件说明

````
{
  "api_key": "openai那边的token",
  "access_token": "接口请求的token（请求头中的xxx值，Authorization: Bearer xxx）",
  "allow_origin": "*",
  "port": 8080,
  "bot_desc": "以下是与AI助手的对话。助手乐于助人，富有创造力，聪明且非常友好。",
  "max_tokens": 1024,
  "model": "text-davinci-003",
  "temperature": 0.9,
  "top_p": 1,
  "frequency_penalty": 0.0,
  "presence_penalty": 0.6
}

api_key：openai api_key
access_token:接口请求的token（请求头中的xxx值，Authorization: Bearer xxx），默认为空，可自定义值
allow_origin:允许访问的域名，多个用英文逗号分隔，默认为*，允许所有
bot_desc：AI特征，非常重要，功能等同给与AI一个身份设定（功能风格），默认为空，可自定义值
port: http服务端口
max_tokens: GPT响应字符数，最大2048，默认值512。max_tokens会影响接口响应速度，字符越大响应越慢。
model: GPT选用模型，默认text-davinci-003，具体选项参考官网训练场
temperature: GPT热度，0到1，默认0.9。数字越大创造力越强，但更偏离训练事实，越低越接近训练事实
top_p: 使用温度采样的替代方法称为核心采样，其中模型考虑具有top_p概率质量的令牌的结果。因此，0.1 意味着只考虑包含前 10% 概率质量的代币。
frequency_penalty: -2.0到2.0之间的数字。正值根据它们在文本中的现有频率惩罚新标记，降低模型逐字重复同一行的可能性。
presence_penalty: 数字介于-2.0和2.0之间。正值根据新标记到目前为止是否出现在文本中来惩罚它们，从而增加模型谈论新主题的可能性。
````
更详细的参数配置请参考：[openai官方文档](https://platform.openai.com/docs/api-reference/completions/create)

【注意】：环境变量的优先级高于config.json，如果二者同时配置，则优先取环境变量的值。


### 接口访问
默认接口请求路径：`/api/chat`
请求方式：POST

接口请求示例请求查看：[openai-api接口示例](https://console-docs.apipost.cn/preview/ecd1aadcde480947/04916b4df98a432b)
[!openai-api-接口请求示例](/screenshots/openai-api-demo.jpg)

### 有疑问请添加微信（备注: openai-api），不定期通过解答

**微信号 SamgeApp **