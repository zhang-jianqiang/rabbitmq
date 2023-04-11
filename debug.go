package rabbitmq

import "log"

// Debug debug log flag
var Debug bool

func debugf(format string, args ...interface{}) {
	if !Debug {
		return
	}
	log.Printf(format, args...)
}
