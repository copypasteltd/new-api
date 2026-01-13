package service

import (
	"github.com/ctrlc-ctrlv-limited/cvai/setting/operation_setting"
	"github.com/ctrlc-ctrlv-limited/cvai/setting/system_setting"
)

func GetCallbackAddress() string {
	if operation_setting.CustomCallbackAddress == "" {
		return system_setting.ServerAddress
	}
	return operation_setting.CustomCallbackAddress
}
