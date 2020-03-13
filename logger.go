/*
Package gwlog is a log package for gamewith that wrap logrus.

gwlog.GetLogger() is a Singleton.

If you set gwlog.GetLogger().SetFormatter or gwlog.GetLogger().SetOutput, all the settings will be inherited.

Usage
	// Setup
	logger := gwlog.GetLogger()
	logger.SetFormatter(&formatter.JSONFormatter{})
	logger.SetOutput(os.Stdout)

	// Simple write log
	logger.Info("aaa")
	// Output: {"level":"INFO","message":"aaa","time":"2000-01-01T00:00:00+09:00"}


	// WithFields
	logger.WithFields(map[string]interface{}{
		"hoge": "hoge"
	}).Info("aaa")
	// Output: {"hoge":"hoge","level":"INFO","message":"aaa","time":"2000-01-01T00:00:00+09:00"}

*/
package gwlog

import (
	"encoding/json"
	"github.com/labstack/gommon/log"
	"github.com/sirupsen/logrus"
	"io"
)

// Logger is gwlog.Logger interface
type Logger interface {
	// SetFormatter is set log format
	//
	// Example:
	// 	gwlog.GetLogger().SetFormatter(&formatter.JSONFormatter{})
	// 	gwlog.GetLogger().Info("hoge")
	// 	> {"message": "hoge", "level": "INFO", "time": "..."}
	SetFormatter(formatter logrus.Formatter)
	// WithFields is SetCustomFields
	//
	// Example:
	// 	gwlog.GetLogger().WithFields(map[string]interface{}{
	// 		"hoge": "dummy",
	// 		"piyo": "dummy2"
	// 	}).Info("aaaa")
	// 	> [INFO][0001] aaaa hoge=dummy piyo=dummy2
	WithFields(fields logrus.Fields) *logrus.Entry
	// Output is get output io
	Output() io.Writer
	// SetOutput is set output io
	//
	// Example:
	// 	gwlog.GetLogger().SetOutput(os.Stdout)
	SetOutput(out io.Writer)
	// Prefix is not used
	Prefix() string
	// SetPrefix is not used
	SetPrefix(_ string)
	// Level is get log level
	Level() log.Lvl
	// SetLevel is set log level
	SetLevel(v log.Lvl)
	// SetHeader is not used
	SetHeader(h string)
	// Print writes to log at log level INFO
	Print(i ...interface{})
	// Println writes to log with a line break at log level INFO
	Println(i ...interface{})
	// Printf writes to log with a format at log level INFO
	Printf(format string, i ...interface{})
	// Printj writes to json log at log level INFO
	Printj(j log.JSON)
	// Debug writes to log at log level DEBUG
	Debug(i ...interface{})
	// Debugln writes to log with a line break at log level DEBUG
	Debugln(i ...interface{})
	// Debugf writes to log with a format at log level DEBUG
	Debugf(format string, i ...interface{})
	// Debugj writes to json log at log level DEBUG
	Debugj(j log.JSON)
	// Info writes to log at log level INFO
	Info(i ...interface{})
	// Infoln writes to log with a line break at log level INFO
	Infoln(i ...interface{})
	// Infof writes to log with a format at log level INFO
	Infof(format string, i ...interface{})
	// Infoj writes to json log at log level INFO
	Infoj(j log.JSON)
	// Warn writes to log at log level WARN
	Warn(i ...interface{})
	// Warnln writes to log with a line break at log level WARN
	Warnln(i ...interface{})
	// Warnf writes to log with a format at log level WARN
	Warnf(format string, i ...interface{})
	// Warnj writes to json log at log level WARN
	Warnj(j log.JSON)
	// Error writes to log at log level ERROR
	Error(i ...interface{})
	// Errorln writes to log with a line break at log level ERROR
	Errorln(i ...interface{})
	// Errorf writes to log with a format at log level ERROR
	Errorf(format string, i ...interface{})
	// Errorj writes to json log at log level ERROR
	Errorj(j log.JSON)
	// Fatal writes to log at log level FATAL
	Fatal(i ...interface{})
	// Fatalln writes to log with a line break at log level FATAL
	Fatalln(i ...interface{})
	// Fatalf writes to log with a format at log level FATAL
	Fatalf(format string, i ...interface{})
	// Fatalj writes to json log at log level FATAL
	Fatalj(j log.JSON)
	// Panic writes to log at log level PANIC
	Panic(i ...interface{})
	// Panicln writes to log with a line break at log level PANIC
	Panicln(i ...interface{})
	// Panicf writes to log with a format at log level PANIC
	Panicf(format string, i ...interface{})
	// Panicj writes to json log at log level PANIC
	Panicj(j log.JSON)
}

