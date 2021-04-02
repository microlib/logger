package multi

import (
	"fmt"
	"log"
)

const (
	INFO  = "info"
	DEBUG = "debug"
	WARN  = "warn"
	ERROR = "error"
	TRACE = "trace"
	PLAIN = "plain"
	HTML  = "html"
	COLOR = "color"
)

type colorRegistry struct {
	Red     string
	Green   string
	Blue    string
	Yellow  string
	Magento string
}

func NewLogger(mode, level string) *Logger {
	var cr *colorRegistry
	switch mode {
	case "color":
		cr = &colorRegistry{
			Red:     "\x1b[1;31m [%s] \x1b[0m\t: %s",
			Green:   "\x1b[1;36m [%s] \x1b[0m\t: %s",
			Blue:    "\x1b[1;32m [%s] \x1b[0m\t: %s",
			Yellow:  "\x1b[1;33m [%s] \x1b[0m\t: %s",
			Magento: "\x1b[1;35m [%s] \x1b[0m\t: %s",
		}
		break
	case "html":
		cr = &colorRegistry{
			Red:     "<span style=\"color:#eb4034\">&nbsp;[%s]</span>&nbsp;:&nbsp;%s</br>",
			Green:   "<span style=\"color:#34eb40\">&nbsp;[%s]</span>&nbsp;:&nbsp;%s</br>",
			Blue:    "<span style=\"color:#3498eb\">&nbsp;[%s]</span>&nbsp;:&nbsp;%s</br>",
			Yellow:  "<span style=\"color:#eba434\">&nbsp;[%s]</span>&nbsp;:&nbsp;%s</br>",
			Magento: "<span style=\"color:#d234eb\">&nbsp;[%s]</span>&nbsp;:&nbsp;%s</br>",
		}
		break
	case "plain":
		cr = &colorRegistry{
			Red:     " [%s]\t: %s",
			Green:   " [%s]\t: %s",
			Blue:    " [%s]\t: %s",
			Yellow:  " [%s]\t: %s",
			Magento: " [%s]\t: %s",
		}
		break
	}
	return &Logger{Level: level, Colors: cr}
}

type Logger struct {
	Level  string
	Colors *colorRegistry
}

func (logger Logger) Debug(message string) {
	if logger.Level == DEBUG || logger.Level == TRACE {
		log.Printf(logger.Colors.Green, "DEBUG", message)
	}
}

func (logger Logger) Error(message string) {
	log.Printf(logger.Colors.Red, "ERROR", message)
}

func (logger Logger) Warn(message string) {
	if logger.Level == WARN || logger.Level == DEBUG || logger.Level == INFO || logger.Level == TRACE {
		log.Printf(logger.Colors.Yellow, "WARN", message)
	}
}

func (logger Logger) Trace(message string) {
	if logger.Level == TRACE {
		log.Printf(logger.Colors.Magento, "TRACE", message)
	}
}

func (logger Logger) Info(message string) {
	if logger.Level == INFO || logger.Level == DEBUG || logger.Level == TRACE {
		log.Printf(logger.Colors.Blue, "INFO", message)
	}
}

func (logger Logger) Errorf(msg string, val ...interface{}) {
	logger.Error(fmt.Sprintf(msg, val...))
}

func (logger Logger) Debugf(msg string, val ...interface{}) {
	logger.Debug(fmt.Sprintf(msg, val...))
}

func (logger Logger) Warnf(msg string, val ...interface{}) {
	logger.Warn(fmt.Sprintf(msg, val...))
}

func (logger Logger) Tracef(msg string, val ...interface{}) {
	logger.Trace(fmt.Sprintf(msg, val...))
}

func (logger Logger) Infof(msg string, val ...interface{}) {
	logger.Info(fmt.Sprintf(msg, val...))
}
