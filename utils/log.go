package utils

import (
	"github.com/zjxpcyc/tinylogger"
)

var logger tinylogger.LogService

// SetLogger 实例化 logger
// 该 logger 会在整个 SDK 中被使用
func SetLogger(l tinylogger.LogService) {
	logger = l
}

// GetLogger 获取 tinylogger.LogService 实例
func GetLogger() tinylogger.LogService {
	if logger == nil {
		logger = new(tinylogger.Logger)
	}

	return logger
}
