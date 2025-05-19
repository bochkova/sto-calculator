package logging

import "github.com/sirupsen/logrus"

type Entry interface {
	Debug(args ...any)
	Debugf(format string, args ...any)
	Info(args ...any)
	Infof(format string, args ...any)
	Warn(args ...any)
	Warnf(format string, args ...any)
	Error(args ...any)
	Errorf(format string, args ...any)
	Fatal(args ...any)
	Fatalf(format string, args ...any)
	Panic(args ...any)
	Panicf(format string, args ...any)
	WithFields(fields Fields) Entry
}

type entryImpl struct {
	entry *logrus.Entry
}

func NewEntry(entry *logrus.Entry) Entry {
	return &entryImpl{entry}
}

func (e *entryImpl) Debug(args ...interface{}) {
	e.entry.Debug(args...)
}

func (e *entryImpl) Debugf(format string, args ...interface{}) {
	e.entry.Debugf(format, args...)
}

func (e *entryImpl) Info(args ...interface{}) {
	e.entry.Info(args...)
}

func (e *entryImpl) Infof(format string, args ...interface{}) {
	e.entry.Infof(format, args...)
}

func (e *entryImpl) Warn(args ...interface{}) {
	e.entry.Warn(args...)
}

func (e *entryImpl) Warnf(format string, args ...interface{}) {
	e.entry.Warnf(format, args...)
}

func (e *entryImpl) Error(args ...interface{}) {
	e.entry.Error(args...)
}

func (e *entryImpl) Errorf(format string, args ...interface{}) {
	e.entry.Errorf(format, args...)
}

func (e *entryImpl) Fatal(args ...interface{}) {
	e.entry.Fatal(args...)
}

func (e *entryImpl) Fatalf(format string, args ...interface{}) {
	e.entry.Fatalf(format, args...)
}

func (e *entryImpl) Panic(args ...interface{}) {
	e.entry.Panic(args...)
}

func (e *entryImpl) Panicf(format string, args ...interface{}) {
	e.entry.Panicf(format, args...)
}

func (e *entryImpl) WithFields(fields Fields) Entry {
	return &entryImpl{e.entry.WithFields(logrus.Fields(fields))}
}
