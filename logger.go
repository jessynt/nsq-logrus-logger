package nsq_logrus_logger

import (
	"strings"

	"github.com/nsqio/go-nsq"
	log "github.com/sirupsen/logrus"
)

var (
	nsqDebugLevel = nsq.LogLevelDebug.String()
	nsqInfoLevel  = nsq.LogLevelInfo.String()
	nsqWarnLevel  = nsq.LogLevelWarning.String()
	nsqErrLevel   = nsq.LogLevelError.String()
)

type NSQLogrusLogger struct{}

func NewNSQLogrusLogger(l log.Level) (NSQLogrusLogger, nsq.LogLevel) {
	level := nsq.LogLevelWarning
	switch l {
	case log.DebugLevel:
		level = nsq.LogLevelDebug
	case log.InfoLevel:
		level = nsq.LogLevelInfo
	case log.WarnLevel:
		level = nsq.LogLevelWarning
	case log.ErrorLevel:
		level = nsq.LogLevelError
	}
	return NSQLogrusLogger{}, level
}

func (n NSQLogrusLogger) Output(_ int, s string) error {
	if len(s) > 3 {
		msg := strings.TrimSpace(s[3:])
		switch s[:3] {
		case nsqDebugLevel:
			log.Debugln(msg)
		case nsqInfoLevel:
			log.Infoln(msg)
		case nsqWarnLevel:
			log.Warnln(msg)
		case nsqErrLevel:
			log.Errorln(msg)
		default:
			log.Infoln(msg)
		}
	}
	return nil
}
