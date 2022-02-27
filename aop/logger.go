package aop

import (
	"github.com/fcfcqloow/go-advance/log"
)

func InitLogger(logLevel string) {
	log.SetLevelOrDefault(logLevel, log.LOG_LEVEL_INFO)
}
