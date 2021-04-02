package main

import (
	"fmt"

	"github.com/microlib/logger/pkg/multi"
	"github.com/microlib/simple"
)

// main - needs no explanation :)
func main() {
	// older version
	var log = &simple.Logger{Level: "trace"}
	fmt.Println("Test")
	log.Info("Test")

	// initial new mutli functional logger
	var logger = multi.NewLogger(multi.COLOR, multi.TRACE)

	logger.Error("This is an error message")
	logger.Info("This is an info message")
	logger.Debug("This is a debug message")
	logger.Warn("This is a warn message")
	logger.Trace("This is a trace message")
	logger.Errorf("This is an error message %s", "hello reedge")
	logger.Infof("This is an info message %s", "hello reedge")
	logger.Debugf("This is a debug message %s", "hello reedge")
	logger.Warnf("This is a warn message %s", "hello reedge")
	logger.Tracef("This is a trace message %s", "hello reedge")

	logger = multi.NewLogger(multi.PLAIN, multi.TRACE)
	logger.Error("This is an error message")
	logger.Info("This is an info message")
	logger.Debug("This is a debug message")
	logger.Warn("This is a warn message")
	logger.Trace("This is a trace message")
	logger.Errorf("This is an error message %s", "hello reedge")
	logger.Infof("This is an info message %s", "hello reedge")
	logger.Debugf("This is a debug message %s", "hello reedge")
	logger.Warnf("This is a warn message %s", "hello reedge")
	logger.Tracef("This is a trace message %s", "hello reedge")

	logger = multi.NewLogger(multi.HTML, multi.TRACE)
	logger.Error("This is an error message")
	logger.Info("This is an info message")
	logger.Debug("This is a debug message")
	logger.Warn("This is a warn message")
	logger.Trace("This is a trace message")
	logger.Errorf("This is an error message %s", "hello reedge")
	logger.Infof("This is an info message %s", "hello reedge")
	logger.Debugf("This is a debug message %s", "hello reedge")
	logger.Warnf("This is a warn message %s", "hello reedge")
	logger.Tracef("This is a trace message %s", "hello reedge")

}