type logger struct {
	*logrus.Logger
}

var loggerInstance = &logger{
	Logger: logrus.New(),
}

// GetLogger is get gw logger.
func GetLogger() Logger {
	return loggerInstance
}

func (l *logger) WithFields(fields logrus.Fields) *logrus.Entry {
	return l.Logger.WithFields(fields)
}

func (l *logger) SetFormatter(formatter logrus.Formatter) {
	l.Logger.SetFormatter(formatter)
}

func (l *logger) Output() io.Writer {
	return l.Out
}

func (l *logger) SetOutput(out io.Writer) {
	l.Logger.SetOutput(out)
}

func (l *logger) Prefix() string {
	return ""
}

func (l *logger) SetPrefix(_ string) {
}

func (l *logger) Level() log.Lvl {
	switch l.Logger.Level {
	case logrus.DebugLevel:
		return log.DEBUG
	case logrus.InfoLevel:
		return log.INFO
	case logrus.WarnLevel:
		return log.WARN
	case logrus.ErrorLevel:
		return log.ERROR
	}
	return log.OFF
}

func (l *logger) SetLevel(level log.Lvl) {
	switch level {
	case log.DEBUG:
		l.Logger.Level = logrus.DebugLevel
	case log.INFO:
		l.Logger.Level = logrus.InfoLevel
	case log.WARN:
		l.Logger.Level = logrus.WarnLevel
	case log.ERROR:
		l.Logger.Level = logrus.ErrorLevel
	default:
		l.Logger.Level = logrus.InfoLevel
	}
}

func (l *logger) SetHeader(_ string) {
}

func (l *logger) Print(i ...interface{}) {
	l.Logger.Print(i...)
}

func (l *logger) Println(i ...interface{}) {
	l.Logger.Println(i...)
}

func (l *logger) Printf(format string, i ...interface{}) {
	l.Logger.Printf(format, i...)
}

func (l *logger) Printj(j log.JSON) {
	l.Logger.Println(l.json(j))
}

func (l *logger) Debug(i ...interface{}) {
	l.Logger.Debug(i...)
}

func (l *logger) Debugln(i ...interface{}) {
	l.Logger.Debugln(i...)
}

func (l *logger) Debugf(format string, i ...interface{}) {
	l.Logger.Debugf(format, i...)
}

func (l *logger) Debugj(j log.JSON) {
	l.Logger.Debugln(l.json(j))
}

func (l *logger) Info(i ...interface{}) {
	l.Logger.Info(i...)
}

func (l *logger) Infoln(i ...interface{}) {
	l.Logger.Infoln(i...)
}

func (l *logger) Infof(format string, i ...interface{}) {
	l.Logger.Infof(format, i...)
}

func (l *logger) Infoj(j log.JSON) {
	l.Logger.Infoln(l.json(j))
}

func (l *logger) Warn(i ...interface{}) {
	l.Logger.Warn(i...)
}

func (l *logger) Warnln(i ...interface{}) {
	l.Logger.Warnln(i...)
}

func (l *logger) Warnf(format string, i ...interface{}) {
	l.Logger.Warnf(format, i...)
}

func (l *logger) Warnj(j log.JSON) {
	l.Logger.Warnln(l.json(j))
}

func (l *logger) Error(i ...interface{}) {
	l.Logger.Error(i...)
}

func (l *logger) Errorln(i ...interface{}) {
	l.Logger.Errorln(i...)
}

func (l *logger) Errorf(format string, i ...interface{}) {
	l.Logger.Errorf(format, i...)
}

func (l *logger) Errorj(j log.JSON) {
	l.Logger.Errorln(l.json(j))
}

func (l *logger) Fatal(i ...interface{}) {
	l.Logger.Fatal(i...)
}

func (l *logger) Fatalln(i ...interface{}) {
	l.Logger.Fatalln(i...)
}

func (l *logger) Fatalf(format string, i ...interface{}) {
	l.Logger.Fatalf(format, i...)
}

func (l *logger) Fatalj(j log.JSON) {
	l.Logger.Fatalln(l.json(j))
}

func (l *logger) Panic(i ...interface{}) {
	l.Logger.Panic(i...)
}

func (l *logger) Panicln(i ...interface{}) {
	l.Logger.Panicln(i...)
}

func (l *logger) Panicf(format string, i ...interface{}) {
	l.Logger.Panicf(format, i...)
}

func (l *logger) Panicj(j log.JSON) {
	l.Logger.Panicln(l.json(j))
}

func (l *logger) json(j log.JSON) string {
	b, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}
	return string(b)
}
