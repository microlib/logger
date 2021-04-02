package multi

import (
	"testing"

	"github.com/microlib/simple"
)

func TestAll(t *testing.T) {

	var log = &simple.Logger{Level: "debug"}

	t.Run("Logging : should pass", func(t *testing.T) {
		log.Info("Testing new logger - this is the old version")

		// plainno color
		l := NewLogger(PLAIN, TRACE)
		l.Info("This is the new version")

		l.Info("Log Info")
		l.Debug("Log Debug")
		l.Trace("Log Trace")
		l.Error("Log Error")
		l.Warn("Log Warn")
		l.Infof("Log Info %s", "test")
		l.Debugf("Log Debug %s", "test")
		l.Tracef("Log Trace %s", "test")
		l.Errorf("Log Error %s", "test")
		l.Warnf("Log Warn %s", "test")

		// with terminal color
		l = NewLogger(COLOR, TRACE)
		l.Info("Log Info")
		l.Debug("Log Debug")
		l.Trace("Log Trace")
		l.Error("Log Error")
		l.Warn("Log Warn")
		l.Infof("Log Info %s", "test")
		l.Debugf("Log Debug %s", "test")
		l.Tracef("Log Trace %s", "test")
		l.Errorf("Log Error %s", "test")
		l.Warnf("Log Warn %s", "test")

		// with html decoration
		l = NewLogger(HTML, TRACE)
		l.Info("Log Info")
		l.Debug("Log Debug")
		l.Trace("Log Trace")
		l.Error("Log Error")
		l.Warn("Log Warn")
		l.Infof("Log Info %s", "test")
		l.Debugf("Log Debug %s", "test")
		l.Tracef("Log Trace %s", "test")
		l.Errorf("Log Error %s", "test")
		l.Warnf("Log Warn %s", "test")

	})
}
