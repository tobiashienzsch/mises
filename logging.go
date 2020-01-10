package mises

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/gobuffalo/pop/logging"
	"github.com/sirupsen/logrus"
)

// Logger wraps a logrus logger & pop logger
type Logger struct {
	*logrus.Logger
	Debug bool
	Color bool
}

// NewLogger returns a configured logger
func NewLogger() *Logger {
	logger := logrus.New()
	// logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.InfoLevel)
	return &Logger{Logger: logger}
}

func (l *Logger) createDatabaseLogger() func(lvl logging.Level, s string, args ...interface{}) {
	return func(lvl logging.Level, s string, args ...interface{}) {
		if !l.Debug && lvl <= logging.Debug {
			return
		}
		if lvl == logging.SQL {
			if len(args) > 0 {
				xargs := make([]string, len(args))
				for i, a := range args {
					switch a.(type) {
					case string:
						xargs[i] = fmt.Sprintf("%q", a)
					default:
						xargs[i] = fmt.Sprintf("%v", a)
					}
				}
				s = fmt.Sprintf("%s - %s | %s", lvl, s, xargs)
			} else {
				s = fmt.Sprintf("%s - %s", lvl, s)
			}
		} else {
			s = fmt.Sprintf(s, args...)
			s = fmt.Sprintf("%s - %s", lvl, s)
		}
		if l.Color {
			s = color.YellowString(s)
		}
		l.Logger.Infof(s)
	}
}
