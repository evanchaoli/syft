package logger

type nopLogger struct{}

func (l *nopLogger) Errorf(format string, args ...interface{}) {}
func (l *nopLogger) Infof(format string, args ...interface{})  {}
func (l *nopLogger) Info(args ...interface{})                  {}
func (l *nopLogger) Debugf(format string, args ...interface{}) {}
func (l *nopLogger) Debug(args ...interface{})                 {}

// func (l *nopLogger) WithFields(fields map[string]interface{}) Logger { return &nopLogger{} }
