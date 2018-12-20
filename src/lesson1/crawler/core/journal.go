package core

import (
	"log"
)

type Logger interface {
	Log(level Level, args ...interface{})
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
}

type Level int

func (level Level) String() string {
	switch level {
	case Debug:
		return "debug"
	case Info:
		return "info"
	case Warn:
		return "wain"
	case Error:
		return "error"
	case Fatal:
		return "fatal"
	default:
		return ""
	}
}

const (
	Debug Level = 0
	Info  Level = 1
	Warn  Level = 2
	Error Level = 3
	Fatal Level = 4
)

type Journal struct {
	level Level
}

func (j *Journal) Log(level Level, args ...interface{}) {
	switch level {
	case Debug:
		j.Debug(args)
	case Info:
		j.Info(args)
	case Warn:
		j.Warn(args)
	case Error:
		j.Error(args)
	case Fatal:
		j.Fatal(args)
	default:
		return
	}
}

func (j *Journal) Debug(args ...interface{}) {
	if j.level <= Debug {
		log.Println(args)
	}
}

func (j *Journal) Info(args ...interface{}) {
	if j.level <= Info {
		log.Println(args)
	}
}

func (j *Journal) Warn(args ...interface{}) {
	if j.level <= Warn {
		log.Println(args)
	}
}

func (j *Journal) Error(args ...interface{}) {
	if j.level <= Error {
		log.Println(args)
	}
}

func (j *Journal) Fatal(args ...interface{}) {
	if j.level <= Fatal {
		log.Println(args)
	}
}

func NewJournal(level Level) *Journal {
	return &Journal{
		level: level,
	}
}
