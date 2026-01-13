package service

import (
	"github.com/ctrlc-ctrlv-limited/cvai/service/openaicompat"
	"github.com/ctrlc-ctrlv-limited/cvai/setting/model_setting"
)

func ShouldChatCompletionsUseResponsesPolicy(policy model_setting.ChatCompletionsToResponsesPolicy, channelID int, model string) bool {
	return openaicompat.ShouldChatCompletionsUseResponsesPolicy(policy, channelID, model)
}

func ShouldChatCompletionsUseResponsesGlobal(channelID int, model string) bool {
	return openaicompat.ShouldChatCompletionsUseResponsesGlobal(channelID, model)
}
